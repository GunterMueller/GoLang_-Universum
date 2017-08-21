package fig

// (c) murus.org  v. 170817 - license see murus.go

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

func init1 (v, n []vect.Vector, x, y, z float64) {
  v[0].Set3 (x, y, z)
  n[0].Set3 (0, 0, 1)
}

func init2 (v, n []vect.Vector, x, y, z, x1, y1, z1 float64) {
  init1 (v, n, x, y, z)
  v[1].Set3 (x1, y1, z1)
  n[1].Set3 (0, 0, 1)
}

func init3 (v, n []vect.Vector, x, y, z, x1, y1, z1, x2, y2, z2 float64) {
  init2 (v, n, x, y, z, x1, y1, z1)
  v[2].Set3 (x2, y2, z2)
}

func init4 (v, n []vect.Vector, x, y, z, x1, y1, z1, x2, y2, z2, x3, y3, z3 float64) {
  init3 (v, n, x, y, z, x1, y1, z1, x2, y2, z2)
  v[3].Set3 (x3, y3, z3)
}

func vectors (a uint) ([]vect.Vector, []vect.Vector) {
  v, n := make ([]vect.Vector, a), make ([]vect.Vector, a)
  for i := uint(0); i < a; i++ { v[i], n[i] = vect.New(), vect.New() }
  return v, n
}

func ready (p pts.Points) {
  v, n := vectors (1)
  p.Ins (pt.None, v, n, col.Black)
}

func start (p pts.Points, x, y, z, xf, yf, zf float64) {
  v, n := vectors (1)
  init1 (v, n, x, y, z) // Auge
  n[0].Set3 (xf, yf, zf) // Fokus
  p.Ins (pt.Start, v, n, col.Black)
}

func light (p pts.Points, l uint, x, y, z float64, ca, cd col.Colour) {
  v, n := vectors (1)
  v[0].Set3 (x, y, z)
  n[0].Set3 (0, 0, 1)
//  r, g, b := col.LongFloat (ca)
  n[0].Set3 (col.LongFloat (ca)) // r, g, b // ambient colour
  p.InsLight (l, v, n, cd)
}


func point (p pts.Points, x, y, z float64, c col.Colour) {
  v, n := vectors (1)
  init1 (v, n, x, y, z)
  p.Ins (pt.Points, v, n, c)
}

func segment (p pts.Points, x, y, z, x1, y1, z1 float64, c col.Colour) {
  v, n := vectors (2)
  init2 (v, n, x, y, z, x1, y1, z1)
  p.Ins (pt.Lines, v, n, c)
  ready (p)
}

func triangle (p pts.Points, x, y, z, x1, y1, z1, x2, y2, z2 float64, c col.Colour) {
  v, n := vectors (3)
  init3 (v, n, x, y, z, x1, y1, z1, x2, y2, z2)
  n[1].Diff (v[1], v[0])
  n[2].Diff (v[2], v[0])
  n[0].Ext (n[1], n[2])
  n[0].Norm()
  for i := 1; i <= 2; i++ { n[i].Copy (n[0]) }
  p.Ins (pt.Triangles, v, n, c)
}

// func triangleFan (p pts.Points, x, y, z []float64, c col.Colour)

// func triangleStrip (p pts.Points, x, y, z []float64, c col.Colour)

func quad (p pts.Points, x, y, z, x1, y1, z1, x2, y2, z2, x3, y3, z3 float64, c col.Colour) {
  v, n := vectors (4)
  init4 (v, n, x, y, z, x1, y1, z1, x2, y2, z2, x3, y3, z3)
  n[1].Diff (v[1], v[0])
  n[2].Diff (v[3], v[0])
  n[0].Ext (n[1], n[2])
  n[0].Norm()
  for i := 1; i <= 3; i++ {
    n[i].Copy (n[0])
  }
  p.Ins (pt.Quads, v, n, c)
}

// func quadStrip (p pts.Points, x, y, z, x1, y1, z1 []float64, c col.Colour)

func horRectangle (p pts.Points, x, y, z, x1, y1 float64, up bool, c col.Colour) {
  v, n := vectors (4)
  init4 (v, n, x, y, z, x1, y, z, x1, y1, z, x, y1, z)
  goUp := 1.; if ! up { goUp = -1 }
  for i := 0; i <= 3; i++ {
    n[i].Set3 (0, 0, goUp)
  }
  p.Ins (pt.Quads, v, n, c)
}

