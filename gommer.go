package gommer

import (
	"golang.org/x/mobile/event/touch"
)

type CallbackFn func(event touch.Event)

var fn CallbackFn

func On(eventType string, callback CallbackFn) {
	fn = callback
}

func Off(eventType string) {

}

func Emit(eventType string, e touch.Event) {
	fn(e)
}
