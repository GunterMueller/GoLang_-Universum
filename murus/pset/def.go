package pset

// (c) murus.org  v. 161216 - license see murus.go

import
  . "murus/obj"
type
  PersistentSet interface {

  Collector
  ExGeq (a Object) bool
  Trav (o Op)

  Persistor
}
// Returns a new empty persistent set for objects of the type of a.
func New(a Object) PersistentSet { return newPset(a) }
