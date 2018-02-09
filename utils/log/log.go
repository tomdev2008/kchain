package log

import (
	l "github.com/apex/log"
	"github.com/apex/log/handlers/text"
	j "github.com/apex/log/handlers/json"
	"os"
)

type LogConfig struct {
	Handler string
	Level   string
}

type KLog struct {
	cfg    *LogConfig
	log    *l.Logger
	isInit bool
}

func (kl *KLog)Init(cfg *LogConfig) *KLog {
	kl.cfg = cfg
	kl.isInit = true
	return kl
}

func (kl *KLog)GetLog(module string) *l.Logger {
	if !kl.isInit {
		kl.DefaultLog().WithFields(l.Fielder{
			"log_config":kl.cfg,
			"module":module,
		}).Error("log服务没有初始化,使用默认配置")
	}

	return kl.log.WithField("module", module)
}

func GetDefault() *KLog {
	_kl := &KLog{
		cfg:&LogConfig{
			Handler:"text",
			Level:"DEBUG",
		},
	}
	_kl.DefaultLog()
	return _kl

}

func (kl *KLog)DefaultConfig() *LogConfig {
	return &LogConfig{
		Handler:"text",
		Level:"DEBUG",
	}
}

func (kl *KLog) DefaultLog() *l.Logger {
	kl.Init(kl.DefaultConfig())
	kl.initLog()
	return kl.log
}

func (kl *KLog) initLog() *l.Logger {

	var (
		cfg = kl.cfg
		f *os.File
		handler *l.Handler
	)
	f = os.Stdout

	switch cfg.Handler {
	case "json":
		handler = j.New(f)
	case "text":
		handler = text.New(f)
	default:
		handler = text.New(f)
	}

	l.SetLevelFromString(cfg.Level)
	l.SetHandler(handler)
	kl.log = l.Log

	return kl.log
}


