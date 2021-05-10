package persaddr

// (c) Christian Maurer   v. 210510 - license see µU.go

import
  . "µU/obj"
type
  PersonAddress interface { // pairs (person, address)

  Indexer
  Rotator
}

// Returns a new empty pair of Person and Address.
func New() PersonAddress { return new_() }
