package gl

// (c) Christian Maurer   v. 221113 - license see µU.go

import (
  "math"
  "µU/ker"
  "µU/obj"
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

func fail() {
  ker.Panic ("wrong number of float64's")
}

func point (x ...float64) {
  if len(x) != 3 { fail() }
  begin (Points)
  vertex (x[0], x[1], x[2])
  end()
}

func line (x ...float64) {
  if len(x) != 6 { fail() }
  for i := 0; i < 6; i++ { if math.IsNaN(x[i]) { return } }
  begin (Lines)
  vertex (x[0], x[1], x[2])
  vertex (x[3], x[4], x[5])
  end()
}

func lineloop (x ...float64) {
  n := len(x)
  if n % 3 != 0 || n < 3 * 3 { fail() }
  begin (LineLoop)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i+1], x[i+2])
  }
  end()
}

func linestrip (x ...float64) {
  n := len(x)
  if n % 3 != 0 || n < 3 * 3 { fail() }
  begin (LineStrip)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i+1], x[i+2])
  }
  end()
}

func triangle (x ...float64) {
  n := len(x)
  if n != 9 { fail() }
  begin (Triangles)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i+1], x[i+2])
  }
  end()
}

func trianglestrip (x ...float64) {
  n := len(x)
  if n % 3 != 0 || n < 12 { fail() }
  begin (TriangleStrip)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i+1], x[i+2])
  }
  end()
}

func trianglefan (x ...float64) {
  n := len(x)
  if n % 3 != 0 || n < 12 { fail() }
  begin (TriangleFan)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i+1], x[i+2])
  }
  end()
}

func quad (x ...float64) {
  if len(x) != 12 { fail() }
  begin (Quads)
  for i := 0; i < 12; i += 3 {
    vertex (x[i], x[i+1], x[i+2])
  }
  end()
}

func quadstrip (x ...float64) {
  n := len(x)
  if n % 3 != 0 || n < 6 * 3 { fail() }
  begin (QuadStrip)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i+1], x[i+2])
  }
  end()
}

func horRectangle (x, y, z, x1, y1 float64, up bool) {
  if up {
    begin (Quads)
    vertex (x,  y,  z)
    vertex (x1, y,  z)
    vertex (x1, y1, z)
    vertex (x,  y1, z)
    end()
  } else {
    begin (Quads)
    vertex (x,  y1, z)
    vertex (x1, y1, z)
    vertex (x1, y,  z)
    vertex (x,  y,  z)
    end()
  }
}

func vertRectangle (x ...float64) {
  if len(x) != 6 { fail() }
  if x[2] == x[5] { ker.PrePanic() }
//  if x[0] == x[3] && x[1] == x[4] { ker.PrePanic() }
  begin (Quads)
  vertex (x[0], x[1], x[2])
  vertex (x[3], x[4], x[2])
  vertex (x[3], x[4], x[5])
  vertex (x[0], x[1], x[5])
  end()
}

func parallelogram (x ...float64) {
  if len(x) != 9 { fail() }
  begin (Quads)
  vertex (x[0],               x[1],               x[2])
  vertex (x[0] + x[3],        x[1] + x[4],        x[2] + x[5])
  vertex (x[0] + x[3] + x[6], x[1] + x[4] + x[7], x[2] + x[5] + x[8])
  vertex (x[0] + x[6],        x[1] + x[7],        x[2] + x[8])
  end()
}

func polygon (x ...float64) {
  n := len(x)
  if n % 3 != 0 || n < 3 * 3 { fail() }
  begin (Polygons)
  for i := 0; i < n; i += 3 {
    vertex (x[i], x[i+1], x[i+2])
  }
  end()
}

func curve (f obj.Ft2xyz, t0, t1 float64) {
  dt := (t1 - t0) / 100
  begin (Lines)
  for t := t0; t <= t1; t += dt {
    x, y, z := f (t)
    x1, y1, z1 := f (t + dt)
    line (x, y, z, x1, y1, z1)
  }
  end()
}

func plane (a, b, c, wx, wy float64) {
  if wx <= 0 || wy <= 0 { ker.PrePanic() }
  surface (func (x, y float64) float64 { return a * x + b * y + c }, wx, wy)
}

