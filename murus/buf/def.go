package buf

// (c) murus.org  v. 161216 - license see murus.go

import (
  . "murus/obj"
  "murus/qu"
)
type
  Buffer interface { // Queues of bounded capacity (could also be named "bqu").

  qu.Queue

// Returns true, iff x is filled up to its capacity.
  Full() bool

// Pre: ! x.Full().
//  Ins (a Any)
}
// Pre: a is atomic or of a type implementing Object.
// Returns an empty buffer of capacity n for objects of the type of a.
func New (a Any, n uint) Buffer { return newBuf(a,n) }
