package audio

// (c) Christian Maurer   v. 170919 - license see µu.go

import
  . "µu/obj"
type
  Audio interface {

  Indexer
}

func New() Audio { return new_() }