func cube (x, y, z, a float64) {
  if a == 0 { ker.PrePanic() }
  Cuboid (x - a / 2, y - a / 2, z - a / 2, x + a / 2, y + a / 2, z + a / 2)
}

func cubeC (c []col.Colour, x, y, z, a float64) {
  if a == 0 { ker.PrePanic() }
  CuboidC (c, x - a / 2, y - a / 2, z - a / 2, x + a / 2, y + a / 2, z + a / 2)
}

func cuboid (x ...float64) {
  if len(x) != 6 { fail() }
  if x[0] == x[3] || x[1] == x[4] || x[2] == x[5] { ker.PrePanic() }
  VertRectangle (x[0], x[1], x[2], x[3], x[1], x[5]) // front
  VertRectangle (x[3], x[1], x[2], x[3], x[4], x[5]) // right
  VertRectangle (x[3], x[4], x[2], x[0], x[4], x[5]) // back
  VertRectangle (x[0], x[4], x[2], x[0], x[1], x[5]) // left
  horRectangle (x[0], x[1], x[5], x[3], x[4], true)  // top
  horRectangle (x[0], x[1], x[2], x[3], x[4], false) // bottom
}

func cuboidC (c []col.Colour, x ...float64) {
  if len(x) != 6 || len(c) != 6 { fail() }
  if x[0] == x[3] || x[1] == x[4] || x[2] == x[5] { ker.PrePanic() }
  colour (c[0])
  VertRectangle (x[0], x[1], x[2], x[3], x[1], x[5])
  colour (c[1])
  VertRectangle (x[3], x[1], x[2], x[3], x[4], x[5])
  colour (c[2])
  VertRectangle (x[3], x[4], x[2], x[0], x[4], x[5])
  colour (c[3])
  VertRectangle (x[0], x[4], x[2], x[0], x[1], x[5])
  colour (c[4])
  horRectangle (x[0], x[1], x[5], x[3], x[4], true)
  colour (c[5])
  horRectangle (x[0], x[1], x[2], x[3], x[4], false)
}

func cuboid1 (x, y, z, dx, dy, dz, a float64) {
  s, c := math.Sin (a * pi_180), math.Cos (a * pi_180)
  x0, y0 := x  + dx * c, y  + dx * s
  x1, y1 := x0 - dy * s, y0 + dy * c
  x2, y2 := x  - dy * s, y  + dy * c
  z1 := z + dz
  VertRectangle (x,  y,  z, x0, y0, z1)
  VertRectangle (x0, y0, z, x1, y1, z1)
  VertRectangle (x1, y1, z, x2, y2, z1)
  VertRectangle (x2, y2, z, x,  y,  z1)
  begin (Quads)
  vertex (x,  y,  z1)
  vertex (x0, y0, z1)
  vertex (x1, y1, z1)
  vertex (x2, y2, z1)
  end()
}

func prism (x ...float64) {
  n := len(x)
  if n % 3 != 0 || n < 12 { fail() }
  y := make([]float64, n)
  for i := 0; i < n; i++ {
    y[i] = x[i] + x[i%3]
  }
  begin (Quads)
  for i := 3; i < n - 3; i += 3 {
    vertex (x[i],   x[i+1], x[i+2])
    vertex (x[i+3], x[i+4], x[i+5])
    vertex (y[i+3], y[i+4], y[i+5])
    vertex (y[i],   y[i+1], y[i+2])
  }
  vertex (x[n-3], x[n-2], x[n-1])
  vertex (x[3],   x[4],   x[5])
  vertex (y[3],   y[4],   y[5])
  vertex (y[n-3], y[n-2], y[n-1])
  end()
}

func prismC (c []col.Colour, x ...float64) {
  n := len(x)
  if n % 3 != 0 || n < 12 { fail() }
  y := make([]float64, n)
  for i := 0; i < n; i++ {
    y[i] = x[i] + x[i%3]
  }
  begin (Quads)
  for i := 3; i < n - 3; i += 3 {
    colour (c[i/3-1])
    vertex (x[i],   x[i+1], x[i+2])
    vertex (x[i+3], x[i+4], x[i+5])
    vertex (y[i+3], y[i+4], y[i+5])
    vertex (y[i],   y[i+1], y[i+2])
  }
  colour (c[n/3-2])
  vertex (x[n-3], x[n-2], x[n-1])
  vertex (x[3],   x[4],   x[5])
  vertex (y[3],   y[4],   y[5])
  vertex (y[n-3], y[n-2], y[n-1])
  end()
}

