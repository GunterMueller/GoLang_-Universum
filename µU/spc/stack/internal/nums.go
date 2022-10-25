package nums

// (c) Christian Maurer   v. 221021 - license see µU.go

import
  . "µU/obj"
const
  n = 9
type
  numbers struct {
               uint
             s []float64
               }

func new_(n uint) Numbers {
  x := new(numbers)
  x.uint = n
  x.s = make([]float64, x.uint)
  return x
}

func (s *numbers) imp (y any) *numbers {
  t, ok := y.(*numbers)
  if ! ok { TypeNotEqPanic (s, y) }
  return t
}

func (x *numbers) Eq (Y any) bool {
  y := x.imp(Y)
  for i := uint(0); i < x.uint; i++ {
    if x.s[i] != y.s[i] {
      return false
    }
  }
  return true
}

func (x *numbers) Copy (Y any) {
  y := x.imp(Y)
  for i := uint(0); i < x.uint; i++ {
    x.s[i] = y.s[i]
  }
}

func (x *numbers) Clone() any {
  y := new_(x.uint).(*numbers)
  y.Copy (x)
  return y
}

func (x *numbers) Less (Y any) bool {
  return false
}

func (x *numbers) Leq (Y any) bool {
  return false
}

func (x *numbers) Empty() bool {
  for i := uint(0); i < x.uint; i++ {
    if x.s[i] != 0 {
      return false
    }
  }
  return true
}

func (x *numbers) Clr() {
  for i := uint(0); i < x.uint; i++ {
    x.s[i] = 0
  }
}

func (x *numbers) Codelen() uint {
  return x.uint * 8
}

func (x *numbers) Encode() Stream {
  s := make(Stream, x.Codelen())
  i, a := uint(0), uint(8)
  for j := uint(0); j < x.uint; j++ {
    copy (s[i:i+a], Encode (x.s[j]))
    i += a
  }
  return s
}

func (x *numbers) Decode (b Stream) {
  i, a := uint(0), uint(8)
  for j := uint(0); j < x.uint; j++ {
    x.s[j] = Decode (x.s[j], b[i:i+a]).(float64)
    i += a
  }
}

func (x *numbers) Set (r ...float64) {
  if uint(len(r)) != x.uint { println ("ganz große Scheiße") }
  for i := uint(0); i < x.uint; i++ {
    x.s[i] = r[i]
  }
}

func (x *numbers) Get() []float64 {
  a := make([]float64, x.uint)
  for i := uint(0); i < x.uint; i++ {
    a[i] = x.s[i]
  }
  return a
}
