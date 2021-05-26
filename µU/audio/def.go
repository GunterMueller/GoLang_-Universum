package audio

// (c) Christian Maurer   v. 210525 - license see µU.go

import
  . "µU/obj"
type
  Audio interface {

  Indexer
  TeXer

  Sub (Y Any) bool
}

func New() Audio { return new_() }
