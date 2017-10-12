package nchan

// (c) Christian Maurer   v. 170923 - license see µu.go

import
  . "µu/ker"

func port (n, i, j, a uint) uint16 {
  if a > 0 { Panic("a > 0") }
  const p0 = uint16(50000)
//  k := uint16(n * (n + 1)/ 2)
  if i > j { i, j = j, i } // i <= j
  return p0 + uint16(n * i - i * (i + 1) / 2 + j) // + uint16(a) * k
}

func nPorts (n, a uint) uint {
  return 1 * n * (n + 1) / 2
}
