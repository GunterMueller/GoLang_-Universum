package lock

// (c) Christian Maurer   v. 161216 - license see nU.go

// >>> Algorithm of Udding

import
   "sync"
type
  udding struct {
    mutex, door,
          queue sync.Mutex
          n0, n uint
                }

func newUdding() Locker {
  x:= new (udding)
  x.door.Lock()
  return x
}

func (x *udding) Lock() {
  x.mutex.Lock()
  x.n0++
  x.mutex.Unlock()
  x.queue.Lock()
  x.mutex.Lock()
  x.n++
  x.n0--
  if x.n0 > 0 {
    x.mutex.Unlock()
  } else { // x.n0 == 0
    x.door.Unlock()
  }
  x.queue.Unlock()
  x.door.Lock()
  x.n--
}

func (x *udding) Unlock() {
  if x.n > 0 {
    x.door.Unlock()
  } else { // x.n == 0
    x.mutex.Unlock()
  }
}
