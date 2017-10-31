package lockn

// (c) Christian Maurer   v. 171022 - license see µU.go

// >>> Tiebreaker-Algorithm of Peterson

import
  . "µU/obj"
type
  tiebreaker struct {
                    uint "number of processes"
     achieved, last []uint
                    }

func newTb (n uint) LockerN {
  if n < 2 { return nil }
  x := new(tiebreaker)
  x.uint = n
  x.achieved = make([]uint, n)
  x.last = make([]uint, n)
  return x
}

func (x *tiebreaker) Lock (i uint) {
  for j := uint(1); j < x.uint; j++ {
    x.achieved[i] = j
    x.last[j] = uint(i)
    for k := uint(0); k < x.uint; k++ {
      if i != k {
        for j <= x.achieved[k] && i == x.last[j] {
          Gothing()
        }
      }
    }
  }
}

func (x *tiebreaker) Unlock (i uint) {
  x.achieved[i] = 0
}
