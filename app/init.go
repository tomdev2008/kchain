package app

import (
	kcfg "kchain/types/cfg"
	klog "kchain/utils/log"
)

var (
	cfg = kcfg.GetConfig()
	logger = klog.GetLogWithKeyVals("module", "app")
)
