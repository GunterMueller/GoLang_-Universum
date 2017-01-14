package euro

// (c) murus.org  v. 161216 - license see murus.go

import
  . "murus/obj"
type
  Euro interface {

  Editor
  Stringer
  Printer
  Valuator
  Val2 () (uint, uint)
  Set2 (uint, uint) bool
  RealVal () float64
  SetReal (r float64) bool
  Adder
  Operate (Factor, Divisor uint)
  ChargeInterest (p, n uint)
  Round (E Euro)
}
// Returns a new empty Euro.
func New() Euro { return newEuro() }
