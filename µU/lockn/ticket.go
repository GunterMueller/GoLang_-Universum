package lockn

// (c) Christian Maurer   v. 171024 - license see µU.go

// >>> Ticket-Algorithm using FetchAndAddUint32

import (
  . "µU/obj"
  . "µU/atomic"
)
type
  ticket struct {
   turn, ticket uint32
                }

func newT (n uint) LockerN {
  return new(ticket)
}

func (x *ticket) Lock (p uint) {
  t := FetchAndIncrement (&x.ticket)
  for t != x.turn {
    Gothing()
  }
}

func (x *ticket) Unlock (p uint) {
  FetchAndIncrement (&x.turn)
}
