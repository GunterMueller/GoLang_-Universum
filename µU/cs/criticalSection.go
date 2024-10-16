package cs

// (c) Christian Maurer   v. 171019 - license see µU.go

import (
  "sync"
  . "µU/obj"
  "µU/perm"
)
type
  criticalSection struct {
                         uint "number of process classes"
                         sync.Mutex "the baton"
                       s []sync.Mutex "on which processes are blocked, if ! CondSpectrum"
                      ns []uint "numbers of processes, that are blocked on these semaphores"
                         CondSpectrum "conditions to enter the critical section"
                      in NFuncSpectrum "functions in the enter protocols"
                     out StmtSpectrum "statements in the leave protocols"
                         perm.Permutation "random permutation"
                         }

func new_(n uint, c CondSpectrum, e NFuncSpectrum, l StmtSpectrum) CriticalSection {
  if n == 0 { return nil }
  x := new (criticalSection)
  x.uint = n
  x.s = make ([]sync.Mutex, x.uint)
  x.ns = make ([]uint, x.uint)
  for k := uint(0); k < x.uint; k++ {
    x.s[k].Lock()
  }
  x.CondSpectrum, x.in, x.out = c, e, l
  x.Permutation = perm.New (x.uint)
  return x
}

func (x *criticalSection) vAll() {
  x.Permutation.Permute()
  for i := uint(0); i < x.uint; i++ {
    k := x.Permutation.F (i)
    if x.CondSpectrum (k) && x.ns[k] > 0 {
      x.ns[k]--
      x.s[k].Unlock()
      return
    }
  }
  x.Mutex.Unlock()
}

func (x *criticalSection) Blocked (i uint) bool {
  if i >= x.uint { return false }
  return x.ns[i] > 0
}

func (x *criticalSection) Enter (i uint) uint {
  if i >= x.uint { return uint(0) }
  x.Mutex.Lock()
  if ! x.CondSpectrum (i) {
    x.ns[i]++
    x.Mutex.Unlock()
    x.s[i].Lock()
  }
  defer x.vAll()
  return x.in (i)
}

func (x *criticalSection) Leave (i uint) {
  if i >= x.uint { return }
  x.Mutex.Lock()
  x.out (i)
  x.vAll()
}
