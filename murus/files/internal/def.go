package internal

// (c) murus.org  v. 150122 - license see murus.go

import
  . "murus/obj"
type
  Pair interface {

  Object

  Name () string
  Set (s string, b byte)
  Typ () byte
}

func New() Pair { return new_() }
