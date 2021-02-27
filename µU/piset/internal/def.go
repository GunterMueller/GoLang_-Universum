package internal

// (c) Christian Maurer   v. 210221 - license see µU.go

import
  . "µU/obj"
type
  Index interface { // TODO detailed explanations

  Equaler
  Comparer

// x has a Clone of a as indexobject and position n.
  Set (a Any, n uint)

// Returns a Clone of the indexobject of x.
  Get() Any

// Returns the position of x.
  Pos() uint
}

func New (a Any) Index { return new_(a) }
