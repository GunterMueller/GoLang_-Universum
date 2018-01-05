package lockn

// (c) Christian Maurer   v. 171231 - license see nU.go

// >>> Tiebreaker-Algorithm of Peterson

type tiebreaker struct {
  uint "Anzahl der beteiligten Prozesse"
  achieved, last []uint
}

func newTb (n uint) LockerN {
  if n < 2 { return nil }
  x := new(tiebreaker)
  x.uint = n
  x.achieved = make([]uint, n)
  x.last = make([]uint, n)
  return x
}

func (x *tiebreaker) Lock (p uint) {
  if p >= x.uint { return }
  for e := uint(0); e < x.uint - 1; e++ {
    x.achieved[p] = e
    x.last[e] = p
    for a := uint(0); a < x.uint; a++ {
      if p != a {
        for e <= x.achieved[a] && p == x.last[e] {
          nothing()
        }
      }
    }
  }
}

func (x *tiebreaker) Unlock (p uint) {
  if p >= x.uint { return }
  x.achieved[p] = 0
}
