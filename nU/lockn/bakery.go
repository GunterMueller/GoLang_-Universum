package lockn

// (c) Christian Maurer   v. 171231 - license see nU.go

// >>> Bakery-Algorithm of Lamport

type lockerNBakery struct {
  uint "Anzahl der beteiligten Prozesse"
  number []uint
  draws []bool
}

func (x *lockerNBakery) max() uint {
  m := uint(0)
  for i := uint(1); i <= x.uint; i++ {
    if x.number[i] > m {
      m = x.number[i]
    }
  }
  return m
}

func (x *lockerNBakery) less (i, k uint) bool {
  if x.number[i] < x.number[k] {
    return true
  }
  if x.number[i] == x.number[k] {
    return i < k
  }
  return false
}

func newB (n uint) LockerN {
  if n < 2 { return nil }
  x := new(lockerNBakery)
  x.uint = n
  x.number = make([]uint, n)
  x.draws = make([]bool, n)
  return x
}

func (x *lockerNBakery) Lock (p uint) {
  if p >= x.uint { return }
  x.draws[p] = true
  x.number[p] = x.max() + 1
  x.draws[p] = false
  for a := uint(1); a <= x.uint; a++ {
    for x.draws[a] {
      nothing()
    }
    for x.number[a] > 0 && x.less (a, p) {
      nothing()
    }
  }
}

func (x *lockerNBakery) Unlock (p uint) {
  if p >= x.uint { return }
  x.number[p] = 0
}
