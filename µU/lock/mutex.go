package lock

// (c) Christian Maurer   v. 230207 - license see µU.go

// >>> Implementation with Mutex

import
  . "sync"
type
  mutex struct {
               Mutex
               }

func newMutex() Locker {
  return new(mutex)
}
