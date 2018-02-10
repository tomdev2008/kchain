package abci

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/tendermint/abci/types"
	crypto "github.com/tendermint/go-crypto"
	"github.com/tendermint/iavl"
	cmn "github.com/tendermint/tmlibs/common"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	"kchain/types/code"
	"kchain/types/transaction"
)

const (
	ValidatorSetChangePrefix string = "val:"
	AccountSetChangePrefix string = "act:"
)

//-----------------------------------------

var _ types.Application = (*PersistentApplication)(nil)

type PersistentApplication struct {
	types.BaseApplication
	state             *iavl.VersionedTree
	ValUpdates        []*types.Validator
	GenesisValidators []*types.Validator
	logger            log.Logger
}

func Run() *PersistentApplication {
	return NewPersistentApplication("kchain", cfg().Config.DBDir(), logger())
}

func NewPersistentApplication(name, dbDir string, log1 log.Logger) *PersistentApplication {
	db, err := dbm.NewGoLevelDB(name, dbDir)
	if err != nil {
		panic(err.Error())
	}

	stateTree := iavl.NewVersionedTree(0, db)
	stateTree.Load()

	return &PersistentApplication{
		state: stateTree,
		logger: log1,
	}
}

func (app *PersistentApplication) SetLogger(l log.Logger) {
	app.logger = l
}

func (app *PersistentApplication) Info(req types.RequestInfo) (res types.ResponseInfo) {
	res.Data = fmt.Sprintf("{\"size\":%v}", app.state.Size())
	res.LastBlockHeight = int64(app.state.LatestVersion())
	res.LastBlockAppHash = app.state.Hash()
	return
}

func (app *PersistentApplication) SetOption(req types.RequestSetOption) types.ResponseSetOption {
	return types.ResponseSetOption{Code: types.CodeTypeOK}
}

// tx is either "val:pubkey/power" or "key=value" or just arbitrary bytes
func (app *PersistentApplication) DeliverTx(txBytes []byte) types.ResponseDeliverTx {
	tx := &transaction.Transaction{}
	tx.FromBytes(txBytes)

	switch tx.Type {
	case transaction.DbSet:
		db, _ := tx.ToDb()
		app.state.Set([]byte(db.Key), []byte(db.Value))

	case transaction.AccountSet:
		account, _ := tx.ToAccount()
		app.state.Set([]byte(AccountSetChangePrefix + account.PubKey), []byte(strconv.Itoa(account.Power)))

	case transaction.ValidatorSet:
		val, _ := tx.ToValidator()
		key := []byte(ValidatorSetChangePrefix + string(val.PubKey))

		if val.Power == 0 {
			// remove validator
			if !app.state.Has(key) {
				return types.ResponseDeliverTx{
					Code: code.CodeTypeUnauthorized.Code,
					Log:  fmt.Sprintf("Cannot remove non-existent validator %X", key)}
			}
			app.state.Remove(key)
		} else {
			// add or update validator
			value := bytes.NewBuffer(make([]byte, 0))
			if err := types.WriteMessage(&types.Validator{val.PubKey, val.Power}, value); err != nil {
				return types.ResponseDeliverTx{
					Code: code.CodeTypeEncodingError.Code,
					Log:  fmt.Sprintf("Error encoding validator: %v", err)}
			}
			app.state.Set(key, value.Bytes())
		}

		// we only update the changes array if we successfully updated the tree
		app.ValUpdates = append(app.ValUpdates, &types.Validator{val.PubKey, val.Power})



	default:
		return types.ResponseDeliverTx{
			Code:code.CodeTypeEncodingError.Code,
			Log:"unknown transaction type",
		}
	}

	return types.ResponseDeliverTx{Code: code.Ok.Code}

	// if it starts with "val:", update the validator set
	// format is "val:pubkey/power"
	if isValidatorTx(tx) {
		// update validators in the merkle tree
		// and in app.ValUpdates
		return app.execValidatorTx(tx)
	}

	// otherwise, update the key-value store
	return types.ResponseDeliverTx{Code: code.Ok.Code}
}

func (app *PersistentApplication) CheckTx(txBytes []byte) types.ResponseCheckTx {
	tx := &transaction.Transaction{}
	if err := tx.FromBytes(txBytes); err != nil {
		return types.ResponseDeliverTx{
			Code:code.CodeTypeEncodingError.Code,
			Log:err.Error(),
		}
	}

	switch tx.Type {
	case transaction.DbSet:
		if _, err := tx.ToDb(); err != nil {
			return types.ResponseDeliverTx{
				Code:code.CodeTypeEncodingError.Code,
				Log:err.Error(),
			}
		}
	case transaction.AccountSet:
		if _, err := tx.ToAccount(); err != nil {
			return types.ResponseDeliverTx{
				Code:code.CodeTypeEncodingError.Code,
				Log:err.Error(),
			}
		}
	case transaction.ValidatorSet:
		if val, err := tx.ToValidator(); err != nil {
			return types.ResponseDeliverTx{
				Code:code.CodeTypeEncodingError.Code,
				Log:err.Error(),
			}
		} else {
			if _, err = crypto.PubKeyFromBytes([]byte(val.PubKey)); err != nil {
				return types.ResponseDeliverTx{
					Code: code.CodeTypeEncodingError.Code,
					Log:  fmt.Sprintf("Pubkey (%X) is invalid go-crypto encoded", val.PubKey)}
			}
		}
	default:
		return types.ResponseDeliverTx{
			Code:code.CodeTypeEncodingError.Code,
			Log:"unknown transaction type",
		}
	}
	return types.ResponseCheckTx{Code: code.Ok.Code}
}

