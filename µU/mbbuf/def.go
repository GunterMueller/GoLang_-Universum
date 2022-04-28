package mbbuf

// (c) Christian Maurer   v. 220420 - license see µU.go

type
  MBoundedBuffer interface { // Synchronized buffers of bounded capacity.
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
// Returns an empty buffer of capacity n for objects of the type of a
// to be used by concurrent processes.
// Classical implementation with explicit synchronisation per Mutex.
func New (a any, n uint) MBoundedBuffer { return new_(a,n) }

// Implementation using synchronisation of µU/buf.
func New1 (a any, n uint) MBoundedBuffer { return new1(a,n) }

// Implementation with a monitor
func NewM (a any, n uint) MBoundedBuffer { return newM(a,n) }

// Implementation with a conditioned monitor a la Go.
func NewGo (a any, n uint) MBoundedBuffer { return newgo(a,n) }

// Implementation with asynchronous message passing
func NewCh (a any, n uint) MBoundedBuffer { return newCh(a,n) }

// Implementation with synchronous message passing
func NewCh1 (a any, n uint) MBoundedBuffer { return newCh1(a,n) }

// Implementation with asynchronous message passing and guarded selective waiting
func NewGS (a any, n uint) MBoundedBuffer { return newgs(a,n) }

// Implementation with a far monitor
func NewFM (a any, n uint, h string, p uint16, s bool) MBoundedBuffer { return newfm(a,n,h,p,s) }
