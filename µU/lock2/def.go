package lock2

// (c) Christian Maurer   v. 240930 - license see µU.go

// Protocols for critical sections.
// The functions Lock and Unlock cannot be interrupted
// by calls of these functions by other processes.

type
  Locker2 interface {

// Pre: p < 2.
//      The calling process is not in the critical section.
// It is the only one in the critical section.
// It might have been delayed, until this was possible.
  Lock (p uint)

// Pre: p < 2.
//      The calling process is in the critical section.
// It is not in the critical section.
  Unlock (p uint)
}

// All constructors return a new lock for 2 processes
// with an implementation revealed by their names.
func NewDekker() Locker2 { return newDekker() }
func NewPeterson() Locker2 { return newPeterson() }
func NewDoranThomas() Locker2 { return newDoranThomas() }
func NewKessels() Locker2 { return newKessels() }
