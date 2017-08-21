package fig

// (c) murus.org  v. 170821 - license see murus.go

// >>> still a lot of work TODO

import (
  "math"
  "murus/ker"
  "murus/rand"
  "murus/col"
  "murus/scr"
  "murus/vect"
  "murus/pt"
  "murus/pts"
)

const (
  angle = 3 // Grad
  N = 360 / angle
  pi_180 = math.Pi / 180
)
var
  sin, cos []float64

func init() {
  sin, cos = make ([]float64, N + 2), make ([]float64, N + 2)
  sin[0], cos[0] = 0, 1
  w := 2 * math.Pi / float64 (N)
  for i := 1; i < N; i++ {
    sin[i] = math.Sin (float64(i) * w)
    cos[i] = math.Cos (float64(i) * w)
  }
  sin[N], cos[N] = 0, 1
  sin[N+1], cos[N+1] = sin[1], cos[1]
}

func sort (x, y, z, x1, y1, z1 *float64) {
  if *x1 < *x { *x, *x1 = *x1, *x }
  if *y1 < *y { *y, *y1 = *y1, *y }
  if *z1 < *z { *z, *z1 = *z1, *z }
}

func v_(x, y, z float64) vect.Vector {
  return vect.New3 (x, y, z)
}

func init_(v, n []vect.Vector, x ...float64) {
  k := len(x)
  for i := 0; i < k / 3; i++ {
    v[i].Set3 (x[3 * i], x[3 * i + 1], x[3 * i + 2])
  }
  n[0].Set3 (0, 0, 1)
  if k > 3 {
    n[1].Set3 (0, 0, 1)
  }
}

func init_v (v, n []vect.Vector, x ...vect.Vector) {
  k := len(v)
  if k != len(x) { ker.Oops() }
  for i := 0; i < k; i++ {
    v[i].Copy (x[i])
  }
  n[0].Set3 (0, 0, 1)
  if k > 3 {
    n[1].Set3 (0, 0, 1)
  }
}

func vectors (a uint) ([]vect.Vector, []vect.Vector) {
  v, n := make ([]vect.Vector, a), make ([]vect.Vector, a)
  for i := uint(0); i < a; i++ { v[i], n[i] = vect.New(), vect.New() }
  return v, n
}

func ready (p pts.Points) {
  v, n := vectors(1)
  p.Ins (pt.Undef, v, n, col.Black)
}

func start (p pts.Points, x, y, z, xf, yf, zf float64) {
  v, n := vectors(1)
  init_(v, n, x, y, z) // Auge
  n[0].Set3 (xf, yf, zf) // Fokus
  p.Ins (pt.Start, v, n, col.Black)
}

// GL classes //////////////////////////////////////////////////////////////////

func figure (class pt.Class, p pts.Points, c col.Colour, x ...vect.Vector) {
  v, n := vectors(uint(len(x)))
  switch class {
  case pt.Points:
    init_v (v, n, x[0])
    p.Ins (pt.Points, v, n, c)
  case pt.Lines:
    init_v (v, n, x[0], x[1])
    p.Ins (pt.Lines, v, n, c)
//  case pt.LineLoop:
//  case pt.LineStrip:
  case pt.Triangles:
    init_v (v, n, x[0], x[1], x[2])
    n[1].Diff (v[1], v[0])
    n[2].Diff (v[2], v[0])
    n[0].Ext (n[1], n[2])
    n[0].Norm()
    for i := 1; i < 3; i++ { n[i].Copy (n[0]) }
    p.Ins (pt.Triangles, v, n, c)
//  case pt.TriangleStrip:
//  case pt.TriangleFan:
  case pt.Quads:
    init_v (v, n, x[0], x[1], x[2], x[3])
    n[1].Diff (v[1], v[0])
    n[2].Diff (v[3], v[0])
    n[0].Ext (n[1], n[2])
    n[0].Norm()
    for i := 1; i < 4; i++ { n[i].Copy (n[0]) }
    p.Ins (pt.Quads, v, n, c)
//  case pt.QuadStrip:
//  case pt.Polygon:
  default:
    ker.Panic ("class not yet implemented")
  }
  ready (p)
}

// light ///////////////////////////////////////////////////////////////////////

