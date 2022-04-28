package texts

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/font"
  "µU/text"
)
type
  texts struct {
               uint "number of texts"
             t []text.Text
             n []uint
               }

func new_(n []uint) Texts {
  x := new(texts)
  x.uint = uint(len(n))
  for i := uint(0); i < x.uint; i++ {
    x.n[i] = n[i]
    x.t[i] = text.New (x.n[i])
  }
  return x
}

func (x *texts) imp(Y any) *texts {
  y, ok := Y.(*texts)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *texts) Len() []uint {
  return x.n
}

func (x *texts) Empty() bool {
  for i := uint(0); i < x.uint; i++ {
    if ! x.t[i].Empty() {
      return false
    }
  }
  return true
}

func (x *texts) Clr() {
  for i := uint(0); i < x.uint; i++ {
    x.t[i].Clr()
  }
}

func (x *texts) Copy (Y any) {
  y := x.imp (Y)
  for i := uint(0); i < x.uint; i++ {
    x.n[i] = y.n[i]
    x.t[i].Copy (y.t[i])
  }
}

func (x *texts) Clone() any {
  y := new_(x.n)
  y.Copy (x)
  return y
}

func (x *texts) Eq (Y any) bool {
  y := x.imp (Y)
  for i := uint(0); i < x.uint; i++ {
    if ! x.t[i].Eq (y.t[i]) {
      return false
    }
  }
  return true
}

func (x *texts) Less (Y any) bool {
  y := x.imp (Y)
  for i := uint(0); i < x.uint; i++ {
    if ! x.t[i].Less (y.t[i]) {
      return false
    }
  }
  return true
}

func (x *texts) Leq (Y any) bool {
  y := x.imp (Y)
  return x.Eq(y) || x.Less(y)
}

func (x *texts) Colours (f, b col.Colour) {
  for i := uint(0); i < x.uint; i++ {
    x.t[i].Colours (f, b)
  }
}

func (x *texts) SetFont (f font.Font) {
  for i := uint(0); i < x.uint; i++ {
    x.t[i].SetFont (f)
  }
}

func (x *texts) Print (l, c uint) {
  for i := uint(0); i < x.uint; i++ {
    x.t[i].Print (l, c)
  }
}

func (x *texts) Codelen() uint {
  l := C0
  for i := uint(0); i < x.uint; i++ {
    l += x.n[i]
  }
  return l
}

func (x *texts) Encode() Stream {
  s := make(Stream, x.Codelen())
  i, a := uint(0), C0
  copy (s[i:a], Encode(x.uint))
  i+= a
  for j := uint(0); j < x.uint; j++ {
    a = x.n[i]
    copy (s[i:i+a], x.t[i].Encode())
    i+= a
  }
  return s
}

func (x *texts) Decode (s Stream) {
  i, a := uint(0), C0
  x.uint = Decode(uint(0), s[i:a]).(uint)
  i+= a
  for i := uint(0); i < x.uint; i++ {
    a = x.n[i]
    x.t[i].Decode (s[i:i+a])
    i+= a
  }
}

func (x *texts) Write (l, c uint) {
  for i := uint(0); i < x.uint; i++ {
    x.t[i].Write (l, c)
  }
}

func (x *texts) Edit (l, c uint) {
  for i := uint(0); i < x.uint; i++ {
    x.t[i].Edit (l, c)
  }
}
