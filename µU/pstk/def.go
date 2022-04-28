package pstk

// (c) Christian Maurer   v. 220420 - license see µU.go

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
func New (a any) PersistentStack { return new_(a) }
