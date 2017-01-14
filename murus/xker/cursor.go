package xker

// (c) murus.org  v. 140615 - license see murus.go

import (
  "runtime"
  . "murus/shape"
  "murus/ker"
)
var
  finished bool

func (x *window) blink() {
  var
    shape Shape
  for {
    x.blinkMutex.Lock()
    if x.cursorShape == Off {
      shape = x.blinkShape
    } else {
      shape = Off
    }
    x.cursor (x.blinkX, x.blinkY, shape)
    x.blinkMutex.Unlock()
    if finished {
      break
    }
    ker.Msleep (250)
  }
  runtime.Goexit()
}

func (x *window) doBlink() {
  if x.blinking { return }
  x.blinking = true
  go x.blink()
}

func (X *window) cursor (x, y uint, s Shape) {
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

func (x *window) Warp (l, c uint, s Shape) {
  x.WarpGr (uint(x.wd1) * c, uint(x.ht1) * l, s)
}

func (X *window) WarpGr (x, y uint, s Shape) {
  X.blinkMutex.Lock()
  X.blinkX, X.blinkY = x, y
  X.blinkShape = s
  X.cursor (x, y, X.blinkShape)
  X.blinkMutex.Unlock()
}