func light (p pts.Points, l uint, x, y, z float64, ca, cd col.Colour) {
  v, n := vectors(1)
  v[0].Set3 (x, y, z)
  n[0].Set3 (0, 0, 1)
//  r, g, b := col.LongFloat (ca)
  n[0].Set3 (col.LongFloat (ca)) // r, g, b // ambient colour
  p.InsLight (l, v, n, cd)
}

// extended figures ////////////////////////////////////////////////////////////

func horRectangle (p pts.Points, x, y, z, x1, y1 float64, up bool, c col.Colour) {
  v, n := vectors(4)
  init_(v, n, x, y, z, x1, y, z, x1, y1, z, x, y1, z)
  goUp := 1.; if ! up { goUp = -1 }
  n[0].Set3 (0, 0, goUp)
  for i := 1; i < 4; i++ {
    n[i].Set3 (0, 0, 0)
  }
  p.Ins (pt.Quads, v, n, c)
  ready (p)
}

// func vertRectangle (p pts.Points, v, v1 vect.Vector, c col.Colour) { // TODO
func vertRectangle (p pts.Points, x, y, z, x1, y1, z1 float64, c col.Colour) {
  v, n := vectors(4)
  if z == z1 { ker.Oops() } // horRectangle
  init_(v, n, x, y, z, x1, y1, z, x1, y1, z1, x, y, z1)
  n[1].Diff (v[1], v[0])
  n[2].Diff (v[3], v[0])
  n[0].Ext (n[1], n[2])
  for i := 1; i < 4; i++ { n[i].Copy (n[0]) }
  p.Ins (pt.Quads, v, n, c)
  ready (p)
}

func parallelogram (p pts.Points, c col.Colour, v, v1, v2 vect.Vector) {
  x, y, z := v.Coord3()
  x1, y1, z1 := v1.Coord3()
  x2, y2, z2 := v1.Coord3()
  figure (pt.Quads, p, c, v_(x, y, z), v_(x1, y1, z1),
                          v_(x1 + x2 - x, y1 + y2 - y, z1 + z2 - z), v_(x2, y2, z2))
}

func cube (p pts.Points, v vect.Vector, a float64, c col.Colour) {
  v1 := v.Clone().(vect.Vector)
  v1.Translate (a)
  cuboid (p, v, v1, c)
}

func cubeC (p pts.Points, v vect.Vector, a float64, c []col.Colour) {
  v1 := vect.New3 (v.Coord(0) + a, v.Coord(1) + a, v.Coord(2) + a)
  cuboidC (p, v, v1, c)
}

func cuboid (p pts.Points, v, v1 vect.Vector, c col.Colour) {
  x, y, z := v.Coord3()
  x1, y1, z1 := v1.Coord3()
  sort (&x, &y, &z, &x1, &y1, &z1)
/*
  vertRectangle (p, x,  y,  z,  x1, y,  z1, c)
  vertRectangle (p, x1, y,  z,  x1, y1, z1, c)
  vertRectangle (p, x1, y1, z,  x,  y1, z1, c)
  vertRectangle (p, x,  y1, z,  x,  y,  z1, c)
*/
  vv, nn := vectors(2 * (4 + 1))
  vv[0].Set3 (x,  y,  z )
  vv[1].Set3 (x,  y,  z1)
  vv[2].Set3 (x1, y,  z )
  vv[3].Set3 (x1, y,  z1)
  vv[4].Set3 (x1, y1, z )
  vv[5].Set3 (x1, y1, z1)
  vv[6].Set3 (x,  y1, z )
  vv[7].Set3 (x,  y1, z1)
  vv[8].Set3 (x,  y,  z )
  vv[9].Set3 (x,  y,  z1)
  nn[0].Set3 (-1, -1, 0)
  nn[1].Set3 (-1, -1, 0)
  nn[2].Set3 ( 1, -1, 0)
  nn[3].Set3 ( 1, -1, 0)
  nn[4].Set3 ( 1,  1, 0)
  nn[5].Set3 ( 1,  1, 0)
  nn[6].Set3 (-1,  1, 0)
  nn[7].Set3 (-1,  1, 0)
  nn[8].Set3 (-1, -1, 0)
  nn[9].Set3 (-1, -1, 0)
  p.Ins (pt.QuadStrip, vv, nn, c)

  horRectangle (p, x,  y,  z1, x1, y1, true, c)
  horRectangle (p, x,  y,  z,  x1, y1, false, c)
}

