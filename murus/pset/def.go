package pset

// (c) Christian Maurer   v. 170424 - license see murus.go

import
  . "murus/obj"
type
  PersistentSet interface {

  Clearer
  Iterator // XXX most methods not yet implemented
  ExGeq (a Object) bool // not yet Ordered, Sort => not yet Sorter
  Persistor
}

// Returns a new empty persistent set for objects of the type of a.
func New(a Object) PersistentSet { return new_(a) }
