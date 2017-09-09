package sort

// (c) Christian Maurer   v. 121125 - license see murus.go

import
  . "murus/obj"

// Pre: a is a slice of atomic variables or of objects of type Object.
// TODO Spec
func Sort (a []Any) { sort_(a) }
