package pair

// (c) Christian Maurer   v. 220420 - license see µU.go

import
  . "µU/obj"
type
  Pair interface {

  Equaler
  Comparer

  Pos() uint
  Index() any
  TeX() string
}

func New (a any, n uint) Pair { return new_(a,n) }
