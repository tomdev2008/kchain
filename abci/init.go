package abci

import (
	kcfg "kchain/types/cfg"
	klog "kchain/utils/log"
	"github.com/json-iterator/go"
)

var (
	logger = klog.GetLogWithKeyVals("module", "abci")
	cfg = kcfg.GetConfig()
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)
