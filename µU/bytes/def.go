package bytes

// (c) Christian Maurer   v. 170121 - license see µU.go
//
// >>> Just for fun, most likely completely worthless

import
  . "µU/obj"
type
  ByteSequence interface {

  Object
}

func New (n uint) ByteSequence { return new_(n) }
