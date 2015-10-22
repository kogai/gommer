package gommer

import (
	"fmt"
)

type CallbackFn func(event string)

func On(eventType string, secondEventType CallbackFn) {
	fmt.Println(eventType)
	secondEventType("pinchi-out")
}

func Off(eventType string) {

}
