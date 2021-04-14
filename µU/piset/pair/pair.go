package pair

// (c) Christian Maurer   v. 210323 - license see µU.go

import
  . "µU/obj"
type
  pair struct {
              Any "index"
              uint "position"
              }

func new_(a Any, n uint) Pair {
  x := new(pair)
  x.Any = Clone(a)
  x.uint = n
  return x
}

func (x *pair) imp (Y Any) *pair {
  y, ok := Y.(*pair)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *pair) Copy (Y Any) {
  y := x.imp(Y)
  x.Any = Clone (y.Any)
  x.uint = y.uint
}

func (x *pair) Clone() Any {
  y := new_(x.Any, x.uint)
  y.Copy (x)
  return y
}

func (x *pair) Eq (Y Any) bool {
  return Eq (x.Any, x.imp(Y).Any)
}

func (x *pair) Less (Y Any) bool {
  return Less (x.Any, x.imp(Y).Any)
}

func (x *pair) Pos() uint {
  return x.uint
}

func (x *pair) Index() Any {
  return x.Any
}
