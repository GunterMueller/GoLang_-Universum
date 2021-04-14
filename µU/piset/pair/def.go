package pair

// (c) Christian Maurer   v. 210323 - license see µU.go

import
  . "µU/obj"
type
  Pair interface {

  Equaler
  Comparer

  Pos() uint
  Index() Any
}

func New (a Any, n uint) Pair { return new_(a,n) }