func cuboidC (p pts.Points, v, v1 vect.Vector, c []col.Colour) {
  x, y, z := v.Coord3()
  x1, y1, z1 := v1.Coord3()
  sort (&x, &y, &z, &x1, &y1, &z1)
  vertRectangle (p, x,  y,  z,  x1, y,  z1, c[0]) // front
  vertRectangle (p, x1, y,  z,  x1, y1, z1, c[1]) // right
  vertRectangle (p, x1, y1, z,  x,  y1, z1, c[2]) // back
  vertRectangle (p, x,  y1, z,  x,  y,  z1, c[3]) // left
/*
  vv, nn := vectors(2 * (4 + 1))

  vv[0].Set3 (x,  y,  z )
  vv[1].Set3 (x,  y,  z1)
  vv[2].Set3 (x1, y,  z )
  vv[3].Set3 (x1, y,  z1)
  vv[4].Set3 (x1, y1, z )
  vv[5].Set3 (x1, y1, z1)
  vv[6].Set3 (x,  y1, z )
  vv[7].Set3 (x,  y1, z1)
  vv[8].Set3 (x,  y,  z )
  vv[9].Set3 (x,  y,  z1)

  nn[0].Set3 (-1, -1, 0)
  nn[1].Set3 (-1, -1, 0)
  nn[2].Set3 ( 1, -1, 0)
  nn[3].Set3 ( 1, -1, 0)
  nn[4].Set3 ( 1,  1, 0)
  nn[5].Set3 ( 1,  1, 0)
  nn[6].Set3 (-1,  1, 0)
  nn[7].Set3 (-1,  1, 0)
  nn[8].Set3 (-1, -1, 0)
  nn[9].Set3 (-1, -1, 0)
  p.Ins (pt.QuadStrip, vv, nn, c) // XXX
*/
  horRectangle (p, x,  y,  z1, x1, y1, true, c[4])  // top
  horRectangle (p, x,  y,  z,  x1, y1, false, c[5]) // bottom
}

func cuboid1 (p pts.Points, f col.Colour, x, y, z, b, t, h, a float64) {
  s, c := math.Sin (a * pi_180), math.Cos (a * pi_180)
  x1 := x  + b * c; y1 := y  + b * s; z1 := z + h
  x2 := x1 - t * s; y2 := y1 + t * c
  x3 := x  - t * s; y3 := y  + t * c
  figure (pt.Quads, p, f, v_(x, y, z1), v_(x1, y1, z1), v_(x2, y2, z1), v_(x3, y3, z1))
  vertRectangle (p, x,  y,  z,  x1, y1, z1, f)
  vertRectangle (p, x1, y1, z,  x2, y2, z1, f)
  vertRectangle (p, x2, y2, z,  x3, y3, z1, f)
  vertRectangle (p, x3, y3, z,  x,  y,  z1, f)
  figure (pt.Quads, p, f, v_(x, y, z), v_(x3, y3, z), v_(x2, y2, z), v_(x1, y1, z))
}

func prism (p pts.Points, c col.Colour, x, y, z []float64) {
// top missing
  n := uint(len (x))
  if n < 4 { ker.Oops() }
  n-- // top !
  for i := uint(0); i < n - 1; i++ {
    figure (pt.Quads, p, c,
            v_(x[i],          y[i],          z[i]),
            v_(x[i+1],        y[i+1],        z[i+1]),
            v_(x[i+1] + x[n], y[i+1] + y[n], z[i+1] + z[n]),
            v_(x[i]   + x[n], y[i]   + y[n], z[i]   + z[n]))
  }
  i := uint(n - 1)
    figure (pt.Quads, p, c,
            v_(x[i],        y[i],        z[i]),
            v_(x[0],        y[0],        z[0]),
            v_(x[0] + x[n], y[0] + y[n], z[0] + z[n]),
            v_(x[i] + x[n], y[i] + y[n], z[i] + z[n]))
// bottom missing
  ready (p)
}

