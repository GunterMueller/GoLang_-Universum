package pstk

// (c) Christian Maurer   v. 210109 - license see µU.go

import (
  . "µU/obj"
  "µU/stk"
)
type
  PersistentStack interface { // Not to be used by concurrent processes !

  stk.Stack
  Persistor
}

// Returns a new empty persistent stack for objects of type a.
func New (a Any) PersistentStack { return new_(a) }
