package persaddr

// (c) murus.org  v. 161216 - license see murus go

import
  . "murus/obj"
type
  PersonAddress interface { // pairs (person, address)

  Indexer
}

// Returns a new empty pair of Person and Address.
func New() PersonAddress { return new_() }
