package internal

// (c) murus.org  v. 150123 - license see murus.go

import
  . "murus/obj"
type
  Index interface { // TODO detailed explanations

  Object
//  Editor

// x has a Clone of a as indexobject and position n.
  Set(a Any, n uint)

// Returns a Clone of the indexobject of x.
  Get() Any

// Returns the position of x.
  Pos() uint
}
