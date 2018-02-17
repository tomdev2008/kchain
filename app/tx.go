package app

import (
	kt "kchain/types"
	ktx "kchain/types/tx"
	kcfg "kchain/types/cfg"
	"gopkg.in/oauth2.v3/store"
	"github.com/tendermint/abci/types"
)

type Tx struct {
	SignPubKey string        `json:"pubkey,omitempty" binding:"required"`
	Signature  string        `json:"sign,omitempty" binding:"required"`
	ID         string        `json:"id,omitempty" binding:"required"`
	Sync       string        `json:"sync,omitempty" binding:"required"`
	Data       interface{}   `json:"data,omitempty" binding:"required"`
	db         *ktx.Db
	account    *ktx.Account
	validator  *ktx.Validator
	isDone     bool
}

func (tx *Tx) FromBytes(bs []byte) error {
	return json.Unmarshal(bs, tx)
}

func (tx *Tx) ToBytes() []byte {
	d, _ := json.Marshal(tx)
	return d
}

// "健康检查"
func (tx *Tx)Ping() {
	kcfg.DbSet(tx.ID, kt.Result{"data":"pong"})
}

// 节点状态
func (tx *Tx) Status() {
	if res, err := kcfg.Abci().Status(); err != nil {
		kcfg.DbSet(tx.ID, []byte(err.Error()))
	} else {
		if _d, err := json.Marshal(res); err != nil {
			kcfg.DbSet(tx.ID, []byte(err.Error()))
		} else {
			kcfg.DbSet(tx.ID, _d)
		}
	}
}

func (tx *Tx)DbSet() {

	//lgvd := kt.LoadOrGenPrivValidatorFS(cfg().Config.PrivValidatorFile())
	//lgvd.Sign(&ktx.Db{
	//	Key:nil,
	//	Value:nil,
	//}.ToSortBytes())

	if d, err := json.MarshalToString(ktx.Tx{Type:tx.Event, Data:tx.Params, SignPubKey:cfg().Node.PrivValidator().GetPubKey().KeyString(), Signature:nil}); err != nil {
		store.SetErr(tx.ID, err)
	} else {
		if res, err := kcfg.Abci().BroadcastTxCommit(types.Tx(d)); err != nil {
			store.SetErr(tx.ID, err)
		} else {
			store.Set(tx.ID, res)
		}
	}
}

func (tx *Tx)DbGet() {
	if d, err := json.MarshalToString(tx.Params); err != nil {
		store.SetErr(tx.ID, err)
	} else {
		if res, err := abci.ABCIQuery("DbGet", tdata.Bytes(d)); err != nil {
			store.SetErr(tx.ID, err)
		} else {
			store.Set(tx.ID, res.Response)
		}
	}
}
