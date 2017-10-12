package scr

// (c) Christian Maurer   v. 170814 - license see µu.go

import
  . "µu/shape"

func (x *screen) Warp (l, c uint, s Shape) {
  if underX {
    x.XWindow.Warp (l, c, s)
  } else {
    x.Console.Warp (l, c, s)
  }
}

func (X *screen) WarpGr (x, y uint, s Shape) {
  if underX {
    X.XWindow.WarpGr (x, y, s)
  } else {
    X.Console.WarpGr (x, y, s)
  }
}
