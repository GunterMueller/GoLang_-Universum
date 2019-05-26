package lockn

// (c) Christian Maurer   v. 190321 - license see nU.go

// >>> Bakery-Algorithm of Lamport

import (
  . "nU/atomic"
  . "nU/obj"
)
type
  bakery struct {
                uint "number of processes"
         number,
          draws []uint
                }

func (x *bakery) max() uint {
  m := uint(0)
  for i := uint(0); i < x.uint; i++ {
    if x.number[i] > m {
      m = x.number[i]
    }
  }
  return m
}

func (x *bakery) less (i, k uint) bool {
  if x.number[i] < x.number[k] {
    return true
  }
  if x.number[i] == x.number[k] {
    return i < k
  }
  return false
}

func newBakery (n uint) LockerN {
  x := new(bakery)
  x.uint = uint(n)
  x.number = make([]uint, n)
  x.draws = make([]uint, n)
  return x
}

func (x *bakery) Lock (p uint) {
  x.draws[p] = 1
  Store (&x.number[p], x.max() + 1)
  x.draws[p] = 0
  for j := uint(0); j < x.uint; j++ {
    for x.draws[j] == 1 {
      Nothing()
    }
    for x.number[j] > 0 && x.less (j, p) {
      Nothing()
    }
  }
}

func (x *bakery) Unlock (p uint) {
  Store (&x.number[p], 0)
}
