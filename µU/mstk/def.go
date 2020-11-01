package mstk

// (c) Christian Maurer   v. 201030 - license see µU.go

import (
  . "µU/obj"
  "µU/stk"
)
type
  MStack interface { // To be used by concurrent processes, but:
                     // do not call Empty as it's value is not reliable !
                     // A process calling Pop or Top is blocked, until x is not empty.
  stk.Stack
}

// Pre: a is atomic or of a type implementing Object.
// Returns a new stack for elements of the type of a
// to be used by concurrent processes.
func New(a Any) MStack { return new_(a) }
