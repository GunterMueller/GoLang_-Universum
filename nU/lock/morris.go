package lock

// (c) Christian Maurer   v. 171024 - license see nU.go

// >>> Algorithm of Morris

import "sync"

type morris struct {
  door0, door,
  mutex sync.Mutex // to protect n0
  n0,    // number of processes blocked on door0
  n uint // number of processes blocked on door
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
