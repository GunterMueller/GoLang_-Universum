package scr

// (c) murus.org  v. 170814 - license see murus.go

import
  . "murus/shape"

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
