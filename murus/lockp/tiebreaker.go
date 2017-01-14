package lockp

// (c) murus.org  v. 161212 - license see murus.go

// >>> Tiebreaker-Algorithm of Peterson

type
  tiebreaker struct {
         nProcesses uint
     achieved, last []uint
                    }

func newTb (n uint) LockerP {
  if n < 2 { return nil }
  L := new (tiebreaker)
  L.nProcesses = n
  L.achieved = make ([]uint, n)
  L.last = make ([]uint, n)
  return L
}

func (L *tiebreaker) Lock (p uint) {
  if p >= L.nProcesses { return }
  for e := uint(0); e < L.nProcesses - 1; e++ {
    L.achieved[p] = e
    L.last[e] = p
    for a := uint(0); a < L.nProcesses; a++ {
      if p != a {
        for e <= L.achieved[a] && p == L.last[e] { /* do nothing */ }
      }
    }
  }
}

func (L *tiebreaker) Unlock (p uint) {
  if p >= L.nProcesses { return }
  L.achieved[p] = 0
}
