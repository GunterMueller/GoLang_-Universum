package lock

// (c) Christian Maurer   v. 190811 - license see µU.go

import (
  "µU/obj"
  . "µU/atomic"
)
type
  dec struct {
             int
             }

func newDEC() Locker {
  return &dec { int: 1 }
}

func (x *dec) Lock() {
  for Decrement (&x.int) {
    obj.Nothing()
  }
}

func (x *dec) Unlock() {
  x.int = 1
}
