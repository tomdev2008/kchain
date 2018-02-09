package log

import (
	tlog "github.com/tendermint/tmlibs/log"
	kcfg "kchain/types/cfg"
	tmflags "github.com/tendermint/tmlibs/cli/flags"
	"os"
)

type Klog struct {
	log tlog.Logger
}

var (
	cfg = kcfg.GetConfig()
)

func Init() *Klog {
	instance, err := tmflags.ParseLogLevel(
		cfg().Config.LogLevel,
		tlog.NewTMLogger(tlog.NewSyncWriter(os.Stderr)),
		"error",
	)
	if err != nil {
		panic(err.Error())
	}

	return &Klog{log:instance}
}

func GetLog() tlog.Logger {
	return Init().log
}

func GetLogWithKeyVals(keyvals ...interface{}) tlog.Logger {
	return Init().log.With(keyvals...)
}
