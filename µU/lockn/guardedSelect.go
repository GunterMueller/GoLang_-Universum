package lockn

// (c) Christian Maurer   v. 171105 - license see µU.go

// >>> Ticket-Algorithm using FetchAndAddUint32

import
  "µU/sem"
type
   guardedSelect struct {
                        sem.Semaphore
                        }

func newGS (n uint) LockerN {
  x := new(guardedSelect)
  x.Semaphore = sem.NewGSel (n)
  return x
}

func (x *guardedSelect) Lock (p uint) {
  x.Semaphore.P()
}

func (x *guardedSelect) Unlock (p uint) {
  x.Semaphore.V()
}
