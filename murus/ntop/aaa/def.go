package aaa

// (c) murus.org  v. 151121 - license see murus.go

import
  . "murus/obj"

func New (n uint) AAA { return new_(n) }

type
  AAA interface {

  Editor

  Put (i uint, a Any)

  Add (y AAA)
}
