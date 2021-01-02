package ker

// (c) Christian Maurer   v. 201226 - license see µU.go

import
  "math"

func binom (n, k uint) uint {
  if k == 1 { return n }
  if n < k { return 0 }
  if n == k { return 1 }
  if n - k < k { k = n - k }
  b := uint(1)
  for i := uint(1); i <= k; i++ {
    b *= n
    b /= i
    n --
  }
  return b
}

func bezier (x, y []int, m, n, i uint) (int, int) {
  a, b := float64 (i) / float64 (n), float64 (n - i) / float64 (n)
  t, t1 := make ([]float64, m), make ([]float64, m)
  t[0], t1[0] = 1.0, 1.0
  for k:= uint(1); k < m; k++ {
    t[k], t1[k] = a * t[k - 1], b * t1[k - 1]
  }
  for k:= uint(0); k < m; k++ {
    w := float64(Binom (m - 1, k)) * t[k] * t1[m - 1 - k]
    a += w * float64 (x[k])
    b += w * float64 (y[k])
  }
  return int(a + 0.5), int(b + 0.5)
}

func arcLen (xs, ys []int) uint {
  var n, dx, dy float64
  for i := 1; i < len(xs); i++ {
    dx, dy = float64(xs[i]) - float64(xs[i-1]), float64(ys[i]) - float64(ys[i-1])
    n += math.Sqrt (float64 (dx * dx + dy * dy))
  }
  return uint (n + 0.5)
}
