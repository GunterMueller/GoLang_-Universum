package gl

// (c) Christian Maurer   v. 220827 - license see µU.go

import (
  "math"
  "µU/col"
)
const (
  angle = 3 // degrees
  N = 360 / angle
  pi_180 = math.Pi / 180
)
var
  sin, cos []float64

func init() {
  sin, cos = make([]float64, N + 2), make([]float64, N + 2)
  sin[0], cos[0] = 0, 1
  w := 2 * math.Pi / float64(N)
  for i := 1; i < N; i++ {
    sin[i] = math.Sin(float64(i) * w)
    cos[i] = math.Cos(float64(i) * w)
  }
  sin[N], cos[N] = 0, 1
  sin[N+1], cos[N+1] = sin[1], cos[1]
}

func ok (x float64) bool {
  return ! math.IsNaN (x)
}

func fail() {
  panic("wrong number of float64's")
}

func fail1() {
  panic("some pre not met")
}

func point (x []float64) {
  if len(x) != 3 { fail() }
  begin (Points)
  vertex (x[0], x[1], x[2])
  end()
}

func line1 (x []float64) {
  for i := 0; i < 6; i++ { if math.IsNaN(x[i]) { return } }
  vertex (x[0], x[1], x[2])
  vertex (x[3], x[4], x[5])
}

func line (x []float64) {
  if len(x) != 6 { fail() }
  for i := 0; i < 6; i++ { if math.IsNaN(x[i]) { return } }
  begin (Lines)
  vertex (x[0], x[1], x[2])
  vertex (x[3], x[4], x[5])
  end()
}

