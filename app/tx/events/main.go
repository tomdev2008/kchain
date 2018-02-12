package events

import (
	"github.com/tendermint/tendermint/types"

	kt "kchain/types"

	abci_tx "kchain/types/transaction"
	tdata "github.com/tendermint/go-wire/data"
)

type Tx kt.Tx

func init_events() {
	store := _store()
	abci := _abci()

	e("Ping", "健康检查",
		func(data []byte) {
			tx := &Tx{}
			if err := json.Unmarshal(data, tx, ); err != nil {
				store.SetErr(tx.ID, err)
				return
			}
			store.Set(tx.ID, kt.Result{"data":"pong"})
		},
	)

	e("Status", "健康检查",
		func(data []byte) {
			tx := &Tx{}
			if err := json.Unmarshal(data, tx); err != nil {
				store.SetErr(tx.ID, err)
				return
			}

			if res, err := abci.Status(); err != nil {
				store.SetErr(tx.ID, err)
			} else {
				if _d, err := json.MarshalToString(res); err != nil {
					store.SetErr(tx.ID, err)
				} else {
					store.Set(tx.ID, _d)
				}
			}
			return
		},
	)

	e("DbSet", "健康检查",
		func(data []byte) {
			tx := &Tx{}
			if err := json.Unmarshal(data, tx); err != nil {
				store.SetErr(tx.ID, err)
				return
			}

			tx.Params

			lgvd := kt.LoadOrGenPrivValidatorFS(cfg().Config.PrivValidatorFile())
			lgvd.Sign(&abci_tx.Db{
				Key:nil,
				Value:nil,
			}.ToSortBytes())

			if d, err := json.MarshalToString(abci_tx.Transaction{Type:tx.Event, Data:tx.Params, SignPubKey:cfg().Node.PrivValidator().GetPubKey().KeyString(), Signature:nil}); err != nil {
				store.SetErr(tx.ID, err)
			} else {
				if res, err := abci.BroadcastTxCommit(types.Tx(d)); err != nil {
					store.SetErr(tx.ID, err)
				} else {
					store.Set(tx.ID, res)
				}
			}
		},
	)

	e("DbGet", "健康检查",
		func(data []byte) {
			tx := &Tx{}
			if err := json.Unmarshal(data, tx); err != nil {
				store.SetErr(tx.ID, err)
				return
			}

			if d, err := json.MarshalToString(tx.Params); err != nil {
				store.SetErr(tx.ID, err)
			} else {
				if res, err := abci.ABCIQuery("DbGet", tdata.Bytes(d)); err != nil {
					store.SetErr(tx.ID, err)
				} else {
					store.Set(tx.ID, res.Response)
				}
			}
		},
	)
}