func vertRectangle (p pts.Points, x, y, z, x1, y1, z1 float64, c col.Colour) {
  v, n := vectors (4)
  if z == z1 { ker.Oops() } // horRectangle
  init4 (v, n, x, y, z, x1, y1, z, x1, y1, z1, x, y, z1)
  n[1].Diff (v[1], v[0])
  n[2].Diff (v[3], v[0])
  n[0].Ext (n[1], n[2])
  for i := 1; i <= 3; i++ { n[i].Copy (n[0]) }
  p.Ins (pt.Quads, v, n, c)
}

func parallelogram (p pts.Points, x, y, z, x1, y1, z1, x2, y2, z2 float64, c col.Colour) {
  quad (p, x, y, z, x1, y1, z1, x1 + x2 - x, y1 + y2 - y, z1 + z2 - z, x2, y2, z2, c)
}

func cube (p pts.Points, x, y, z, a float64, c col.Colour) {
  cuboid (p, x, y, z, x + a, y + a, z + a, c)
}

func cubeC (p pts.Points, x, y, z, a float64, c []col.Colour) {
  cuboidC (p, x, y, z, x + a, y + a, z + a, c)
}

func cuboid (p pts.Points, x, y, z, x1, y1, z1 float64, c col.Colour) {
  sort (&x, &y, &z, &x1, &y1, &z1)
  horRectangle (p, x,  y,  z1, x1, y1, true, c)
  vertRectangle (p, x,  y,  z,  x1, y,  z1, c)
  vertRectangle (p, x1, y,  z,  x1, y1, z1, c)
  vertRectangle (p, x1, y1, z,  x,  y1, z1, c)
  vertRectangle (p, x,  y1, z,  x,  y,  z1, c)
/*
  v, n := vectors (2 * (4 + 1))
  v[0].Set3 (x,  y,  z )
  v[1].Set3 (x,  y,  z1)
  v[2].Set3 (x1, y,  z )
  v[3].Set3 (x1, y,  z1)
  v[4].Set3 (x1, y1, z )
  v[5].Set3 (x1, y1, z1)
  v[6].Set3 (x,  y1, z )
  v[7].Set3 (x,  y1, z1)
  v[8].Set3 (x,  y,  z )
  v[9].Set3 (x,  y,  z1)
  n[0].Set3 (-1., -1., 0.)
  n[1].Set3 (-1., -1., 0.)
  n[2].Set3 ( 1., -1., 0.)
  n[3].Set3 ( 1., -1., 0.)
  n[4].Set3 ( 1.,  1., 0.)
  n[5].Set3 ( 1.,  1., 0.)
  n[6].Set3 (-1.,  1., 0.)
  n[7].Set3 (-1.,  1., 0.)
  n[8].Set3 (-1., -1., 0.)
  n[9].Set3 (-1., -1., 0.)
  p.Ins (pt.QuadStrip, v, n, c)
*/
  horRectangle (p, x,  y,  z,  x1, y1, false, c)
  ready (p)
}

func cuboidC (p pts.Points, x, y, z, x1, y1, z1 float64, c []col.Colour) {
  sort (&x, &y, &z, &x1, &y1, &z1)
  horRectangle (p, x,  y,  z1, x1, y1, true, c[4])  // top
  vertRectangle (p, x,  y,  z,  x1, y,  z1, c[0])   // front
  vertRectangle (p, x1, y,  z,  x1, y1, z1, c[1])   // right
  vertRectangle (p, x1, y1, z,  x,  y1, z1, c[2])   // back
  vertRectangle (p, x,  y1, z,  x,  y,  z1, c[3])   // left
/*
  v, n := vectors (2 * (4 + 1))
  v[0].Set3 (x,  y,  z )
  v[1].Set3 (x,  y,  z1)
  v[2].Set3 (x1, y,  z )
  v[3].Set3 (x1, y,  z1)
  v[4].Set3 (x1, y1, z )
  v[5].Set3 (x1, y1, z1)
  v[6].Set3 (x,  y1, z )
  v[7].Set3 (x,  y1, z1)
  v[8].Set3 (x,  y,  z )
  v[9].Set3 (x,  y,  z1)
  n[0].Set3 (-1., -1., 0.)
  n[1].Set3 (-1., -1., 0.)
  n[2].Set3 ( 1., -1., 0.)
  n[3].Set3 ( 1., -1., 0.)
  n[4].Set3 ( 1.,  1., 0.)
  n[5].Set3 ( 1.,  1., 0.)
  n[6].Set3 (-1.,  1., 0.)
  n[7].Set3 (-1.,  1., 0.)
  n[8].Set3 (-1., -1., 0.)
  n[9].Set3 (-1., -1., 0.)
  p.Ins (pt.QuadStrip, v, n, c)
*/
  horRectangle (p, x,  y,  z,  x1, y1, false, c[5]) // bottom
  ready (p)
}

