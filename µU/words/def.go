package words

// (c) Christian Maurer   v. 180815 - license see µU.go

import
  . "µU/obj"
type
  WordSequence interface {

  Object
  Editor
}

func New (n, l uint) WordSequence { return new_(n,l) }
