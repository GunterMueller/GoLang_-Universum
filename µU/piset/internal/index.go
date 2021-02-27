package internal

// (c) Christian Maurer   v. 210221 - license see µU.go

import
  . "µU/obj"
type
  index struct {
         empty Any
               Any "index Object"
               uint32 "position"
               }

func new_(a Any) Index {
  x := new(index)
  x.empty, x.Any = Clone(a), Clone(a)
  return x
}

func (x *index) imp (Y Any) *index {
  y, ok := Y.(*index)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *index) Set (a Any, n uint) {
  x.Any, x.uint32 = Clone (a), uint32(n)
}

func (x *index) Get () Any {
  return Clone (x.Any)
}

func (x *index) Copy (Y Any) {
  y := x.imp(Y)
  x.empty = Clone (y.empty)
  x.Set(y.Any, uint(y.uint32))
}

func (x *index) Clone() Any {
  y := new_(x.empty)
  y.Copy(x)
  return y
}

func (x *index) Eq (Y Any) bool {
  y := x.imp(Y)
  return Eq (x.Any, y.Any) && x.uint32 == y.uint32
}

func (x *index) Less (Y Any) bool {
  return Less (x.Any, x.imp(Y).Any)
}

func (x *index) Pos() uint {
  return uint(x.uint32)
}

func (x *index) Defined (s string) bool {
  return false
}

func (x *index) String() string {
  return x.Any.(Stringer).String()
}
