package euro

// (c) Christian Maurer   v. 220131 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
const
  Limit = uint(10 * 1000 * 1000)
type
  Euro interface {

  Object
  col.Colourer
  Editor
  Stringer
  Printer

// The values are cents.
  Valuator
  RealValuator

  Val2() (uint, uint)

  // Pre: The first parameter is < Limit and the second parameter is < 100.
  SetVal2 (uint, uint)

  Adder
  Operate (Factor, Divisor uint)
  ChargeInterest (p, n uint)
  Round (E Euro)
}

// Returns a new empty Euro.
func New() Euro { return new_() }

// Pre: e < Limit and c < 100.
// Returns a new Euro with Val2() = (e, c).
func New2 (e, c uint) Euro { return new2(e,c) }
