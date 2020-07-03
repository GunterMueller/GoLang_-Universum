package corn

// (c) Christian Maurer   v. 200120 - license see µU.go

import (
  . "µU/obj"
  "µU/buf"
)
type
  Cornet interface {

  buf.Buffer
}

// Pre: a is atomic or of a type implementing Object.
// Returns a new cornet for elements of the type of a.
func New (a Any) Cornet { return new_(a) }
