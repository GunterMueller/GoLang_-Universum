package lockn

// (c) Christian Maurer   v. 171026 - license see µU.go

// >>> Algorithm of Habermann

import
  "time"
type
  habermann struct {
        nProcesses,
          favoured uint
        interested,
          critical []bool
                   }

func newH (n uint) LockerN {
  x := new(habermann)
  x.nProcesses = n
  x.favoured = 0
  x.interested = make([]bool, n)
  x.critical = make([]bool, n)
  return x
}

func (x *habermann) Lock (i uint) {
  for {
    x.interested[i] = true
    for {
      x.critical[i] = false
      f := x.favoured
      otherInterested := false
      for f != i {
        otherInterested = x.interested[f] || otherInterested
        if f + 1 < x.nProcesses {
          f++
        } else {
          f = 0
        }
      }
      if ! otherInterested { break }
      time.Sleep(1) // oder runtime.Gosched()
    }
    x.critical[i] = true
    otherCritical := false
    for j := uint(0); j < x.nProcesses; j++ {
      if j != i {
        otherCritical = otherCritical || x.critical[j]
      }
    }
    if ! otherCritical { break }
    time.Sleep(1) // oder runtime.Gosched()
  }
  x.favoured = i
}

func (x *habermann) Unlock (i uint) {
  f := i
  for {
    f = (f + 1) % x.nProcesses
    if x.interested[i] || f == i { break }
  }
  x.favoured = f
  x.interested[i], x.critical[i] = false, false
}
