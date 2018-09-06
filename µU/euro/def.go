package euro

// (c) Christian Maurer   v. 180901 - license see µU.go

import (
  . "µU/obj"
  . "µU/add"
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
