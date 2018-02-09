package abci

import (
	"os"

	"github.com/tendermint/tmlibs/log"
	kcfg "kchain/types/cfg"
)

var (
	logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With("module", "abci")
	cfg = kcfg.GetConfig()
)
