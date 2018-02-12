package events

import (
	"sync"
)

type event struct {
	Desc    string
	Name    string
	Handler func([]byte)
}

var (
	once sync.Once
	instance map[string]*event
)

func Init() {
	once.Do(func() {
		instance = make(map[string]*event)
		init_events()
	})
}

func GetEvent(name string) *event {
	return instance[name]
}

func e(name string, desc string, handler func([]byte)) {
	logger.Info("register event", "event", name)
	instance[name] = &event{
		Name:name,
		Desc:desc,
		Handler:handler,
	}
}
