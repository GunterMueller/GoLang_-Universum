package lockn

// (c) Christian Maurer   v. 171231 - license see nU.go

// >>> Algorithm of Habermann

type habermann struct {
  nProcesses, favoured uint
  interested, critical []bool
}

func newH (n uint) LockerN {
  if n < 2 { return nil }
  x := new(habermann)
  x.nProcesses = n
  x.favoured = 0
  x.interested = make([]bool, n + 1)
  x.critical = make([]bool, n)
  return x
}

func (x *habermann) Lock (p uint) {
  if p >= x.nProcesses { return }
  var (
    b uint
    someoneElseCritical, someoneElseInterested bool
  )
  for {
    x.interested[p] = true
    for {
      x.critical[p] = false
      b = x.favoured
      someoneElseInterested = false
      for b != p {
        someoneElseInterested = x.interested[b] || someoneElseInterested
        if b < x.nProcesses {
          b++
        } else {
          b = 1
        }
      }
      if ! someoneElseInterested {
        break
      }
    }
    x.critical[p] = true
    someoneElseCritical = false
    for a := uint(1); a <= x.nProcesses; a++ {
      if a != p {
        someoneElseCritical = someoneElseCritical || x.critical[a]
      }
    }
    if ! someoneElseCritical {
      break
    }
  }
  x.favoured = p
}

func (x *habermann) Unlock (p uint) {
  if p >= x.nProcesses { return }
  i := p
  for {
    i = i % x.nProcesses + 1
    if x.interested[i] || i == p {
      break
    }
  }
  x.favoured = i
  x.critical[p] = false
  x.interested[p] = false
}
