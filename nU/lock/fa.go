package lock

// (c) Christian Maurer   v. 190312 - license see nU.go

import ("nU/atomic"; . "nU/obj")

type fa struct {
  uint
}

func newFA() Locker {
  return new (fa)
}

func (x *fa) Lock() {
  for atomic.FetchAndAdd (&x.uint, 1) != 0 {
    Nothing()
  }
}

func (x *fa) Unlock() {
  x.uint = 0
}
