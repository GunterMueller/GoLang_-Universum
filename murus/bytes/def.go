package bytes

// (c) Christian Maurer   v. 170121 - license see murus.go
//
// >>> Just for fun, most likely completely worthless

import
  . "murus/obj"

type
  ByteSequence interface {

  Object
}

func New (n uint) ByteSequence { return new_(n) }
