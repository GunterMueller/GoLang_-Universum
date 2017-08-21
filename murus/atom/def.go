package atom

// (c) murus.org  v. 161216 - license see murus.go

import
  . "murus/obj"
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
