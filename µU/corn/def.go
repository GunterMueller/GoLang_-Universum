package corn

// (c) Christian Maurer   v. 240104 - license see µU.go

import
  "µU/buf"
type
  Cornet interface {

  buf.Buffer

// Returns the list of all objects of x.
  All() []any
}

// Pre: a is atomic or of a type implementing Object.
// Returns a new cornet for elements of the type of a.
func New (a any) Cornet { return new_(a) }

