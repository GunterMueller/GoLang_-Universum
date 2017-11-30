package lockp

// (c) Christian Maurer   v. 171013 - license see nU.go

// >>> Bakery-Algorithm of Lamport

type lockerPBakery struct {
  uint "number of processes involved"
  number []uint
  draws []bool
}

func (x *lockerPBakery) max() uint {
  m := uint(0)
  for i := uint(1); i <= x.uint; i++ {
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

func newB (n uint) LockerP {
  if n < 2 { return nil }
  x := new(lockerPBakery)
  x.uint = n
  x.number = make([]uint, n)
  x.draws = make([]bool, n)
  return x
}

func (x *lockerPBakery) Lock (p uint) {
  if p >= x.uint { return }
  x.draws[p] = true
  x.number[p] = x.max() + 1
  x.draws[p] = false
  for a := uint(1); a <= x.uint; a++ {
    for x.draws[a] { /* Null() */ }
    for x.number[a] > 0 && x.less (a, p) { /* Null() */ }
  }
}

func (x *lockerPBakery) Unlock (p uint) {
  if p >= x.uint { return }
  x.number[p] = 0
}
