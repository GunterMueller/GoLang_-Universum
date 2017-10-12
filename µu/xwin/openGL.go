package xwin

// (c) Christian Maurer   v. 170920 - license see µu.go

// #cgo LDFLAGS: -lGLU
// #include <GL/gl.h>
// #include <GL/glx.h>
// #include <GL/glu.h>
// int etyp (XEvent *e) { return (*e).type; }
// unsigned int kState (XEvent *e) { return (*e).xkey.state; }
// unsigned int kCode (XEvent *e) { return (*e).xkey.keycode; }
import
  "C"
import (
  . "math"
)
const (
  um = Pi / 180; epsilon = 1e-6
  esc = 9; enter = 36; back = 22; tab = 23
  left = 113; right = 114; up = 111; down = 116
  pgUp = 112; pgDown = 117; pgLeft = 113; pgRight = 114
  pos1 = 110; end = 115; ins = 118; del = 119; prt = 107; roll = 78; paus = 127
  f1 = 67; f2 = 68; /* ... */ f10 = 76; f11 = 95; f12 = 96
)
var
  show = true
var
  matrix [3][3]f

func switchShow() { show = ! show }

func (X *xwindow) Start (e0, e1, e2, f0, f1, f2 float) {
  C.glEnable (C.GL_DEPTH_TEST)
  C.glMatrixMode (C.GL_PROJECTION)
  C.glLoadIdentity()
  near, far := f(0.2), f(100)
  p := C.GLdouble(X.wd) / C.GLdouble(X.ht)
//  C.glFrustum (-near, near, -near/p, near/p, near, far
  C.gluPerspective (45, p, near, far)
  X.eye = vector {e0, e1, e2}
  X.focus = vector {f0, f1, f2}
  X.delta = X.distance (X.eye, X.focus)
//  C.glViewport (C.GLint(0), C.GLint(0), C.GLsizei(X.wd), C.GLsizei(X.ht))

//  C.glMatrixMode (C.GL_MODELVIEW)
  X.vec[1] = X.diff (X.focus, X.eye); X.norm (1)
  if eq (e2, f2) {
    X.vec[2] = vector {0, 0, 1}
    X.vec[0] = X.ext (X.vec[1], X.vec[2]); X.norm(0)
println("gleiche Höhe")
  } else { // e2 != f2
    if eq (e0, f0) && eq (e1, f1) {
      X.vec[0] = vector {1, 0, 0}
      if e2 > f2 {
println("von oben")
        X.vec[1] = vector { 0, 0,-1}
        X.vec[2] = vector { 0, 1, 0}
      } else {
println("von unten")
        X.vec[1] = vector { 0, 0, 1}
        X.vec[2] = vector { 0,-1, 0}
      }
    } else {
println("von Seite")
      X.vec[2] = X.clone (X.vec[1])
      v2 := X.vec[1][2]
      if e2 < f2 { v2 = -v2 }
      X.vec[2][2] -= 1 / v2; X.norm(2)
      X.vec[0] = X.ext (X.vec[1], X.vec[2]); X.norm(0)
    }
  }
  X.chk()
}

func (X *xwindow) anfang() {
  return
  C.glMatrixMode (C.GL_MODELVIEW)
  C.glLoadIdentity()
  for i := 0; i < 3; i++ {
    matrix[i][0] = f(X.vec[0][i])
    matrix[i][1] = f(X.vec[1][i])
    matrix[i][2] = f(-X.vec[2][i])
  }
  C.glMultMatrixd (&matrix[0][0])
  C.glTranslated (f(-X.eye[0]), f(-X.eye[1]), f(-X.eye[2]))
}

func (X *xwindow) Draw (draw func()) {
  C.glClear (C.GL_COLOR_BUFFER_BIT + C.GL_DEPTH_BUFFER_BIT)
  C.glMatrixMode (C.GL_MODELVIEW)
  C.glLoadIdentity()
  C.glTranslated (0, 0, -5)
/*
  C.gluLookAt (f(X.eye[0]),    f(X.eye[1]),    f(X.eye[2]),
               f(X.focus[0]),  f(X.focus[1]),  f(X.focus[2]),
               f(X.vec[2][0]), f(X.vec[2][1]), f(X.vec[2][2]))
*/
//               0, 0, 1)
// C.glTranslated (f(X.vec[1][0]), f(X.vec[1][1]), f(X.vec[1][2]))
  X.anfang()
  draw()
  C.glXSwapBuffers (dpy, C.GLXDrawable(X.win))
}

func (X *xwindow) rotate (d int, a float) {
  n := (d + 1) % 3
  X.vec[n] = X.rot (X.vec[d], X.vec[n], a)
  X.vec[n] = X.normv (X.vec[n])
  X.vec[(d + 2) % 3] = X.ext (X.vec[d], X.vec[n])
}

func (X *xwindow) adjustEye() {
  X.eye = X.diff (X.focus, X.dilate (X.delta, X.vec[1]))
//  C.glTranslated (0, -f(X.delta), 0)
}

func (X *xwindow) TurnAroundFocus (d int, a float) {
  X.rotate (d, a)
  X.chk(); println ("rotated", d)
}

func (X *xwindow) Move (d int, a float) {
/*
  temp := X.vec[2]
  X.vec[2] = X.ext (X.vec[0], X.vec[1]); X.norm(2)
  if ! X.eq (temp, X.vec[2]) { println ("Käse") }
*/
  X.eyeOld = X.clone (X.eye)
  X.eye = X.sum (X.eye, X.dilate (a, X.vec[d]))
  X.delta = X.distance (X.eye, X.focus)
//  X.adjustFocus()
  X.chk(); println ("moved")
}

func (X *xwindow) Look (draw func()) {}
