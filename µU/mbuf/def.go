package mbuf

// (c) Christian Maurer   v. 220420 - license see µU.go

type
  MBuffer interface { // Synchronized buffers.
                      // The exported functions cannot be interrupted
                      // by calls of these functions of other goroutines.

// a is inserted as last object into x.
// The calling process was blocked, until x was not full.
  Ins (a any)

// Returns the first object of x and that object is removed from x.
// The calling process was blocked, until x was not empty.
  Get() any
}
// Pre: a is atomic or of a type implementing Object.
// Returns an empty queue for objects of the type of a
// to be used by concurrent processes.
func New (a any) MBuffer { return new_(a) }

// TODO Spec
func NewFarMonitor (a any, h string, p uint16, s bool) MBuffer { return newfm(a,h,p,s) }
