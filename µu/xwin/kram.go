package xwin

import
  . "math"

func (X *xwindow) Focus (d float64) {
  if d < epsilon { return }
  X.delta = d
  X.adjustEye()
}

func (X *xwindow) adjustFocus() {
  for i := 0; i < 3; i++ {
    X.vec[1][i] *= X.delta
    X.focus[i] += X.eye[i]
  }
}

func (X *xwindow) DistanceFrom (x, y, z float64) float64 {
  return X.distance (X.eye, vector {x, y, z})
}

func (X *xwindow) Distance() float64 {
  if Abs (X.distance (X.eye, X.focus) - X.delta) > epsilon { X.adjustFocus() }
  return X.delta
}
