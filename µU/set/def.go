package set

// (c) Christian Maurer   v. 240318 - license see µU.go

import
  . "µU/obj"
type
  Set interface {

  Equaler
  Collector

// My work is so secret, that even I don't know what I'm doing.
  Write (x0, x1, y, dy uint, f func (any) string)
  Write1 (f func (any) string)
}

// Pre: a is atomic or of a type implementing Object.
// Returns a new empty set for objects of the type of a.
func New (a any) Set { return new_(a) }
