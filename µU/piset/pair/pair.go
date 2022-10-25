package pair

// (c) Christian Maurer   v. 221021 - license see µU.go

import
  . "µU/obj"
type
  pair struct {
              any "index"
              uint "position"
              }

func new_(a any, n uint) Pair {
  x := new(pair)
  x.any = Clone(a)
  x.uint = n
  return x
}

func (x *pair) imp (Y any) *pair {
  y, ok := Y.(*pair)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *pair) Copy (Y any) {
  y := x.imp(Y)
  x.any = Clone (y.any)
  x.uint = y.uint
}

func (x *pair) Clone() any {
  y := new_(x.any, x.uint)
  y.Copy (x)
  return y
}

func (x *pair) Eq (Y any) bool {
  return Eq (x.any, x.imp(Y).any)
}

func (x *pair) Less (Y any) bool {
  return Less (x.any, x.imp(Y).any)
}

func (x *pair) Leq (Y any) bool {
  return Leq (x.any, x.imp(Y).any)
}

func (x *pair) Pos() uint {
  return x.uint
}

func (x *pair) Index() any {
  return x.any
}

func (x *pair) TeX() string {
  return x.any.(TeXer).TeX()
}
