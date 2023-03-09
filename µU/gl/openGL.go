package gl

// (c) Christian Maurer   v. 230225 - license see µU.go

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
var (
  lightSource [MaxL][3]float64
  lightColour /* diffus */ [MaxL]col.Colour
  initialized [MaxL]bool
  lightVis bool
  sinL, cosL [MaxL+2]D
  lma, mad, md [4]F
  amb, diff [MaxL][4]F
)

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
  for i := 0; i < 3; i++ { lma[i] = F(0.2) }; lma[3] = F(1) // defaults: 0.2, 1
  C.glLightModelfv (C.GL_LIGHT_MODEL_AMBIENT, &mad[0])
  for i := 0; i < 3; i++ { mad[i] = F(0.2) }; mad[3] = F(1) // defaults: 0.2, 1
  C.glMaterialfv (FrontAndBack, AmbientAndDiffuse, &mad[0])
  clearDepth (1)
  for i := 0; i < 3; i++ { md[i] = F(0.8) }; md[3] = F(1) // defaults: 0.8, 1
  C.glMaterialfv (FrontAndBack, Diffuse, &md[0])
  colorMaterial (FrontAndBack, Diffuse)
  colorMaterial (Front, Ambient)
  colorMaterial (FrontAndBack, AmbientAndDiffuse)
  enable (ColourMaterial)
  enable (Lighting)
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

func colorMaterial (face, mode int) {
  C.glColorMaterial (e(face), e(mode))
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

// Light ///////////////////////////////////////////////////////////////////

func initLight (n uint, x, y, z, a0, a1, a2 float64, r, g, b byte) {
  if n >= MaxL { return }
  if initialized[n] { return }
  initialized[n] = true
  lightVis = true
  amb[n][0], amb[n][1], amb[n][2], amb[n][3] = F(a0), F(a1), F(a2), F(1)
  lightColour[n] = col.New3 (r, g, b)
  diff[n][0], diff[n][1], diff[n][2], diff[n][3] = F(r), F(g), F(b), F(1)
  posLight (n, x, y, z)
  actLight (n)
}

func posLight (n uint, x, y, z float64) {
  if n >= MaxL { return }
  lightSource[n][0], lightSource[n][1], lightSource[n][2] = x, y, z
}

func actLight (n uint) {
  if n >= MaxL { return }
  var L [3]float64
  L[0], L[1], L[2] = lightSource[n][0], lightSource[n][1], lightSource[n][2]
  var l [4]F
  for i := 0; i < 3; i++ { l[i] = F(L[i]) }; l[3] = F(1)
  C.glLightfv (e(n + Light0), Position, &l[0])
  C.glLightfv (e(n + Light0), Ambient, &amb[n][0])
  C.glLightfv (e(n + Light0), Diffuse, &diff[n][0])
  enable (n + Light0)
}

func lamp (n uint) {
  if ! initialized[n] { return }
  xx, yy, zz := lightSource[n][0], lightSource[n][1], lightSource[n][2]
  x, y, z := D(xx), D(yy), D(zz)
  r := D(0.1)
  begin (TriangleFan)
  C.glColor3ub (C.GLubyte(lightColour[n].R()), C.GLubyte(lightColour[n].G()), C.GLubyte(lightColour[n].B()))
  C.glNormal3d (D(0), D(0), D(-1))
  C.glVertex3d (D(x), D(y), D(z + r))
  r0, z0 := r * sinL[1], z + r * cosL[1]
  for l := 0; l <= MaxL; l++ {
    C.glNormal3d (-sinL[1] * cosL[l], -sinL[1] * sinL[l], -cosL[1])
    C.glVertex3d (x + r0 * cosL[l],   y + r0 * sinL[l],   z0)
  }
  end()
  begin (QuadStrip)
  var r1, z1 D
  for b := 1; b <= MaxL / 2 - 2; b++ {
    r0, z0 = r * sinL[b],   z + r * cosL[b]
    r1, z1 = r * sinL[b+1], z + r * cosL[b+1]
    for l := 0; l <= MaxL; l++ {
      C.glNormal3d (-sinL[b+1] * cosL[l], -sinL[b+1] * sinL[l], -cosL[b+1])
      C.glVertex3d (x + r1 * cosL[l], y + r1 * sinL[l], z1)
      C.glNormal3d (-sinL[b] * cosL[l], -sinL[b] * sinL[l], -cosL[b])
      C.glVertex3d (x + r0 * cosL[l], y + r0 * sinL[l], z0)
    }
  }
  end()
  begin (TriangleFan)
  C.glNormal3d (0., 0., 1.)
  C.glVertex3d (x, y, z - r)
  r0, z0 = r * sinL[1], z - r * cosL[1]
  b := MaxL / 2 - 1
  for l := 0; l <= MaxL; l++ {
    C.glNormal3d (-sinL[b] * cosL[l], -sinL[b] * sinL[l], -cosL[b])
    C.glVertex3d (x + r0 * cosL[l], y + r0 * sinL[l], z0)
  }
  end()
}