func parallelepiped (x ...float64) {
  if len(x) != 12 { fail() }
  colour (col.Red())
  quad (x[3], x[4], x[5], x[3] + x[9], x[4] + x[10], x[5] + x[11],
        x[3] + x[6] + x[9] - x[0], x[4] + x[7] + x[10] - x[1],
        x[5] + x[8] + x[11] - x[2], x[3] + x[6], x[4] + x[7], x[5] + x[8])
  prism (x[3] - x[0], x[4] - x[1], x[5] - x[2], x[0], x[1], x[2],
         x[9], x[10], x[11], x[6] + x[9] - x[0], x[7] + x[10] - x[1],
         x[8] + x[11] - x[2], x[6], x[7], x[8])
  quad (x[0], x[1], x[2], x[6], x[7], x[8], x[6] + x[9] - x[0],
        x[7] + x[10] - x[1], x[8] + x[11] - x[2], x[9], x[10], x[11])
}

func parallelepipedC (c []col.Colour, x ...float64) {
  if len(x) != 12 || len(c) != 6 { fail() }
  colour (c[1])
  quad (x[3], x[4], x[5],
        x[3] + x[9], x[4] + x[10], x[5] + x[11],
        x[3] + x[6] + x[9] - x[0], x[4] + x[7] + x[10] - x[1],
        x[5] + x[8] + x[11] - x[2], x[3] + x[6], x[4] + x[7], x[5] + x[8])
  prismC ([]col.Colour { c[0], c[4], c[2], c[5] },
          x[3] - x[0], x[4] - x[1], x[5] - x[2], x[0], x[1], x[2],
          x[9], x[10], x[11], x[6] + x[9] - x[0], x[7] + x[10] - x[1],
          x[8] + x[11] - x[2], x[6], x[7], x[8])
  colour (c[3])
  quad (x[0], x[1], x[2], x[6], x[7], x[8], x[6] + x[9] - x[0],
        x[7] + x[10] - x[1], x[8] + x[11] - x[2], x[9], x[10], x[11])
}

func pyramid (x, y, z, a, h float64) {
  if h == 0 || a <= 0 { ker.PrePanic() }
  begin (TriangleFan)
  vertex (x,     y,     z + h)
  vertex (x + a, y + a, z)
  vertex (x - a, y + a, z)
  vertex (x - a, y - a, z)
  vertex (x + a, y - a, z)
  end()
  horRectangle (x - a, y - a, z, x + a, y + a, false)
}

func pyramidC (c []col.Colour, x, y, z, a, h float64) {
  if h == 0 || a <= 0 || len(c) != 5 { ker.PrePanic() }
  begin (TriangleFan)
  colour (c[0])
  vertex (x,     y,     z + h)
  vertex (x - a, y - a, z)
  vertex (x + a, y - a, z)
  colour (c[1])
  vertex (x + a, y + a, z)
  colour (c[2])
  vertex (x - a, y + a, z)
  colour (c[3])
  vertex (x - a, y - a, z)
  end()
  colour (c[4])
  horRectangle (x - a, y - a, z, x + a, y + a, false)
}

func multipyramid (x, y, z, h float64, c...float64) {
  n := len(c)
  if n % 2 != 0 || n < 6 { fail() }
  begin (TriangleFan)
  vertex (x, y, z + h)
  for i := 0; i < n; i += 2 {
    vertex (c[i], c[i+1], 0)
  }
  vertex (c[0], c[1], 0)
  end()
  begin (TriangleFan)
  vertex (x, y, 0)
  for i := 0; i < n; i += 2 {
    vertex (c[i], c[i+1], 0)
  }
  vertex (c[0], c[1], 0)
  end()
}

func multipyramidC (f []col.Colour, x, y, z, h float64, c ...float64) {
  n := len(c)
  if n % 2 != 0 || n < 6 || len(f) != n / 2 + 1 { fail() }
  begin (TriangleFan)
  vertex (x, y, z + h)
  for i := 0; i < n; i += 2 {
    vertex (c[i], c[i+1], 0)
    colour (f[i/2])
  }
  vertex (c[0], c[1], 0)
  end()
  begin (TriangleFan)
  colour (f[n/2])
  vertex (x, y, 0)
  for i := 0; i < n; i += 2 {
    vertex (c[i], c[i+1], 0)
  }
  vertex (c[0], c[1], 0)
  end()
}

