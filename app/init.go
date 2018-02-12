package app

import (
	kcfg "kchain/types/cfg"
	tlog "github.com/tendermint/tmlibs/log"
)

var (
	cfg = kcfg.GetConfig()
	logger tlog.Logger = nil
)
