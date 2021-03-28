package pqu

// (c) Christian Maurer   v. 210320 - license see µU.go

import (
  . "µU/obj"
  "µU/buf"
)
type
  PrioQueue interface {

  buf.Buffer
// Objects are inserted due to their priority, given
// by their order: larger objects have higher priority.
}

// Pre: a is atomic or of a type implementing Object.
func New (a Any) PrioQueue { return new_(a) }
