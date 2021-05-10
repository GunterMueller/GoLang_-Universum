package audio

// (c) Christian Maurer   v. 210509 - license see µU.go

import
  . "µU/obj"
type
  Audio interface {

  Indexer
  TeXer
}

func New() Audio { return new_() }
