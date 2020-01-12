package lockn

// (c) Christian Maurer   v. 190815 - license see µU.go

// >>> Tiebreaker-Algorithm of Peterson

import (
  . "µU/obj"
  . "µU/atomic"
)
type
  tiebreaker struct {
                    uint "number of processes"
     achieved, last []uint
                    }

func newTiebreaker (n uint) LockerN {
  if n < 2 { return nil }
  x := new(tiebreaker)
  x.uint = n
  x.achieved = make([]uint, n)
  x.last = make([]uint, n)
  return x
}

func (x *tiebreaker) Lock (p uint) {
  for q := uint(1); q < x.uint; q++ {
    Store (&x.achieved[p], q)
    Store (&x.last[q], p)
    for k := uint(0); k < x.uint; k++ {
      if p != k {
        for q <= x.achieved[k] && p == x.last[q] {
          Nothing()
        }
      }
    }
  }
}

func (x *tiebreaker) Unlock (p uint) {
  Store (&x.achieved[p], 0)
}