func cuboid1 (p pts.Points, x, y, z, b, t, h, a float64, f col.Colour) {
  s, c := math.Sin (a * pi_180), math.Cos (a * pi_180)
  x1 := x  + b * c; y1 := y  + b * s; z1 := z + h
  x2 := x1 - t * s; y2 := y1 + t * c
  x3 := x  - t * s; y3 := y  + t * c
  quad (p, x, y, z1, x1, y1, z1, x2, y2, z1, x3, y3, z1, f)
  vertRectangle (p, x,  y,  z,  x1, y1, z1, f)
  vertRectangle (p, x1, y1, z,  x2, y2, z1, f)
  vertRectangle (p, x2, y2, z,  x3, y3, z1, f)
  vertRectangle (p, x3, y3, z,  x,  y,  z1, f)
  quad (p, x, y, z, x3, y3, z, x2, y2, z, x1, y1, z, f)
}

func prism (p pts.Points, x, y, z []float64, c col.Colour) {
// top missing
  n := uint(len (x))
  if n < 4 { ker.Oops() }
  n -- // top !
  for i := uint(0); i < n - 1; i++ {
    quad (p,
          x[i],          y[i],          z[i],
          x[i+1],        y[i+1],        z[i+1],
          x[i+1] + x[n], y[i+1] + y[n], z[i+1] + z[n],
          x[i]   + x[n], y[i]   + y[n], z[i]   + z[n], c)
  }
  i := uint(n - 1)
  quad (p,
        x[i],        y[i],        z[i],
        x[0],        y[0],        z[0],
        x[0] + x[n], y[0] + y[n], z[0] + z[n],
        x[i] + x[n], y[i] + y[n], z[i] + z[n], c)
// bottom missing
}

func parallelepiped (p pts.Points, x0, y0, z0, x1, y1, z1, x2, y2, z2, x3, y3, z3 float64, c col.Colour) {
  parallelogram (p, x0, y0, z0, x1, y1, z1, x3, y3, z3, c)
  parallelogram (p, x0, y0, z0, x2, y2, z2, x1, y1, z1, c)
  parallelogram (p, x0, y0, z0, x3, y3, z3, x2, y2, z2, c)
  parallelogram (p, x1, y1, z1, x1 + x2 - x0, y1 + y2 - y0, z1 + z2 - z0, x1 + x3 - x0, y1 + y3 - y0, z1 + z3 - z0, c)
  parallelogram (p, x2, y2, z2, x2 + x3 - x0, y2 + y3 - y0, z2 + z3 - z0, x2 + x1 - x0, y2 + y1 - y0, z2 + z1 - z0, c)
  parallelogram (p, x3, y3, z3, x3 + x1 - x0, y3 + y1 - y0, z3 + z1 - z0, x3 + x2 - x0, y3 + y2 - y0, z3 + z2 - z0, c)
}

func pyramid (p pts.Points, x, y, z, x1, y1, z1, x2, y2, z2 float64, c col.Colour) {
  triangle (p, x,  y,  z, x1, y,  z1, x2, y2, z2, c)
  triangle (p, x1, y,  z, x1, y1, z,  x2, y2, z2, c)
  triangle (p, x1, y1, z, x,  y1, z1, x2, y2, z2, c)
  triangle (p, x,  y1, z, x,  y,  z,  x2, y2, z2, c)
  horRectangle (p, x, y, z, x1, y1, false, c)
}

func octahedron (p pts.Points, x, y, z, r float64, c col.Colour) {
  d := r * math.Sqrt (2)
  triangle (p, x + r, y + r, z, x - r, y + r, z, x, y, z + d, c)
  triangle (p, x - r, y + r, z, x - r, y - r, z, x, y, z + d, c)
  triangle (p, x - r, y - r, z, x + r, y - r, z, x, y, z + d, c)
  triangle (p, x + r, y - r, z, x + r, y + r, z, x, y, z + d, c)
  triangle (p, x + r, y + r, z, x - r, y + r, z, x, y, z - d, c)
  triangle (p, x - r, y + r, z, x - r, y - r, z, x, y, z - d, c)
  triangle (p, x - r, y - r, z, x + r, y - r, z, x, y, z - d, c)
  triangle (p, x + r, y - r, z, x + r, y + r, z, x, y, z - d, c)
}