func octopus (x ...float64) {
  n := len(x)
  if n % 3 != 0 || n < 12 { fail() }
  begin (TriangleFan)
  vertex (x[0], x[1], x[2])
  for i := 1; i < n / 3; i++ {
    j := 3 * i
    vertex (x[j], x[j+1], x[j+2])
  }
  vertex (x[3], x[4], x[5])
  end()
  begin (TriangleFan)
  colour (col.Red())
  vertex (x[0], x[1], 0)
  for i := 1; i < n / 3; i++ {
    j := 3 * i
    vertex (x[j], x[j+1], x[j+2])
  }
  vertex (x[3], x[4], x[5])
  end()
}

func octopusC (c []col.Colour, x ...float64) {
  n := len(x)
  if n % 3 != 0 || n < 12 || len(c) != n / 3 + 1 { fail() }
  for i := 1; i < n / 3 - 1; i++ {
    j, k := 3 * i, 3 * (i + 1)
    colour (c[i-1])
    Triangle (x[0], x[1],   x[2],
              x[j], x[j+1], x[j+2],
              x[k], x[k+1], x[k+2])
  }
  colour (c[n/3-2])
  Triangle (x[0],   x[1],   x[2],
            x[n-3], x[n-2], x[n-1],
            x[3],   x[4],   x[5])
  begin(TriangleFan)
  colour (c[n/3-1])
  vertex (x[0], x[1], 0)
  for i := 1; i < n / 3; i++ {
    j := 3 * i
    vertex (x[j], x[j+1], x[j+2])
  }
  vertex (x[3], x[4], x[5])
  end()
}

func octahedron (x, y, z, e float64) {
  if e == 0 { ker.PrePanic() }
  d := e * math.Sqrt (2)
  begin (TriangleFan)
  vertex (x,     y,     z + d)
  vertex (x + e, y + e, z)
  vertex (x - e, y + e, z)
  vertex (x - e, y - e, z)
  vertex (x + e, y - e, z)
  vertex (x + e, y + e, z)
  end()
  begin (TriangleFan)
  vertex (x,     y,     z - d)
  vertex (x + e, y - e, z)
  vertex (x - e, y - e, z)
  vertex (x - e, y + e, z)
  vertex (x + e, y + e, z)
  vertex (x + e, y - e, z)
  end()
}

func octahedronC (c []col.Colour, x, y, z, e float64) {
  if e == 0 || len(c) != 8 { ker.PrePanic() }
  d := e * math.Sqrt (2)
  begin (TriangleFan)
  colour (c[0])
  vertex (x,     y,     z + d)
  vertex (x - e, y - e, z)
  vertex (x + e, y - e, z)
  colour (c[1])
  vertex (x + e, y + e, z)
  colour (c[2])
  vertex (x - e, y + e, z)
  colour (c[3])
  vertex (x - e, y - e, z)
  end()
  begin (TriangleFan)
  colour (c[4])
  vertex (x,     y,     z - d)
  vertex (x - e, y - e, z)
  vertex (x + e, y - e, z)
  colour (c[5])
  vertex (x + e, y + e, z)
  colour (c[6])
  vertex (x - e, y + e, z)
  colour (c[7])
  vertex (x - e, y - e, z)
  end()
}

func circle (x, y, z, r float64) {
//  circleSegment (x, y, z, r, 0, 360)
  if r == 0 { ker.PrePanic() }
//  if r < 0 { r = -r } // change orientation // TODO
  begin (TriangleFan)
  vertex (x, y, z)
  for i := 0; i <= N; i++ {
    vertex (x + r * cos[i], y + r * sin[i], z)
  }
  end()
}

func circleSegment (x, y, z, r, a, b float64) {
  if r == 0 { ker.PrePanic() }
//  if r < 0 { r = -r } // change orientation // TODO
  aa := uint(math.Floor (a / float64 (angle) + 0.5))
  bb := uint(math.Floor (b / float64 (angle) + 0.5))
  begin (TriangleFan)
  vertex (x, y, z)
  for i := aa; i <= bb; i++ {
    vertex (x + r * cos[i], y + r * sin[i], z)
  }
  end()
}

