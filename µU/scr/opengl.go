package scr

// (c) Christian Maurer   v. 201031 - license see µU.go

func (X *screen) Go (m int, draw func(), ox, oy, oz, fx, fy, fz, tx, ty, tz float64) {
  if underX {
    X.XWindow.Go (m, draw, ox, oy, oz, fx, fy, fz, tx, ty, tz)
  }
}