func multipyramid (p pts.Points, x, y, z []float64, c col.Colour) {
  n := len (x)
  if n < 4 { ker.Oops() }
  n -- // top !
  for i := 0; i < n - 1; i++ {
    triangle (p, x[i], y[i], z[i], x[i+1], y[i+1], z[i+1], x[n], y[n], z[n], c)
  }
  triangle (p, x[n-1], y[n-1], z[n-1], x[0], y[0], z[0], x[n], y[n], z[n], c)
// bottom missing, because it need not be even
}

func circle (p pts.Points, x, y, z, r float64, c col.Colour) {
  circleSegment (p, x, y, z, r, 0, 360, c)
}

func circleSegment (p pts.Points, x, y, z, r, a, b float64, c col.Colour) {
  if r == 0 {
    point (p, x, y, z, c)
    return
  }
  A := uint(math.Floor (a / float64 (angle) + 0.5))
  B := uint(math.Floor (b / float64 (angle) + 0.5))
  C := B - A + 2
  v, n := vectors (C)
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
    point (p, x, y, z, f)
    return
  }
  s, c := math.Sin (a * pi_180), math.Cos (a * pi_180)
  C := uint(N) + 2
  v, n := vectors (C)
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
  v, n := vectors (N + 2)
  v[0].Set3 (x, y, z + r)
  n[0].Set3 (0, 0, 1)
  r0, z0 := r * sin[1], z + r * cos[1]
  for l := 0; l <= N; l++ {
    v[1 + l].Set3 (x + r0 * cos[l], y + r0 * sin[l], z0)
    n[1 + l].Set3 (sin[1] * cos[l], sin[1] * sin[l], cos[1])
  }
  p.Ins (pt.TriangleFan, v, n, f)

  v, n = vectors (2 * (N + 1))
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
  v, n = vectors (N + 2)
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
  v, n := vectors (N + 2)
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
  v, n := vectors (N + 2)
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
  v, n := vectors (C)
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
  v, n := vectors (C)
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
//    v, n := vectors (2 * N)
    for l := 0; l < N; l++ {
      quad (p,
            x + s0 * cos[l],   y + s0 * sin[l],   z0,
            x + s0 * cos[l+1], y + s0 * sin[l+1], z0,
            x + s1 * cos[l+1], y + s1 * sin[l+1], z1,
            x + s1 * cos[l],   y + s1 * sin[l],   z1, c)
/*
      v[2*l].Set3 (x + s0 * cos[l], y + s0 * sin[l], z0)
      n[2*l].Set3 (1., 1., 1.)
      v[2*l+1].Set3 (x + s0 * cos[l+1], y + s0 * sin[l+1], z0)
      n[2*l+1].Set3 (1, 1, 1)
*/
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
      quad (p,
            x + x0 * c - y00 * s, y + x0 * s + y00 * c, z + s0 * sin[l],
            x + x0 * c - y01 * s, y + x0 * s + y01 * c, z + s0 * sin[l+1],
            x + x1 * c - y11 * s, y + x1 * s + y11 * c, z + s1 * sin[l+1],
            x + x1 * c - y10 * s, y + x1 * s + y10 * c, z + s1 * sin[l], f)
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
      segment (p, x, y, z, x1, y1, z1, c)
    }
  }
}

