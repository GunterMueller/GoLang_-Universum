package lock

// (c) Christian Maurer   v. 171024 - license see µU.go

import (
  . "µU/obj"
  . "µU/atomic"
)
type
  xchg struct {
              uint32 "0 or 1; initially 0"
              }

func newXCHG() Locker {
  return new(xchg)
}

func (x *xchg) Lock() {
  local := uint32(1)
  for Exchange (&x.uint32, local) == 1 {
    Gothing()
  }
}

func (x *xchg) Unlock() {
  x.uint32 = 0
}