func parallelepiped (p pts.Points, c col.Colour, v, v1, v2, v3 vect.Vector) {
  x, y, z := v.Coord3()
  x1, y1, z1 := v1.Coord3()
  x2, y2, z2 := v2.Coord3()
  x3, y3, z3 := v3.Coord3()
  parallelogram (p, c, v, v1, v3)
  parallelogram (p, c, v, v2, v1)
  parallelogram (p, c, v, v3, v2)
  parallelogram (p, c, v1, v_(x1 + x2 - x, y1 + y2 - y, z1 + z2 - z),
                           v_(x1 + x3 - x, y1 + y3 - y, z1 + z3 - z))
  parallelogram (p, c, v2, v_(x2 + x3 - x, y2 + y3 - y, z2 + z3 - z),
                           v_(x2 + x1 - x, y2 + y1 - y, z2 + z1 - z))
  parallelogram (p, c, v3, v_(x3 + x1 - x, y3 + y1 - y, z3 + z1 - z),
                           v_(x3 + x2 - x, y3 + y2 - y, z3 + z2 - z))
  ready (p)
}

func pyramid (p pts.Points, c col.Colour, v, v1, v2 vect.Vector) {
  x, y, z := v.Coord3()
  x1, y1, z1 := v1.Coord3()
  figure (pt.Triangles, p, c, v,             v_(x1, y,  z1), v2)
  figure (pt.Triangles, p, c, v_(x1, y,  z), v_(x1, y1, z),  v2)
  figure (pt.Triangles, p, c, v_(x1, y1, z), v_(x,  y1, z1), v2)
  figure (pt.Triangles, p, c, v_(x,  y1, z), v_(x,  y,  z),  v2)
  horRectangle (p, x, y, z, x1, y1, false, c)
  ready (p)
}

func octahedron (p pts.Points, c col.Colour, x, y, z, r float64) {
  d := r * math.Sqrt (2)
  figure (pt.Triangles, p, c, v_(x + r, y + r, z), v_(x - r, y + r, z), v_(x, y, z + d))
  figure (pt.Triangles, p, c, v_(x - r, y + r, z), v_(x - r, y - r, z), v_(x, y, z + d))
  figure (pt.Triangles, p, c, v_(x - r, y - r, z), v_(x + r, y - r, z), v_(x, y, z + d))
  figure (pt.Triangles, p, c, v_(x + r, y - r, z), v_(x + r, y + r, z), v_(x, y, z + d))
  figure (pt.Triangles, p, c, v_(x + r, y + r, z), v_(x - r, y + r, z), v_(x, y, z - d))
  figure (pt.Triangles, p, c, v_(x - r, y + r, z), v_(x - r, y - r, z), v_(x, y, z - d))
  figure (pt.Triangles, p, c, v_(x - r, y - r, z), v_(x + r, y - r, z), v_(x, y, z - d))
  figure (pt.Triangles, p, c, v_(x + r, y - r, z), v_(x + r, y + r, z), v_(x, y, z - d))
}

func multipyramid (p pts.Points, c col.Colour, x, y, z []float64) {
  n := len (x)
  if n < 4 { ker.Oops() }
  n-- // top !
  v0 := vect.New3(x[0], y[0], z[0])
  vn := vect.New3(x[n], y[n], z[n])
  for i := 0; i < n - 1; i++ {
    figure (pt.Triangles, p, c, v_(x[i], y[i], z[i]), v_(x[i+1], y[i+1], z[i+1]), vn)
  }
  figure (pt.Triangles, p, c, v_(x[n-1], y[n-1], z[n-1]), v0, vn)
// bottom missing, because it need not be even
}

func circle (p pts.Points, x, y, z, r float64, c col.Colour) {
  circleSegment (p, x, y, z, r, 0, 360, c)
}

func circleSegment (p pts.Points, x, y, z, r, a, b float64, c col.Colour) {
  if r == 0 {
    figure (pt.Points, p, c, v_(x, y, z))
    return
  }
  A := uint(math.Floor (a / float64 (angle) + 0.5))
  B := uint(math.Floor (b / float64 (angle) + 0.5))
  C := B - A + 2
  v, n := vectors(C)
  v[0].Set3 (x, y, z)
  n[0].Set3 (0, 0, 1)
  if r < 0. {
    r = -r
    n[0].Dilate (-1)
  }
  for i := A; i <= B; i++ {
    v[1 + i-A].Set3 (x + r * cos[i], y + r * sin[i], z)
    n[1 + i-A].Copy (n[0])
  }
  p.Ins (pt.TriangleFan, v, n, c)
  ready (p)
}

