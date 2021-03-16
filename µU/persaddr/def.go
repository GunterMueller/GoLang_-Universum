package persaddr

// (c) Christian Maurer   v. 210308 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  PersonAddress interface { // pairs (person, address)

  col.Colourer
  Indexer
  Rotator
}

// Returns a new empty pair of Person and Address.
func New() PersonAddress { return new_() }
