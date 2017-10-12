package qu

// (c) Christian Maurer   v. 161216 - license see µu.go

import
  . "µu/obj"
type
  Queue interface { // Fifo-Queues

// Returns true, if there are no objects in x.
  Empty() bool

// Returns the number of objects in x.
  Num() uint

// a is inserted as last object into x.
  Ins (a Any)

// Returns nil, if x.Empty(), otherwise the first object of x
// and that object is removed from x.
  Get() Any
}

// Pre: a is atomic or of a type implementing Object.
// Returns a new empty queue for objects of type a.
func New(a Any) Queue { return new_(a) }
func NewS(a Any) Queue { return news(a) }