func vertCircle (x, y, z, r, a float64) {
  if r == 0 { ker.PrePanic() }
//  if r < 0 { r = -r } // change orientation // TODO
  s, c := math.Sin (a * pi_180), math.Cos (a * pi_180)
  begin (TriangleFan)
  vertex (x, y, z)
  for i := 0; i <= N; i++ {
    vertex (x - r * s * cos[i], y + r * c * cos[i], z + r * sin[i])
  }
  end()
}

func sphere (x, y, z, r float64) {
  if r == 0 { ker.PrePanic() }
  r0, z0 := r * sin[1], z + r * cos[1]
  begin (TriangleFan)
  vertex (x, y, z + r)
  for i := 0; i <= N; i++ {
    vertex (x + r0 * cos[i], y + r0 * sin[i], z0)
  }
  end()
  begin (QuadStrip)
  for i := 1; i <= N / 2 - 2; i++ {
    r0, z0 =  r * sin[i],   z + r * cos[i]
    r1, z1 := r * sin[i+1], z + r * cos[i+1]
    for j := 0; j <= N; j++ {
      s, c := sin[j], cos[j]
      vertex (x + r0 * c, y + r0 * s, z0)
      vertex (x + r1 * c, y + r1 * s, z1)
    }
  }
  end()
  r0, z0 = r * sin[N/2-1], z + r * cos[N/2-1]
  begin (TriangleFan)
  vertex (x, y, z - r)
  for i := N; i >= 0; i -= 1 {
    vertex (x + r0 * cos[i], y + r0 * sin[i], z0)
  }
  end()
}

func cone (x, y, z, r, h float64) {
  if r == 0 || h == 0 { ker.PrePanic() }
  begin (TriangleFan)
  vertex (x, y, z + h)
  for l := 0; l <= N; l++ {
    vertex (x + r * cos[l], y + r * sin[l], z)
  }
  end()
  circle (x, y, z, -r)
}

func doubleCone (x, y, z, r, h float64) {
  if r == 0 || h == 0 { ker.PrePanic() }
  cone (x, y, z - h, r,  h)
  cone (x, y, z + h, r, -h)
}

func cylinder (x, y, z, r, h float64) {
  if r == 0 || h == 0 { ker.PrePanic() }
  begin (QuadStrip)
  for i := 0; i <= N; i++ {
    vertex (x + r * cos[i], y + r * sin[i], z)
    vertex (x + r * cos[i], y + r * sin[i], z + h)
  }
  end()
  circle (x, y, z,    -r)
  circle (x, y, z + h, r)
}

func cylinderC (c []col.Colour, x, y, z, r, h float64) {
  if r == 0 || h == 0 || len(c) != 2 { ker.PrePanic() }
  begin (QuadStrip)
  colour (c[0])
  for i := 0; i <= N; i++ {
    vertex (x + r * cos[i], y + r * sin[i], z)
    vertex (x + r * cos[i], y + r * sin[i], z + h)
  }
  end()
  colour (c[1])
  circle (x, y, z,    -r)
  circle (x, y, z + h, r)
}

func cylinderSegment (x, y, z, r, h, a, b float64) {
  if r == 0 || h == 0 { ker.PrePanic() }
  circleSegment (x, y, z,    -r, a, b)
  circleSegment (x, y, z + h, r, a, b)
  aa := uint(math.Floor (a / float64 (angle) + 0.5))
  bb := uint(math.Floor (b / float64 (angle) + 0.5))
  begin (QuadStrip)
  for i := aa; i <= bb; i++ {
    vertex (x + r * cos[i], y + r * sin[i], z)
    vertex (x + r * cos[i], y + r * sin[i], z + h)
  }
  end()
}

func horCylinder (x, y, z, r, l, a float64) {
  if r == 0 { ker.PrePanic() }
  sa, ca := math.Sin (a * pi_180), math.Cos (a * pi_180)
  dx, dy := l * ca, l * sa
  begin (QuadStrip)
  for i := 0; i <= 2 * N; i += 2 {
    si, ci := sin[i / 2], cos[i / 2]
    x0, y0, z0 := x - r * sa * ci, y + r * ca * ci, z + r * si
    vertex (x0, y0, z0)
    vertex (x0 + dx, y0 + dy, z0)
  }
  end()
  vertCircle (x, y, z, -r, a)
  vertCircle (x + dx, y + dy, z, r, a)
}

