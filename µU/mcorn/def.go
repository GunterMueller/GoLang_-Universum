package mcorn

// (c) Christian Maurer   v. 220420 - license see µU.go

import
  "µU/corn"
type
  MCornet interface { // Do not call Empty or Num, as their values are not reliable !
                      // A process calling Get is blocked, until x is not empty.
  corn.Cornet
}

// Pre: a is atomic or of a type implementing Object.
// Returns a new cornet for elements of the type of a
// to be used by concurrent processes.
func New (a any) MCornet { return new_(a) }
