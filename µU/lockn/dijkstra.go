package lockn

// (c) Christian Maurer   v. 171026 - license see µU.go

// >>> Algorithm of Dijkstra
//     Cooperating Sequential Processes, 0 -> true, 1 -> false

import
  "time"
type
  dijkstra struct {
       nProcesses,
         favoured uint
       interested,
         critical []bool
                  }

func newD (n uint) LockerN {
  x := new(dijkstra)
  x.nProcesses = n
  x.interested, x.critical = make([]bool, n + 1), make([]bool, n)
  x.favoured = x.nProcesses
  return x
}

func (x *dijkstra) LockGoto (i uint) {
  x.interested[i] = true
L:
  if x.favoured != i {
    x.critical[i] = false
    if ! x.interested[x.favoured] {
      x.favoured = i
      goto L
    }
  }
  time.Sleep(1)
  x.critical[i] = true
  time.Sleep(1)
  for j := uint(0); j < x.nProcesses; j++ {
    if j != i && x.critical[j] {
      goto L
    }
  }
}

func (x *dijkstra) UnlockGoto (i uint) {
  x.favoured = (i + 1) % x.nProcesses
  x.interested[i] = false
  x.critical[i] = false
}

func (x *dijkstra) otherCritical (i uint) bool {
  for j := uint(0); j < x.nProcesses; j++ {
    if j != i {
      if x.critical[j] { return true }
    }
  }
  return false
}

func (x *dijkstra) Lock (i uint) {
  x.interested[i] = true
  for {
    for x.favoured != i {
      x.critical[i] = false
      if ! x.interested[x.favoured] {
        x.favoured = i
      }
      time.Sleep(1) // oder runtime.Gosched()
    }
    x.critical[i] = true
    if ! x.otherCritical (i) { break }
  }
}

func (x *dijkstra) Unlock (i uint) {
  x.favoured = (i + 1) % x.nProcesses
  x.interested[i] = false
  x.critical[i] = false
}
