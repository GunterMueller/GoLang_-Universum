package lockn

// (c) Christian Maurer   v. 190323 - license see nU.go

// >>> Algorithm of J. G. DeBruijn, J. G.

import (
  . "nU/atomic"
  . "nU/obj"
)
type
  deBruijn struct {
                  uint "number of processes"
         favoured uint
             flag []uint
                  }

func newDeBruijn (n uint) LockerN {
  x := new(deBruijn)
  x.uint = uint(n)
  x.flag = make([]uint, x.uint)
  return x
}

func (x *deBruijn) test (p uint) bool {
  for q := uint(0); q < x.uint; q++ {
    if q != p {
      if x.flag[q] == active {
        return false
      }
    }
  }
  return true
}

func (x *deBruijn) Lock (p uint) {
  for {
    Store (&x.flag[p], requesting)
    q := x.favoured
    for q != p {
      if x.flag[q] != passive {
        q = x.favoured
      } else {
        q = (q + x.uint - 1) % x.uint
      }
      Nothing()
    }
    Store (&x.flag[p], active)
    if x.test (p) {
      break
    }
    Nothing()
  }
}

func (x *deBruijn) Unlock (p uint) {
  if x.flag[x.favoured] == passive || x.favoured == p {
    Store (&x.favoured, (x.favoured + x.uint - 1) % x.uint)
  }
  Store (&x.flag[p], passive)
}
