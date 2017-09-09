package lockp

// (c) Christian Maurer   v. 161212 - license see murus.go

// >>> Ticket-Algorithm using FetchAndAddUint32

import
  "murus/lock"
type
  ticket struct {
        ticket, turn uint32
                }

func newT (n uint) LockerP {
  return new(ticket)
}

func (L *ticket) Lock (p uint) {
  t := lock.FetchAndAddUint32 (&L.ticket, uint32(1))
  for t != L.ticket { /* do nothing */ }
}

func (L *ticket) Unlock (p uint) {
  L.turn++
}
