package lockn

// (c) Christian Maurer   v. 190326 - license see µU.go

// >>> Algorithm of Habermann

import (
  . "µU/atomic"
  . "µU/obj"
)
type
  habermann struct {
        nProcesses,
          favoured uint
        interested,
          critical []uint
                   }

func newHabermann (n uint) LockerN {
  x := new(habermann)
  x.nProcesses = uint(n)
  x.favoured = 0
  x.interested = make([]uint, n)
  x.critical = make([]uint, n)
  return x
}

func (x *habermann) Lock (p uint) {
  for {
    Store (&x.interested[p], 1)
    for {
      Store (&x.critical[p], 0)
      f := x.favoured
      otherInterested := uint(0)
      for f != p {
        otherInterested += x.interested[f]
        if f + 1 < x.nProcesses {
          f++
        } else {
          f = 0
        }
       }
      if otherInterested == 0 {
        break
      }
      Nothing()
    }
    Store (&x.critical[p], 1)
    otherCritical := uint(0)
    for q := uint(0); q < x.nProcesses; q++ {
      if q != p {
        otherCritical += x.critical[q]
      }
    }
    if otherCritical == 0 {
      break
    }
    Nothing()
  }
  Store (&x.favoured, p)
}

func (x *habermann) Unlock (p uint) {
  f := p
  for {
    f = (f + 1) % x.nProcesses
    if x.interested[p] == 1 || f == p {
      break
    }
  }
  Store (&x.favoured, f)
  Store (&x.interested[p], 0)
  Store (&x.critical[p], 0)
}
