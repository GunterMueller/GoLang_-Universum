package lockn

// (c) Christian Maurer   v. 171024 - license see µU.go

// >>> Bakery-Algorithm of Lamport, corrected version

import
  . "µU/obj"
type
  bakery1 struct {
                 uint "number of processes"
          number []uint
           draws []bool
                 }

func (x *bakery1) max() uint {
  m := uint(0)
  for i := uint(0); i < x.uint; i++ {
    if x.number[i] > m {
      m = x.number[i]
    }
  }
  return m
}

func (x *bakery1) less (i, k uint) bool {
  if x.number[i] < x.number[k] {
    return true
  }
  if x.number[i] == x.number[k] {
    return i < k
  }
  return false
}

func newB1 (n uint) LockerN {
  x := new(bakery1)
  x.uint = n
  x.number = make([]uint, n)
  x.draws = make([]bool, n)
  return x
}

func (x *bakery1) Lock (i uint) {
  x.number[i] = 1
  x.number[i] = x.max() + 1
  for j := uint(0); j < x.uint; j++ {
    if j != i {
      for x.number[j] > 0 && x.less (j, i) {
        Gothing()
      }
    }
  }
}

func (x *bakery1) Unlock (i uint) {
  x.number[i] = 0
}
