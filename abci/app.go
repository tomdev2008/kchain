package abci

import (
	"fmt"

	"github.com/tendermint/abci/types"
	"github.com/tendermint/iavl"
	dbm "github.com/tendermint/tmlibs/db"

	"kchain/types/transaction"

	"kchain/types/code"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var _ types.Application = (*KchainApplication)(nil)

type KchainApplication struct {
	types.BaseApplication

	state *iavl.VersionedTree
}

func NewKchainApplication() *KchainApplication {
	state := iavl.NewVersionedTree(0, dbm.NewMemDB())
	return &KchainApplication{state: state}
}

func (app *KchainApplication) Info(req types.RequestInfo) types.ResponseInfo {
	return types.ResponseInfo{Data: fmt.Sprintf("{\"size\":%v}", app.state.Size())}
}

func (app *KchainApplication) DeliverTx(txBytes []byte) types.ResponseDeliverTx {

	tx := &transaction.Transaction{}
	if err := tx.FromBytes(txBytes); err != nil {
		return types.ResponseDeliverTx{
			Code:code.CodeTypeEncodingError.Code,
			Log:err.Error(),
		}
	}

	logger.Error(string(tx.Type))

	switch tx.Type {
	case transaction.DbSet:
		db := &transaction.Db{}
		if err := tx.ToDb(db); err != nil {
			return types.ResponseDeliverTx{
				Code:code.CodeTypeEncodingError.Code,
				Log:err.Error(),
			}
		} else {
			logger.Error(db.Key)
			logger.Error(db.Value)

			app.state.Set([]byte(db.Key), []byte(db.Value))
		}

	default:
		return types.ResponseDeliverTx{
			Code:code.CodeTypeEncodingError.Code,
			Log:"unknown transaction type",
		}
	}

	return types.ResponseDeliverTx{Code: code.Ok.Code}
}

func (app *KchainApplication) CheckTx(tx []byte) types.ResponseCheckTx {
	return types.ResponseCheckTx{Code: code.Ok.Code}
}

func (app *KchainApplication) Commit() (res types.ResponseCommit) {

	var hash []byte
	var err error
	if app.state.Size() < 0 {
		res.Code = code.CodeTypeBadNonce.Code
		res.Log = "size is small"
		return

	}

	app.state.Hash()
	if hash, err = app.state.SaveVersion(app.state.LatestVersion() + 1); err != nil {
		res.Code = code.CodeTypeBadNonce.Code
		res.Log = err.Error()
	} else {
		res.Code = code.Ok.Code
		res.Data = hash
	}
	return
}

func (app *KchainApplication) Query(reqQuery types.RequestQuery) (resQuery types.ResponseQuery) {
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


