package lock

// (c) Christian Maurer   v. 190323 - license see nU.go

import (. "nU/obj"; . "nU/atomic")

type dec struct {
  int
}

func newDEC() Locker {
  return &dec { int: 1 }
}

func (x *dec) Lock() {
  for Decrement (&x.int) {
    Nothing()
  }
}

func (x *dec) Unlock() {
  x.int = 1
}
