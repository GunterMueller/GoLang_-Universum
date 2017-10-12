package lock

// (c) Christian Maurer   v. 161216 - license see µu.go

// >>> Implementation with sync.Mutex

import
  "sync"
type
  mutex struct {
               sync.Mutex
               }

func newMutex() Locker {
  return new(mutex)
}
