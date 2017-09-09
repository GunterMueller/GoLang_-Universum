package mqu

// (c) Christian Maurer   v. 170218 - license see murus.go

import (
  . "murus/obj"
  "murus/host"
)
type
  MQueue interface { // Synchronized Queues.
                     // The exported functions cannot be interrupted
                     // by calls of these functions of other goroutines.

// a is inserted as last object into x.
// The calling process was blocked, until x was not full.
  Ins (a Any)

// Returns the first object of x and that object is removed from x.
// The calling process was blocked, until x was not empty.
  Get () Any
}
// Pre: a is atomic or of a type implementing Object.
// Returns an empty queue for objects of the type of a
// to be used by concurrent processes.
func New (a Any) MQueue { return new_(a) }

// TODO Spec
func NewFar (a Any, h host.Host, p uint16, s bool) MQueue { return newf(a,h,p,s) }