func lineloop (x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 3 * 3 { fail() }
  begin (LineLoop)
  for i := 0; i < n / 3; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func linestrip (x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 3 * 3 { fail() }
  begin (LineStrip)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func triangle (x ...float64) {
  n := len(x)
  if n != 9 { fail() }
  begin (Triangles)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func trianglestrip (x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 4 * 3 { fail() }
  begin (TriangleStrip)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func trianglefan (x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 4 * 3 { fail() }
  begin (TriangleFan)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func quad (x []float64) {
  if len(x) != 12 { fail() }
  begin (Quads)
  for i := 0; i < 12; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func quadstrip (x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 6 * 3 { fail() }
  begin (QuadStrip)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func horRectangle (x, y, z, x1, y1 float64, up bool) {
  begin (Quads)
  vertex (x,  y,  z)
  vertex (x1, y,  z)
  vertex (x1, y1, z)
  vertex (x,  y1, z)
  end()
}

func vertRectangle (x []float64) {
  if len(x) != 6 { fail() }
  if x[2] == x[5] { fail1() }
  begin (Quads)
  vertex (x[0], x[1], x[2])
  vertex (x[3], x[4], x[2])
  vertex (x[3], x[4], x[5])
  vertex (x[0], x[1], x[5])
  end()
}

func parallelogram (x []float64) {
  if len(x) != 9 { fail() }
  begin (Quads)
  vertex (x[0],               x[1],               x[2])
  vertex (x[0] + x[3],        x[1] + x[4],        x[2] + x[5])
  vertex (x[0] + x[3] + x[6], x[1] + x[4] + x[7], x[2] + x[5] + x[8])
  vertex (x[0] + x[6],        x[1] + x[7],        x[2] + x[8])
  end()
}

func polygon (x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 3 * 3 { fail() }
  begin (Polygons)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func polygon1 (x, y, z []float64) {
  n := len(x)
  begin (Polygons)
  for i := 0; i < n; i += 3 {
    vertex (x[i], y[i], z[i])
  }
  end()
}

const grain = 8 // reasonable compromise between fine grained
                // versus lots of data w.r.t. output efficiency

func curve (w uint, f1, f2, f3 func (float64) float64, t0, t1 float64) {
  mX := float64 (w / grain)
  dt := (t1 - t0) / mX
  begin (Lines)
  for a := t0; a <= t1; a += dt {
    x, y, z := f1 (a), f2 (a), f3 (a)
    a1 := a + dt
    x1, y1, z1 := f1 (a1), f2 (a1), f3 (a1)
    if ok (x) && ok (y) && ok (z) && ok (x1) && ok (y1) && ok (z1) {
      Line (x, y, z, x1, y1, z1)
    }
  }
  end()
}

func cube (x, y, z, a float64) {
  Cuboid (x - a / 2, y - a / 2, z - a / 2, x + a / 2, y + a / 2, z + a / 2)
}

func cubeC (c []col.Colour, x, y, z, a float64) {
  CuboidC (c, x - a / 2, y - a / 2, z - a / 2, x + a / 2, y + a / 2, z + a / 2)
}

func cuboid (x []float64) {
  if len(x) != 6 { fail() }
  VertRectangle (x[0], x[1], x[2], x[3], x[1], x[5])
  VertRectangle (x[3], x[1], x[2], x[3], x[4], x[5])
  VertRectangle (x[3], x[4], x[2], x[0], x[4], x[5])
  VertRectangle (x[0], x[4], x[2], x[0], x[1], x[5])
  horRectangle (x[0], x[1], x[5], x[3], x[4], true)
  horRectangle (x[0], x[1], x[2], x[3], x[4], false)
}

func cuboidC (c []col.Colour, x []float64) {
  colour (c[0])
  VertRectangle (x[0], x[1], x[2], x[3], x[1], x[5]) // front
  colour (c[1])
  VertRectangle (x[3], x[1], x[2], x[3], x[4], x[5]) // right
  colour (c[2])
  VertRectangle (x[3], x[4], x[2], x[0], x[4], x[5]) // back
  colour (c[3])
  VertRectangle (x[0], x[4], x[2], x[0], x[1], x[5]) // left
  colour (c[4])
  horRectangle (x[0], x[1], x[5], x[3], x[4], true)  // top
  colour (c[5])
  horRectangle (x[0], x[1], x[2], x[3], x[4], false) // bottom
}

func cuboid1 (x, y, z, w, d, h, a float64) {
  s, c := math.Sin (a * pi_180), math.Cos (a * pi_180)
  x0, y0 := x  + w * c, y  + w * s
  x1, y1 := x0 - d * s, y0 + d * c
  x2, y2 := x  - d * s, y  + d * c
  z0 := z + h
  VertRectangle (x,  y,  z, x0, y0, z0)
  VertRectangle (x0, y0, z, x1, y1, z0)
  VertRectangle (x1, y1, z, x2, y2, z0)
  VertRectangle (x2, y2, z, x,  y,  z0)
  begin (Quads)
  vertex (x,  y,  z0)
  vertex (x0, y0, z0)
  vertex (x1, y1, z0)
  vertex (x2, y2, z0)
  end()
/*
  n[1].Diff (v[1], v[0])
  n[2].Diff (v[3], v[0])
  n[0].Ext (n[1], n[2]); n[0].Norm()
  for i := 1; i < 4; i++ { n[i].Copy (n[0]) }
*/
//  Quad (f, x, y, z, x2, y2, z, x1, y1, z, x0, y0, z) // XXX
}

func prism (x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 4 * 3 { fail() }
/*
  x1 := make([]float64, n)
  for i := 0; i < n - 3; i++ {
    x1[i] = x[i] + x[i % 3]
  }
*/
  begin (QuadStrip)
  for i := 3; i < n - 3; i += 3 {
    vertex (x[i], x [i + 1], x [i + 2])
    vertex (x[i] + x[0], x[i + 1] + x[1], x[i + 2] + x[2])
  }
  vertex (x[3], x[4], x[5])
  vertex (x[3] + x[0], x[4] + x[1], x[5] + x[2])
  end()
}

func prismC (c []col.Colour, x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 4 * 3 { fail() }
  x1 := make([]float64, n)
  for i := 0; i < n; i++ {
    x1[i] = x[i] + x[i % 3]
  }
  begin (Quads)
  for i := 3; i < n - 3; i += 3 {
    j := i / 3 - 1
    colour (c[j])
    vertex ( x[i]    ,  x[i + 1],  x[i + 2])
    vertex ( x[i + 3],  x[i + 4],  x[i + 5])
    vertex (x1[i + 3], x1[i + 4], x1[i + 5])
    vertex (x1[i],     x1[i + 1], x1[i + 2])
  }
  j := n / 3 - 1
  colour (c[j])
  vertex (x [n - 3], x [n - 2], x [n - 1])
  vertex (x [3],     x [4],     x [5])
  vertex (x1[3],     x1[4],     x1[5])
  vertex (x1[n - 3], x1[n - 2], x1[n - 1])
  end()
}

func parallelepiped (x []float64) { // BUG
  if  len(x) != 12 { fail() }
  prism (x)
}

func octahedron (x, y, z, r float64) {
  d := r * math.Sqrt (2)
  begin (TriangleFan)
  vertex (x,     y,     z + d)
  vertex (x + r, y + r, z)
  vertex (x - r, y + r, z)
  vertex (x - r, y - r, z)
  vertex (x + r, y - r, z)
  vertex (x,     y,     z - d)
  vertex (x + r, y - r, z)
  vertex (x - r, y - r, z)
  vertex (x - r, y + r, z)
  vertex (x + r, y + r, z)
  end()
}

func octahedronC (c []col.Colour, x, y, z, r float64) { // BUG
  d := r * math.Sqrt (2)
  begin (Triangles)
  colour (c[0])
  vertex (x, y, z + d); vertex (x + r, y - r, z); vertex (x + r, y + r, z)
  colour (c[1])
  vertex (x, y, z + d); vertex (x + r, y + r, z); vertex (x - r, y + r, z)
  colour (c[2])
  vertex (x, y, z + d); vertex (x - r, y + r, z); vertex (x - r, y - r, z)
  colour (c[3])
  vertex (x, y, z + d); vertex (x - r, y - r, z); vertex (x + r, y - r, z)
  colour (c[4])
  vertex (x, y, z - d); vertex (x + r, y + r, z); vertex (x + r, y - r, z)
  colour (c[5])
  vertex (x, y, z - d); vertex (x + r, y - r, z); vertex (x - r, y - r, z)
  colour (c[6])
  vertex (x, y, z - d); vertex (x - r, y - r, z); vertex (x - r, y + r, z)
  colour (c[7])
  vertex (x, y, z - d); vertex (x - r, y + r, z); vertex (x + r, y + r, z)
  end()
}

func pyramid (x, y, z, a, b, h float64) {
  begin (TriangleFan)
  vertex (x,     y,     z + h)
  vertex (x + a, y + b, z)
  vertex (x - a, y + b, z)
  vertex (x - a, y - b, z)
  vertex (x + a, y - b, z)
//  vertex (x + a, y + b, z)
  end()
  horRectangle (x - a, y - b, z, x + a, y + b, false)
}

func pyramidC (c []col.Colour, x, y, z, a, b, h float64) {
  begin (TriangleFan)
  colour (c[0])
  vertex (x,     y,     z + h)
  vertex (x + a, y + b, z)
  vertex (x - a, y + b, z)
  colour (c[1])
  vertex (x - a, y - b, z)
  colour (c[2])
  vertex (x + a, y - b, z)
  colour (c[3])
  vertex (x + a, y + b, z)
  end()
  colour (c[4])
  horRectangle (x - a, y - b, z, x + a, y + b, false)
}

func multipyramid (x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 4 * 3 { fail() }
  begin (TriangleFan)
  vertex (x[0], x[1], x[2])
  for i := 1; i < n / 3; i++ {
    j := 3 * i
    vertex (x[j], x[j + 1], x[j + 2])
  }
  vertex (x[3], x[4], x[5])
  end()
}

func multipyramidC (c []col.Colour, x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 4 * 3 { fail() }
  for i := 1; i < n / 3 - 1; i++ {
    j, k := 3 * i, 3 * (i + 1)
    colour (c[i])
    Triangle (x[0], x[1],     x[2],
              x[j], x[j + 1], x[j + 2],
              x[k], x[k + 1], x[k + 2])
  }
  colour (c[n / 3 - 1])
  Triangle (x[0], x[1],         x[2],
            x[n - 3], x[n - 2], x[n - 1],
            x[3], x[4],         x[5])
}

/*/
  func frustum (x, y, z, r, h, h1 float64) { TODO
  if h1 > h { fail() }
  v, n := vectors(N + 2)
  v[0] = vect.New3 (x, y, h)
  n[0] = vect.New3 (0, 0, 1)
  begin (TriangleFan)
  for l := 0; l <= N; l++ {
    v[l+1] = vect.New3 (x + r * cos[l], y + r * sin[l], z)
    n[l+1] = vect.New3 (cos[l], sin[l], r / (h - z))
    n[l+1].Norm()
  }
  end()
  circle (x, y, z, -r)
}
*/
