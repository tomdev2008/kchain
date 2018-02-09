package store

import (
	"github.com/json-iterator/go"

	dbm "github.com/tendermint/tmlibs/db"
	kcfg "kchain/types/cfg"
	kt "kchain/types"
	"sync"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type StoreService struct {
	store *dbm.GoLevelDB
}

func (ss *StoreService) Set(k string, v interface{}) {
	if d, err := json.Marshal(kt.Result{"data":v}); err != nil {
		ss.store.Set([]byte(k), []byte(err.Error()))
	} else {
		ss.store.Set([]byte(k), d)
	}
}

func (ss *StoreService) SetErr(k string, err error) {
	ss.Set(k, err.Error())
}

func (ss *StoreService) Get(k string) interface{} {

	if res := ss.store.Get([]byte(k)); res == nil {
		return ""
	} else {
		res1 := kt.Result{}
		if err := json.Unmarshal(res, &res1); err != nil {
			return err.Error()
		}
		return res1["data"]
	}
}

var (
	once sync.Once
	instance *StoreService
	cfg = kcfg.GetConfig()
)

func GetStoreClient() func() *StoreService {
	return func() *StoreService {
		return InitStoreClient()
	}
}

func InitStoreClient() *StoreService {
	once.Do(func() {
		if store, err := dbm.NewGoLevelDB("app", cfg().Config.DBDir()); err != nil {
			panic(err)
		} else {
			instance = &StoreService{
				store:store,
			}
		}
	})
	return instance
}

