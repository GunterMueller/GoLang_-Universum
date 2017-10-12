package sort

// (c) Christian Maurer   v. 121125 - license see µu.go

import
  . "µu/obj"

// Pre: a is a slice of atomic variables or of objects of type Object.
// TODO Spec
func Sort (a []Any) { sort_(a) }
