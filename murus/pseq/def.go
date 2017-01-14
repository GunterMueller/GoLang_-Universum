package pseq

// (c) murus.org  v. 140430 - license see murus.go

import
  . "murus/obj"
type
  PersistentSequence interface {

  Equaler; Clearer // ! Comparer, ! Coder => ! Object

  Collector
  Iterator
  Seeker

  Persistor
}
// Pre: a is atomic or of a type implementing Object.
// Returns a new empty persistent sequence for objects of the type of a.
func New(a Any) PersistentSequence { return newPseq(a) }
