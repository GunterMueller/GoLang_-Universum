package persaddr

// (c) Christian Maurer   v. 210601 - license see µU.go

import
  . "µU/obj"
type
  PersonAddress interface { // pairs (person, address)

  TeXer
  Indexer
  Rotator

  Sub (y Any) bool
}

// Returns a new empty pair of Person and Address.
func New() PersonAddress { return new_() }
