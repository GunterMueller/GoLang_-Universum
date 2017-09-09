package gl

// (c) Christian Maurer   v. 170908 - license see murus.go

// #cgo LDFLAGS: -lGL
// #include <GL/gl.h> 
import
  "C"
import (
  "math"
  "murus/col"
  "murus/scr" // XXX
)
var (
  initialized bool
  matrix [4][4]C.GLdouble
)

////////////////////////////////////////////////////////////////////////////////////////

func init() {
  matrix[3][3] = 1.
  w := 2.0 * math.Pi / float64 (nLamp)
  C.glDepthFunc (C.GL_LESS) // default
  C.glEnable (C.GL_DEPTH_TEST)
  C.glShadeModel (C.GL_SMOOTH)
  w = 1.0
  C.glClearDepth (C.GLclampd(w))
//  C.glColorMaterial (C.GL_FRONT_AND_BACK, C.GL_AMBIENT_AND_DIFFUSE)
//  C.glEnable (C.GL_COLOR_MATERIAL)
//  C.glEnable (C.GL_LIGHTING)
}

func clear() {
  C.glClear (C.GL_COLOR_BUFFER_BIT | C.GL_DEPTH_BUFFER_BIT)
}

func begin (c C.GLenum) {
  C.glBegin (c)
}

func end() {
  C.glEnd()
}

func colour (r, g, b byte) {
  const m = 1<<8 - 1
  C.glColor3f (C.GLfloat(float32(r) / m), C.GLfloat(float32(g) / m), C.GLfloat(float32(b) / m))
}

func vertex (x, y, z float64) {
  C.glVertex3d (C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
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

func cls (c col.Colour) {
  r, g, b := col.Float (c) // 0.0, 0.0, 0.0
  C.glClearColor (C.GLclampf(r), C.GLclampf(g), C.GLclampf(b), C.GLclampf(0.0))
}

func init0() {
  if initialized { return }
  initialized = true
  C.glViewport (0, 0, C.GLsizei(scr.Wd()), C.GLsizei(scr.Ht()))
//                              XXX                  XXX 
}

func init_(far float64) {
  init0()
  C.glMatrixMode (C.GL_PROJECTION)
  C.glLoadIdentity()
  const D = 2.0 // -fold screen width
  deg := D * math.Atan ((0.5 / D) / scr.Proportion())
//                                  XXX
  deg /= 0.9 // experimental wideangle correction
  var m [4][4]C.GLdouble
  m[1][1] = 1.0 / C.GLdouble(math.Tan (deg)) // Cot
  m[0][0] = m[1][1] / C.GLdouble (scr.Proportion())
//                                XXX
  deg /= 0.9 // experimental wideangle correction
  const near = C.GLdouble(0.2)
//  delta := C.GLdouble(far) - near
//  m[2][2] = - (C.GLdouble(far) + near) / delta
//  m[2][3] = GLdouble(-1.0)
//  m[3][2] = -2. * near * C.GLdouble(far) / delta
  m[2][2] = C.GLdouble(-1.0)
  m[2][3] = C.GLdouble(-1.0)
  m[3][2] = C.GLdouble(-1.0) * near
  C.glMultMatrixd (&m[0][0])
//  q := C.GLdouble(0.75)
//  C.glFrustum (-1.0 * near, 1.0 * near, -q * near, q * near, 1.0 * near, C.GLdouble(far))
  C.glMatrixMode (C.GL_MODELVIEW)
  C.glLoadIdentity()
}
