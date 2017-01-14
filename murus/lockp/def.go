package lockp

// (c) murus.org  v. 161212 - license see murus.go

// Ensures the access to a critical section.
// The functions Lock and Unlock cannot be interrupted
// by calls of Lock or Unlock of other goroutines.

type
  LockerP interface {

// Pre: The calling goroutine is not in the critical section.
// It is the only one in the critical section.
  Lock (p uint)

// Pre: The calling goroutine is in the critical section.
// It is not in the critical section.
  Unlock (p uint)
}
// Return new unlocked locks for 2 processes
// with an implementation revealed by their names.
func NewDekker() LockerP { return newDe() }
func NewDoranThomas() LockerP { return newDT() }
func NewKessels2() LockerP { return newK2() }
func NewPeterson() LockerP { return newP() }
// Return new unlocked locks for n processes
// with an implementation revealed by their names.
func NewBakery (n uint) LockerP { return newB(n) }
func NewBakery1 (n uint) LockerP { return newB1(n) }
func NewDijkstra (n uint) LockerP { return newD(n) }
func NewHabermann (n uint) LockerP { return newH(n) }
func NewKesselsN (n uint) LockerP { return newKN(n) }
func NewTicket (n uint) LockerP { return newT(n) }
func NewTiebreaker (n uint) LockerP { return newTb(n) }
