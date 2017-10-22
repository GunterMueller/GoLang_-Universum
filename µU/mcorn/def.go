package mcorn

// (c) Christian Maurer   v. 170320 - license see µU.go

import (
  . "µU/obj"
  "µU/corn"
)
type
  MCornet interface { // Do not call Empty or Num, as their values are not reliable !
                      // A process caling Get is block, until x is not empty.

  corn.Cornet
}

// Pre: a is atomic or of a type implementieng Object.
// Returns a new cornet for elements of the type of a
// to be used by concurrent processes.
func New (a Any) MCornet { return new_(a) }
