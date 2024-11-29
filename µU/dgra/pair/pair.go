package pair

// (c) Christian Maurer   v. 241014 - license see µU.go

import (
  . "µU/obj"
)
type
  pair struct {
              uint "process number"
              bool
              }

func new_() Pair {
  x := new(pair)
  x.uint = 0
  x.bool = false
  return x
}

func (x *pair) imp (Y any) *pair {
  y, ok := Y.(*pair)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *pair) Eq (Y any) bool {
  y := x.imp (Y)
  return x.uint == y.uint &&
         x.bool == y.bool
}

func (x *pair) Copy (Y any) {
  y := x.imp (Y)
  x.uint, x.bool = y.uint, y.bool 
}

func (x *pair) Clone() any {
  y := new_()
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
  return x.uint == 0 &&
         x.bool == false
}

func (x *pair) Clr() {
  x.uint, x.bool = 0, false
}

func (x *pair) Codelen() uint {
  return C0 + 1
}

func (x *pair) Encode() Stream {
  s := make (Stream, x.Codelen())
  copy (s[0:C0], Encode (x.uint))
  s[5] = '0'
  if x.bool { s[5] = '1' }
  return s
}

func (x *pair) Decode (s Stream) {
  x.uint = Decode (x.uint, s[0:C0]).(uint)
  x.bool = s[5] == '1'
}

func (x *pair) Uint () uint {
  return x.uint
}

func (x *pair) Bool () bool {
  return x.bool
}

func (x *pair) Set (n uint, b bool) {
   x.uint, x.bool = n, b 
}
