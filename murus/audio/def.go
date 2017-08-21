package audio

// (c) murus.org  v. 161216 - license see murus.go

import
  . "murus/obj"
type
  Audio interface {

  Indexer
}

func New() Audio { return new_() }
