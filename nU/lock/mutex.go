package lock

// (c) Christian Maurer   v. 171231 - license see nU.go

import "sync"

type mutex struct {
  sync.Mutex
}

func newMutex() Locker {
  return new(mutex)
}
