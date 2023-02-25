package lock

// (c) Christian Maurer   v. 230207 - license see µU.go

// >>> Algorithm of Morris: A starvation-free Solution to the Mutual Exclusion Problem
//     Inf. Proc. Letters 8 (1979), 76-80

import (
  . "sync"
  . "µU/atomic"
)
type
  morris struct {
    door0, door,
          mutex Mutex // to protect n0
             n0,     // number of processes blocked on door0
              n uint // number of processes blocked on door
                }

func newMorris() Locker {
  x := new (morris)
  x.door.Lock()
  return x
}

func (x *morris) Lock() {
  x.mutex.Lock()
  Store (&x.n0, x.n0 + 1)
  x.mutex.Unlock()
  x.door0.Lock()
  Store (&x.n, x.n + 1)
  x.mutex.Lock()
  Store (&x.n0, x.n0 - 1)
  if x.n0 > 0 {
    x.mutex.Unlock()
    x.door0.Unlock()
  } else { // n0 == 0
    x.mutex.Unlock()
    x.door.Unlock()
  }
  x.door.Lock()
  Store (&x.n, x.n - 1)
}

func (x *morris) Unlock() {
  if x.n > 0 {
    x.door.Unlock()
  } else { // x.n == 0
    x.door0.Unlock()
  }
}
