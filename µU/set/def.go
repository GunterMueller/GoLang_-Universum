package set

// (c) Christian Maurer   v. 201014 - license see µU.go

import
  . "µU/obj"
type
  Set interface { // ordered sets of elements, that are atomic or implement Object.

  Sorter
// Split is not yet implemented!

// Undocumented.
// (My work is so secret, that even I don't know what I'm doing.)
  Write (x0, x1, y, dy uint, f func (Any) string)
  Write1 (f func (Any) string)
}

// Pre: a is atomic or of a type implementing Object.
// Returns a new empty set for objects of the type of a.
func New (a Any) Set { return new_(a) }
