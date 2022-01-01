package ker

// (c) Christian Maurer   v. 211218 - license see µU.go

import
  "math"

func arcLen (xs, ys []int) uint {
  var n, dx, dy float64
  for i := 1; i < len(xs); i++ {
    dx, dy = float64(xs[i]) - float64(xs[i-1]), float64(ys[i]) - float64(ys[i-1])
    n += math.Sqrt (float64 (dx * dx + dy * dy))
  }
  return uint (n + 0.5)
}
