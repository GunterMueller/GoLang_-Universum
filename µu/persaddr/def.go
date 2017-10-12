package persaddr

// (c) Christian Maurer   v. 170919 - license see µu.go

import (
  . "µu/obj"
  "µu/col"
)
type
  PersonAddress interface { // pairs (person, address)

  col.Colourer
  Indexer
}

// Returns a new empty pair of Person and Address.
func New() PersonAddress { return new_() }
