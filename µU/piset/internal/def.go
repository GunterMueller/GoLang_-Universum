package internal

// (c) Christian Maurer   v. 210228 - license see µU.go

import
  . "µU/obj"
type
  Pair interface {

  Equaler
  Comparer

  Pos() uint
}

func New (a Any, n uint) Pair { return new_(a,n) }