func vertCircle (p pts.Points, x, y, z, r, a float64, f col.Colour) {
  if r == 0 {
    figure (pt.Points, p, f, v_(x, y, z))
    return
  }
  s, c := math.Sin (a * pi_180), math.Cos (a * pi_180)
  C := uint(N) + 2
  v, n := vectors(C)
  v[0].Set3 (x, y, z)
  n[0].Set3 (c, s, 0)
  if r < 0 {
    r = -r
    n[0].Dilate (-1)
  }
  for i := 0; i <= N; i++ {
    v[i+1].Set3 (x - r * s * cos[i], y + r * c * cos[i], z + r * sin[i])
    n[i+1].Copy (n[0])
  }
  p.Ins (pt.TriangleFan, v, n, f)
  ready (p)
}

func sphere (p pts.Points, x, y, z, r float64, f col.Colour) {
  v, n := vectors(N + 2)
  v[0].Set3 (x, y, z + r)
  n[0].Set3 (0, 0, 1)
  r0, z0 := r * sin[1], z + r * cos[1]
  for l := 0; l <= N; l++ {
    v[1 + l].Set3 (x + r0 * cos[l], y + r0 * sin[l], z0)
    n[1 + l].Set3 (sin[1] * cos[l], sin[1] * sin[l], cos[1])
  }
  p.Ins (pt.TriangleFan, v, n, f)

  v, n = vectors(2 * (N + 1))
  for b := 1; b <= N / 2 - 2; b++ {
    r0, z0 =     r * sin[b], z + r * cos[b]
    r1, z1 :=     r * sin[b+1], z + r * cos[b+1]
    for l := 0; l <= N; l++ {
      s, c := sin[l], cos[l]
      v[2*l].Set3 (x + r0 * c, y + r0 * s, z0)
      n[2*l].Set3 (sin[b] * c, sin[b] * s, cos[b])
      v[1 + 2*l].Set3 (x + r1 * c, y + r1 * s, z1)
      n[1 + 2*l].Set3 (sin[b+1] * c, sin[b+1] * s, cos[b+1])
    }
    p.Ins (pt.QuadStrip, v, n, f)
  }
  v, n = vectors(N + 2)
  v[0].Set3 (x, y, z - r)
  n[0].Set3 (0, 0, -1)
  b := N / 2 - 1
  r0, z0 = r * sin[b], z + r * cos[b]
  for l := N; l >= 0; l -= 1 {
    v[1 + N-l].Set3 (x + r0 * cos[l], y + r0 * sin[l], z0)
    n[1 + N-l].Set3 (sin[b] * cos[l], sin[b] * sin[l], cos[b])
  }
  p.Ins (pt.TriangleFan, v, n, f)
  ready (p)
}

func cone (p pts.Points, x, y, z, r, h float64, c col.Colour) {
  v, n := vectors(N + 2)
  v[0].Set3 (x, y, z + h)
  n[0].Set3 (0, 0, 1)
  for l := 0; l <= N; l++ {
    v[l+1].Set3 (x + r * cos[l], y + r * sin[l], z)
    n[l+1].Set3 (cos[l], sin[l], r / (h - z))
    n[l+1].Norm()
  }
  p.Ins (pt.TriangleFan, v, n, c)
  ready (p)
  circle (p, x, y, z, -r, c)
}

func frustum (p pts.Points, x, y, z, r, h, h1 float64, c col.Colour) {
  if h1 > h { ker.Oops() }
  v, n := vectors(N + 2)
  v[0].Set3 (x, y, h)
  n[0].Set3 (0, 0, 1)
  for l := 0; l <= N; l++ {
    v[l+1].Set3 (x + r * cos[l], y + r * sin[l], z)
    n[l+1].Set3 (cos[l], sin[l], r / (h - z))
    n[l+1].Norm()
  }
  p.Ins (pt.TriangleFan, v, n, c)
  ready (p)
  circle (p, x, y, z, -r, c)
}

