package stk

// (c) Christian Maurer   v. 201103 - license see µU.go

import
  . "µU/obj"
type
  Stack interface { // Not to be used by concurrent processes !

// Returns true, iff there is no element on x.
  Empty() bool

// Pre: a is atomic type or of a type implementing object.
// a is the element on top of x, the stack below a is x before.
  Push (a Any)

// Returns nil, if x is empty, otherwise a copy of
// the element on top of x. That element is removed,
// i.e. x now equals the stack below x before.
  Pop() Any
}

// Returns a new empty stack for objects of type a.
func New (a Any) Stack { return new_(a) }
