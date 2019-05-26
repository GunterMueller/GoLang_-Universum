package lockn

// (c) Christian Maurer   v. 190323 - license see µU.go

// >>> Algorithm of Burns

import
  . "µU/atomic"
type
  burns struct {
     nProcesses,
       favoured uint
     interested []uint
                }

func newBurns (n uint) LockerN {
  x := new(burns)
  x.nProcesses = uint(n)
  x.interested = make([]uint, n + 1)
  x.favoured = x.nProcesses
  return x
}

func (x *burns) Lock (p uint) {
  x.interested[p] = 1
  Store (&x.favoured, p)
  var q uint
  for {
    for x.favoured != p {
      x.interested[p] = 0
      q = 1
      for q < x.nProcesses && (q == p || x.interested[q] == 0) {
        q++
      }
      if q >= x.nProcesses {
        x.interested[p] = 1
        Store (&x.favoured, p)
      }
    }
    q = 1
    for q < x.nProcesses && (q == p || x.interested[q] == 0) {
      q++
    }
    if q >= x.nProcesses {
      break
    }
  }
}

func (x *burns) Unlock (p uint) {
  x.interested[p] = 0
}
