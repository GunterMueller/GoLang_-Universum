package stk

// (c) murus.org  v. 170316 - license see murus.go

import
  . "murus/obj"
type
  Stack interface {

// a is the object on top of x, the stack below a is x before.
  Push (a Any)

// Returns true, iff there is no object on x.
// In the concurrent case, this value is not reliable,
// as another process could have pushed on object on the stack
// immediately after the call.
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
