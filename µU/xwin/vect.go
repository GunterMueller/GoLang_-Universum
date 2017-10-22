package xwin

// (c) Christian Maurer   v. 170920 - license see µU.go

import
  . "math"


func next (i int) int { return (i + 1) % 3 }
func prev (i int) int { return (i + 2) % 3 }

func (X *xwindow) distance (v, v1 vector) float64 {
  d := 0.
  for i := 0; i < 3; i++ { d += (v1[i] - v[i]) * (v1[i] - v[i]) }
  return Sqrt (d)
}

func (X *xwindow) length (d int) float64 {
  return X.distance (X.vec[d], vector {0, 0, 0})
}

func (X *xwindow) lengthv (v vector) float64 {
  return X.distance (v, vector {0, 0, 0})
}

func (X *xwindow) clone (v vector) vector {
  return vector {v[0], v[1], v[2]}
}

func (X *xwindow) norm (d int) {
  w := X.length (d)
  for i := 0; i < 3; i++ { X.vec[d][i] /= w
    if Abs(X.vec[d][i]) < epsilon { X.vec[d][i] = 0 }
  }
}

func (X *xwindow) normv (v vector) vector {
  w := X.clone (v)
  a := X.lengthv (v)
  for i := 0; i < 3; i++ {
    w[i] /= a
    if Abs(w[i]) < epsilon { w[i] = 0 }
  }
  return w
}

func (X *xwindow) sum (v, w vector) vector {
  return vector {v[0] + w[0], v[1] + w[1], v[2] + w[2]}
}

func (X *xwindow) diff (v, w vector) vector {
  return vector {v[0] - w[0], v[1] - w[1], v[2] - w[2]}
}

func (X *xwindow) dilate (a float64, v vector) vector {
  return vector {v[0] * a, v[1] * a, v[2] * a}
}

func (X *xwindow) inn (v, w vector) float64 {
  y := 0.
  for i := 0; i < 3; i++ { y += v[i] * w[i] }
  return y
}

func (X *xwindow) eq (v, w vector) bool {
  return eq(v[0], w[0]) && eq(v[1], w[1]) && eq(v[2], w[2])
}

func (X *xwindow) ext (v, w vector) vector {
  var u vector
  for i := 0; i < 3; i++ {
    u[i] = v[next(i)] * w[prev(i)] - v[prev(i)] * w[next(i)]
  }
  return u
}

// Pred: e is normed.
// Returns x rotated around e by angle a
func (X *xwindow) rot (e, x vector, a float64) vector {
//  if X.collinear (v, w) { return } // error
  for a <= -180 { a += 360 }
  for a > 180 { a -= 360 }
//  e = e.normv (e) // without Pre
  s, c := Sin (a * um), Cos (a * um)
// return  c * x + (1 - c) * <e, x> * e + s * [e, x]
  return X.sum (X.sum (X.dilate(c, x), X.dilate ((1 - c) * X.inn (e, x), e)), X.dilate (s, X.ext (e, x)))
}

func (X *xwindow) test() {
  v := X.normv (vector {1, 1, 1})
  a := real(120)
  X.vec[0] = X.rot (v, X.vec[0], a)
  X.vec[1] = X.rot (v, X.vec[1], a)
  X.vec[2] = X.rot (v, X.vec[2], a)
  X.chk(); println("rotated")
}
