package pqu

// (c) murus.org  v. 130115 - license see murus.go

import
  "murus/qu"
type
  PrioQueue interface {

  qu.Queue
// where Objects are inserted due to their priority, given by their Order.
// Lower Objects have higher priority.
}
