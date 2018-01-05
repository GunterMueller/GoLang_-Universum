package lock

// (c) Christian Maurer   v. 171231 - license see nU.go

import "nU/atomic"

type xchg struct {
  uint32 "0 or 1; initially 1"
}

func newXCHG() Locker {
  return &xchg { 1 }
}

func (x *xchg) Lock() {
  local := uint32(0)
  for atomic.Exchange (&x.uint32, local) == 0 {
    nothing()
  }
}

func (x *xchg) Unlock() {
  x.uint32 = 1
}
