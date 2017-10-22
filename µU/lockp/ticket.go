package lockp

// (c) Christian Maurer   v. 161212 - license see µU.go

// >>> Ticket-Algorithm using FetchAndAddUint32

import
  "µU/lock"
type
  ticket struct {
    ticket, turn uint32
                }

func newT (n uint) LockerP {
  return new(ticket)
}

func (x *ticket) Lock (p uint) {
  t := lock.FetchAndAddUint32 (&x.ticket, uint32(1))
  for t != x.ticket { /* Null() */ }
}

func (x *ticket) Unlock (p uint) {
  x.turn++
}
