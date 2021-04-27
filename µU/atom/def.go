package atom

// (c) Christian Maurer   v. 210415 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
const (
  Enumerator = uint(iota)
  TruthValue
  Text
  Natural
  Real
  Clocktime
  Calendarday
  Euro
  Person
  PhoneNumber
  Address
  Country
  Ntypes
)
type
  Atom interface {

  Object
  Formatter
  col.Colourer
  Editor
  Printer
  Stringer

// Returns true, iff x and y have the same type.
  Equiv (y Any) bool
}

// Returns a new Atom of the type of o, where o is an object
// of a type corresponding to one of the above constants.
func New (o Object) Atom { return new_(o) }
