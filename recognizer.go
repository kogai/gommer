package gommer

import (
	"fmt"
	"golang.org/x/mobile/event/touch"
	"math"
	"time"
)

// type Recognizer struct {
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

/*
- [x] hold(short/long)
- [x] swipe(direction)
- [x] dobble-tap
*/

const timeout time.Duration = 500 * time.Millisecond

var (
	start  time.Time
	touchX float32
	touchY float32
)

func Recognize(e touch.Event) (err error) {

	switch e.Type.String() {
	case "begin":
		touchX = e.X
		touchY = e.Y
		state = STATE_BEGAN

		/*
			var onTouchStart time.Time = time.Now()
			go detectHold(onTouchStart, touchX, touchY)
		*/
	case "move":
		state = STATE_CHANGED
		/*
			go detectSwipe(touchX, touchY, e.X, e.Y)
		*/
	case "end":
		state = STATE_RECOGNIZED
		go detectDoubleTap(touchX, touchY)
	}

	if state >= STATE_ENDED {
		fmt.Println("end event detected.\n")
		state = STATE_POSSIBLE
		return nil
	}
	return nil
}

func detectDoubleTap(touchX, touchY float32) {
	var onTouchStart time.Time = time.Now()
	for {
		var now time.Time = time.Now()
		// 一定時間が経過していない
		if now.Sub(onTouchStart) < timeout {
			// 次のイベントが開始した
			if state == STATE_BEGAN {
				fmt.Printf("double tap event detected. state: %d\n", state)
				break
			}
		} else {
			break
		}
	}
}

func detectSwipe(srcX, srcY, destX, destY float32) {
	var degree float32 = getDegree(srcX, srcY, destX, destY)
	var direction string = getDirection(degree)
	state = STATE_RECOGNIZED
	// TODO event publish
	fmt.Printf("[%g]swipe event direction to: %s\n", degree, direction)
}

func detectHold(onTouchStart time.Time, touchX, touchY float32) {
	for {
		var now time.Time = time.Now()
		// 一定時間が経過した
		if now.Sub(onTouchStart) > timeout {
			// イベントが終了していない
			if state > STATE_POSSIBLE {
				// タップ位置が動いていない
				// TODO event publish
				fmt.Printf("hold event detected. state: %d\n", state)
				state = STATE_RECOGNIZED
			}
			break
		}
	}
}

func getDegree(srcX, srcY, destX, destY float32) float32 {
	var distanseX float64 = float64(destX - srcX)
	var distanseY float64 = float64(destY - srcY)
	radian := math.Atan2(distanseX, distanseY)
	degree := radian * 180 / math.Pi
	if degree < 0 {
		degree = degree * -1
		degree = 360 - degree
	}
	return float32(degree)
}

func getDirection(degree float32) string {
	/*
		135 ~ 225 up
		45 ~ 135 right
		315 ~ 45 down
		225 ~ 315 left
	*/
	var direction string
	if degree > 135 && degree < 225 {
		direction = "up"
	}
	if degree > 45 && degree < 135 {
		direction = "right"
	}
	if degree > 315 || degree < 45 {
		direction = "down"
	}
	if degree > 225 && degree < 315 {
		direction = "left"
	}
	return direction
}

/*
func reset() {

}
*/
