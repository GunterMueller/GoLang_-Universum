package bytes

// (c) murus.org  v. 170121 - license see murus.go
//
// >>> Just for fun, most likely completely worthless

import
  . "murus/obj"

type
  ByteSequence interface {

  Object
}

func New (n uint) ByteSequence { return new_(n) }
