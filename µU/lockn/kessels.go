package lockn

// (c) Christian Maurer   v. 190815 - license see µU.go

// >>> Algorithm of Kessels

import (
  . "µU/obj"
  . "µU/atomic"
)
type
  kessels struct {
                 uint "number of processes"
      interested,
        favoured [][2]uint // < 2
            edge []uint // < 2
                 }

func newKessels (n uint) LockerN {
  x := new(kessels)
  x.uint = n
  x.interested = make([][2]uint, n)
  x.favoured = make([][2]uint, n)
  x.edge = make([]uint, n)
  return x
}

func (x *kessels) Lock (p uint) {
  for n := x.uint + p; n > 1; n /= 2 {
    Store (&x.interested[n/2][n%2], 1)
    k, m := n / 2, n % 2
    Store (&x.favoured[k][m], (x.favoured[k][1-m] + m) % 2)
    for x.interested[k][1-m] == 1 && (x.favoured[k][1-m] + m) % 2 == x.favoured[k][m] {
      Nothing()
    }
    Store (&x.edge[k], m)
  }
}

func (x *kessels) Unlock (p uint) {
  n := uint(1)
  for n < x.uint {
    n = 2 * n + x.edge[n]
    Store (&x.interested[n/2][n%2], 0)
  }
}
