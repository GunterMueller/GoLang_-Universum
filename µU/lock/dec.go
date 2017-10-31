package lock

// (c) Christian Maurer   v. 171024 - license see µU.go

import (
  . "µU/obj"
  . "µU/atomic"
)
type
  dec struct {
             int32
             }

func newDEC() Locker {
  return &dec { int32: 1 }
}

func (x *dec) Lock() {
  for Decrement (&x.int32) {
    Gothing()
  }
}

func (x *dec) Unlock() {
  x.int32 = 1
}
