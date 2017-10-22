package audio

// (c) Christian Maurer   v. 170919 - license see µU.go

import
  . "µU/obj"
type
  Audio interface {

  Indexer
}

func New() Audio { return new_() }
