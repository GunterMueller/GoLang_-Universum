package ieee

// (c) Christian Maurer   v. 201204 - license see µU.go

import (
  "unsafe"
  "µU/ker"
  "µU/bn"
)
const (
  M = 64 // number of bits for sign, exponent and mantissa
  E = 11 // number of bits for exponent
  A = M - (1 + E) // number of bits for mantissa
  bias = 1023 // 2^(E - 1) - 1 // 
)
type
  ieee struct {
              float64
              string
              }

func new_() IEEE {
  return new(ieee)
}

func dual (n uint) string {
  k := bn.New (bn.M)
  k.SetVal (n)
  return k.Dual()
}

func (x *ieee) SetFloat64 (f float64) {
  x.float64 = f
  p := uint(*(*uint64)(unsafe.Pointer (&f)))
  s := dual(p)
  for len(s) < M { s = "0" + s }
  if f < 0 { s = "1" + s[1:] }
  x.string = s
}

func (x *ieee) String() string {
  if x.float64 == 0 {
    return "0"
  }
  return x.string
}

func decimal (s string) uint {
  n := bn.New (bn.M)
  n.Decimal (s)
  return n.Val()
}

func (x *ieee) SetString (s string) {
  if len(s) != M { ker.PrePanic() }
  x.string = s
  n := decimal (s)
  x.float64 = *(*float64)(unsafe.Pointer (&n))
}

func (x *ieee) Float64() float64 {
  if x.string == "" { ker.PrePanic() }
  return x.float64
}
