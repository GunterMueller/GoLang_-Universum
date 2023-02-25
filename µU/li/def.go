package li

// (c) Christian Maurer   v. 230217 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  LongInteger interface {

  Editor
  col.Colourer
  Stringer
  Printer
  Adder
  Multiplier

  Odd() bool
  Geq0() bool
  ChSign()
  Val() int
  SumDigits() uint

  Inc()
  Dec()
  Abs() LongInteger
  Mod (y LongInteger)
  Div2 (y, r LongInteger)
  MulMod (y, m LongInteger)
  PowMod (y, m LongInteger)
  ProbabylPrime (n int) bool

// Specs see µU.n
  Gcd (y LongInteger)
  Lcm (y LongInteger)
  Pow (y LongInteger)
  Fak (n uint)
  Binom (n, k uint)
  LowFak (n, k uint)
  Stirl2 (n, k uint)
}

// Returns a new long integer with Val(n) == n.
func New (n int) LongInteger { return new_(n) }
