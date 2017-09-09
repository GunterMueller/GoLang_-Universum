package piset

// (c) Christian Maurer   v. 170424 - license see murus.go

import
  . "murus/obj"
type
  PersistentIndexedSet interface {

  Clearer
  Sorter
  Iterator
  Persistor
}

// Returns a new empty persistent indexed set for objects of the type of o.
func New(o Object, f Func) PersistentIndexedSet { return new_(o, f) }
