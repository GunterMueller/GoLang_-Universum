package euro

// (c) Christian Maurer   v. 201009 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Euro interface {

  Object
  col.Colourer
  Editor
  Stringer
  Printer
  Valuator
  Val2() (uint, uint)
  Set2 (uint, uint) bool
  RealVal () float64
  SetReal (r float64) bool
  Adder
  Operate (Factor, Divisor uint)
  ChargeInterest (p, n uint)
  Round (E Euro)
}

// Returns a new empty Euro.
func New() Euro { return new_() }

// Returns a new Euro with Val2 = (e,c).
func New2 (e, c uint) Euro { return new2(e,c) }
