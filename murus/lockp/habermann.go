package lockp

// (c) murus.org  v. 161212 - license see murus.go

// >>> Algorithm of Habermann

type
  habermann struct {
        nProcesses,
          favoured uint
         willenter,
          critical []bool
                   }

func newH (n uint) LockerP {
  if n < 2 { return nil }
  L := new (habermann)
  L.nProcesses = n
  L.favoured = 0
  L.willenter = make ([]bool, n + 1) // n + 1 ???
  L.critical = make ([]bool, n)
  return L
}

func (L *habermann) Lock (p uint) {
  if p >= L.nProcesses { return }
  var (
    b uint
    andererKritisch, andererEintrittswillig bool
  )
  for {
    L.willenter[p] = true
    for {
      L.critical[p] = false
      b = L.favoured
      andererEintrittswillig = false
      for b != p {
        andererEintrittswillig = L.willenter[b] || andererEintrittswillig
        if b < L.nProcesses {
          b++
        } else {
          b = 1
        }
      }
      if ! andererEintrittswillig {
        break
      }
    }
    L.critical[p] = true
    andererKritisch = false
    for a := uint(1); a <= L.nProcesses; a++ {
      if a != p {
        andererKritisch = andererKritisch || L.critical[a]
      }
    }
    if ! andererKritisch {
      break
    }
  }
  L.favoured = p
}

func (L *habermann) Unlock (p uint) {
  if p >= L.nProcesses { return }
  i := p
  for {
    i = i % L.nProcesses + 1
    if L.willenter[i] || i == p {
      break
    }
  }
  L.favoured = i
  L.critical[p] = false
  L.willenter[p] = false
}
