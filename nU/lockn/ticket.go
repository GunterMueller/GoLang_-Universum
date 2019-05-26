package lockn

// (c) Christian Maurer   v. 190323 - license see nU.go

// >>> Ticket-Algorithm using FetchAndIncrement

import (
  . "nU/obj"
  . "nU/atomic"
)
type
  ticket struct {
   turn, ticket uint
                }

func newTicket (n uint) LockerN {
  return new(ticket)
}

func (x *ticket) Lock (p uint) {
  t := FetchAndIncrement (&x.ticket)
  for t != x.turn {
    Nothing()
  }
}

func (x *ticket) Unlock (p uint) {
  FetchAndIncrement (&x.turn)
}
