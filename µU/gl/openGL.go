package gl

// (c) Christian Maurer   v. 191018 - license see µU.go

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
  f = C.GLfloat
)

func init() {
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
  C.glClearColor (C.GLclampf(r), C.GLclampf(g), C.GLclampf(b), C.GLclampf(0.0))
}

func colour (c col.Colour) {
  r, g, b := c.Float32()
  C.glColor3f (f(r), f(g), f(b))
}

func clearColor (c col.Colour) {
  r, g, b := c.Float64()
  C.glClearColor (C.GLclampf(r), C.GLclampf(g), C.GLclampf(b), C.GLclampf(0));
}

func begin (f Figure) {
//  C.glBegin (C.uint(f))
  C.glBegin (C.GLenum(f))
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
  C.glMatrixMode (C.GLenum(m));
}

func translate (x, y, z float64) {
  C.glTranslatef (C.GLfloat(x), C.GLfloat(y), C.GLfloat(z));
}

func rotate (a, x, y, z float64) {
  C.glRotatef (C.GLfloat(a), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z));
}

func scale (x, y, z float64) {
  C.glScalef (C.GLfloat(x), C.GLfloat(y), C.GLfloat(z));
}

func pushMatrix() {
  C.glPushMatrix();
}

func popMatrix() {
  C.glPopMatrix();
}

func enable (i uint) {
  C.glEnable (C.GLenum(i));
}

func shadeModel (m uint) {
  C.glShadeModel (C.GLenum(m));
}
