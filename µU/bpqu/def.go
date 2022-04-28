package bpqu

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "µU/pqu"
)
type
  BoundedPrioQueue interface {

  pqu.PrioQueue // priority queue with bounded capacity

// Returns true, iff x is filled up to its capacity.
  Full() bool
}

// Pre: a is atomic or of a type implementing Object; m > 0.
// Returns a new empty priority queue for objects of type a
// with maximal capacity m.
func New (a any, m uint) BoundedPrioQueue { return new_(a,m) }
