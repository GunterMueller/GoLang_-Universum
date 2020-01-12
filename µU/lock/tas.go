package lock

// (c) Christian Maurer   v. 190312 - license see µU.go

import (
  "µU/obj"
  . "µU/atomic"
)
type
  tas struct {
             bool "true, iff locked"
             }

func newTAS() Locker {
  return new(tas)
}

func (x *tas) Lock() {
  for TestAndSet (&x.bool) {
    obj.Nothing()
  }
}

func (x *tas) Unlock() {
  x.bool = false
}
