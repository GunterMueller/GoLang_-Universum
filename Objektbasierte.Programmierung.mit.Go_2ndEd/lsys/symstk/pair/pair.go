package pair

// (c) Christian Maurer   v. 230303 - license see µU.go

import
  . "µU/obj"
type
  pair struct {
              Symbol
              uint
              }

func new_() Pair {
  return new(pair)
}

func (x *pair) imp (Y any) *pair {
  y, ok := Y.(*pair)
  if ! ok { TypeNotEqPanic (Y, y) }
  return y
}

func (x *pair) Eq (Y any) bool {
  y := x.imp (Y)
  return y.Symbol == x.Symbol && y.uint == x.uint
}

func (x *pair) Copy (Y any) {
  y := x.imp(Y)
  x.Symbol, x.uint = y.Symbol, y.uint
}

func (x *pair) Clone() any {
  y := new_().(*pair)
  y.Copy (x)
  return y
}

func (x *pair) Less (Y any) bool {
  return false
}

func (x *pair) Leq (Y any) bool {
  return false
}

func (x *pair) Empty() bool {
  return x.Symbol == 0 && x.uint == 0
}

func (x *pair) Clr() {
  x.Symbol, x.uint = 0, 0
}

func (x *pair) Codelen() uint {
  return 1 + C0
}

func (x *pair) Encode() Stream {
  s := make(Stream, x.Codelen())
  s[0] = x.Symbol
  copy (s[1:], Encode (x.uint))
  return s
}

func (x *pair) Decode (s Stream) {
  x.Symbol = s[0]
  x.uint = Decode (uint(0), s[1:]).(uint)
}

func (x *pair) Set (s Symbol, i uint) {
  x.Symbol, x.uint = s, i
}

func (x *pair) Get() (Symbol, uint) {
  return x.Symbol, x.uint
}
