package gommer

import (
	"fmt"
	"golang.org/x/mobile/event/touch"
)

// type TouchEvent struct {
// }

/*
 * If the recognizer has the state FAILED, CANCELLED or RECOGNIZED (equals ENDED), it is reset to
 * POSSIBLE to give it another change on the next cycle.
 *
 *               Possible
 *                  |
 *            +-----+---------------+
 *            |                     |
 *      +-----+-----+               |
 *      |           |               |
 *   Failed      Cancelled          |
 *                          +-------+------+
 *                          |              |
 *                      Recognized       Began
 *                                         |
 *                                      Changed
 *                                         |
 *                                  Ended/Recognized
 */

var STATE_POSSIBLE int = 1
var STATE_BEGAN int = 2
var STATE_CHANGED int = 4
var STATE_ENDED int = 8
var STATE_RECOGNIZED int = STATE_ENDED
var STATE_CANCELLED int = 16
var STATE_FAILED int = 32

// シングルトンと化してる
// 後でstruct作って入れる
var state int = STATE_POSSIBLE

func Recognize(e touch.Event) (err error) {
	switch e.Type.String() {
	case "begin":
		state = STATE_BEGAN
	case "move":
		state = STATE_CHANGED
	case "end":
		state = STATE_RECOGNIZED
	}

	if state >= STATE_ENDED {
		fmt.Println("end event detected.")
		state = STATE_POSSIBLE
		return nil
	}
	return nil
}

/**
 * update the recognizer
 * @param {Object} inputData
 */
/*
func recognize(inputData) {
     // make a new copy of the inputData
     // so we can change the inputData without messing up the other recognizers
     var inputDataClone = extend({}, inputData);

     // is is enabled and allow recognizing?
     if (!boolOrFn(this.options.enable, [this, inputDataClone])) {
         this.reset();
         this.state = STATE_FAILED;
         return;
     }

     // reset when we've reached the end
     if (this.state & (STATE_RECOGNIZED | STATE_CANCELLED | STATE_FAILED)) {
         this.state = STATE_POSSIBLE;
     }

     this.state = this.process(inputDataClone);

     // the recognizer has recognized a gesture
     // so trigger an event
     if (this.state & (STATE_BEGAN | STATE_CHANGED | STATE_ENDED | STATE_CANCELLED)) {
         this.tryEmit(inputDataClone);
     }
 }
*/
