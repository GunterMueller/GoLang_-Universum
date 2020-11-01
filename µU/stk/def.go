package stk

// (c) Christian Maurer   v. 201030 - license see µU.go

import
  . "µU/obj"
type
  Stack interface { // Not to be used by concurrent processes !

// a is the object on top of x, the stack below a is x before.
  Push (a Any)

// Returns true, iff there is no object on x.
  Empty() bool

// If x was empty, nothing has happened.
// The object on top of x is removed;
// i.e. x now equals the stack below x before.
  Pop()

// Returns nil, if x is empty, otherwise a copy of the object
// on top of x. x is not changed.
  Top() Any
}

// Returns a new empty stack for objects of type a.
func New(a Any) Stack { return new_(a) }