func doubleCone (p pts.Points, x, y, z, r, h float64, c col.Colour) {
  cone (p, x, y, z - h, r, h, c)
  cone (p, x, y, z + h, r, -h, c)
}

func cylinder (p pts.Points, x, y, z, r, h float64, c col.Colour) {
  cylinderSegment (p, x, y, z, r, h, 0, 360, c)
}

func cylinderSegment (p pts.Points, x, y, z, r, h, a, b float64, c col.Colour) {
  circleSegment (p, x, y, z, -r, a, b, c)
  circleSegment (p, x, y, z + h, r, a, b, c)
  A := uint(math.Floor (a / float64 (angle) + 0.5))
  B := uint(math.Floor (b / float64 (angle) + 0.5))
  C := 2 * (B - A) + 2
  v, n := vectors(C)
  for l := A; l <= B; l++ {
    v[2*(l-A)].Set3 (x + r * cos[l], y + r * sin[l], z)
    n[2*(l-A)].Set3 (cos[l], sin[l], 0)
    v[2*(l-A)+1].Set3 (x + r * cos[l], y + r * sin[l], z + h)
    n[2*(l-A)+1].Set3 (cos[l], sin[l], 0)
  }
  p.Ins (pt.QuadStrip, v, n, c)
  ready (p)
}

func horCylinder (p pts.Points, x, y, z, r, l, a float64, f col.Colour) {
  if r == 0 {
    vertCircle (p, x, y, z, r, a, f)
    return
  }
  s, c := math.Sin (a * pi_180), math.Cos (a * pi_180)
  dx, dy := l * c, l * s
  vertCircle (p, x, y, z, -r, a, f)
  vertCircle (p, x + dx, y + dy, z, r, a, f)
  C := 2 * (uint(N) + 1)
  v, n := vectors(C)
  for i := 0; i <= 2 * N; i += 2 {
    si, ci := sin[i / 2], cos[i / 2]
    sci, cci := s * ci, c * ci
    x0, y0, z0 := x - r * sci, y + r * cci, z + r * si
    v[i].Set3 (x0, y0, z0)
    n[i].Set3 (- sci, cci, si)
    v[i+1].Set3 (x0 + dx, y0 + dy, z0)
    n[i+1].Copy (n[i])
  }
  p.Ins (pt.QuadStrip, v, n, f)
  ready (p)
}

func torus (p pts.Points, x, y, z, R, r float64, c col.Colour) {
  if r <= 0 || R <= 0 { ker.Oops() }
  for b := 0; b < N; b++ {
    s0, s1 := R + r * cos[b], R + r * cos[b+1]
    z0, z1 := z + r * sin[b], z + r * sin[b+1]
//    v, n := vectors(2 * N)
    for l := 0; l < N; l++ {
      figure (pt.Quads, p, c,
              v_(x + s0 * cos[l],   y + s0 * sin[l],   z0),
              v_(x + s0 * cos[l+1], y + s0 * sin[l+1], z0),
              v_(x + s1 * cos[l+1], y + s1 * sin[l+1], z1),
              v_(x + s1 * cos[l],   y + s1 * sin[l],   z1))
//      v[2*l].Set3 (x + s0 * cos[l], y + s0 * sin[l], z0)
//      n[2*l].Set3 (1., 1., 1.)
//      v[2*l+1].Set3 (x + s0 * cos[l+1], y + s0 * sin[l+1], z0)
//      n[2*l+1].Set3 (1, 1, 1)
    }
//    p.Ins (pt.QuadStrip, 2 * N, v, n, c)
  }
  ready (p)
}

