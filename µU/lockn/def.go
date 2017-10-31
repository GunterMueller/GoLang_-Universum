package lockn

// (c) Christian Maurer   v. 171025 - license see µU.go

// Ensures the access to a critical section.
// The functions Lock and Unlock cannot be interrupted
// by calls of Lock or Unlock of other goroutines.

type
  LockerN interface {

// Pre: The calling goroutine is not in the critical section.
// It is the only one in the critical section.
  Lock (p uint)

// Pre: The calling goroutine is in the critical section.
// It is not in the critical section.
  Unlock (p uint)
}

// Return new unlocked locks for 2 processes
// with an implementation revealed by their names.
func NewDekker() LockerN { return newDe() }
func NewPeterson() LockerN { return newP() }
func NewDoranThomas() LockerN { return newDT() }
func NewKessels2() LockerN { return newK2() }

// Return new unlocked locks for n processes
// with an implementation revealed by their names.
func NewDijkstra (n uint) LockerN { return newD(n) }
func NewHabermann (n uint) LockerN { return newH(n) }
func NewBakery (n uint) LockerN { return newB(n) }
func NewBakery1 (n uint) LockerN { return newB1(n) }
func NewTicket (n uint) LockerN { return newT(n) }
func NewTiebreaker (n uint) LockerN { return newTb(n) }
func NewKesselsN (n uint) LockerN { return newKN(n) }
func NewSzymanski (n uint) LockerN { return newS(n) }
func NewKnuth (n uint) LockerN { return newK(n) }
func NewDeBruijn (n uint) LockerN { return newDB(n) }
func NewBurns (n uint) LockerN { return newBu(n) }
