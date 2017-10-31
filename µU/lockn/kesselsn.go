package lockn

// (c) Christian Maurer   v. 171024 - license see µU.go

// >>> Tournament-Algorithm of Kessels due to Taubenfeld, p. 41

import (
  . "µU/obj"
)
type
  kesselsn struct {
                  uint "number of processes"
       interested [][2]bool
         favoured [][2]uint
                b []uint
                  }

func newKN (n uint) LockerN {
  x := new(kesselsn)
  x.uint = n
  x.b = make([]uint, n)
  x.interested = make([][2]bool, n)
  x.favoured = make([][2]uint, n)
  return x
}

func (x *kesselsn) Lock (i uint) {
  n := i + x.uint
  for n > 1 {
    j := n % 2
    n /= 2
    x.interested[n][j] = true
    local := (x.favoured[n][1 - j] + j) % 2
    x.favoured[n][j] = local
    for x.interested[n][1 - j] && local == (x.favoured[n][1 - j] + j) % 2 {
      Gothing()
    }
    x.b[n] = j
  }
}

func (x *kesselsn) Unlock (i uint) {
  n := uint(1)
  for n < x.uint {
    x.interested[n][x.b[n]] = false
    n = 2 * n + x.b[n]
  }
}
