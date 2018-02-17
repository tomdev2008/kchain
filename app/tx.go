package app

import (
	kt "kchain/types"
	kcfg "kchain/types/cfg"
)

type Tx struct {
	SignPubKey string        `json:"pubkey,omitempty" binding:"required"`
	Signature  string        `json:"sign,omitempty" binding:"required"`
	ID         string        `json:"id,omitempty" binding:"required"`
	Data       interface{}   `json:"data,omitempty" binding:"required"`
}

// "健康检查"
func (tx *Tx) Ping() {
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

//lgvd := kt.LoadOrGenPrivValidatorFS(cfg().Config.PrivValidatorFile())
//lgvd.Sign(&ktx.Db{
//	Key:nil,
//	Value:nil,
//}.ToSortBytes())