func horTorus (p pts.Points, x, y, z, R, r, a float64, f col.Colour) {
  if r <= 0 || R <= 0 { ker.Oops() }
  for a <= -180 { a += 180 }
  for a >=  180 { a -= 180 }
  s, c := math.Sin (a * pi_180), math.Cos (a * pi_180)
  for b := 0; b < N; b++ {
    s0, s1 := R + r * cos[b], R + r * cos[b+1]
    x0, x1 := r * sin[b], r * sin[b+1]
    for l := 0; l < N; l++ { //  x -> x * c - y * s, y -> x * s + y * c
      y00, y01 := s0 * cos[l], s0 * cos[l+1]
      y10, y11 := s1 * cos[l], s1 * cos[l+1]
      figure (pt.Quads, p, f,
              v_(x + x0 * c - y00 * s, y + x0 * s + y00 * c, z + s0 * sin[l]),
              v_(x + x0 * c - y01 * s, y + x0 * s + y01 * c, z + s0 * sin[l+1]),
              v_(x + x1 * c - y11 * s, y + x1 * s + y11 * c, z + s1 * sin[l+1]),
              v_(x + x1 * c - y10 * s, y + x1 * s + y10 * c, z + s1 * sin[l]))
    }
  }
  ready (p)
}

// func paraboloid (p pts.Points, x, y, z, p float64, c col.Colour)

// func horParaboloid (p pts.Points, x, y, z, p, a float64, c col.Colour)

func ok (x float64) bool {
  return ! math.IsNaN (x)
}

const grain = 8 // reasonable compromise between fine grained
                // versus lots of data w.r.t. output efficiency

func curve (p pts.Points, f1, f2, f3 RealFunc, t0, t1 float64, c col.Colour) {
  mX := float64 (scr.Wd() / grain)
  dt := (t1 - t0) / mX
  for a := t0; a <= t1; a += dt {
    x, y, z := f1 (a), f2 (a), f3 (a)
    a1 := a + dt
    x1, y1, z1 := f1 (a1), f2 (a1), f3 (a1)
    if ok (x) && ok (y) && ok (z) && ok (x1) && ok (y1) && ok (z1) {
      figure (pt.Lines, p, c, v_(x, y, z), v_(x1, y1, z1))
    }
  }
}

func surface (p pts.Points, f RealFunc2, X, Y, Z, X1, Y1, Z1 float64, c col.Colour) {
// XXX
  if X == X1 || Y == Y1 || Z == Z1 { return }
  if X1 < X { X, X1 = X1, X }
  if Y1 < Y { Y, Y1 = Y1, Y }
  if Z1 < Z { Z, Z1 = Z1, Z }
  dx, dy := (X1 - X) / float64 (scr.Wd() / grain), (Y1 - Y) / float64 (scr.Ht() / grain)
  for x := X; x <= X1; x += dx {
//    y := Y
//    n := uint(0)
//    for y <= Y1 {
//      n ++
//      y += dy
//    }
// die Anwendung der OpenGL-Ausgabe in gl von TriangleFan ist noch fehlerhaft
    x1, x0 := x + dx, x + dx / 2
//    temp, temp1 := vect.New(), vect.New()
//    v, n := vectors(2 * n) // (2 * n + 1)                 ? ? ? ? ? ? ? ? ? ? ? ? ?
    for y := Y; y <= Y1; y += dy {
//    for i := uint(0); i < n; i++ { // oder i <= n            ? ? ? ? ? ? ? ? ? ? ? ? ?
//      v[2 * i].Set3 (x, y,  z)
//      v[2 * i + 1].Set3 (x, y1, z1)
//      if i == 0 { // ?
//        n[0].Set3 (1, 1, 1)
//        n[0].Norm()
//      } else {
//        temp.Diff (v[2 * i - 2], v[2 * i - 1])
//        temp1.Diff (v[2 * i], v[2 * i - 1])
//        n[2 * i - 1].Ext (temp, temp1)
//        n[2 * i - 1].Norm()
//        temp.Diff (v[2 * i + 1], v[2 * i])
//        temp1.Diff (v[2 * i - 1], v[2 * i])
//        n[2 * i].Cross (temp, temp1)
//        n[2 * i].Dilate (-1)
//        n[2 * i].Norm()
//      }
//      i ++
      y1, y0 := y + dy, y + dy / 2
      z, z1, z2, z3 := f (x, y), f (x1, y), f (x1, y1), f (x, y1)
      z0 := f (x0, y0)
      b0 := true // ok (z)
      b1, b2, b3 := true, true, true // ok (z1), ok (z2), ok (z3)
      c0 := Z < z && z < Z1
      c1, c2, c3 := Z < z1 && z1 < Z1, Z < z2 && z2 < Z1, Z < z3 && z3 < Z1
      if ok (z0) && Z < z0 && z0 < Z1 {
        if b0 && b1 && c0 && c1 {
          figure (pt.Triangles, p, c, v_(x,  y,  z), v_(x1, y,  z1), v_(x0, y0, z0))
        }
        if b1 && b2 && c1 && c2 {
          figure (pt.Triangles, p, c, v_(x1, y,  z1), v_(x1, y1, z2), v_(x0, y0, z0))
        }
        if b2 && b3 && c2 && c3 {
          figure (pt.Triangles, p, c, v_(x1, y1, z2), v_(x,  y1, z3), v_(x0, y0, z0))
        }
        if b3 && b0 && c3 && c0 {
          figure (pt.Triangles, p, c, v_(x,  y1, z3), v_(x,  y,  z),  v_(x0, y0, z0))
        }
      }
    }
//    p.Ins (pt.TriangleFan, v, n, c)
  }
}

