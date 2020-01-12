package glu

// (c) Christian Maurer   v. 191012 - license see µU.go

// #cgo LDFLAGS: -lGLU
// #include <GL/glu.h>
import
  "C"
type
  d = C.GLdouble

func perspective (fovy, aspect, zNear, zFar float64) {
  C.gluPerspective (d(fovy), d(aspect), d(zNear), d(zFar))
}

func lookAt (ex, ey, ez, cx, cy, cz, ux, uy, uz float64) {
  C.gluLookAt (d(ex), d(ey), d(ez), d(cx), d(cy), d(cz), d(ux), d(uy), d(uz))
}
