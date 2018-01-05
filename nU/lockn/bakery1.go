package lockn

// (c) Christian Maurer   v. 171231 - license see nU.go

// >>> Bakery-Algorithm of Lamport, corrected version

type bakery1 struct {
  uint "Anzahl der beteiligten Prozesse"
  number []uint
  draws []bool
}

func (x *bakery1) max() uint {
  m := uint(0)
  for i := uint(1); i <= x.uint; i++ {
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
  if n < 2 { return nil }
  x := new(bakery1)
  x.uint = n
  x.number = make([]uint, n)
  x.draws = make([]bool, n)
  return x
}

func (x *bakery1) Lock (p uint) {
  if p >= x.uint { return }
  x.number[p] = 1
  x.number[p] = x.max() + 1
  for a := uint(1); a <= x.uint; a++ {
    if a != p {
      for x.number[a] > 0 && x.less (a, p) {
        nothing()
      }
    }
  }
}

func (x *bakery1) Unlock (p uint) {
  if p >= x.uint { return }
  x.number[p] = 0
}
