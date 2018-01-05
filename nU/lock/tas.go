package lock

// (c) Christian Maurer   v. 171231 - license see nU.go

import "nU/atomic"

type tas struct {
  bool "locked"
}

func newTAS() Locker {
  return new(tas)
}

func (x *tas) Lock() {
  for atomic.TestAndSet (&x.bool) {
    nothing()
  }
}

func (x *tas) Unlock() {
  x.bool = false
}
