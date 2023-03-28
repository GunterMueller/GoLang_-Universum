package gl

// (c) Christian Maurer   v. 230322 - license see µU.go

// #cgo LDFLAGS: -lGL -lGLU
// #include <GL/gl.h>
import
  "C"
const (
  MaxL = 8
  Front = C.GL_FRONT
  FrontAndBack = C.GL_FRONT_AND_BACK
  Diffuse = C.GL_DIFFUSE
  AmbientAndDiffuse = C.GL_AMBIENT_AND_DIFFUSE
  LmTwoSide = C.GL_LIGHT_MODEL_TWO_SIDE
  LmAmbient = C.GL_LIGHT_MODEL_AMBIENT
)
var (
  mt, lm [4]F
  sh [MaxL]F
  am, df, sp, em [MaxL][4]F
  ps [MaxL][3]F
  switched [MaxL]bool
)

func init() {
  mt[0], mt[1], mt[2], mt[3] = F(0.2), F(0.2), F(0.2), F(1)
  lm[0], lm[1], lm[2], lm[3] = F(0.2), F(0.2), F(0.2), F(1)
  clearDepth (1)
}

// Pre: n < MaxL.
// Light n is defined with the default values for
// ambience, diffusion, specularity and emission.
func SetLight (n uint) {
  if n >= MaxL { return }
  Lma (0.2, 0.2, 0.2)
  Amb (n, 1.0, 1.0, 1.0)
  Dif (n, 1.0, 1.0, 1.0)
  Spe (n, 1.0, 1.0, 1.0)
  Emi (n, 0.0, 0.0, 0.0)
  Shi (n, 0.0)
}

//
func ColorMaterial (f, m int) {
//  C.glMaterialfv (e(face), e(pname), &mat[0])
}

func mat (a, b, c float64) {
  mt[0], mt[1], mt[2], mt[3] = F(a), F(b), F(c), F(1)
}

func material (face, pname int) {
  C.glMaterialfv (e(face), e(pname), &mt[0])
}

func Lma (a, b, c float64) {
  lm[0], lm[1], lm[2], lm[3] = F(a), F(b), F(c), F(1)
  C.glLightModelfv (C.GL_LIGHT_MODEL_AMBIENT, &lm[0])
}

// Pre: n < MaxL. SetLight (n) was called.
// The ambience of light n has the values a, b and c.
func Amb (n uint, a, b, c float64) {
  if n >= MaxL { return }
  am[n][0], am[n][1], am[n][2], am[n][3] = F(a), F(b), F(c), F(1)
  C.glLightfv (e(C.GL_LIGHT0 + n), C.GL_AMBIENT, &am[n][0])
}

// Pre: n < MaxL. SetLight (n) was called.
// Light n has the ambience values a, b and c.
func Dif (n uint, a, b, c float64) {
  if n >= MaxL { return }
  df[n][0], df[n][1], df[n][2], df[n][3] = F(a), F(b), F(c), F(1)
  C.glLightfv (e(C.GL_LIGHT0 + n), C.GL_DIFFUSE, &df[n][0])
}

// Pre: n < MaxL. SetLight (n) was called.
// Light n has the specular values a, b and c.
func Spe (n uint, a, b, c float64) {
  if n >= MaxL { return }
  sp[n][0], sp[n][1], sp[n][2], sp[n][3] = F(a), F(b), F(c), F(1)
  C.glLightfv (e(C.GL_LIGHT0 + n), C.GL_SPECULAR, &sp[n][0])
}

// Pre: n < MaxL. SetLight (n) was called.
// Light n has the emission values a, b and c.
func Emi (n uint, a, b, c float64) {
  if n >= MaxL { return }
  em[n][0], em[n][1], em[n][2], em[n][3] = F(a), F(b), F(c), F(1)
  C.glLightfv (e(C.GL_LIGHT0 + n), C.GL_EMISSION, &em[n][0])
}

// Pre: s <= 128.
func Shi (n uint, s float64) {
  if n >= MaxL { return }
  sh[n] = F(s)
  C.glLightfv (e(C.GL_LIGHT0 + n), C.GL_SHININESS, &sh[n])
}

// Pre: n < MaxL. SetLight (n) was called.
// Light n has position x, y, z.
func Pos (n uint, x, y, z float64) {
  if n >= MaxL { return }
  ps[n][0], ps[n][1], ps[n][2] = F(x), F(y), F(z)
  C.glLightfv (e(C.GL_LIGHT0 + n), C.GL_POSITION, &ps[n][0])
}

// Pre: n < MaxL.
// For on == true light n is switched on,
// otherwise it is switched off.
func SwitchLight (n uint, on bool) {
  if n >= MaxL || switched[n] { return }
  if on {
    C.glEnable (C.GL_LIGHT0 + e(n))
    C.glEnable (C.GL_LIGHTING)
  } else {
    C.glDisable (C.GL_LIGHT0 + e(n))
    C.glDisable (C.GL_LIGHTING)
  }
}

// Pre: n < MaxL.
func Lamp (n uint) {
/*/
  if ! initialized[n] { return }
  xx, yy, zz := pos[n][0], pos[n][1], pos[n][2]
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
/*/
}
