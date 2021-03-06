package gl

// (c) Christian Maurer   v. 170910 - license see µu.go

import (
  "math"
  "µu/col"
)
const (
  angle = 3 // Grad
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

func sort (x, y, z, x1, y1, z1 *float64) {
  if *x1 < *x { *x, *x1 = *x1, *x }
  if *y1 < *y { *y, *y1 = *y1, *y }
  if *z1 < *z { *z, *z1 = *z1, *z }
}

// light ///////////////////////////////////////////////////////////////////////

/* XXX
func light (l uint, x, y, z float64, ca, cd col.Colour) {
  v := vect.New3 (x, y, z)
//  r, g, b := col.LongFloat (ca)
  n := vect.New3 (col.LongFloat (ca)) // r, g, b // ambient colour
  InsLight (l, v, n, cd)
}
*/

/*
func clear() {
  Clear()
}
*/

// extended figures ////////////////////////////////////////////////////////////

func fail() {
  panic("wrong number of float64's")
}

func point (c col.Colour, x []float64) {
  if len(x) != 3 { fail() }
  begin (POINTS)
  vertexC (x[0], x[1], x[2], c)
  end()
}

func line (c col.Colour, x []float64) {
  if len(x) != 6 { fail() }
  begin (LINES)
  colour (c)
  vertex (x[0], x[1], x[2])
  vertex (x[3], x[4], x[5])
  end()
}

func lineloop (c col.Colour, x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 3 * 3 { fail() }
  begin (LINE_LOOP)
  colour (c)
  for i := 0; i < n / 3; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func linestrip (c col.Colour, x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 3 * 3 { fail() }
  begin (LINE_STRIP)
  colour (c)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func triangle (c col.Colour, x []float64) {
  n := len(x)
  if n != 9 { fail() }
  begin (TRIANGLES)
  colour (c)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func trianglestrip (c col.Colour, x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 4 * 3 { fail() }
  begin (TRIANGLE_STRIP)
  colour (c)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func trianglefan (c col.Colour, x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 4 * 3 { fail() }
  begin (TRIANGLE_FAN)
  colour (c)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func quad (c col.Colour, x []float64) {
  if len(x) != 12 { fail() }
  begin (QUADS)
  colour (c)
  for i := 0; i < 12; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func quadstrip (c col.Colour, x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 6 * 3 { fail() }
  begin (QUAD_STRIP)
  colour (c)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func polygon (c col.Colour, x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 3 * 3 { fail() }
  begin (POLYGON)
  colour (c)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i + 1], x[i + 2])
  }
  end()
}

func horRectangle (x, y, z, x1, y1 float64, up bool) {
  vertex (x,  y,  z)
  vertex (x1, y,  z)
  vertex (x1, y1, z)
  vertex (x,  y1, z)
}

func horRectangleC (c col.Colour, x, y, z, x1, y1 float64, up bool) {
  colour (c)
  horRectangle (x, y, z, x1, y1, up)
}

func vertRectangle (x []float64) {
  if len(x) != 6 || x[2] == x[5] { fail() }
  vertex (x[0], x[1], x[2])
  vertex (x[3], x[4], x[2])
  vertex (x[3], x[4], x[5])
  vertex (x[0], x[1], x[5])
}

func vertRectangleC (c col.Colour, x []float64) {
  colour (c)
  vertRectangle (x)
}

func parallelogram (c col.Colour, x []float64) {
  if len(x) != 9 { fail() }
  begin (QUADS)
  colour (c)
  vertex (x[0],               x[1],               x[2])
  vertex (x[0] + x[3],        x[1] + x[4],        x[2] + x[5])
  vertex (x[0] + x[3] + x[6], x[1] + x[4] + x[7], x[2] + x[5] + x[8])
  vertex (x[0] + x[6],        x[1] + x[7],        x[2] + x[8])
  end()
}

func cube (c col.Colour, x, y, z, a float64) {
  Cuboid (c, x - a / 2, y - a / 2, z - a / 2, x + a / 2, y + a / 2, z + a / 2)
}

func cubeC (c []col.Colour, x, y, z, a float64) {
  CuboidC (c, x - a / 2, y - a / 2, z - a / 2, x + a / 2, y + a / 2, z + a / 2)
}

func cuboid (c col.Colour, x []float64) {
  if len(x) != 6 { fail() }
//  sort (&x, &y, &z, &x1, &y1, &z1)
  begin (QUADS)
  colour (c)
  VertRectangle (x[0], x[1], x[2], x[3], x[1], x[5])
  VertRectangle (x[3], x[1], x[2], x[3], x[4], x[5])
  VertRectangle (x[3], x[4], x[2], x[0], x[4], x[5])
  VertRectangle (x[0], x[4], x[2], x[0], x[1], x[5])
  horRectangle (x[0], x[1], x[5], x[3], x[4], true)
  horRectangle (x[0], x[1], x[2], x[3], x[4], false)
  end()
}

func cuboidC (c []col.Colour, x []float64) {
//  sort (&x, &y, &z, &x1, &y1, &z1)
  begin (QUADS)
  VertRectangleC (c[0], x[0], x[1], x[2], x[3], x[1], x[5]) // front
  VertRectangleC (c[1], x[3], x[1], x[2], x[3], x[4], x[5]) // right
  VertRectangleC (c[2], x[3], x[4], x[2], x[0], x[4], x[5]) // back
  VertRectangleC (c[3], x[0], x[4], x[2], x[0], x[1], x[5]) // left
  horRectangleC (c[4], x[0], x[1], x[5], x[3], x[4], true)  // top
  horRectangleC (c[5], x[0], x[1], x[2], x[3], x[4], false) // bottom
  end()
}

func cuboid1 (f col.Colour, x, y, z, b, t, h, a float64) {
  s, c := math.Sin (a * pi_180), math.Cos (a * pi_180)
  x1 := x  + b * c; y1 := y  + b * s; z1 := z + h
  x2 := x1 - t * s; y2 := y1 + t * c
  x3 := x  - t * s; y3 := y  + t * c
  begin (QUADS)
  colour (f)
  VertRectangle (x,  y,  z, x1, y1, z1)
  VertRectangle (x1, y1, z, x2, y2, z1)
  VertRectangle (x2, y2, z, x3, y3, z1)
  VertRectangle (x3, y3, z, x,  y,  z1)
  vertex (x,  y,  z1)
  vertex (x1, y1, z1)
  vertex (x2, y2, z1)
  vertex (x3, y3, z1)
  end()
/*
  n[1].Diff (v[1], v[0])
  n[2].Diff (v[3], v[0])
  n[0].Ext (n[1], n[2]); n[0].Norm()
  for i := 1; i < 4; i++ { n[i].Copy (n[0]) }
*/
//  Quad (f, x, y, z, x3, y3, z, x2, y2, z, x1, y1, z) // XXX
}

func prism (c col.Colour, x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 4 * 3 { fail() }
  x1 := make([]float64, n)
  for i := 3; i < n; i++ {
    x1[i] = x[i] + x[i % 3]
  }
  begin (QUAD_STRIP)
  colour (c)
  for i := 3; i < n - 3; i += 3 {
    vertex (x [i], x [i + 1], x [i + 2])
    vertex (x1[i], x1[i + 1], x1[i + 2])
  }
  vertex (x [3], x [4], x [5])
  vertex (x1[3], x1[4], x1[5])
  end()
}

func prismC (c []col.Colour, x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 4 * 3 { fail() }
  x1 := make([]float64, n)
  for i := 0; i < n; i++ {
    x1[i] = x[i] + x[i % 3]
  }
  begin (QUADS)
  for i := 3; i < n - 3; i += 3 {
    j := i / 3 - 1
    vertexC ( x[i]    ,  x[i + 1],  x[i + 2], c[j])
    vertexC ( x[i + 3],  x[i + 4],  x[i + 5], c[j])
    vertexC (x1[i + 3], x1[i + 4], x1[i + 5], c[j])
    vertexC (x1[i],     x1[i + 1], x1[i + 2], c[j])
  }
  j := n / 3 - 1
  vertexC (x [n - 3], x [n - 2], x [n - 1], c[j])
  vertexC (x [3],     x [4],     x [5],     c[j])
  vertexC (x1[3],     x1[4],     x1[5],     c[j])
  vertexC (x1[n - 3], x1[n - 2], x1[n - 1], c[j])
  end()
}

func parallelepiped (c col.Colour, x []float64) { // BUG
  if  len(x) != 12 { fail() }
  prism (c, x)
}

func pyramid (c col.Colour, x, y, z, a, b, h float64) {
  begin (TRIANGLE_FAN)
  colour (c)
  vertex (x,     y,     z + h)
  vertex (x + a, y + b, z)
  vertex (x - a, y + b, z)
  vertex (x - a, y - b, z)
  vertex (x + a, y - b, z)
//  vertex (x + a, y + b, z)
  end()
  begin (QUADS)
  horRectangle (x - a, y - b, z, x + a, y + b, false)
  end()
}

func pyramidC (c []col.Colour, x, y, z, a, b, h float64) {
  begin (TRIANGLE_FAN)
  vertexC (x,     y,     z + h, c[0])
  vertexC (x + a, y + b, z,     c[0])
  vertexC (x - a, y + b, z,     c[0])
  vertexC (x - a, y - b, z,     c[1])
  vertexC (x + a, y - b, z,     c[2])
  vertexC (x + a, y + b, z,     c[3])
  end()
  begin (QUADS)
  horRectangleC (c[4], x - a, y - b, z, x + a, y + b, false)
  end()
}

func octahedron (c col.Colour, x, y, z, r float64) {
  d := r * math.Sqrt (2)
  begin (TRIANGLE_FAN)
  colour (c)
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
  begin (TRIANGLES)
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

func multipyramid (c col.Colour, x []float64) {
  n := len(x)
  if n % 3 != 0 || n < 4 * 3 { fail() }
  begin (TRIANGLE_FAN)
  colour (c)
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
    Triangle (c[i],
              x[0], x[1],     x[2],
              x[j], x[j + 1], x[j + 2],
              x[k], x[k + 1], x[k + 2])
  }
  Triangle (c[n / 3 - 1],
            x[0], x[1],         x[2],
            x[n - 3], x[n - 2], x[n - 1],
            x[3], x[4],         x[5])
}

func circle (c col.Colour, x, y, z, r float64) {
//  circleSegment (c, x, y, z, r, 0, 360)
  if r == 0 {
    Point (c, x, y, z)
    return
  }
//  if r < 0 { r = -r } // change orientation
  begin (TRIANGLE_FAN)
  colour (c)
  vertex (x, y, z)
  for i := 0; i <= N; i++ {
    vertex (x + r * cos[i], y + r * sin[i], z)
  }
  end()
}

func circleSegment (c col.Colour, x, y, z, r, a, b float64) {
  if r == 0 {
    Point (c, x, y, z)
    return
  }
//  if r < 0 { r = -r } // change orientation
  aa := uint(math.Floor (a / float64 (angle) + 0.5))
  bb := uint(math.Floor (b / float64 (angle) + 0.5))
  begin (TRIANGLE_FAN)
  colour (c)
  vertex (x, y, z)
  for i := aa; i <= bb; i++ {
    vertex (x + r * cos[i], y + r * sin[i], z)
  }
  end()
}

func vertCircle (f col.Colour, x, y, z, r, a float64) {
  if r == 0 {
    Point (f, x, y, z)
    return
  }
//  if r < 0 { r = -r } // change orientation
  s, c := math.Sin (a * pi_180), math.Cos (a * pi_180)
  begin (TRIANGLE_FAN)
  colour (f)
  vertex (x, y, z)
  for i := 0; i <= N; i++ {
    vertex (x - r * s * cos[i], y + r * c * cos[i], z + r * sin[i])
  }
  end()
}

func sphere (f col.Colour, x, y, z, r float64) {
  colour (f)
  r0, z0 := r * sin[1], z + r * cos[1]
  begin (TRIANGLE_FAN)
  vertex (x, y, z + r)
  for i := 0; i <= N; i++ {
    vertex (x + r0 * cos[i], y + r0 * sin[i], z0)
  }
  end()
  begin (QUAD_STRIP)
  for i := 1; i <= N / 2 - 2; i++ {
    r0, z0 =  r * sin[i],     z + r * cos[i]
    r1, z1 := r * sin[i + 1], z + r * cos[i + 1]
    for j := 0; j <= N; j++ {
      s, c := sin[j], cos[j]
      vertex (x + r0 * c, y + r0 * s, z0)
      vertex (x + r1 * c, y + r1 * s, z1)
    }
  }
  end()
  r0, z0 = r * sin[N / 2 - 1], z + r * cos[N / 2 - 1]
  begin (TRIANGLE_FAN)
  vertex (x, y, z - r)
  for i := N; i >= 0; i -= 1 {
    vertex (x + r0 * cos[i], y + r0 * sin[i], z0)
  }
  end()
/*
  n := vectors(N + 2)
  n[0] = vect.New3 (0, 0, 1)
  for l := 0; l <= N; l++ {
    n[1 + l] = vect.New3 (sin[1] * cos[l], sin[1] * sin[l], cos[1])
  }
  n = vectors(2 * (N + 1))
  for i := 1; i <= N / 2 - 2; i++ {
    for j := 0; j <= N; j++ {
      n[2 * j] = vect.New3 (sin[i] * c, sin[i] * s, cos[i])
      n[2 * j + 1].Set3 (sin[i+1] * c, sin[i+1] * s, cos[i+1])
    }
  }
  n = vectors(N + 2)
  n[0] = vect.New3 (0, 0, -1)
  for l := N; l >= 0; l -= 1 {
    n[1 + N-l] = vect.New3 (sin[b] * cos[l], sin[b] * sin[l], cos[b])
  }
*/
}

func cone (c col.Colour, x, y, z, r, h float64) {
  colour (c)
  begin (TRIANGLE_FAN)
  vertex (x, y, z + h)
  for l := 0; l <= N; l++ {
    vertex (x + r * cos[l], y + r * sin[l], z)
/*
  n := vectors(N + 2)
  n[0] = vect.New3 (0, 0, 1)
  for l := 0; l <= N; l++ {
    n[l+1] = vect.New3 (cos[l], sin[l], r / (h - z))
    n[l+1].Norm()
*/
  }
  end()
  circle (c, x, y, z, -r)
}

func frustum (c col.Colour, x, y, z, r, h, h1 float64) {
/* XXX
  if h1 > h { fail() }
  v, n := vectors(N + 2)
  v[0] = vect.New3 (x, y, h)
  n[0] = vect.New3 (0, 0, 1)
  begin (TRIANGLE_FAN)
  for l := 0; l <= N; l++ {
    v[l+1] = vect.New3 (x + r * cos[l], y + r * sin[l], z)
    n[l+1] = vect.New3 (cos[l], sin[l], r / (h - z))
    n[l+1].Norm()
  }
  end()
  circle (x, y, z, -r, c)
*/
}

func doubleCone (c col.Colour, x, y, z, r, h float64) {
  cone (c, x, y, z - h, r,  h)
  cone (c, x, y, z + h, r, -h)
}

func cylinder (c col.Colour, x, y, z, r, h float64) {
//  cylinderSegment (c, x, y, z, r, h, 0, 360)
  circle (c, x, y, z,    -r)
  circle (c, x, y, z + h, r)
  begin (QUAD_STRIP)
  colour (c)
  for i := 0; i <= N; i++ {
    vertex (x + r * cos[i], y + r * sin[i], z)
    vertex (x + r * cos[i], y + r * sin[i], z + h)
  }
  end()
}

func cylinderSegment (c col.Colour, x, y, z, r, h, a, b float64) {
  circleSegment (c, x, y, z,    -r, a, b)
  circleSegment (c, x, y, z + h, r, a, b)
  aa := uint(math.Floor (a / float64 (angle) + 0.5))
  bb := uint(math.Floor (b / float64 (angle) + 0.5))
  begin (QUAD_STRIP)
  for i := aa; i <= bb; i++ {
    vertex (x + r * cos[i], y + r * sin[i], z)
    vertex (x + r * cos[i], y + r * sin[i], z + h)
  }
  end()
}

func horCylinder (c col.Colour, x, y, z, r, l, a float64) {
  if r == 0 {
    vertCircle (c, x, y, z, r, a)
    return
  }
  sa, ca := math.Sin (a * pi_180), math.Cos (a * pi_180)
  dx, dy := l * ca, l * sa
  vertCircle (c, x, y, z, -r, a)
  vertCircle (c, x + dx, y + dy, z, r, a)
  begin (QUAD_STRIP)
  colour (c)
  for i := 0; i <= 2 * N; i += 2 {
    si, ci := sin[i / 2], cos[i / 2]
    sci, cci := sa * ci, ca * ci
    x0, y0, z0 := x - r * sci, y + r * cci, z + r * si
    vertex (x0, y0, z0)
    vertex (x0 + dx, y0 + dy, z0)
  }
  end()
}

func torus (c col.Colour, x, y, z, R, r float64) {
  if r <= 0 || R <= 0 { fail() }
  colour (c)
  begin (QUAD_STRIP)
  for i := 0; i < N; i++ {
    s0, s1 := R + r * cos[i], R + r * cos[i+1]
    z0, z1 := z + r * sin[i], z + r * sin[i+1]
    for j := 0; j < N; j++ {
      vertex (x + s0 * cos[j],     y + s0 * sin[j],     z0)
      vertex (x + s0 * cos[j + 1], y + s0 * sin[j + 1], z0)
      vertex (x + s1 * cos[j + 1], y + s1 * sin[j + 1], z1)
      vertex (x + s1 * cos[j],     y + s1 * sin[j],     z1)
      vertex (x + s0 * cos[j],     y + s0 * sin[j],     z0)
//      n[2*j].Set3 (1., 1., 1.)
      vertex (x + s0 * cos[j + 1], y + s0 * sin[j + 1], z0)
//      n[2*j+1].Set3 (1, 1, 1)
    }
  }
  end()
}

func horTorus (c col.Colour, x, y, z, R, r, a float64) { // XXX a ?
  if r <= 0 || R <= 0 { fail() }
  for a <= -180 { a += 180 }
  for a >=  180 { a -= 180 }
  sa, ca := math.Sin (a * pi_180), math.Cos (a * pi_180)
  begin (QUAD_STRIP)
  colour (c)
  for i := 0; i < N; i++ {
    s0, s1 := R + r * cos[i], R + r * cos[i + 1]
    x0, x1 := r * sin[i], r * sin[i + 1]
    for j := 0; j < N; j++ { //  x -> x * c - y * sa, y -> x * sa + y * ca
      y00, y01 := s0 * cos[j], s0 * cos[j + 1]
      y10, y11 := s1 * cos[j], s1 * cos[j + 1]
      vertex (x + x0 * ca - y00 * sa, y + x0 * sa + y00 * ca, z + s0 * sin[j])
      vertex (x + x0 * ca - y01 * sa, y + x0 * sa + y01 * ca, z + s0 * sin[j + 1])
      vertex (x + x1 * ca - y11 * sa, y + x1 * sa + y11 * ca, z + s1 * sin[j + 1])
      vertex (x + x1 * ca - y10 * sa, y + x1 * sa + y10 * ca, z + s1 * sin[j])
    }
  }
  end()
}

func paraboloid (c col.Colour, x, y, z, p float64) {
// XXX
}

func horParaboloid (c col.Colour, x, y, z, p, a float64) {
// XXX
}

//func hyperboloid (c col.Colour, x, y, z, ___ float64) { // XXX

func ok (x float64) bool {
  return ! math.IsNaN (x)
}

const grain = 8 // reasonable compromise between fine grained
                // versus lots of data w.r.t. output efficiency

func curve (c col.Colour, w uint, f1, f2, f3 func (float64) float64,
            t0, t1 float64) {
  mX := float64 (w / grain)
  dt := (t1 - t0) / mX
  begin (LINES)
  colour (c)
  for a := t0; a <= t1; a += dt {
    x, y, z := f1 (a), f2 (a), f3 (a)
    a1 := a + dt
    x1, y1, z1 := f1 (a1), f2 (a1), f3 (a1)
    if ok (x) && ok (y) && ok (z) && ok (x1) && ok (y1) && ok (z1) {
      Line (c, x, y, z, x1, y1, z1)
    }
  }
  end()
}

func surface (c col.Colour, f func (float64, float64) float64,
              X, Y, Z, X1, Y1, Z1 float64, wd, ht, g uint) {
  if X == X1 || Y == Y1 || Z == Z1 { fail() }
  if X1 < X { X, X1 = X1, X }
  if Y1 < Y { Y, Y1 = Y1, Y }
  if Z1 < Z { Z, Z1 = Z1, Z }
  if g < 4 { g = 4 }
  dx, dy := (X1 - X) / float64(wd / g), (Y1 - Y) / float64(ht / g)
//  v, n := vectors(3)
//  w1, w2 := vect.New(), vect.New()
  begin (TRIANGLE_STRIP)
  colour (c)
  for y := Y; y < Y1; y += dy {
    for x := X; x < X1; x += dx {
      x1, y1 := x + dx, y + dy
//
//    1
//    | \
//    |   \
//    0 --- 2
//
      z0, z1, z2 := f(x, y), f(x, y1), f(x1, y)
      if ok(z0) && ok(z1) && ok(z2) {
        vertex (x, y,  z0)
        vertex (x, y,  z1)
        vertex (x, y1, z2)
//        w1.Diff (v[2], v[0])
//        w2.Diff (v[1], v[0])
//        n[0].Ext (w1, w2); n[0].Norm()
//        n[1].Copy (n[0])
//        n[2].Copy (n[0])
      }
//
//    0 --- 2
//      \   |
//        \ |
//          1
//
      z0, z1, z2 = f(x, y1), f(x1, y), f(x1, y1)
      if ok(z0) && ok(z1) && ok(z2) {
        vertex (x, y1,  z0)
        vertex (x1, y,  z1)
        vertex (x1, y1, z2)
//        w1.Diff (v[1], v[0])
//        w2.Diff (v[2], v[0])
//        n[0].Ext (w1, w2); n[0].Norm()
//        n[1].Copy (n[0])
//        n[2].Copy (n[0])
      }
    }
  }
  end()
}

func coSy (X, Y, Z float64, mit bool) {
/*
  const N = 0.
  cX, cY, cZ := col.LightRed, col.LightGreen, col.LightBlue
  if mit {
    parallelogram (cX, N,-Y,-Z), N, Y,-Z), N,-Y, Z))
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
        figure (POINTS, cX, N, y, z))
      } else {
        Octahedron (cX, N, y, z, R)
      }
      z += 1 // muß gekörnt werden
    }
    y += 1
  }
  figure (LINES, cX, -G1, N, N), G1, N, N))
  sphere (G1, N, N, R1, cX)
  if mit {
    parallelogram (cY, -X, N,-Z), X, N,-Z), -X, N, Z))
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
        figure (POINTS, cY, x, N, z))
      } else {
        octahedron (cY, x, N, z, R)
      }
      z += 1
    }
    x += 1
  }
  figure (LINES, cY, N,-G1, N), N, G1, N))
  sphere (N, G1, N, R1, cY)
  if mit {
    parallelogram (cZ, -X,-Y, N), X,-Y, N), -X, Y, N))
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
        figure (POINTS, cZ, x, y, N))
      } else {
        octahedron (cZ, x, y, N, R)
      }
      y += 1
    }
    x += 1
  }
  figure (LINES, cZ, N, N,-G1), N, N, G1))
  sphere (N, N, G1, R1, cZ)
*/
}
