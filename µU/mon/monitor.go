package mon

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "sync"
  . "µU/obj"
  "µU/perm"
)
type
  monitor struct {
                 uint "number of monitor functions"
                 sync.Mutex "monitor entry queue"
               s []sync.Mutex "condition variable queues"
              ns []uint "numbers of goroutines blocked on s"
               u sync.Mutex "urgent queue"
              nu uint "number of goroutines blocked on urgent"
                 FuncSpectrum "monitor functions"
                 perm.Permutation "indeterminism"
                 }

func new_(n uint, f FuncSpectrum) Monitor {
  if n == 0 { return nil }
  x := new(monitor)
  x.uint = n
  x.s = make([]sync.Mutex, x.uint)
  for i := uint(0); i < x.uint; i++ {
    x.s[i].Lock()
  }
  x.ns = make([]uint, x.uint)
  x.u.Lock()
  x.FuncSpectrum = f
  x.Permutation = perm.New (x.uint)
  return x
}

func (x *monitor) chk (s string, i uint) {
  if i >= x.uint { WrongUintParameterPanic (s, x, i) }
}

func (x *monitor) Wait (i uint) {
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

func (x *monitor) Signal (i uint) {
  x.chk ("Signal", i)
  if x.ns[i] > 0 {
    x.nu++
    x.s[i].Unlock()
    x.u.Lock()
    x.nu--
  }
}

func (x *monitor) SignalAll (i uint) {
  x.chk ("SignalAll", i)
  for x.ns[i] > 0 {
    x.nu++
    x.s[i].Unlock()
    x.u.Lock()
    x.nu--
  }
}

func (x *monitor) F (a any, i uint) any {
  x.chk ("F", i)
  x.Mutex.Lock()
  y := x.FuncSpectrum (a, i)
  if x.nu > 0 {
    x.u.Unlock()
  } else {
    x.Mutex.Unlock()
  }
  return y
}
