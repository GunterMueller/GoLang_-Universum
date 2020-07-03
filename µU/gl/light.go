package gl

// (c) Christian Maurer   v. 200126 - license see µU.go

// #cgo LDFLAGS: -lGL
// #include <GL/gl.h> 
import
  "C"
import (
  "math"
  "µU/col"
  "µU/vect"
)
var (
  lmAmb, mAmbi, mDiff [4]C.GLfloat
  sinL, cosL [MaxL+2]C.GLdouble
  lightSource [MaxL][3]float64
  lightColour /* diffus */ [MaxL]col.Colour
  initialized [MaxL]bool
  lightVis bool
  aa, dd [MaxL][4]C.GLfloat
)

func init() {
  sinL[0], cosL[0] = C.GLdouble(0.0), C.GLdouble(1.0)
  w := 2 * math.Pi / float64 (MaxL)
  for g := 1; g < MaxL; g++ {
    sinL[g] = C.GLdouble(math.Sin (float64 (g) * w))
    cosL[g] = C.GLdouble(math.Cos (float64 (g) * w))
  }
  sinL[MaxL], cosL[MaxL] = C.GLdouble(0), C.GLdouble(1)
  sinL[MaxL+1], cosL[MaxL+1] = sinL[1], cosL[1]
  C.glDepthFunc (C.GL_LESS) // default
  C.glEnable (C.GL_DEPTH_TEST)
  C.glShadeModel (C.GL_SMOOTH)
  for i := 0; i < 3; i++ { lmAmb[i] = C.GLfloat(0.2) } // default: 0.2
  lmAmb[3] = C.GLfloat(1.0) // default: 1.0
  C.glLightModelfv (C.GL_LIGHT_MODEL_AMBIENT, &lmAmb[0])
  for i := 0; i < 3; i++ { mAmbi[i] = C.GLfloat(0.2) } // default: 0.2
  mAmbi[3] = C.GLfloat(1.0) // default: 1.0
  for i := 0; i < 3; i++ { mDiff[i] = C.GLfloat(0.8) } // default: 0.8
  mDiff[3] = C.GLfloat(1.0) // default: 1.0
  C.glMaterialfv (C.GL_FRONT_AND_BACK, C.GL_AMBIENT_AND_DIFFUSE, &mAmbi[0])
  w = 1.0
  C.glClearDepth (C.GLclampd(w))
  C.glMaterialfv (C.GL_FRONT_AND_BACK, C.GL_DIFFUSE, &mDiff[0])
  C.glColorMaterial (C.GL_FRONT_AND_BACK, C.GL_DIFFUSE)
  C.glColorMaterial (C.GL_FRONT, C.GL_AMBIENT)
  C.glColorMaterial (C.GL_FRONT_AND_BACK, C.GL_AMBIENT_AND_DIFFUSE)
  C.glEnable (C.GL_COLOR_MATERIAL)
  C.glEnable (C.GL_LIGHTING)
}

func initLight (n uint, v, h vect.Vector, r, g, b byte) {
  if n >= MaxL { panic ("gl.initLight: too many lights") }
  if initialized[n] { return }
  initialized[n] = true
  lightVis = true
  x, x1, x2 := h.Coord3()
  aa[n][0], aa[n][1], aa[n][2] = C.GLfloat(x), C.GLfloat(x1), C.GLfloat(x2)
  aa[n][3] = C.GLfloat(1)
  lightColour[n] = col.New3 ("", r, g, b)
  dd[n][0], dd[n][1], dd[n][2] = C.GLfloat(r), C.GLfloat(g), C.GLfloat(b)
  dd[n][3] = C.GLfloat(1)
  lightSource[n][0], lightSource[n][1], lightSource[n][2] = v.Coord3()
  actLight (n)
}

func posLight (n uint, v vect.Vector) {
  if n < MaxL {
    x, x1, x2 := v.Coord3()
    lightSource[n][0], lightSource[n][1], lightSource[n][2] = x, x1, x2
  }
}

func actLight (n uint) {
  if n >= MaxL { panic("oops") }
  var L [4]float64
  L[0], L[1], L[2] = lightSource[n][0], lightSource[n][1], lightSource[n][2]
  var l [4]C.GLfloat
  for i := 0; i < 3; i++ { l[i] = C.GLfloat(L[i]) }; l[3] = C.GLfloat(1.0)
  C.glLightfv (C.GL_LIGHT0 + C.GLenum(n), C.GL_POSITION, &l[0])
  C.glLightfv (C.GL_LIGHT0 + C.GLenum(n), C.GL_AMBIENT, &aa[n][0])
  C.glLightfv (C.GL_LIGHT0 + C.GLenum(n), C.GL_DIFFUSE, &dd[n][0])
  C.glEnable (C.GL_LIGHT0 + C.GLenum(n))
}

func lamp (n uint) {
  if ! initialized[n] { return }
  xx, yy, zz := lightSource[n][0], lightSource[n][1], lightSource[n][2]
  x, y, z := C.GLdouble(xx), C.GLdouble(yy), C.GLdouble(zz)
  r := C.GLdouble(0.1)
  begin (TriangleFan)
  C.glColor3ub (C.GLubyte(lightColour[n].R()), C.GLubyte(lightColour[n].G()), C.GLubyte(lightColour[n].B()))
  C.glNormal3d (C.GLdouble(0.0), C.GLdouble(0.0), C.GLdouble(-1.0))
  C.glVertex3d (C.GLdouble(x), C.GLdouble(y), C.GLdouble(z + r))
  r0, z0 := r * sinL[1], z + r * cosL[1]
  for l := 0; l <= MaxL; l++ {
    C.glNormal3d (-sinL[1] * cosL[l], -sinL[1] * sinL[l], -cosL[1])
    C.glVertex3d (x + r0 * cosL[l],   y + r0 * sinL[l],   z0)
  }
  end()
  begin (QuadStrip)
  var r1, z1 C.GLdouble
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
