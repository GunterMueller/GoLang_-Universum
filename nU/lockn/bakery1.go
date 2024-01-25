package lockn

// (c) Christian Maurer   v. 190321 - license see nU.go

// >>> Bakery-Algorithm of Lamport, corrected version

import
  . "nU/obj"
type
  bakery1 struct {
                 uint "Anzahl der beteiligten Prozesse"
          number,
           draws []uint
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

func newBakery1 (n uint) LockerN {
  x := new(bakery1)
  x.uint = uint(n)
  x.number = make([]uint, n)
  x.draws = make([]uint, n)
  return x
}

func (x *bakery1) Lock (p uint) {
  x.number[p] = 1
  x.number[p] = x.max() + 1
  for q := uint(0); q < x.uint; q++ {
    if q != p {
      for x.number[q] > 0 && x.less (q, p) {
        Nothing()
      }
    }
  }
}

func (x *bakery1) Unlock (p uint) {
  x.number[p] = 0
}
