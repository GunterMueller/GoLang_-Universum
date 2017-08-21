package gl

// (c) murus.org  v. 170820 - license see murus.go

// #cgo LDFLAGS: -lGL
// #include <GL/gl.h> 
import
  "C"
import (
  "math"
  "murus/ker"
  "murus/spc"
  "murus/env"
  "murus/col"
  "murus/scr"
  "murus/vect"
)
const
  nLamp = 12
var (
  lmAmb, mAmbi, mDiff [4]C.GLfloat
  sin, cos [nLamp + 2]C.GLdouble
  lightSource [MaxL]vect.Vector
  lightColour /* diffus */ [MaxL]col.Colour
  lightInitialized [MaxL]bool
  initialized, lightVis bool
  nn uint
  aa, dd [MaxL][4]C.GLfloat
  yyy [3]C.GLdouble
  fff Class = POINTS
  matrix [4][4]C.GLdouble
  right, front, top, eye [3]float64
//  counter uint
  text = []string {"UNDEF", "POINTS", "LINES", "LINE_LOOP", "LINE_STRIP", "TRIANGLES",
                   "TRIANGLE_STRIP", "TRIANGLE_FAN", "QUADS", "QUAD_STRIP", "POLYGON", "LIGHT"}
)

func init() {
  right[0], front[1], top[2] = 1.0, 1.0, 1.0
  matrix[3][3] = 1.
  for l := 0; l < MaxL; l++ { lightSource[l] = vect.New() }
  w := 2.0 * math.Pi / float64 (nLamp)
  sin[0], cos[0] = C.GLdouble(0.0), C.GLdouble(1.0)
  for g := 1; g < nLamp; g++ {
    sin[g] = C.GLdouble(math.Sin (float64 (g) * w))
    cos[g] = C.GLdouble(math.Cos (float64 (g) * w))
  }
  sin[nLamp], cos[nLamp] = C.GLdouble(0.0), C.GLdouble(1.0)
  sin[nLamp+1], cos[nLamp+1] = sin[1], cos[1]
  C.glDepthFunc (C.GL_LESS) // default
  C.glEnable (C.GL_DEPTH_TEST)
  C.glShadeModel (C.GL_SMOOTH)
  for i := 0; i < 3; i++ { lmAmb[i] = C.GLfloat(0.2) } // default: 0.2
  lmAmb[3] = C.GLfloat(1.0) // default: 1.0
  C.glLightModelfv (C.GL_LIGHT_MODEL_AMBIENT, &lmAmb[0])
//  C.glLightModelfv (C.GL_LIGHT_MODEL_TWO_SIDE, 1)
  for i := 0; i < 3; i++ { mAmbi[i] = C.GLfloat(0.2) } // default: 0.2
  mAmbi[3] = C.GLfloat(1.0) // default: 1.0
  for i := 0; i < 3; i++ { mDiff[i] = C.GLfloat(0.8) } // default: 0.8
  mDiff[3] = C.GLfloat(1.0) // default: 1.0
  C.glMaterialfv (C.GL_FRONT_AND_BACK, C.GL_AMBIENT_AND_DIFFUSE, &mAmbi[0])
  w = 1.
  C.glClearDepth (C.GLclampd(w))
//  C.glMaterialfv (C.GL_FRONT_AND_BACK, C.GL_DIFFUSE, mDiff)
//  C.glColorMaterial (C.GL_FRONT_AND_BACK, C.GL_DIFFUSE)
//  C.glColorMaterial (C.GL_FRONT, C.GL_AMBIENT)
  C.glColorMaterial (C.GL_FRONT_AND_BACK, C.GL_AMBIENT_AND_DIFFUSE)
  C.glEnable (C.GL_COLOR_MATERIAL)
  C.glEnable (C.GL_LIGHTING)
}

func cls (c col.Colour) {
  r, g, b := col.Float (c) // 0.0, 0.0, 0.0
  C.glClearColor (C.GLclampf(r), C.GLclampf(g), C.GLclampf(b), C.GLclampf(0.0))
}

func init0() {
  if initialized { return }
  initialized = true
  C.glViewport (0, 0, C.GLsizei(scr.Wd()), C.GLsizei(scr.Ht()))
}

