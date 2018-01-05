package lock

// (c) Christian Maurer   v. 171231 - license see nU.go

import "nU/atomic"

type dec struct {
  int32
}

func newDEC() Locker {
  return &dec { int32: 1 }
}

func (x *dec) Lock() {
  for atomic.Decrement (&x.int32) {
    nothing()
  }
}

func (x *dec) Unlock() {
  x.int32 = 1
}
