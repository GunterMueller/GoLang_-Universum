package scr

// (c) murus.org  v. 140527 - license see murus.go

import
  . "murus/shape"

func (x *screen) Warp (l, c uint, s Shape) {
  if underX {
    x.Window.Warp (l, c, s)
  } else {
    x.Console.Warp (l, c, s)
  }
}

func (X *screen) WarpGr (x, y uint, s Shape) {
  if underX {
    X.Window.WarpGr (x, y, s)
  } else {
    X.Console.WarpGr (x, y, s)
  }
}
