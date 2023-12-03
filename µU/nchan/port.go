package nchan

// (c) Christian Maurer   v. 231108 - license see µU.go

import
  . "µU/ker"

func port (n, i, j, a uint) uint16 {
  if a > 1 { Panic("nchan.Port: a > 1") }
  k := uint16(0)
  if a == 1 {
    k = uint16(n * (n + 1) / 2)
  }
  if i > j { i, j = j, i } // i <= j
  return uint16(n * i - i * (i + 1) / 2 + j) + k
}

func port1 (n, i, j, a uint) uint16 {
  return port (n, i, j, a) + uint16(n * (n + 1)) / 2 - 2
}

func nPorts (n, a uint) uint {
  return (1 + a) * n * (n + 1) / 2
}

