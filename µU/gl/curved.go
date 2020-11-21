package gl

// (c) Christian Maurer   v. 201102 - license see µU.go

import
  "math"

func circle (x, y, z, r float64) {
//  circleSegment (x, y, z, r, 0, 360)
  if r == 0 {
    Point (x, y, z)
    return
  }
//  if r < 0 { r = -r } // change orientation
  begin (TriangleFan)
  vertex (x, y, z)
  for i := 0; i <= N; i++ {
    vertex (x + r * cos[i], y + r * sin[i], z)
  }
  end()
}

func circleSegment (x, y, z, r, a, b float64) {
  if r == 0 {
    Point (x, y, z)
    return
  }
//  if r < 0 { r = -r } // change orientation
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
  if r == 0 {
    Point (x, y, z)
    return
  }
//  if r < 0 { r = -r } // change orientation
  s, c := math.Sin (a * pi_180), math.Cos (a * pi_180)
  begin (TriangleFan)
  vertex (x, y, z)
  for i := 0; i <= N; i++ {
    vertex (x - r * s * cos[i], y + r * c * cos[i], z + r * sin[i])
  }
  end()
}

func sphere (x, y, z, r float64) {
  r0, z0 := r * sin[1], z + r * cos[1]
  begin (TriangleFan)
  vertex (x, y, z + r)
  for i := 0; i <= N; i++ {
    vertex (x + r0 * cos[i], y + r0 * sin[i], z0)
  }
  end()
  begin (QuadStrip)
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
  begin (TriangleFan)
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

func cone (x, y, z, r, h float64) {
  begin (TriangleFan)
  vertex (x, y, z + h)
  for l := 0; l <= N; l++ {
    vertex (x + r * cos[l], y + r * sin[l], z)
  }
  end()
  circle (/* c, */ x, y, z, -r)
}

func doubleCone (x, y, z, r, h float64) {
  cone (x, y, z - h, r,  h)
  cone (x, y, z + h, r, -h)
}

func cylinder (x, y, z, r, h float64) {
//  cylinderSegment (c, x, y, z, r, h, 0, 360)
  circle (/* c, */ x, y, z,    -r)
  circle (/* c, */ x, y, z + h, r)
  begin (QuadStrip)
  for i := 0; i <= N; i++ {
    vertex (x + r * cos[i], y + r * sin[i], z)
    vertex (x + r * cos[i], y + r * sin[i], z + h)
  }
  end()
}

func cylinderSegment (x, y, z, r, h, a, b float64) {
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
  if r == 0 {
    vertCircle (x, y, z, r, a)
    return
  }
  sa, ca := math.Sin (a * pi_180), math.Cos (a * pi_180)
  dx, dy := l * ca, l * sa
  vertCircle (x, y, z, -r, a)
  vertCircle (x + dx, y + dy, z, r, a)
  begin (QuadStrip)
  for i := 0; i <= 2 * N; i += 2 {
    si, ci := sin[i / 2], cos[i / 2]
    sci, cci := sa * ci, ca * ci
    x0, y0, z0 := x - r * sci, y + r * cci, z + r * si
    vertex (x0, y0, z0)
    vertex (x0 + dx, y0 + dy, z0)
  }
  end()
}

func torus (x, y, z, R, r float64) {
  if r <= 0 || R <= 0 { fail() }
  begin (QuadStrip)
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

func verTorus (x, y, z, R, r, a float64) {
  if r <= 0 || R <= 0 { fail() }
  for a <= -180 { a += 180 }
  for a >=  180 { a -= 180 }
  sa, ca := math.Sin (a * pi_180), math.Cos (a * pi_180)
  begin (QuadStrip)
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

func surface (f func (float64, float64) float64,
              X, Y, Z, X1, Y1, Z1 float64, wd, ht, g uint) {
  if X == X1 || Y == Y1 || Z == Z1 { fail() }
  if X1 < X { X, X1 = X1, X }
  if Y1 < Y { Y, Y1 = Y1, Y }
  if Z1 < Z { Z, Z1 = Z1, Z }
  if g < 4 { g = 4 }
  dx, dy := (X1 - X) / float64(wd / g), (Y1 - Y) / float64(ht / g)
  for y := Y; y < Y1; y += dy {
    for x := X; x < X1; x += dx {
      x1, y1 := x + dx, y + dy
      z0, z1, z2, z3 := f(x, y), f(x1, y), f(x1, y1), f(x, y1)
      if ok(z0) && ok(z1) && ok(z2) && ok (z3) {
        triangle (x, y, z0, x1, y1, z2, x, y1, z3)
        triangle (x, y, z0, x1, y, z1, x1, y1, z2)
      }
    }
  }
}
