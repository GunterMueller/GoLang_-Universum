package lock

// (c) Christian Maurer   v. 171231 - license see nU.go

import "nU/atomic"

type fa struct {
  uint32
}

func newFA() Locker {
  return new (fa)
}

func (x *fa) Lock() {
  for atomic.FetchAndAdd (&x.uint32, 1) != 0 {
    nothing()
  }
}

func (x *fa) Unlock() {
  x.uint32 = 0
}
