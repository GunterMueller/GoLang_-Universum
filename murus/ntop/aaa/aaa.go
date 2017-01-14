package aaa

// (c) murus.org  v. 151121 - license see murus.go

import (
  . "murus/obj"
  "murus/col"
  "murus/scr"
)
type
  aaa struct {
             uint
           n []uint32
        f, b col.Colour
             }

func new_(n uint) AAA {
  x:= new (aaa)
  x.uint = n
  x.n = make ([]uint32, n)
  return x
}

func (x *aaa) imp (Y Any) *aaa {
  y, ok:= Y.(*aaa)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *aaa) Eq (Y Any) bool {
  y:= x.imp (Y)
  for i:= uint(0); i < x.uint; i++ {
    if x.n[i] != y.n[i] { return false }
  }
  return true
}

func (x *aaa) Copy (Y Any) {
  y:= x.imp (Y)
  for i:= uint(0); i < x.uint; i++ {
    x.n[i] = y.n[i]
  }
}

func (x *aaa) Clone() Any {
  y:= New (x.uint)
  y.Copy (x)
  return y
}

func (x *aaa) Less (Y Any) bool {
  return false
}

func (x *aaa) Empty() bool {
  for i:= uint(0); i < x.uint; i++ {
    if x.n[i] > 0 { return false }
  }
  return true
}

func (x *aaa) Clr() {
  for i:= uint(0); i < x.uint; i++ {
    x.n[i] = 0
  }
}

func (x *aaa) Codelen() uint {
  return x.uint * 4
}

func (x *aaa) Encode() []byte {
  bs:= make([]byte, x.Codelen())
  for i:= uint(0); i < x.uint; i++ {
    copy (bs[4*i:4*(i+1)], Encode (x.n[i]))
  }
  return bs
}

func (x *aaa) Decode (bs []byte) {
  for i:= uint(0); i < x.uint; i++ {
     x.n[i] = Decode (uint32(0), bs[4*i:4*(i+1)]).(uint32)
  }
}

func (x *aaa) Colours (f, b col.Colour) {
  x.f, x.b = f, b
}

func (x *aaa) Write (l, c uint) {
  scr.Colours (x.f, x.b)
  for i:= uint(0); i < x.uint; i++ {
    scr.WriteNat (uint(x.n[i]), l + i, c)
  }
}

func (x *aaa) Edit (l, c uint) {
  scr.Colours (x.f, x.b)
  for i:= uint(0); i < x.uint; i++ {
    // TODO
  }
}

func (x *aaa) Put (i uint, a Any) {
  if i >= x.uint { return }
  x.n[i] = a.(uint32)
}

func (x *aaa) Add (Y AAA) {
  y:= x.imp(Y)
  for i:= uint(0); i < x.uint; i++ {
    if y.n[i] > 0 {
      x.n[i] = y.n[i]
    }
  }
}
