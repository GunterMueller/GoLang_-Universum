package gl

// (c) Christian Maurer   v. 211127 - license see µU.go

// #cgo LDFLAGS: -lGL
// #include <GL/gl.h> 
import
  "C"
import (
  "math"
  "µU/col"
//  "µU/vect"
)
var (
  lmAmb, mAmbi, mDiff [4]C.GLfloat
  sinL, cosL [MaxL+2]C.GLdouble
  lightSource [MaxL][3]float64
  lightColour /* diffus */ [MaxL]col.Colour
  initialized [MaxL]bool
  lightVis bool
  amb, diff [MaxL][4]C.GLfloat
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
  C.glDepthFunc (C.GL_LESS) // default
  C.glEnable (C.GL_DEPTH_TEST)
  C.glShadeModel (C.GL_SMOOTH)
  for i := 0; i < 3; i++ { lmAmb[i] = C.GLfloat(0.2) } // default: 0.2
  lmAmb[3] = C.GLfloat(1) // default: 1
  C.glLightModelfv (C.GL_LIGHT_MODEL_AMBIENT, &lmAmb[0])
  for i := 0; i < 3; i++ { mAmbi[i] = C.GLfloat(0.2) } // default: 0.2
  mAmbi[3] = C.GLfloat(1) // default: 1
  for i := 0; i < 3; i++ { mDiff[i] = C.GLfloat(0.8) } // default: 0.8
  mDiff[3] = C.GLfloat(1) // default: 1
  C.glMaterialfv (C.GL_FRONT_AND_BACK, C.GL_AMBIENT_AND_DIFFUSE, &mAmbi[0])
  w = 1
  C.glClearDepth (C.GLclampd(w))
  C.glMaterialfv (C.GL_FRONT_AND_BACK, C.GL_DIFFUSE, &mDiff[0])
  C.glColorMaterial (C.GL_FRONT_AND_BACK, C.GL_DIFFUSE)
  C.glColorMaterial (C.GL_FRONT, C.GL_AMBIENT)
  C.glColorMaterial (C.GL_FRONT_AND_BACK, C.GL_AMBIENT_AND_DIFFUSE)
  C.glEnable (C.GL_COLOR_MATERIAL)
  C.glEnable (C.GL_LIGHTING)
}

func initLight (n uint, x, y, z, a0, a1, a2 float64, r, g, b byte) {
  if n >= MaxL { return }
  if initialized[n] { return }
  initialized[n] = true
  lightVis = true
  amb[n][0], amb[n][1], amb[n][2] = C.GLfloat(a0), C.GLfloat(a1), C.GLfloat(a2)
  amb[n][3] = C.GLfloat(1)
  lightColour[n] = col.New3 (r, g, b)
  diff[n][0], diff[n][1], diff[n][2] = C.GLfloat(r), C.GLfloat(g), C.GLfloat(b)
  diff[n][3] = C.GLfloat(1)
  lightSource[n][0], lightSource[n][1], lightSource[n][2] = x, y, z
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
  var l [4]C.GLfloat
  for i := 0; i < 3; i++ { l[i] = C.GLfloat(L[i]) }; l[3] = C.GLfloat(1)
  C.glLightfv (C.GL_LIGHT0 + C.GLenum(n), C.GL_POSITION, &l[0])
  C.glLightfv (C.GL_LIGHT0 + C.GLenum(n), C.GL_AMBIENT, &amb[n][0])
  C.glLightfv (C.GL_LIGHT0 + C.GLenum(n), C.GL_DIFFUSE, &diff[n][0])
  C.glEnable (C.GL_LIGHT0 + C.GLenum(n))
}

func lamp (n uint) {
  if ! initialized[n] { return }
  xx, yy, zz := lightSource[n][0], lightSource[n][1], lightSource[n][2]
  x, y, z := d(xx), d(yy), d(zz)
  r := d(0.1)
  begin (TriangleFan)
  C.glColor3ub (C.GLubyte(lightColour[n].R()), C.GLubyte(lightColour[n].G()), C.GLubyte(lightColour[n].B()))
  C.glNormal3d (d(0), d(0), d(-1))
  C.glVertex3d (d(x), d(y), d(z + r))
  r0, z0 := r * sinL[1], z + r * cosL[1]
  for l := 0; l <= MaxL; l++ {
    C.glNormal3d (-sinL[1] * cosL[l], -sinL[1] * sinL[l], -cosL[1])
    C.glVertex3d (x + r0 * cosL[l],   y + r0 * sinL[l],   z0)
  }
  end()
  begin (QuadStrip)
  var r1, z1 d
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
