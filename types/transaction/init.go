package transaction

import (
	klog "kchain/utils/log"
	kcfg "kchain/types/cfg"
	"github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	logger = klog.GetLogWithKeyVals("module", "abci.tx")
	cfg = kcfg.GetConfig()
)
