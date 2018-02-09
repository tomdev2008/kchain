package store

import (
	dbm "github.com/tendermint/tmlibs/db"
	kcfg "kchain/types/cfg"
	"sync"
)

type StoreService struct {
	store *dbm.GoLevelDB
}

func (ss *StoreService) Set(k, v string) {
	ss.store.Set([]byte(k), []byte(v))
}

func (ss *StoreService) Get(k string) string {
	if res := ss.store.Get([]byte(k)); res == nil {
		return ""
	} else {
		return string(res)
	}
}

var (
	once sync.Once
	instance *StoreService
)

func GetStoreClient() *StoreService {
	once.Do(func() {
		cfg := kcfg.GetConfig()
		if store, err := dbm.NewGoLevelDB("app", cfg.Config.DBDir()); err != nil {
			panic(err)
		} else {
			instance = &StoreService{
				store:store,
			}
		}
	})
	return instance
}

func InitStoreClient() *StoreService {
	return GetStoreClient()
}

