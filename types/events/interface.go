package events

import (
	"sync"
	"fmt"
	"github.com/tendermint/tmlibs/log"
	"os"
)

type Event struct {
	Desc    string
	Name    string
	Handler func([]byte)
}

var (
	once sync.Once
	instance map[string]*Event
)

func Init() {
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With("module", "main")
	logger.Info("init")
	once.Do(func() {
		instance = make(map[string]*Event)
	})
	init_events()

	for k, _ := range instance {
		fmt.Println(k)
	}
}

func GetEvent(name string) *Event {
	return instance[name]
}

func e(name string, desc string, handler func([]byte)) {
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With("module", "main")
	logger.Info("registry", "value", name)
	instance[name] = &Event{
		Name:name,
		Desc:desc,
		Handler:handler,
	}
}
