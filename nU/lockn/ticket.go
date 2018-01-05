package lockn

// (c) Christian Maurer   v. 171231 - license see nU.go

// >>> Ticket-Algorithm using FetchAndAddUint32

import "nU/atomic"

type ticket struct {
  ticket, turn uint32
}

func newT (n uint) LockerN {
  return new(ticket)
}

func (x *ticket) Lock (p uint) {
  t := atomic.FetchAndAdd (&x.ticket, uint32(1))
  for t != x.ticket {
    nothing()
  }
}

func (x *ticket) Unlock (p uint) {
  x.turn++
}
