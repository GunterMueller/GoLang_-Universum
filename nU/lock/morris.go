package lock

// (c) Christian Maurer   v. 171231 - license see nU.go

import "sync"

type morris struct {
  door0, door, mutex sync.Mutex
  n0, n uint
}

func newMorris() Locker {
  x := new (morris)
  x.door.Lock()
  return x
}

func (x *morris) Lock() {
  x.mutex.Lock()
  x.n0++
  x.mutex.Unlock()
  x.door0.Lock()
  x.n++
  x.mutex.Lock()
  x.n0--
  if x.n0 > 0 {
    x.mutex.Unlock()
    x.door0.Unlock()
  } else { // n0 == 0
    x.mutex.Unlock()
    x.door.Unlock()
  }
  x.door.Lock()
  x.n--
}

func (x *morris) Unlock() {
  if x.n > 0 {
    x.door.Unlock()
  } else { // x.n == 0
    x.door0.Unlock()
  }
}
