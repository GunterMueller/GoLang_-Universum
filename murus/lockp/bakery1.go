package lockp

// (c) murus.org  v. 161212 - license see murus.go

// >>> Bakery-Algorithm of Lamport, corrected version

type
  bakery1 struct {
      nProcesses uint
          number []uint
           draws []bool
                 }

func (L *bakery1) max() uint {
  m := uint(0)
  for i := uint(1); i <= L.nProcesses; i++ {
    if L.number[i] > m {
      m = L.number[i]
    }
  }
  return m
}

func (L *bakery1) less (i, k uint) bool {
  if L.number[i] < L.number[k] {
    return true
  }
  if L.number[i] == L.number[k] {
    return i < k
  }
  return false
}

func newB1 (n uint) LockerP {
  if n < 2 { return nil }
  L := new (bakery1)
  L.nProcesses = n
  L.number = make ([]uint, n)
  L.draws = make ([]bool, n)
  return L
}

func (L *bakery1) Lock (p uint) {
  if p >= L.nProcesses { return }
  L.number[p] = 1
  L.number[p] = L.max() + 1
  for a := uint(1); a <= L.nProcesses; a++ {
    if a != p {
      for L.number[a] > 0 && L.less (a, p) { /* nichts */ }
    }
  }
}

func (L *bakery1) Unlock (p uint) {
  if p >= L.nProcesses { return }
  L.number[p] = 0
}
