package lockn

// (c) Christian Maurer   v. 171024 - license see µU.go

// >>> Bakery-Algorithm of Lamport

import
  . "µU/obj"
type
  lockerPBakery struct {
                       uint "number of processes"
                number []uint
                 draws []bool
                       }

func (x *lockerPBakery) max() uint {
  m := uint(0)
  for i := uint(0); i < x.uint; i++ {
    if x.number[i] > m {
      m = x.number[i]
    }
  }
  return m
}

func (x *lockerPBakery) less (i, k uint) bool {
  if x.number[i] < x.number[k] {
    return true
  }
  if x.number[i] == x.number[k] {
    return i < k
  }
  return false
}

func newB (n uint) LockerN {
  x := new(lockerPBakery)
  x.uint = n
  x.number = make([]uint, n)
  x.draws = make([]bool, n)
  return x
}

func (x *lockerPBakery) Lock (i uint) {
  x.draws[i] = true
  x.number[i] = x.max() + 1
  x.draws[i] = false
  for j := uint(0); j < x.uint; j++ {
    for x.draws[j] {
      Gothing()
    }
    for x.number[j] > 0 && x.less (j, i) {
      Gothing()
    }
  }
}

func (x *lockerPBakery) Unlock (i uint) {
  x.number[i] = 0
}
