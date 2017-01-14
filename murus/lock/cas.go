package lock

// (c) murus.org  v. 161216 - license see murus.go

import
  . "sync/atomic"
type
  cas struct {
           n uint32
             }

func newCAS() Locker {
  return new(cas)
}

func (L *cas) Lock() {
  for ! CompareAndSwapUint32(&L.n, 0, 1) {
    null()
  }
}

func (L *cas) Unlock() {
  L.n = 0
}
