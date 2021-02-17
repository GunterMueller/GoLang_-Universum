package set

// (c) Christian Maurer   v. 210212 - license see µU.go

import
  . "µU/obj"
type
  Set interface { // ordered sets of elements, that are atomic or implement Object.

  Collector

  ExGeq (a Any) bool
}

// Pre: a is atomic or of a type implementing Object.
// Returns a new empty set for objects of the type of a.
func New (a Any) Set { return new_(a) }
