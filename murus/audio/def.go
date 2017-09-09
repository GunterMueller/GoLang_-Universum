package audio

// (c) Christian Maurer   v. 161216 - license see murus.go

import
  . "murus/obj"
type
  Audio interface {

  Indexer
}

func New() Audio { return new_() }
