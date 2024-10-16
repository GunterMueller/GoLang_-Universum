package cmon

// (c) Christian Maurer   v. 171019 - license see µU.go

import (
  "sync"
  . "µU/obj"
  "µU/perm"
  "µU/errh"
)
type
  monitor struct {
                 uint "number of monitor functions"
                 sync.Mutex "monitor entry queue"
               s []sync.Mutex "condition variable queues"
              ns []uint "numbers of processes blocked on s"
               u sync.Mutex "urgent queue"
              nu uint "number of processes blocked on urgent"
                 NFuncSpectrum "monitor functions"
                 CondSpectrum "conditions"
                 perm.Permutation "indeterminism"
                 }

func new_(n uint, f NFuncSpectrum, c CondSpectrum) Monitor {
  if n == 0 { return nil }
  x := new(monitor)
  x.uint = n
  x.s = make([]sync.Mutex, x.uint)
  for i := uint(0); i < x.uint; i++ {
    x.s[i].Lock()
  }
  x.ns = make([]uint, x.uint)
  x.u.Lock()
  x.NFuncSpectrum, x.CondSpectrum = f, c
  x.Permutation = perm.New (x.uint)
  return x
}

func (x *monitor) chk (s string, i uint) {
  if i >= x.uint {
    errh.Error2 ("i", i, "x.uint", x.uint)
    return
    WrongUintParameterPanic (s, x, i)
  }
}

func (x *monitor) wait (i uint) {
  x.chk ("Wait", i)
  x.ns[i]++
  if x.nu > 0 {
    x.u.Unlock()
  } else {
    x.Mutex.Unlock()
  }
  x.s[i].Lock()
  x.ns[i]--
}

func (x *monitor) Blocked (i uint) uint {
  x.chk ("Blocked", i)
  return x.ns[i]
}

func (x *monitor) Awaited (i uint) bool {
  x.chk ("Awaited", i)
  return x.ns[i] > 0
}

func (x *monitor) signal (i uint) {
  x.chk ("Signal", i)
  if x.ns[i] > 0 {
    x.nu++
    x.s[i].Unlock()
    x.u.Lock()
    x.nu--
  }
}

func (x *monitor) F (i uint) uint {
  x.chk ("F", i)
  x.Mutex.Lock()
  if ! x.CondSpectrum (i) {
    x.wait (i)
  }
  y := x.NFuncSpectrum (i)
  x.Permutation.Permute()
  for j := uint(0); j < x.uint; j++ {
    n := x.Permutation.F(j)
    if x.CondSpectrum (n) {
      x.signal (n)
    }
  }
  if x.nu > 0 {
    x.u.Unlock()
  } else {
    x.Mutex.Unlock()
  }
  return y
}
