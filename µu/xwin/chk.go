package xwin

// (c) Christian Maurer   v. 170921 - license see µu.go

// #include <GL/gl.h>
import
  "C"
import (
  . "math"
  "strconv"
)
type (
  float = float64
  f = C.GLdouble
)

func st (a float64) string {
  s := strconv.FormatFloat (a, 'f', 2, 64)
  if a >= 0 { s = " " + s }
  return s
}

func (X *xwindow) chk() {
//  print ("eye foc ", st(X.eye[0]),    " ", st(X.eye[1]),    " ", st(X.eye[2]),    "   ")
//  print (            st(X.focus[0]),  " ", st(X.focus[1]),  " ", st(X.focus[2]));  println()
  println ("eye  ", st(X.eye[0]),    " ", st(X.eye[1]),    " ", st(X.eye[2]))
  println ("rft  ", st(X.vec[0][0]), " ", st(X.vec[0][1]), " ", st(X.vec[0][2]), "   ",
                    st(X.vec[1][0]), " ", st(X.vec[1][1]), " ", st(X.vec[1][2]), "   ",
                    st(X.vec[2][0]), " ", st(X.vec[2][1]), " ", st(X.vec[2][2]))
  if ! X.eq (X.focus, vector {0, 0, 0}) { println ("Mist") }
}

func (X *xwindow) chkef() {
  e := X.distance (X.eye, X.focus)
  if! eq (X.delta, Abs(e)) { println ("Kacke", e,X.delta) }
}

func eq (x, y float) bool {  return Abs (x - y) < epsilon }
