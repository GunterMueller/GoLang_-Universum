package internal

// (c) Christian Maurer   v. 150122 - license see µu.go

import
  . "µu/obj"
type
  Heap interface {

// Pre: n == number of objects in x > 0.
// a is inserted as n-th node in x; returns x.
  Ins (a Any, n uint) Heap

// Pre: n == number of objects in x.
// If n <= 1, nothing has happened, otherwise the n-th
// object of x is lifted, until the heap invariant is restored.
  Lift (n uint)

// Pre: n == number of objects in x.
// The n-th object of x moved to the top of x;
// returns (x, former top of x).
  Del (n uint) (Heap, Any)

// Pre: n == number of objects in x > 0.
// The top of x is dropped down, until the heap invariant is restored.
  Sift (n uint)

// Returns nil, iff x == nil, otherwise a copy of the top of x.
  Get () Any
}

func New() Heap { return new_() }
