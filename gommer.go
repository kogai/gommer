package gommer

import (
	// "github.com/mattn/go-pubsub"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/touch"
)

/**
var gom gommer.Gommer = &{
	//特定の領域のtouch.Eventを送ってくる
	target: target,

	// 必要なら
	options: options,
}

var gom gommer.Gommer = gommer.New()

gom.On("hold", func(event touch.Event){
	// targetのイベントを判定
})

**/

func New(a app.App) Manager {
	m := Manager{
		target: a,
	}
	go m.listen()
	return m
}

type Manager struct {
	target   app.App
	handlers map[string]CallbackFns
	// c chan touch.Event
}

type CallbackFn func(event touch.Event)
type CallbackFns []CallbackFn

func (m Manager) listen() {
	for e := range m.target.Events() {
		switch e := app.Filter(e).(type) {
		case touch.Event:
			// 	gommer.Recognize(e)
		}
	}
}

func (m Manager) On(eventType string, callback CallbackFn) {
	m.handlers[eventType] = append(m.handlers[eventType], CallbackFn)
}

func (m Manager) Off(eventType string) {
	delete(m.handlers[eventType])
}

func (m Manager) emit(eventType string, event touch.Event) {
	handlers := m.handlers[eventType]
	for _, handler := range handlers {
		handler(event)
	}
}
