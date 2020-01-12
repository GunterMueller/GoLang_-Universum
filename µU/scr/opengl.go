package scr

// (c) Christian Maurer   v. 191019 - license see µU.go

import (
  "µU/ker"
  "µU/show"
)

func (X *screen) Start (m show.Mode, draw func(), ex, ey, ez, fx, fy, fz, nx, ny, nz float64) {
  if underX {
    X.XWindow.Start (m, draw, ex, ey, ez, fx, fy, fz, nx, ny, nz)
  } else {
    ker.Panic ("no GUI")
  }
}

func (X *screen) Go() {
  if underX {
    X.XWindow.Go()
  } else {
    ker.Panic ("no GUI")
  }
}
