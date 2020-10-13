package pseq

// (c) Christian Maurer   v. 201011 - license see µU.go

import
  . "µU/obj"
type
  PersistentSequence interface {

  Equaler
  Clearer
  Sorter
  Seeker
  Persistor
}

// Pre: a is atomic or of a type implementing Equaler and Coder.
// Returns a new empty persistent sequence for objects of the type of a,
// that is ordered iff o.
func New (a Any, o bool) PersistentSequence { return new_(a,o) }
