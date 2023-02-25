package lock

// (c) Christian Maurer   v. 230207 - license see µU.go

import (
  . "sync"
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
