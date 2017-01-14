package lock

// (c) murus.org  v. 161216 - license see murus.go

// >>> Algorithm of Morris: A starvation-free Solution to the Mutual Exclusion Problem
//     Inf. Proc. Letters 8 (1979), 76-80

import
  "sync"
type
  morris struct {
    door0, door,
          mutex sync.Mutex // to protect n0
             n0, // number of processes ready to pass through gate a
              n uint // number of processes, that have passed through
                }    // gate a, but not yet through gate m

func newMorris() Locker {
  x:= new (morris)
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
