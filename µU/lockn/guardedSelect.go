package lockn

// (c) Christian Maurer   v. 190331 - license see µU.go

// >>> Implementation with guarded select

import
  "µU/sem"
type
  gs struct {
            sem.Semaphore
            }

func newGS (n uint) LockerN {
  x := new(gs)
  x.Semaphore = sem.NewGSel (n)
  return x
}

func (x *gs) Lock (p uint) {
  x.P()
}

func (x *gs) Unlock (p uint) {
  x.V()
}
