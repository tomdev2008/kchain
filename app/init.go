package app

import (
	kcfg "kchain/types/cfg"
	tlog "github.com/tendermint/tmlibs/log"

	"github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	cfg = kcfg.GetConfig()
	logger tlog.Logger
)
