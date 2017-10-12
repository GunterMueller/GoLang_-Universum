package corn

// (c) Christian Maurer   v. 170316 - license see µu.go

import (
  . "µu/obj"
  "µu/qu"
)
type
  Cornet interface {

  qu.Queue
}

func New (a Any) Cornet { return new_(a) }
