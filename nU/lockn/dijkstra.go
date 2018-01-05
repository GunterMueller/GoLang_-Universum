package lockn

// (c) Christian Maurer   v. 171231 - license see nU.go

// >>> Algorithm of Dijkstra

type dijkstra struct {
  nProcesses, favoured uint
  critical, interested []bool
}

func newD (n uint) LockerN {
  if n < 2 { return nil }
  x := new(dijkstra)
  x.nProcesses = n
  x.critical, x.interested = make([]bool, n), make([]bool, n)
  return x
}

func (x *dijkstra) Lock (p uint) {
  if p >= x.nProcesses { return }
  x.critical[p] = true
  for {
    x.interested[p] = false
    for x.favoured != p {
      if ! x.critical[x.favoured] {
        x.favoured = p
      }
    }
    x.interested[p] = true
    someoneElseInterested := false
    for q := uint(1); q <= x.nProcesses; q++ {
      if q != p {
        someoneElseInterested = someoneElseInterested || x.interested[q]
      }
    }
    if ! someoneElseInterested {
      break
    }
  }
}

func (x *dijkstra) Unlock (p uint) {
  if p >= x.nProcesses { return }
  x.critical[p] = false
  x.interested[p] = false
  x.favoured = 0
}
