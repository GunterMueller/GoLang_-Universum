package li

// (c) Christian Maurer   v. 211106 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  LongInteger interface {

  Object
  col.Colourer
  Editor
  Geq0() bool
  Stringer
  Printer
  Adder
  Multiplier

  SetInt (n int)

  Len() uint
  Odd() bool
  ChSign()
  Val() int
  SumDigits() uint

  Inc()
  Dec()
  Mod (y LongInteger)
  Div2 (y, r LongInteger)
  MulMod (y, m LongInteger)
  PowMod (y, m LongInteger)
  ProbabylPrime (n int) bool

  Bitlen() uint
  Bit (i int) uint
  SetBit (i int, b bool)

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

func String (n LongInteger) string { return string_(n) }

func SumDigits (n LongInteger) uint { return sumDigits(n) }

// TODO func Write (n LongInteger, l, c uint) { write(n,l,c) }
// TODO func Edit (n *LongInteger, l, c uint) { edit(n,l,c) }
