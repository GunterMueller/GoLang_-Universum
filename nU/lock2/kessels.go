package lock2

// (c) Christian Maurer   v. 190325 - license see nU.go
//
// >>> Algorithm of Kessels

import (
  . "nU/atomic"
  . "nU/obj"
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
