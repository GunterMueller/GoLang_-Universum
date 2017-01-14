package lock

// (c) murus.org  v. 161216 - license see murus.go

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