func CoSy (p pts.Points, X, Y, Z float64, mit bool) {
  const N = 0.
  cX, cY, cZ := col.LightRed, col.LightGreen, col.LightBlue
  if mit {
    parallelogram (p, cX, v_(N,-Y,-Z), v_(N, Y,-Z), v_(N,-Y, Z))
  }
  R := X / 128
  R1 := X / 16
  G := X
  G1 := G + 2
  var x float64
  fein := X <= 10
  y := -Y
//  var c0 col.Colour
  for y < Y {
    z := -Z
    for z < Z {
//      if y = 0 {
//        c0 = cY
//      } else if z = 0 {
//        c0 = cZ
//      } else {
//        c0 = cX
//      }
      if fein {
        figure (pt.Points, p, cX, v_(N, y, z))
      } else {
        Octahedron (p, cX, N, y, z, R)
      }
      z += 1 // muß gekörnt werden
    }
    y += 1
  }
  figure (pt.Lines, p, cX, v_(-G1, N, N), v_(G1, N, N))
  sphere (p, G1, N, N, R1, cX)
  if mit {
    parallelogram (p, cY, v_(-X, N,-Z), v_(X, N,-Z), v_(-X, N, Z))
  }
  x = -X
  for x < X {
    z := - Z
    for z < Z {
//      if x = 0 {
//        c0 = cX
//      } else if z = 0 {
//        c0 = cZ
//      } else {
//        c0 = cY
//      }
      if fein {
        figure (pt.Points, p, cY, v_(x, N, z))
      } else {
        octahedron (p, cY, x, N, z, R)
      }
      z += 1
    }
    x += 1
  }
  figure (pt.Lines, p, cY, v_(N,-G1, N), v_(N, G1, N))
  sphere (p, N, G1, N, R1, cY)
  if mit {
    parallelogram (p, cZ, v_(-X,-Y, N), v_(X,-Y, N), v_(-X, Y, N))
  }
  x = -X
  for x < X {
    y := -Y
    for y < Y {
//      if x = 0 {
//        c0 = cX
//      } else if y = 0 {
//        c0 = cY
//      } else {
//        c0 = cZ
//      }
      if fein {
        figure (pt.Points, p, cZ, v_(x, y, N))
      } else {
        octahedron (p, cZ, x, y, N, R)
      }
      y += 1
    }
    x += 1
  }
  figure (pt.Lines, p, cZ, v_(N, N,-G1), v_(N, N, G1))
  sphere (p, N, N, G1, R1, cZ)
}

func Tree (p pts.Points, x, y, z, r float64, c col.Colour) {
  v, _ := vectors(2)
  v[0].Set3 (x, y, z)
  for b := 1; b < N / 2; b++ {
    for l := 0; l < N; l++ {
//      rz := r * rand.LongFloat()
//      r0 :=     rz * sin[b]
//      z0 := z + rz * cos[b]
      v[1].SetPolar (x, y, z, r * rand.Real(), float64 (b * angle), float64 (l * angle))
//      v[1].Set3 (x + r0 * cos[l], y + r0 * sin[l], z0)
//      v[1].Inc (v[0])
      p.Ins1 (pt.LineStrip, v, c)
    }
  }
}
