package set

// (c) Christian Maurer   v. 170424 - license see µU.go

import
  . "µU/obj"
type
  Set interface {

  Object
  Sorter
  Iterator

// Undocumented.
// (My work is so secret, that even I don't know what I'm doing.)
  Write (x0, x1, y, dy uint, f func (Any) string)
  Write1 (f func (Any) string)
}

// Pre: a is atomic or of a type implementing Object.
// Returns a new empty set for objects of the type of a.
func New(a Any) Set { return new_(a) }
