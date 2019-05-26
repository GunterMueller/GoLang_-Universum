package lockn

// (c) Christian Maurer   v. 190331 - license see µU.go

// >>> Algorithm of Kessels

import (. "nU/obj"; . "nU/atomic")

type kessels struct {
  uint "number of processes"
  interested, favoured [][2]uint
  e []uint // < 2
}

func newKessels (n uint) LockerN {
  x := new(kessels)
  x.uint = uint(n)
  x.interested = make([][2]uint, n)
  x.favoured = make([][2]uint, n)
  x.e = make([]uint, n)
  return x
}

func (x *kessels) Lock (p uint) {
  for n := x.uint + p; n > 1; n /= 2 {
    k, m := n / 2, n % 2
    Store (&x.interested[k][m], 1)
    m1 := 1 - m
    Store (&x.favoured[k][m], (x.favoured[k][m1] + m) % 2)
    for x.interested[k][m1] == 1 && x.favoured[k][m] == (x.favoured[k][m1] + m) % 2 {
      Nothing()
    }
    x.e[k] = m
  }
}

func (x *kessels) Unlock (p uint) {
  n := uint(1)
  for n < x.uint {
    n = 2 * n + x.e[n]
    Store (&x.interested[n/2][n%2], 0)
  }
}
