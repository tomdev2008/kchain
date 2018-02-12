package abci

import (
	kcfg "kchain/types/cfg"
	"github.com/json-iterator/go"
)

var (
	cfg = kcfg.GetConfig()
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)
