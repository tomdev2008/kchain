package events

import (
	"github.com/json-iterator/go"

	kcfg "kchain/types/cfg"
	klog "kchain/utils/log"
	kstore "kchain/types/store"
	kabci "kchain/types/abci"
)

var (
	cfg = kcfg.GetConfig()
	logger = klog.GetLogWithKeyVals("module", "events")
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	_store = kstore.GetStoreClient()
	_abci = kabci.GetAbciClient()
)
