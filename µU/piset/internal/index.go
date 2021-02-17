package internal

// (c) Christian Maurer   v. 201014 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
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

func (x *index) Empty() bool {
  return Eq (x.Any, x.empty)
}

func (x *index) Clr() {
  x.Any, x.uint32 = Clone (x.empty), 0
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

func editor (X Any) Editor {
  x, ok := X.(Editor)
  if ! ok { TypeNotEqPanic(x, X) }
  return x
}

func colourer (X Any) col.Colourer {
  x, ok := X.(col.Colourer)
  if ! ok { TypeNotEqPanic(x, X) }
  return x
}

func (x *index) Colours (f, b col.Colour) {
  colourer(x.Any).Colours (f, b)
}

func (x *index) Write (l, c uint) {
  editor(x.Any).Write (l, c)
}

func (x *index) Edit (l, c uint) {
  e := editor(x.Any)
  e.Edit(l, c)
  x.Any = Clone(e)
}

func (x *index) Defined (s string) bool {
  return false
}

func (x *index) String() string {
  return x.Any.(Stringer).String()
}

func (x *index) Codelen() uint {
  return Codelen(x.Any) + 4
}

func (x *index) Encode() Stream {
  b := make(Stream, x.Codelen())
  n := uint(Codelen(x.Any))
  copy(b[:n], Encode(x.Any))
  copy(b[n:n+4], Encode(x.uint32))
  return b
}

func (x *index) Decode (b Stream) {
  n := Codelen(x.Any)
  Decode(x.Any, b[:n])
  x.uint32 = Decode(uint32(0), b[n:n+4]).(uint32)
}
