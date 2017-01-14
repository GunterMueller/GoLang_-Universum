package piset

// (c) murus.org  v. 161216 - license see murus.go

import
  . "murus/obj"
type
  PersistentIndexedSet interface {

  Clearer
  Sorter
  Collector
  Iterator
  Persistor
}
// Returns a new empty persistent indexed set for objects of the type of o.
func New(o Object, f Func) PersistentIndexedSet { return newPiset(o, f) }
