package pqu

// (c) Christian Maurer   v. 171104 - license see µU.go

import (
  . "µU/obj"
  "µU/buf"
)
type
  PrioQueue interface {

  buf.Buffer
// where Objects are inserted due to their priority, given by their Order.
// Lower Objects have higher priority.
}

// Pre: a is atomic or of a type implementing Object (a != nil).
func New(a Any) PrioQueue { return new_(a) }
