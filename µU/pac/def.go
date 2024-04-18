package pac

// (c) Christian Maurer   v. 240408 - license see µU.go

import
  . "µU/obj"
type
  PersonAddressContact interface { // person, address and contact

  Editor
  Stringer 
  TeXer
  Rotator
// Pre: y is of type PersonAddressContact
// Returns true, iff x is a part of y.
  Sub (y any) bool
}

// Returns a new empty triple of person, address and contact.
func New() PersonAddressContact { return new_() }
