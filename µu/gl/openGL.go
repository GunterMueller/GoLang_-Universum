package gl

// (c) Christian Maurer   v. 170918 - license see µu.go

// #cgo LDFLAGS: -lGL
// #include <GL/gl.h> 
import
  "C"
import (
  "math"
  "µu/col"
)
var (
  initialized bool
  matrix [4][4]C.GLdouble
)

func init() {
  matrix[3][3] = 1
  C.glDepthFunc (C.GL_LESS) // default
  C.glEnable (C.GL_DEPTH_TEST)
  C.glShadeModel (C.GL_SMOOTH)
  w := 1.0
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

func colour (c col.Colour) {
  r, g, b := c.Double()
  C.glColor3f (C.GLfloat(r), C.GLfloat(g), C.GLfloat(b))
}

func vertex (x, y, z float64) {
  C.glVertex3d (C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

func vertexC (x, y, z float64, c col.Colour) {
  colour (c)
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

func cls (r, g, b byte) {
  C.glClearColor (C.GLclampf(r), C.GLclampf(g), C.GLclampf(b), C.GLclampf(0.0))
}

func init0 (w, h uint) {
  if initialized { return }
  initialized = true
  C.glViewport (0, 0, C.GLsizei(w), C.GLsizei(h))
}

func init_(far float64, w, h uint) {
  init0 (w, h)
  C.glMatrixMode (C.GL_PROJECTION)
  C.glLoadIdentity()
  p := float64(w) / float64(h)
  const D = 2 // -fold screen width
  deg := D * math.Atan ((0.5 / D) / p)
  deg /= 0.9 // experimental wideangle correction
  var m [4][4]C.GLdouble
  m[1][1] = 1 / C.GLdouble(math.Tan (deg)) // Cot
  m[0][0] = m[1][1] / C.GLdouble (p)
  deg /= 0.9 // experimental wideangle correction
  const near = C.GLdouble(0.2)
//  delta := C.GLdouble(far) - near
//  m[2][2] = - (C.GLdouble(far) + near) / delta
//  m[2][3] = GLdouble(-1.0)
//  m[3][2] = -2. * near * C.GLdouble(far) / delta
  m[2][2] = C.GLdouble(-1)
  m[2][3] = C.GLdouble(-1)
  m[3][2] = C.GLdouble(-1) * near
  C.glMultMatrixd (&m[0][0])
//  q := C.GLdouble(0.75) // obviously effectless
//  C.glFrustum (-near, near, -q * near, q * near, near, C.GLdouble(far))
  C.glMatrixMode (C.GL_MODELVIEW)
  C.glLoadIdentity()
}

/*
func actualize (r, f, t, a vect.Vector) {
  right[0], right[1], right[2] = r.Coord3()
  front[0], front[1], front[2] = f.Coord3()
  top[0], top[1], top[2] = t.Coord3()
  eye[0], eye[1], eye[2] = a.Coord3()
}
*/
