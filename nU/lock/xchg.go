package lock

// (c) Christian Maurer   v. 190323 - license see nU.go

import (
  . "nU/obj"
  . "nU/atomic"
)
type
  xchg struct {
              uint "0 or 1; initially 0"
              }

func newXCHG() Locker {
  return new(xchg)
}

func (x *xchg) Lock() {
  local := uint(1)
  for Exchange (&x.uint, local) == 1 {
    Nothing()
  }
}

func (x *xchg) Unlock() {
  x.uint = 0
}
