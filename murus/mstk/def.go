package mstk

// (c) murus.org  v. 130328 - license see murus.go

import
  . "murus/obj"
type
  MStack interface {

// x.Top() equals a and the result of x.Push(); x.Pop() is x; i.e.
// a is the object on top of x, the stack below a is x before.
  Push (a Any)

// The object on top of x is removed;
// i.e. x now equals the stack below that object before.
// The calling process was blocked, until x was not empty.
  Pop()

// Returns a copy of the object on top of x. x is not changed.
// The calling process was blocked, until x was not empty.
  Top() Any
}
// Pre: a is atomic or of a type implementing Object.
// Returns a new stack for elements of the type of a
// to be used by concurrent processes.
func New(a Any) MStack { return newMstk(a) }
