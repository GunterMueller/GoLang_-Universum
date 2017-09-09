package scr

// (c) Christian Maurer   v. 170818 - license see murus.go

import
  "murus/ker"

func (X *screen) WriteGlx() {
  if underX {
    X.XWindow.WriteGlx()
  } else {
    ker.Panic ("no GUI")
  }
}
