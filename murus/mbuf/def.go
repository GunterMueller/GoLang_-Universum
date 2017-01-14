package mbuf

// (c) murus.org  v. 161216 - license see murus.go

import
  . "murus/obj"
type
  MBuffer interface { // Synchronized Queues of bounded capacity.
                      // The exported functions cannot be interrupted
                      // by calls of these functions of other goroutines.

// a is inserted as last object into x.
// The calling process was blocked, until x was not full.
  Ins (a Any)

// Returns the first object of x and that object is removed from x.
// The calling process was blocked, until x was not empty.
  Get() Any
}
// Pre: a is atomic or of a type implementing Object.
// Returns an empty buffer of capacity n for objects of the type of a
// to be used by concurrent processes.
func New (a Any, n uint) MBuffer { return newMbuf(a,n) }
