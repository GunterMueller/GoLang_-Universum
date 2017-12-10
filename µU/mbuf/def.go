package mbuf

// (c) Christian Maurer   v. 171127 - license see µU.go

import
  . "µU/obj"
type
  MBuffer interface { // Synchronized buffers.
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
// Returns an empty queue for objects of the type of a
// to be used by concurrent processes.
func New (a Any) MBuffer { return new_(a) }

// TODO Spec
func NewFarMonitor (a Any, h string, p uint16, s bool) MBuffer { return newfm(a,h,p,s) }
