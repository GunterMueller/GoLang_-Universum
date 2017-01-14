package set

// (c) murus.org  v. 161216 - license see murus.go

import
  . "murus/obj"
type
  Set interface {

  Object
  Sorter
  Collector
  Iterator

// Undocumented.
// (My work is so secret, that even I don't know what I'm doing.)
  Write (x0, x1, y, dy uint, f func (Any) string)
  Write1 (f func (Any) string)
}
// Pre: a is atomic or of a type implementing Object.
// Returns a new empty set for objects of the type of a.
func New(a Any) Set { return newSet(a) }
