package gommer

import (
	// "fmt"
	// "github.com/mattn/go-pubsub"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/touch"
)

func New(a app.App) Manager {
	m := Manager{
		target: a,
	}
	// go m.listen()
	return m
}

type Manager struct {
	target   app.App
	handlers map[string]callbackFns
	// c chan touch.Event
}

type callbackFn func(event touch.Event)
type callbackFns []callbackFn

/*
func (m Manager) listen() {
	for e := range m.target.Events() {
		switch e := app.Filter(e).(type) {
		case touch.Event:
			// 	gommer.Recognize(e)
		}
	}
}
*/

func (m Manager) On(eventType string, callback callbackFn) {
	var callbacks callbackFns

	if m.handlers[eventType] != nil {
		callbacks = m.handlers[eventType]
	}

	callbacks = append(callbacks, callback)
	m.handlers = map[string]callbackFns{
		eventType: callbacks,
	}
	// m.emit(eventType)
}

/*
func (m Manager) Off(eventType string) {
	delete(m.handlers, eventType)
}
*/

func (m Manager) emit(eventType string, event touch.Event) {
	handlers := m.handlers[eventType]
	for _, handler := range handlers {
		handler(event)
	}
}
