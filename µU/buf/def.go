package buf

// (c) Christian Maurer   v. 170621 - license see µU.go

import (
  . "µU/obj"
  "µU/qu"
)
type
  Buffer interface { // Bounded buffers = queues of bounded capacity.

  qu.Queue

// Returns true, iff x is filled up to its capacity.
// ! x.Full() is a precondition for a call of x.Ins(a).
  Full() bool
}

// Pre: a is atomic or of a type implementing Object. 
// Returns an empty buffer of capacity n for objects of the type of a.
func New (a Any, n uint) Buffer { return new_(a,n) }
