package bbuf

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "µU/buf"
)
type
  BoundedBuffer interface {

  buf.Buffer

// Returns true, iff x is filled up to its capacity.
// ! x.Full() is a precondition for a call of x.Ins(a).
  Full() bool
}

// Pre: a is atomic or of a type implementing Object. 
// Returns an empty buffer of capacity n for objects of the type of a.
func New (a any, n uint) BoundedBuffer { return new_(a,n) }
