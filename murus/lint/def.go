package lint

// (c) Christian Maurer   v. 161216 - license see murus.go

import
  . "murus/obj"
type
  LongInteger interface {

  Editor
  Geq0() bool
  Stringer
  Printer
  Adder
  Multiplier

  SetInt(n int)
  SetInt64(n int64)

  Len() uint
  Odd() bool
  ChSign()
  Val() int
  Val64() int64
  SumDigits() uint

// Specs see murus.nat
  Inc()
  Dec()
  Mod(y LongInteger)
  MulMod(y, m LongInteger)
  Div2(y, r LongInteger)
  Gcd(y LongInteger)
  Lcm(y LongInteger)
  Pow(y LongInteger)
  PowMod(y, m LongInteger)
  Fak(n uint)
  Binom(n, k uint)
  LowFak(n, k uint)
  Stirl2(n, k uint)
  ProbabylPrime (n int) bool

  Bitlen() uint
  Bit(i int) uint
  SetBit(i int, b bool)
}
// Return the long integer with Val(n) == n resp. Val64(n) == n.
func New(n int) LongInteger { return newInt(n) }
func New64(n int64) LongInteger { return new64(n) }
