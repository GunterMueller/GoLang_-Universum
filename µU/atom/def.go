package atom

// (c) Christian Maurer   v. 210410 - license see µU.go

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
//  PersonAddress
  Ntypes
)
type
  Atom interface {

  Formatter
  Object
  col.Colourer
  Editor
  Printer
  Stringer

  Size() (uint, uint)

// Returns true, iff x and Y have the same type.
  Equiv (Y Any) bool

// Returns the type of x (<= Ntypes).
  Type() uint

// Returns the object given in the call of New.
  Obj() Object

  Selected (l, c uint) bool
}

// Returns a new Atom of the type of o, where o is an object
// of one of the types corresponding to the above constants.
func New (o Object) Atom { return new_(o) }
