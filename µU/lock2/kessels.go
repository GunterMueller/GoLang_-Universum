package lock2

// (c) Christian Maurer   v. 190815 - license see µU.go
//
// >>> Algorithm of Kessels

import (
  . "µU/obj"
  . "µU/atomic"
)
type
  kessels struct {
      interested,
            turn [2]uint // < 2
                 }

func newKessels() Locker2 {
  return new(kessels)
}

func (x *kessels) Lock (p uint) {
  Store (&x.interested[p], 1)
  Store (&x.turn[p], (p + x.turn[1-p]) % 2)
  for x.interested[1-p] == 1 && x.turn[p] == (p + x.turn[1-p]) % 2 {
    Nothing()
  }
}

func (x *kessels) Unlock (p uint) {
  Store (&x.interested[p], 0)
}
