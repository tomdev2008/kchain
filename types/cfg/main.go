package cfg

import (
	cfg "github.com/tendermint/tendermint/config"
	"sync"
)

type AppConfig struct {
	Name string
	Addr string
}

type services struct {
	Config *cfg.Config
	App    *AppConfig
}

var (
	once sync.Once
	instance *services
)

func initServices() *services {
	return &services{
		Config:cfg.DefaultConfig(),
		App:&AppConfig{
			Name:"kchain",
			Addr:":9000",
		},
	}
}

func GetConfig() *services {
	once.Do(func() {
		instance = initServices()
	})
	return instance
}