// Commit will panic if InitChain was not called
func (app *PersistentApplication) Commit() (res types.ResponseCommit) {
	// Save a new version for next height
	height := app.state.LatestVersion() + 1
	if appHash, err := app.state.SaveVersion(height); err != nil {
		panic(err)
	} else {
		app.logger.Info("Commit block", "height", height, "root", appHash)
		return types.ResponseCommit{Code: code.Ok.Code, Data: appHash}
	}
}

func (app *PersistentApplication) Query(reqQuery types.RequestQuery) (resQuery types.ResponseQuery) {
	var (
		data = reqQuery.Data
		path = reqQuery.Path
	)

	switch path {
	case transaction.DbGet:

		db := &transaction.Db{}
		if err := json.Unmarshal(data, db); err != nil {
			resQuery.Code = code.CodeTypeBadNonce.Code
			resQuery.Log = err.Error()
			return
		}
		logger.Error(db.Key, "search", "abci")
		index, value := app.state.Get([]byte(db.Key))

		logger.Error(string(value), "search", "abci")

		resQuery.Index = int64(index)
		resQuery.Key = []byte(db.Key)
		resQuery.Value = value

		if value != nil {
			resQuery.Log = "exists"
		} else {
			resQuery.Log = "does not exist"
		}
	default:
		resQuery.Code = code.CodeTypeBadNonce.Code
		resQuery.Log = "wrong path"
	}
	return
}

// Save the validators in the merkle tree
func (app *PersistentApplication) InitChain(req types.RequestInitChain) types.ResponseInitChain {
	for _, v := range req.Validators {
		r := app.updateValidator(v)
		if r.IsErr() {
			app.logger.Error("Error updating validators", "r", r)
		} else {
			// 把创世验证者添加进去
			app.GenesisValidators = append(app.GenesisValidators, v)
		}
	}
	return types.ResponseInitChain{}
}

// Track the block hash and header information
func (app *PersistentApplication) BeginBlock(req types.RequestBeginBlock) types.ResponseBeginBlock {
	// reset valset changes
	app.ValUpdates = make([]*types.Validator, 0)
	app.GenesisValidators = make([]*types.Validator, 0)
	return types.ResponseBeginBlock{}
}

// Update the validator set
func (app *PersistentApplication) EndBlock(req types.RequestEndBlock) types.ResponseEndBlock {
	return types.ResponseEndBlock{ValidatorUpdates: app.ValUpdates}
}

//---------------------------------------------
// update validators
func (app *PersistentApplication) Validators() (validators []*types.Validator) {
	app.state.Iterate(func(key, value []byte) bool {
		if isValidatorTx(key) {
			validator := new(types.Validator)
			err := types.ReadMessage(bytes.NewBuffer(value), validator)
			if err != nil {
				panic(err)
			}
			validators = append(validators, validator)
		}
		return false
	})
	return
}

func MakeValSetChangeTx(pubkey []byte, power int64) []byte {
	return []byte(cmn.Fmt("val:%X/%d", pubkey, power))
}

func isValidatorTx(tx []byte) bool {
	return strings.HasPrefix(string(tx), ValidatorSetChangePrefix)
}

// format is "val:pubkey1/power1,addr2/power2,addr3/power3"tx
func (app *PersistentApplication) execValidatorTx(tx []byte) types.ResponseDeliverTx {
	tx = tx[len(ValidatorSetChangePrefix):]

	//get the pubkey and power
	pubKeyAndPower := strings.Split(string(tx), "/")
	if len(pubKeyAndPower) != 2 {
		return types.ResponseDeliverTx{
			Code: code.CodeTypeEncodingError.Code,
			Log:  fmt.Sprintf("Expected 'pubkey/power'. Got %v", pubKeyAndPower)}
	}
	pubkeyS, powerS := pubKeyAndPower[0], pubKeyAndPower[1]

	// decode the pubkey, ensuring its go-crypto encoded
	pubkey, err := hex.DecodeString(pubkeyS)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: code.CodeTypeEncodingError.Code,
			Log:  fmt.Sprintf("Pubkey (%s) is invalid hex", pubkeyS)}
	}
	_, err = crypto.PubKeyFromBytes(pubkey)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: code.CodeTypeEncodingError.Code,
			Log:  fmt.Sprintf("Pubkey (%X) is invalid go-crypto encoded", pubkey)}
	}

	// decode the power
	power, err := strconv.ParseInt(powerS, 10, 64)
	if err != nil {
		return types.ResponseDeliverTx{
			Code: code.CodeTypeEncodingError.Code,
			Log:  fmt.Sprintf("Power (%s) is not an int", powerS)}
	}

	// update
	return app.updateValidator(&types.Validator{pubkey, power})
}

// add, update, or remove a validator
func (app *PersistentApplication) updateValidator(v *types.Validator) types.ResponseDeliverTx {
	key := []byte("val:" + string(v.PubKey))
	if v.Power == 0 {
		// remove validator
		if !app.state.Has(key) {
			return types.ResponseDeliverTx{
				Code: code.CodeTypeUnauthorized.Code,
				Log:  fmt.Sprintf("Cannot remove non-existent validator %X", key)}
		}
		app.state.Remove(key)
	} else {
		// add or update validator
		value := bytes.NewBuffer(make([]byte, 0))
		if err := types.WriteMessage(v, value); err != nil {
			return types.ResponseDeliverTx{
				Code: code.CodeTypeEncodingError.Code,
				Log:  fmt.Sprintf("Error encoding validator: %v", err)}
		}
		app.state.Set(key, value.Bytes())
	}

	// we only update the changes array if we successfully updated the tree
	app.ValUpdates = append(app.ValUpdates, v)

	return types.ResponseDeliverTx{Code: code.Ok.Code}
}
