package lock

// (c) Christian Maurer   v. 230207 - license see µU.go

import (
  . "sync"
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
