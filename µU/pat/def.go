package pat

// (c) Christian Maurer   v. 221003 - license see µU.go

import
  . "µU/obj"
type
  PersonAddressTelMail interface { // Person, Adress and TelMail

  TeXer
  Rotator

  Sub (y any) bool
}

// Returns a new empty triple of Person, Address and TelMail.
func New() PersonAddressTelMail { return new_() }
