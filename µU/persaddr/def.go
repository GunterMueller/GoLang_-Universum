package persaddr

// (c) Christian Maurer   v. 170919 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  PersonAddress interface { // pairs (person, address)

  col.Colourer
  Indexer
  Orderer
}

// Returns a new empty pair of Person and Address.
func New() PersonAddress { return new_() }
