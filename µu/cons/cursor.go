package cons

// (c) Christian Maurer   v. 140615 - license see µu.go

import (
  "runtime"
  . "µu/shape"
  "µu/ker"
)
var
  finished bool

func (x *console) blink() {
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

func (x *console) doBlink() {
  if x.blinking { return }
  go x.blink()
  x.blinking = true
}

func (X *console) cursor (x, y uint, s Shape) {
  y0, y1 := Cursor (x, y, X.ht1, X.cursorShape, s)
  if y0 + y1 == 0 { return }
  X.cursorShape = s
//  Lock() // weg ?
  x -= uint(X.x); y -= uint(X.y)
  X.RectangleFullInv (int(x), int(y + y0), int(x + X.wd1) - 1, int(y + y1))
//  Unlock() // weg ?
}

func (x *console) Warp (l, c uint, s Shape) {
  x.WarpGr (x.wd1 * c, x.ht1 * l, s)
}

func (X *console) WarpGr (x, y uint, s Shape) {
  x += uint(X.x); y += uint(X.y)
  X.blinkMutex.Lock()
  X.blinkX, X.blinkY = x, y
  X.blinkShape = s
  X.cursor (x, y, X.blinkShape)
  X.blinkMutex.Unlock()
}
