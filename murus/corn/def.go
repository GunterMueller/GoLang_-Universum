package corn

// (c) Christian Maurer   v. 170316 - license see murus.go

import (
  . "murus/obj"
  "murus/qu"
)
type
  Cornet interface {

  qu.Queue
}

func New (a Any) Cornet { return new_(a) }
