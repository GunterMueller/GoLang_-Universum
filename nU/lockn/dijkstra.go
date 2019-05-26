package lockn

// (c) Christian Maurer   v. 190323 - license see nU.go

// >>> Algorithm of Dijkstra
//     Cooperating Sequential Processes, 0 -> true, 1 -> false

import (
  . "nU/atomic"
  . "nU/obj"
)
type
  dijkstra struct {
       nProcesses,
         favoured uint
       interested,
         critical []uint
                  }

func newDijkstra (n uint) LockerN {
  x := new(dijkstra)
  x.nProcesses = uint(n)
  x.interested, x.critical = make([]uint, n + 1), make([]uint, n)
  x.favoured = x.nProcesses
  return x
}

func (x *dijkstra) LockGoto (p uint) {
  Store (&x.interested[p], 1)
L:
  if x.favoured != p {
    Store (&x.critical[p], 0)
    if x.interested[x.favoured] == 0 {
      x.favoured = p
      goto L
    }
  }
  Nothing()
  Store (&x.critical[p], 1)
  Nothing() // iff more cpus are working, otherwise no mutual exclusion
  for j := uint(0); j < x.nProcesses; j++ {
    if j != p && x.critical[j] == 1 {
      goto L
    }
  }
}

func (x *dijkstra) UnlockGoto (p uint) {
  x.Unlock (p)
}

func (x *dijkstra) otherCritical (p uint) bool {
  for j := uint(0); j < x.nProcesses; j++ {
    if j != p && x.critical[j] == 1 {
      return true
    }
  }
  return false
}

func (x *dijkstra) Lock (p uint) {
  Store (&x.interested[p], 1)
  for {
    for x.favoured != p {
      Store (&x.critical[p], 0)
      if x.interested[x.favoured] == 0 {
        Store (&x.favoured, p)
      }
      Nothing()
    }
    Store (&x.critical[p], 1)
    if ! x.otherCritical (p) {
      break
    }
  }
}

func (x *dijkstra) Unlock (p uint) {
  Store (&x.favoured, (p + 1) % x.nProcesses)
  Store (&x.interested[p], 0)
  Store (&x.critical[p], 0)
}
