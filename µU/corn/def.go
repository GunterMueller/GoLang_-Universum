package corn

// (c) Christian Maurer   v. 170316 - license see µU.go

import (
  . "µU/obj"
  "µU/qu"
)
type
  Cornet interface {

  qu.Queue
}

func New (a Any) Cornet { return new_(a) }
