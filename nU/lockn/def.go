package lockn

// (c) Christian Maurer   v. 190325 - license see nU.go

// Ensures the access to a critical section.
// The functions Lock and Unlock cannot be interrupted
// by calls of Lock or Unlock of other goroutines.

type
  LockerN interface {

// Pre: p < number of processes defined by the constructor.
//      The calling goroutine is not in the critical section.
// It is the only one in the critical section.
  Lock (p uint)

// Pre: p < number of processes defined by the constructor.
//      The calling goroutine is in the critical section.
// It is not in the critical section.
  Unlock (p uint)
}

// Return new unlocked locks for n processes
// with an implementation revealed by their names.
func NewDijkstra (n uint) LockerN { return newDijkstra(n) }
func NewHabermann (n uint) LockerN { return newHabermann(n) }
func NewBakery (n uint) LockerN { return newBakery(n) }
func NewBakery1 (n uint) LockerN { return newBakery1(n) }
func NewTicket (n uint) LockerN { return newTicket(n) }
func NewTiebreaker (n uint) LockerN { return newTiebreaker(n) }
func NewKessels (n uint) LockerN { return newKessels(n) } // Pre: n is a power of 2.
func NewSzymanski (n uint) LockerN { return newSzymanski(n) }
func NewKnuth (n uint) LockerN { return newKnuth(n) }
func NewDeBruijn (n uint) LockerN { return newDeBruijn(n) }
func NewChannel (n uint) LockerN { return newChannel(n) }
