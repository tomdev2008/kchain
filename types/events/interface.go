package events

import (
	"sync"
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
	once.Do(func() {
		instance = make(map[string]*Event)
		init_events()
	})
}

func GetEvent(name string) *Event {
	return instance[name]
}

func e(name string, desc string, handler func([]byte)) {
	logger.Info("register event", "event", name)
	instance[name] = &Event{
		Name:name,
		Desc:desc,
		Handler:handler,
	}
}
