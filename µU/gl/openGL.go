package gl

// #cgo LDFLAGS: -lGL -lGLU
// #include <GL/gl.h>
// #include <GL/glu.h>
// // include <stdio.h>
import
  "C"
import (
  "math"
  "µU/col"
)
type (
  d = C.GLdouble
  e = C.GLenum
  l = C.GLclampf
)
var
  sinL, cosL [MaxL+2]D

func init() {
  sinL[0], cosL[0] = d(0), d(1)
  w := 2 * math.Pi / float64 (MaxL)
  for g := 1; g < MaxL; g++ {
    sinL[g] = d(math.Sin (float64(g) * w))
    cosL[g] = d(math.Cos (float64(g) * w))
  }
  sinL[MaxL], cosL[MaxL] = d(0), d(1)
  sinL[MaxL+1], cosL[MaxL+1] = sinL[1], cosL[1]
  C.glDepthFunc (C.GL_LESS)
  enable (Depthtest)
  shadeModel (Smooth)
}

func clear() {
  C.glClear (C.GL_COLOR_BUFFER_BIT | C.GL_DEPTH_BUFFER_BIT)
}

func clearDepth (d float64) {
  C.glClearDepth (C.GLclampd(d))
}

func cls (c col.Colour) {
  r, g, b := c.R(), c.G(), c.B()
  C.glClearColor (l(r), l(g), l(b), l(0))
}

func colour (c col.Colour) {
  r, g, b := c.Float32()
  C.glColor3f (F(r), F(g), F(b))
}

func clearColour (c col.Colour) {
  r, g, b := c.Float64()
  C.glClearColor (l(r), l(g), l(b), l(0))
}

func linewidth (w float64) {
  C.glLineWidth (F(w))
}

func begin (f Figure) {
  C.glBegin (e(f))
}

func end() {
  C.glEnd()
}

func vertex (x, y, z float64) {
  C.glVertex3d (D(x), D(y), D(z))
}

func newList (n uint) {
  C.glNewList (C.GLuint(n), C.GL_COMPILE_AND_EXECUTE)
}

func endList() {
  C.glEndList()
}

func callList (n uint) {
  C.glCallList(C.GLuint(n))
}

func frustum (l, r, b, t, n, f float64) {
  C.glFrustum (D(l), D(r), D(b), D(t), D(n), D(f))
}

func ortho (l, r, b, t, n, f float64) {
  C.glOrtho (D(l), D(r), D(b), D(t), D(n), D(f))
}

func loadIdentity() {
  C.glLoadIdentity()
}

func viewport (x, y int, w, h uint) {
  C.glViewport (C.GLint(x), C.GLint(y), C.GLsizei(w), C.GLsizei(h))
}

func matrixMode (m uint) {
  C.glMatrixMode (e(m))
}

func translate (x, y, z float64) {
  C.glTranslated (D(x), D(y), D(z))
}

func rotate (a, x, y, z float64) {
  C.glRotated (D(a), D(x), D(y), D(z))
}

func scale (x, y, z float64) {
  C.glScaled (D(x), D(y), D(z))
}

func pushMatrix() {
  C.glPushMatrix()
}

func popMatrix() {
  C.glPopMatrix()
}

func enable (i uint) {
  C.glEnable (e(i))
}

func shadeModel (m uint) {
  C.glShadeModel (e(m))
}

func flush() {
  C.glFlush()
}

func error() string {
  switch C.glGetError() {
  case C.GL_INVALID_ENUM:
    return "invalid enum"
  case C.GL_INVALID_VALUE:
    return "invalid value"
  case C.GL_INVALID_OPERATION:
    return "invalid operation"
  case C.GL_STACK_OVERFLOW:
    return "stack overflow"
  case C.GL_STACK_UNDERFLOW:
    return "stack underflow"
  case C.GL_OUT_OF_MEMORY:
    return "out of memory"
  }
  return ""
}
