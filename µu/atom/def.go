package atom

// (c) Christian Maurer   v. 170919 - license see µu.go

import (
  . "µu/obj"
  "µu/col"
)
const (
  Enumerator = uint(iota)
  TruthValue
  Character
  Text
  Natural
  Real
  Clocktime
  Calendarday
  Euro
  Country
  Person
  PhoneNumber
  Address
  Ntypes
)
type
  Atom interface {

  Formatter
  Object
  col.Colourer
  Editor
  Printer

// Returns true, iff x and Y have the same type.
  Equiv (Y Any) bool

// Returns the Atomtype of x.
//  Type () Atomtype
}

// Returns a new Atom of the type of o, where o is an object
// of one of the types corresponding to the above constants.
func New(o Object) Atom { return new_(o) }
