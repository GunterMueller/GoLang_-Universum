package internal

// (c) murus.org  v. 150122 - license see murus.go

import (
  . "murus/obj"
  "murus/str"
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

func (x *pair) imp (Y Any) *pair {
  y, ok := Y.(*pair)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *pair) Variabel() bool {
  return false
}

func (x *pair) Eq (Y Any) bool {
  y := x.imp (Y)
  return x.name == y.name && x.typ == y.typ
}

func (x *pair) Less (Y Any) bool {
  return false
}

func (x *pair) Copy (Y Any) {
  y := x.imp (Y)
  x.name, x.typ = y.name, y.typ
}

func (x *pair) Clone() Any {
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

func (x *pair) Encode() []byte {
  b := make ([]byte, x.Codelen())
  n := uint(len (x.name))
  copy (b[0:n], x.name)
  b[n] = x.typ
  return b
}

func (x *pair) Decode (b []byte) {
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
