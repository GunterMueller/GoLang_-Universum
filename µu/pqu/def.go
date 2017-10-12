package pqu

// (c) Christian Maurer   v. 130316 - license see µu.go

import (
  . "µu/obj"
  "µu/qu"
)
type
  PrioQueue interface {

  qu.Queue
// where Objects are inserted due to their priority, given by their Order.
// Lower Objects have higher priority.
}

func New(a Any) PrioQueue { return new_(a) }