func init_(far float64) {
  init0()
  C.glMatrixMode (C.GL_PROJECTION)
  C.glLoadIdentity()
  const D = 2.0 // -fold screen width
  deg := D * math.Atan ((0.5 / D) / scr.Proportion())
  deg /= 0.9 // experimental wideangle correction
  var m [4][4]C.GLdouble
  m[1][1] = 1.0 / C.GLdouble(math.Tan (deg)) // Cot
  m[0][0] = m[1][1] / C.GLdouble (scr.Proportion())
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

func initLight (n uint, v, h vect.Vector, c col.Colour) {
  if lightInitialized[n] { return }
  lightVis = true
  var a [4]float64
  a[0], a[1], a[2] = h.Coord3()
  // Arbeitsdrumrum, weil die Punkte bisher nur eine Farbe transportieren, hier die diffuse.
  // In L wird die ambiente Farbe geliefert.
  for i := 0; i < 3; i++ { aa[n][i] = C.GLfloat(a[i]) }; aa[n][3] = C.GLfloat(1.0)
  lightColour[n] = c
  d0, d1, d2 := col.Float (c)
  dd[n][0], dd[n][1], dd[n][2] = C.GLfloat(d0), C.GLfloat(d1), C.GLfloat(d2)
  dd[n][3] = C.GLfloat(1.0)
  lightSource[n].Copy (v)
  ActualizeLight (n)
  lightInitialized[n] = true
}

func posLight (n uint, v vect.Vector) {
  if lightInitialized[n] { lightSource[n].Copy (v) }
}

func actLight (n uint) { // n < MaxL
  var L [4]float64
  L[0], L[1], L[2] = lightSource[n].Coord3()
  var l [4]C.GLfloat
  for i := 0; i < 3; i++ { l[i] = C.GLfloat(L[i]) }; l[3] = C.GLfloat(1.0)
  C.glLightfv (C.GL_LIGHT0 + C.GLenum(n), C.GL_POSITION, &l[0])
  C.glLightfv (C.GL_LIGHT0 + C.GLenum(n), C.GL_AMBIENT, &aa[n][0])
  C.glLightfv (C.GL_LIGHT0 + C.GLenum(n), C.GL_DIFFUSE, &dd[n][0])
  C.glEnable (C.GL_LIGHT0 + C.GLenum(n))
}

func write0() {
  init0()
  if ! scr.UnderX() { ker.Panic (env.Par(0) + " must not be called in a console") }
  C.glEnable (C.GL_DEPTH_TEST)
  C.glMatrixMode (C.GL_MODELVIEW)
  C.glLoadIdentity()
  for i := 0; i < 3; i++ {
    matrix[i][0] = C.GLdouble(right[i])
    matrix[i][1] = C.GLdouble(top[i])
    matrix[i][2] = C.GLdouble(-front[i])
  }
//  ker.Mess ("glw00")
  C.glMultMatrixd (&matrix[0][0])
  C.glTranslated (C.GLdouble(-eye[0]), C.GLdouble(-eye[1]), C.GLdouble(-eye[2]))
  C.glClear (C.GL_COLOR_BUFFER_BIT + C.GL_DEPTH_BUFFER_BIT)
//  ker.Mess ("glw01")
  for n := uint(0); n < MaxL; n++ {
    if lightInitialized[n] { ActualizeLight (n) }
  }
//  ker.Mess ("glw02")
  C.glBegin (POINTS)
// println ("glBegin", "POINTS")
  nn = 0
//  counter = 1
}

func vector2yyy (v vect.Vector) {
  for i := 0; i < 3; i++ { yyy[i] = C.GLdouble(v.Coord (spc.Direction(i))) }
}

func write (class Class, a uint, V, N []vect.Vector, c col.Colour) {
//  if counter % 10 == 0 { ker.Mess ("glw") }
//  println ("glw", counter); counter++
  switch class {
  case UNDEF:
    nn = 0 // forces glEnd / glBegin
    return
  case LIGHT:
    if a >= MaxL { ker.Panic ("gl.Write: 2nd parameter >= MaxL") }
    InitLight (a, V[0], N[0], c)
    nn = 0
    return
  }
//  case POINTS, LINES, LINE_LOOP, LINE_STRIP, TRIANGLES,
//       TRIANGLE_STRIP, TRIANGLE_FAN, QUADS, QUAD_STRIP, POLYGON:
  if class != fff || a != nn || nn == 0 {
//            ^^^ at start: POINTS
    fff = class
    nn = a
    C.glEnd()
    C.glBegin (C.GLenum(class))
println (a, "glBegin", text[class])
  }
  C.glColor3ub (C.GLubyte(c.R), C.GLubyte(c.G), C.GLubyte(c.B))
println ("class", text[class])
  for i := uint(0); i < a; i++ {
    vector2yyy (V[i])
println (i, V[i].String())
    C.glVertex3dv (&yyy[0])
//    vector2yyy (N[i]); C.glNormal3dv (&yyy[0])
  }
/*
  tmp := vect.New()
  C.glEnd()
  for i := uint(0); i < a; i++ {
    C.glBegin (LINES)
    C.glColor3ub (C.GLubyte(0), C.GLubyte(0), C.GLubyte(0))
    vector2yyy (V[i]); C.glVertex3dv (&yyy[0])
    tmp.Add (V[i], N[i])
    vector2yyy (tmp); C.glVertex3dv (&yyy[0])
    C.glEnd()
  }
  nn = 0
  C.glBegin (POINTS)
*/

}

func lamp (n uint) {
  if ! lightInitialized[n] { return }
  xx, yy, zz := lightSource[n].Coord3()
  x, y, z := C.GLdouble(xx), C.GLdouble(yy), C.GLdouble(zz)
  r := C.GLdouble(0.1)
  C.glBegin (TRIANGLE_FAN)
  C.glColor3ub (C.GLubyte(lightColour[n].R), C.GLubyte(lightColour[n].G), C.GLubyte(lightColour[n].B))
  C.glNormal3d (C.GLdouble(0.0), C.GLdouble(0.0), C.GLdouble(-1.0))
  C.glVertex3d (C.GLdouble(x), C.GLdouble(y), C.GLdouble(z + r))
  r0, z0 := r * sin[1], z + r * cos[1]
  for l := 0; l <= nLamp; l++ {
    C.glNormal3d (-sin[1] * cos[l], -sin[1] * sin[l], -cos[1])
    C.glVertex3d (x + r0 * cos[l],   y + r0 * sin[l],   z0)
  }
  C.glEnd()
  C.glBegin (QUAD_STRIP)
  var r1, z1 C.GLdouble
  for b := 1; b <= nLamp / 2 - 2; b++ {
    r0, z0 = r * sin[b],   z + r * cos[b]
    r1, z1 = r * sin[b+1], z + r * cos[b+1]
    for l := 0; l <= nLamp; l++ {
      C.glNormal3d (-sin[b+1] * cos[l], -sin[b+1] * sin[l], -cos[b+1])
      C.glVertex3d (x + r1 * cos[l], y + r1 * sin[l], z1)
      C.glNormal3d (-sin[b] * cos[l], -sin[b] * sin[l], -cos[b])
      C.glVertex3d (x + r0 * cos[l], y + r0 * sin[l], z0)
    }
  }
  C.glEnd()
  C.glBegin (TRIANGLE_FAN)
  C.glNormal3d (0., 0., 1.)
  C.glVertex3d (x, y, z - r)
  r0, z0 = r * sin[1], z - r * cos[1]
  b := nLamp / 2 - 1
  for l := 0; l <= nLamp; l++ {
    C.glNormal3d (-sin[b] * cos[l], -sin[b] * sin[l], -cos[b])
    C.glVertex3d (x + r0 * cos[l], y + r0 * sin[l], z0)
  }
  C.glEnd()
}

func write1() {
  C.glEnd()
println("glEnd")
  if lightVis { for n := uint(0); n < MaxL; n++ { lamp (n) } }
  scr.WriteGlx()
}

func act (r, v, o, a vect.Vector) {
  right[0], right[1], right[2] = r.Coord3()
  front[0], front[1], front[2] = v.Coord3()
  top[0], top[1], top[2] = o.Coord3()
  eye[0], eye[1], eye[2] = a.Coord3()
}

/*
func hold() {
  C.glPushMatrix()
}

func continue() {
  C.glPopMatrix()
}
*/