func horCylinderC (c []col.Colour, x, y, z, r, l, a float64) {
  if r == 0 || len(c) != 2 { ker.PrePanic() }
  sa, ca := math.Sin (a * pi_180), math.Cos (a * pi_180)
  dx, dy := l * ca, l * sa
  colour (c[0])
  begin (QuadStrip)
  for i := 0; i <= 2 * N; i += 2 {
    si, ci := sin[i / 2], cos[i / 2]
    x0, y0, z0 := x - r * sa * ci, y + r * ca * ci, z + r * si
    vertex (x0, y0, z0)
    vertex (x0 + dx, y0 + dy, z0)
  }
  end()
  colour (c[1])
  vertCircle (x, y, z, -r, a)
  vertCircle (x + dx, y + dy, z, r, a)
}

func torus (x, y, z, R, r float64) {
  if r <= 0 || R <= 0 { fail() }
  begin (QuadStrip)
  for i := 0; i < N; i++ {
    s0, s1 := R + r * cos[i], R + r * cos[i+1]
    z0, z1 := z + r * sin[i], z + r * sin[i+1]
    for j := 0; j < N; j++ {
      vertex (x + s0 * cos[j],   y + s0 * sin[j],   z0)
      vertex (x + s0 * cos[j+1], y + s0 * sin[j+1], z0)
      vertex (x + s1 * cos[j+1], y + s1 * sin[j+1], z1)
      vertex (x + s1 * cos[j],   y + s1 * sin[j],   z1)
      vertex (x + s0 * cos[j],   y + s0 * sin[j],   z0)
      vertex (x + s0 * cos[j+1], y + s0 * sin[j+1], z0)
    }
  }
  end()
}

func verTorus (x, y, z, R, r, a float64) {
  if r <= 0 || R <= 0 { ker.PrePanic() }
  for a <= -180 { a += 180 }
  for a >=  180 { a -= 180 }
  sa, ca := math.Sin (a * pi_180), math.Cos (a * pi_180)
  begin (QuadStrip)
  for i := 0; i < N; i++ {
    s0, s1 := R + r * cos[i], R + r * cos[i+1]
    x0, x1 := r * sin[i], r * sin[i+1]
    for j := 0; j < N; j++ { //  x -> x * c - y * sa, y -> x * sa + y * ca
      y00, y01 := s0 * cos[j], s0 * cos[j+1]
      y10, y11 := s1 * cos[j], s1 * cos[j+1]
      vertex (x + x0 * ca - y00 * sa, y + x0 * sa + y00 * ca, z + s0 * sin[j])
      vertex (x + x0 * ca - y01 * sa, y + x0 * sa + y01 * ca, z + s0 * sin[j+1])
      vertex (x + x1 * ca - y11 * sa, y + x1 * sa + y11 * ca, z + s1 * sin[j+1])
      vertex (x + x1 * ca - y10 * sa, y + x1 * sa + y10 * ca, z + s1 * sin[j])
    }
  }
  end()
}

func paraboloid (x0, y0, z0, a, wx, wy float64) {
  if a == 0 || wx <= 0 || wy <= 0 { ker.PrePanic() }
  surface (func (x, y float64) float64 {
             dx, dy := x - x0, y - y0
             return a * a * (dx * dx + dy * dy) + z0
           }, wx, wy)
}

func surface (f obj.Fxy2z, wx, wy float64) {
  if wx <= 0 || wy <= 0 { ker.PrePanic() }
  dx, dy := wx / 100, wy / 100
  for x := -wx; x < wx; x += dx {
    for y := -wy; y < wy; y += dy {
      x1, y1 := x + dx, y + dy
      z0, z1, z2, z3 := f(x, y), f(x1, y), f(x1, y1), f(x, y1)
      triangle (x, y, z0, x1, y1, z2, x, y1, z3)
      triangle (x, y, z0, x1, y, z1, x1, y1, z2)
    }
  }
}
