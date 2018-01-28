package xwin

// (c) Christian Maurer   v. 171217 - license see µU.go

import (
  "runtime"
  "time"
  . "µU/shape"
)
var
  finished bool

func (X *xwindow) blink() {
  var
    shape Shape
  for {
    X.blinkMutex.Lock()
    if X.cursorShape == Off {
      shape = X.blinkShape
    } else {
      shape = Off
    }
    X.cursor (X.blinkX, X.blinkY, shape)
    X.blinkMutex.Unlock()
    if finished {
      break
    }
    time.Sleep (250 * 1e6)
  }
  runtime.Goexit()
}

func (X *xwindow) doBlink() {
  if X.blinking { return }
  X.blinking = true
  go X.blink()
}

func (X *xwindow) cursor (x, y uint, s Shape) {
  y0, y1:= Cursor (x, y, uint(X.ht1), X.cursorShape, s)
  if y0 + y1 == 0 { return }
  X.cursorShape = s
//  Lock() // weg ?
  xx, yy:= x, y + y0
  xx1, yy1:= xx + uint(X.wd1) - 1, y + y1
  X.RectangleFullInv (int(xx), int(yy), int(xx1), int(yy1))
  X.Flush()
//  Unlock() // weg ?
}

func (X *xwindow) Warp (l, c uint, s Shape) {
  X.WarpGr (uint(X.wd1) * c, uint(X.ht1) * l, s)
}

func (X *xwindow) WarpGr (x, y uint, s Shape) {
  X.blinkMutex.Lock()
  X.blinkX, X.blinkY = x, y
  X.blinkShape = s
  X.cursor (x, y, X.blinkShape)
  X.blinkMutex.Unlock()
}
