package lock

// (c) Christian Maurer   v. 171021 - license see µU.go

import (
  . "µU/obj"
  . "µU/atomic"
)
type
  cas struct {
             uint32 "0 or 1, initially 0"
             }

func newCAS() Locker {
  return new(cas)
}

func (x *cas) Lock() {
  for ! CompareAndSwap (&x.uint32, 0, 1) {
    Gothing()
  }
}

func (x *cas) Unlock() {
  x.uint32 = 0
}
