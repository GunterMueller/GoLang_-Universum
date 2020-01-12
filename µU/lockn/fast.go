package lockn

// (c) Christian Maurer   v. 190815 - license see µU.go

// >>> fast algorithm
// see Lamport, L.: A fast mutual exclusion algorithm. ACM TOCS 5 (1987), 1-11

import (
  . "µU/obj"
  . "µU/atomic"
)

type fast struct {
  uint "Anzahl der beteiligten Prozesse"
  interested []uint
  gate, gate1 uint
}

func newFast (n uint) LockerN {
  x := new(fast)
  x.uint = n
  x.interested = make([]uint, n)
  x.gate, x.gate1 = x.uint, x.uint
  return x
}

func (x *fast) Lock (p uint) {
  start:
  Store (&x.interested[p], 1)
  Store (&x.gate, p)
  if x.gate1 != x.uint {
    Store (&x.interested[p], 0)
    for x.gate1 != x.uint { Nothing() }
    goto start
  }
  Store (&x.gate1, p)
  if x.gate != p {
    Store (&x.interested[p], 0)
    for i := uint(0); i < x.uint; i++ {
      for x.interested[i] == 1 { Nothing() }
    }
    if x.gate1 != p {
      for x.gate1 != x.uint { Nothing() }
      goto start
    }
  }
}

func (x *fast) Unlock (p uint) {
  Store (&x.gate1, x.uint)
  Store (&x.interested[p], 0)
}
