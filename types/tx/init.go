package tx

import (
	"github.com/tendermint/tmlibs/log"
	kcfg "kchain/types/cfg"
	"os"
)

var (
	logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With("module", "app.tx")
	cfg = kcfg.GetConfig()
)