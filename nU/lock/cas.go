package lock

// (c) Christian Maurer   v. 171013 - license see nU.go

import
  . "sync/atomic"
type
  cas struct {
             uint32 "0 or 1, initially 0"
             }

func newCAS() Locker {
  return new(cas)
}

func (x *cas) Lock() {
  for ! CompareAndSwapUint32 (&x.uint32, 0, 1) {
    null()
  }
}

func (x *cas) Unlock() {
  x.uint32 = 0
}
