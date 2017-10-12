package pseq

// (c) Christian Maurer   v. 170509 - license see µu.go

import
  . "µu/obj"
type
  PersistentSequence interface {

  Equaler
  Clearer
// not yet Comparer, Coder => not yet Object
  Iterator
  Seeker
  Persistor
}

// Pre: a is atomic or of a type implementing Equaler and Coder.
// Returns a new empty persistent sequence for objects of the type of a.
func New(a Any) PersistentSequence { return new_(a) }
