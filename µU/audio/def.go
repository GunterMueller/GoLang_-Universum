package audio

// (c) Christian Maurer   v. 220420 - license see µU.go

import
  . "µU/obj"
type
  Audio interface {

  Indexer
  TeXer

  Sub (Y any) bool
}

func New() Audio { return new_() }
