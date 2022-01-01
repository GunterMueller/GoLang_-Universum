package ker

// (c) Christian Maurer   v. 211218 - license see µU.go

func bezier (x, y []int, m, n, i uint) (int, int) {
  a, b := float64 (i) / float64 (n), float64 (n - i) / float64 (n)
  t, t1 := make ([]float64, m), make ([]float64, m)
  t[0], t1[0] = 1.0, 1.0
  for j := uint(1); j < m; j++ {
    t[j], t1[j] = a * t[j - 1], b * t1[j - 1]
  }
  for j := uint(0); j < m; j++ {
    w := float64(Binom (m - 1, j)) * t[j] * t1[m - 1 - j]
    a += w * float64 (x[j])
    b += w * float64 (y[j])
  }
  return int(a + 0.5), int(b + 0.5)
}
