package lock

// (c) Christian Maurer   v. 190312 - license see nU.go

import (. "nU/obj"; . "nU/atomic")

type tas struct {
  bool "true, iff locked"
}

func newTAS() Locker {
  return new(tas)
}

func (x *tas) Lock() {
  for TestAndSet (&x.bool) {
    Nothing()
  }
}

func (x *tas) Unlock() {
  x.bool = false
}
