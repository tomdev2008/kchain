package events

import (
	"github.com/tendermint/tendermint/types"

	kstore "kchain/types/store"
	"github.com/json-iterator/go"

	kabci "kchain/types/abci"
	kt "kchain/types"

	abci_tx "kchain/types/transaction"
	tdata "github.com/tendermint/go-wire/data"

	"github.com/tendermint/tmlibs/log"
	"os"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type Tx kt.Tx

func init_events() {

	var (
		logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With("module", "app.tx")
		store = kstore.GetStoreClient()
		abci = kabci.GetAbciClient()
	)

	e("Ping", "健康检查",
		func(data []byte) {
			tx := &Tx{}
			if err := json.Unmarshal(data, tx); err != nil {
				logger.Error(err.Error())
				return
			}
			store.Set(tx.ID, "pong")
		},
	)

	e("Status", "健康检查",
		func(data []byte) {
			tx := &Tx{}
			if err := json.Unmarshal(data, tx); err != nil {
				logger.Error(err.Error())
				return
			}

			if res, err := abci.Status(); err != nil {
				store.Set(tx.ID, err.Error())
			} else {
				if _d, err := json.MarshalToString(res); err != nil {
					store.Set(tx.ID, err.Error())
				} else {
					store.Set(tx.ID, _d)
					logger.Info(store.Get(tx.ID))
				}
			}
			return
		},
	)

	e("DbSet", "健康检查",
		func(data []byte) {
			tx := &Tx{}
			if err := json.Unmarshal(data, tx); err != nil {
				logger.Error(err.Error())
				return
			}

			if d, err := json.MarshalToString(abci_tx.Transaction{Type:tx.Event, Data:tx.Params}); err != nil {
				store.Set(tx.ID, err.Error())
			} else {
				if res, err := abci.BroadcastTxCommit(types.Tx(d)); err != nil {
					store.Set(tx.ID, err.Error())
				} else {
					d, _ := json.MarshalToString(res)
					store.Set(tx.ID, d)
				}
			}
		},
	)

	e("DbGet", "健康检查",
		func(data []byte) {
			tx := &Tx{}
			if err := json.Unmarshal(data, tx); err != nil {
				logger.Error(err.Error())
				return
			}

			if d, err := json.MarshalToString(tx.Params); err != nil {
				store.Set(tx.ID, err.Error())
			} else {
				if res, err := abci.ABCIQuery("DbGet", tdata.Bytes(d)); err != nil {
					store.Set(tx.ID, err.Error())
				} else {
					d, _ := json.MarshalToString(res.Response)
					store.Set(tx.ID, d)
				}
			}
		},
	)
}