func surface (p pts.Points, f RealFunc2, X, Y, Z, X1, Y1, Z1 float64, c col.Colour) {
  if X == X1 || Y == Y1 || Z == Z1 { return }
  if X1 < X { X, X1 = X1, X }
  if Y1 < Y { Y, Y1 = Y1, Y }
  if Z1 < Z { Z, Z1 = Z1, Z }
  dx, dy := (X1 - X) / float64 (scr.Wd() / grain), (Y1 - Y) / float64 (scr.Ht() / grain)
  for x := X; x <= X1; x += dx {
/*
    y := Y
    n := uint(0)
    for y <= Y1 {
      n ++
      y += dy
    }
*/
// die Anwendung der OpenGL-Ausgabe in gl von TriangleFan ist noch fehlerhaft
    x1, x0 := x + dx, x + dx / 2
//    temp, temp1 := vect.New(), vect.New()
//    v, n := vectors (2 * n) // (2 * n + 1)                 ? ? ? ? ? ? ? ? ? ? ? ? ?
    for y := Y; y <= Y1; y += dy {
/*
    for i := uint(0); i < n; i++ { // oder i <= n            ? ? ? ? ? ? ? ? ? ? ? ? ?
      v[2 * i].Set3 (x, y,  z)
      v[2 * i + 1].Set3 (x, y1, z1)
      if i == 0 { // ?
        n[0].Set3 (1, 1, 1)
        n[0].Norm()
      } else {
        temp.Diff (v[2 * i - 2], v[2 * i - 1])
        temp1.Diff (v[2 * i], v[2 * i - 1])
        n[2 * i - 1].Ext (temp, temp1)
        n[2 * i - 1].Norm()
        temp.Diff (v[2 * i + 1], v[2 * i])
        temp1.Diff (v[2 * i - 1], v[2 * i])
        n[2 * i].Cross (temp, temp1)
        n[2 * i].Dilate (-1)
        n[2 * i].Norm()
      }
      i ++
*/
      y1, y0 := y + dy, y + dy / 2
      z, z1, z2, z3 := f (x, y), f (x1, y), f (x1, y1), f (x, y1)
      z0 := f (x0, y0)
      b0 := true // ok (z)
      b1, b2, b3 := true, true, true // ok (z1), ok (z2), ok (z3)
      c0 := Z < z && z < Z1
      c1, c2, c3 := Z < z1 && z1 < Z1, Z < z2 && z2 < Z1, Z < z3 && z3 < Z1
      if ok (z0) && Z < z0 && z0 < Z1 {
        if b0 && b1 && c0 && c1 {
          triangle (p, x,  y,  z,  x1, y,  z1, x0, y0, z0, c)
        }
        if b1 && b2 && c1 && c2 {
          triangle (p, x1, y,  z1, x1, y1, z2, x0, y0, z0, c)
        }
        if b2 && b3 && c2 && c3 {
          triangle (p, x1, y1, z2, x,  y1, z3, x0, y0, z0, c)
        }
        if b3 && b0 && c3 && c0 {
          triangle (p, x,  y1, z3, x,  y,  z,  x0, y0, z0, c)
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
    parallelogram (p, N,-Y,-Z, N, Y,-Z, N,-Y, Z, cX)
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
/*
      if y = 0 {
        c0 = cY
      } else if z = 0 {
        c0 = cZ
      } else {
        c0 = cX
      }
*/
      if fein {
        point (p, N, y, z, /* c0 */ cX)
      } else {
        Octahedron (p, N, y, z, R, /* c0 */ cX)
      }
      z += 1 // muß gekörnt werden
    }
    y += 1
  }
  segment (p, -G1, N, N, G1, N, N, cX)
  sphere (p, G1, N, N, R1, cX)
  if mit {
    parallelogram (p, -X, N,-Z, X, N,-Z,-X, N, Z, cY)
  }
  x = -X
  for x < X {
    z := - Z
    for z < Z {
/*
      if x = 0 {
        c0 = cX
      } else if z = 0 {
        c0 = cZ
      } else {
        c0 = cY
      }
*/
      if fein {
        point (p, x, N, z, /* c0 */ cY)
      } else {
        octahedron (p, x, N, z, R, /* c0 */ cY)
      }
      z += 1
    }
    x += 1
  }
  segment (p, N,-G1, N, N, G1, N, cY)
  sphere (p, N, G1, N, R1, cY)
  if mit {
    parallelogram (p, -X,-Y, N, X,-Y, N,-X, Y, N, cZ)
  }
  x = -X
  for x < X {
    y := -Y
    for y < Y {
/*
      if x = 0 {
        c0 = cX
      } else if y = 0 {
        c0 = cY
      } else {
        c0 = cZ
      }
*/
      if fein {
        point (p, x, y, N, /* c0 */ cZ)
      } else {
        octahedron (p, x, y, N, R, /* c0 */ cZ)
      }
      y += 1
    }
    x += 1
  }
  segment (p, N, N,-G1, N, N, G1, cZ)
  sphere (p, N, N, G1, R1, cZ)
}

func Tree (p pts.Points, x, y, z, r float64, c col.Colour) {
  v, _ := vectors (2)
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
