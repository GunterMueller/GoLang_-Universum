package bpqu

// (c) Christian Maurer   v. 171104 - license see µU.go

import (
  . "µU/obj"
  "µU/pqu"
)
type
  BoundedPrioQueue interface {

// A bounded priority queue is a
  pqu.PrioQueue // with bounded capacity.

// Returns true, iff x is filled up to its capacity.
  Full() bool
}

// Pre: a is atomic or of a type implementing Object.
// Returns a new empty priority queue for objects of type a
// with maximal capacity m. m > 0.
func New (a Any, m uint) BoundedPrioQueue { return new_(a,m) }
