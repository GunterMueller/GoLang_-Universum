package pair

// (c) Christian Maurer   v. 221021 - license see µU.go

import (
  . "µU/obj"
  "µU/str"
)
const
  pack = "files/internal"
type
  pair struct {
         name string
          typ byte
              }

func new_() Pair {
  return new(pair)
}

func (x *pair) imp (Y any) *pair {
  y, ok := Y.(*pair)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *pair) Variabel() bool {
  return false
}

func (x *pair) Eq (Y any) bool {
  y := x.imp (Y)
  return x.name == y.name && x.typ == y.typ
}

func (x *pair) Less (Y any) bool {
  return false
}

func (x *pair) Leq (Y any) bool {
  return false
}

func (x *pair) Copy (Y any) {
  y := x.imp (Y)
  x.name, x.typ = y.name, y.typ
}

func (x *pair) Clone() any {
  y := New()
  y.Copy (x)
  return y
}

func (x *pair) Empty() bool {
  return str.Empty (x.name)
}

func (x *pair) Clr() {
  x.name = ""
  x.typ = 0
}

func (x *pair) Codelen() uint {
  return uint(len (x.name)) + 1
}

func (x *pair) Encode() Stream {
  b := make (Stream, x.Codelen())
  n := uint(len (x.name))
  copy (b[0:n], x.name)
  b[n] = x.typ
  return b
}

func (x *pair) Decode (b Stream) {
  n := uint(len (b))
  x.name = string(b[0:n])
  x.typ = b[n]
}

func (x *pair) Set (s string, b byte)  {
  x.name, x.typ = s, b
}

func (x *pair) Name() string {
  return x.name
}

func (x *pair) Typ() byte {
  return x.typ
}
