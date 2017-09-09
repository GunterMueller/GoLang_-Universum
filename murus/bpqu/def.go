package bpqu

// (c) Christian Maurer   v. 170218 - license see murus.go

import (
  . "murus/obj"
  "murus/pqu"
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
// with maximal capacity m.
func New (a Any, m uint) BoundedPrioQueue { return new_(a,m) }
