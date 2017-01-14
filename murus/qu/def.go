package qu

// (c) murus.org  v. 161216 - license see murus.go

import
  . "murus/obj"
type
  Queue interface { // Fifo-Queues

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
func New(a Any) Queue { return newQu(a) }
