package buf

// (c) Christian Maurer   v. 171106 - license see µU.go

import
  . "µU/obj"
type
  Buffer interface { // FIFO-Queues

// Returns true, if there are no objects in x.
  Empty() bool

// Returns the number of objects in x.
  Num() uint

// a is inserted as last object into x.
  Ins (a Any)

// Returns the pattern object of x, if x.Empty().
// Returns otherwise the first object of x
// and that object is removed from x.
  Get() Any
}

// Pre: a is atomic or of a type implementing Object.
// Returns a new empty queue for objects of the type of a.
// a is the pattern object of this buffer.
func New (a Any) Buffer { return new_(a) }
func NewS (a Any) Buffer { return newS(a) }
