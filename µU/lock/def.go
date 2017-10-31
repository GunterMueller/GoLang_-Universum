package lock

// (c) Christian Maurer   v. 171020 - license see µU.go

// Secures the access to a critical section.
// The functions Lock and Unlock cannot be interrupted
// by calls of Lock or Unlock of other goroutines.

type
  Locker interface {

// Pre: The calling goroutine is not in the critical section.
// It is the only one in the critical section.
  Lock()

// Pre: The calling goroutine is in the critical section.
// It is not in the critical section.
  Unlock()
}

// Return new unlocked locks
// with an implementation revealed by their names.
func NewCAS() Locker { return newCAS() }
func NewChannel() Locker { return newChan() }
func NewDEC() Locker { return newDEC() }
func NewMorris() Locker { return newMorris() }
func NewMutex() Locker { return newMutex() }
func NewTAS() Locker { return newTAS() }
func NewUdding() Locker { return newUdding() }
func NewXCHG() Locker { return newXCHG() }
