package lockp

// (c) murus.org  v. 161212 - license see murus.go

// >>> Algorithm of Dijkstra

type
  dijkstra struct {
       nProcesses,
         favoured uint
         critical,
       interested []bool
                  }

func newD (n uint) LockerP {
  if n < 2 { return nil }
  L := new (dijkstra)
  L.nProcesses = n
  L.critical, L.interested = make ([]bool, n), make ([]bool, n)
  return L
}

func (L *dijkstra) Lock (p uint) {
  if p >= L.nProcesses { return }
  L.critical[p] = true
  for {
    L.interested[p] = false
    for L.favoured != p {
      if ! L.critical[L.favoured] {
        L.favoured = p
      }
    }
    L.interested[p] = true
    someoneElseInterested := false
    for q := uint(1); q <= L.nProcesses; q++ {
      if q != p {
        someoneElseInterested = someoneElseInterested || L.interested[q]
      }
    }
    if ! someoneElseInterested {
      break
    }
  }
}

func (L *dijkstra) Unlock (p uint) {
  if p >= L.nProcesses { return }
  L.critical[p] = false
  L.interested[p] = false
  L.favoured = 0
}
