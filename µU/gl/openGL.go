package gl

// (c) Christian Maurer   v. 210504 - license see µU.go

// #cgo LDFLAGS: -lGL -lGLU
// #include <GL/gl.h>
// #include <GL/glu.h>
// #include <stdio.h>
//// void printerror (GLubyte *es) { printf ("%s\n", es); }
import
  "C"
import
  "µU/col"
type (
  d = C.GLdouble
  l = C.GLclampf
  e = C.GLenum
)

func init() { // TODO
/*
  C.glDepthFunc (C.GL_LESS) // default
  C.glEnable (C.GL_DEPTH_TEST)
  C.glShadeModel (C.GL_SMOOTH)
  C.glClearDepth (C.GLclampd(1.0))
//  C.glColorMaterial (C.GL_FRONT_AND_BACK, C.GL_AMBIENT_AND_DIFFUSE)
//  C.glEnable (C.GL_COLOR_MATERIAL)
//  C.glEnable (C.GL_LIGHTING)
*/
}

func clear() {
  C.glClear (C.GL_COLOR_BUFFER_BIT | C.GL_DEPTH_BUFFER_BIT)
}

func cls (c col.Colour) {
  r, g, b := c.R(), c.G(), c.B()
  C.glClearColor (l(r), l(g), l(b), l(0))
}

func colour (c col.Colour) {
  r, g, b := c.Float32()
  C.glColor3f (C.GLfloat(r), C.GLfloat(g), C.GLfloat(b))
}

func clearColour (c col.Colour) {
  r, g, b := c.Float64()
  C.glClearColor (l(r), l(g), l(b), l(0));
}

func begin (f Figure) {
  C.glBegin (e(f))
}

func end() {
  C.glEnd()
}

func vertex (x, y, z float64) {
  C.glVertex3d (d(x), d(y), d(z))
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
  C.glFrustum (d(l), d(r), d(b), d(t), d(n), d(f))
}

func loadIdentity() {
  C.glLoadIdentity()
}

func viewport (x, y int, w, h uint) {
  C.glViewport (C.GLint(x), C.GLint(y), C.GLsizei(w), C.GLsizei(h));
}

func matrixMode (m uint) {
  C.glMatrixMode (e(m));
}

func translate (x, y, z float64) {
  C.glTranslated (d(x), d(y), d(z))
}

func rotate (a, x, y, z float64) {
  C.glRotated (d(a), d(x), d(y), d(z))
}

func scale (x, y, z float64) {
  C.glScaled (d(x), d(y), d(z))
}

func pushMatrix() {
  C.glPushMatrix();
}

func popMatrix() {
  C.glPopMatrix();
}

func enable (i uint) {
  C.glEnable (e(i));
}

func shadeModel (m uint) {
  C.glShadeModel (e(m));
}

func error() string {
  switch C.glGetError() {
  case 0x500:
    return "invalid enum"
  case 0x501:
    return "invalid value"
  case 0x502:
    return "operation illegal"
  case 0x503:
    return "stack overflow"
  case 0x504:
    return "stack underflow"
  case 0x505:
    return "out of memory"
  }
  return ""
}
