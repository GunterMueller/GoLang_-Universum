package lockn

// (c) Christian Maurer   v. 190815 - license see µU.go

// >>> Algorithm of Dijkstra
//     Cooperating Sequential Processes, 0 -> true, 1 -> false

import (
  . "µU/obj"
  . "µU/atomic"
)
type
  dijkstraGoto struct {
           nProcesses,
             favoured uint
           interested,
             critical []uint
                      }

func newDijkstraGoto (n uint) LockerN {
  x := new(dijkstraGoto)
  x.nProcesses = uint(n)
  x.interested, x.critical = make([]uint, n + 1), make([]uint, n)
  x.favoured = x.nProcesses
  return x
}

func (x *dijkstraGoto) Lock (p uint) {
  Store (&x.interested[p], 1)
L:
  if x.favoured != p {
    Store (&x.critical[p], 0)
    if x.interested[x.favoured] == 0 {
      Store (&x.favoured, p)
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

func (x *dijkstraGoto) Unlock (p uint) {
  Store (&x.favoured, (p + 1) % x.nProcesses)
  Store (&x.interested[p], 0)
  Store (&x.critical[p], 0)
}
