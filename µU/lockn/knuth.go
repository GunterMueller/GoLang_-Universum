package lockn

// (c) Christian Maurer   v. 190816 - license see µU.go

// >>> Knuth, D. E.: Additional Comments on a Problem in Concurrent Programming Control.
//     CACM 9 (1966) 321-322

import (
  . "µU/obj"
  . "µU/atomic"
)
const (
  passive = iota
  requesting
  active
)
type
  knuth struct {
               uint "number of processes"
      favoured uint
          flag []uint
               }

func newKnuth (n uint) LockerN {
  x := new(knuth)
  x.uint = n
  x.favoured = x.uint
  x.flag = make([]uint, x.uint + 1)
  return x
}

func (x *knuth) test (p uint) bool {
  for q := uint(0); q < x.uint; q++ {
    if q != p {
      if x.flag[q] == active {
        return false
      }
    }
  }
  return true
}

func (x *knuth) Lock (p uint) {
  for {
    Store (&x.flag[p], requesting)
    q := x.favoured
    for q != p {
      if x.flag[q] == passive {
        q = (q + x.uint - 1) % x.uint
      } else {
        q = x.favoured
      }
      Nothing()
    }
    Store (&x.flag[p], active)
    if x.test (p) {
      break
    }
    Nothing()
  }
  Store (&x.favoured, p)
}

func (x *knuth) Unlock (p uint) {
  Store (&x.favoured, (p + x.uint - 1) % x.uint)
  Store (&x.flag[p], passive)
}
