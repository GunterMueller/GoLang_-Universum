package lock2

// (c) Christian Maurer   v. 190325 - license see µU.go

// Ensures the access to a critical section.
// The functions Lock and Unlock cannot be interrupted
// by calls of Lock or Unlock of other goroutines.

type
  Locker2 interface {

// Pre: p < 2.
//      The calling goroutine is not in the critical section.
// It is the only one in the critical section.
  Lock (p uint)

// Pre: p < 2.
//      The calling goroutine is in the critical section.
// It is not in the critical section.
  Unlock (p uint)
}

// Return new unlocked locks for 2 processes
// with an implementation revealed by their names.
func NewDekker() Locker2 { return newDekker() }
func NewPeterson() Locker2 { return newPeterson() }
func NewDoranThomas() Locker2 { return newDoranThomas() }
func NewKessels() Locker2 { return newKessels() }
