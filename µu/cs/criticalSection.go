package cs

// (c) Christian Maurer   v. 170407 - license see µu.go

// >>> Synchronization with Locker - what happens behind the scene, depends
//     on the choice of the implementation of Locker by the call of the
//     constructor New... (Mutex, Channel, CAS, TAS, XCHG, ... see µu/lock)
//     One might of course get around this step of indirection and replace
//     lock.Locker e.g. by sync.Mutex and then directy call Lock/Unlock.

//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 87 ff.
import (
  . "µu/obj"
  "µu/perm"
  "µu/lock"
)
type
  criticalSection struct {
                         uint "number of process classes"
                         lock.Locker "the baton"
                       s []lock.Locker "on which goroutines are blocked, if ! CondSpectrum"
                nBlocked []uint "numbers of goroutines, that are blocked on these semaphores"
                         CondSpectrum "conditions to enter the critical section"
                 in, out OpSpectrum "operations in the entry and exit protocols"
                         perm.Permutation "random permutation"
                         }

func new_(n uint, c CondSpectrum, e, l OpSpectrum) CriticalSection {
  if n == 0 { return nil }
  x := new (criticalSection)
  x.uint = n
  x.Locker = lock.NewMutex ()
  x.s = make ([]lock.Locker, x.uint)
  x.nBlocked = make ([]uint, x.uint)
  for k := uint(0); k < x.uint; k++ {
    x.s[k] = lock.NewMutex ()
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
    if x.CondSpectrum (k) && x.nBlocked[k] > 0 {
      x.nBlocked[k]--
      x.s[k].Unlock()
      return
    }
  }
  x.Locker.Unlock()
}

func (x *criticalSection) Blocked (k uint) bool {
  if k >= x.uint { return false }
  return x.nBlocked[k] > 0
}

func (x *criticalSection) Enter (k uint, a Any) {
  if k >= x.uint { return }
  x.Locker.Lock()
  if ! x.CondSpectrum (k) {
    x.nBlocked[k]++
    x.Locker.Unlock()
    x.s[k].Lock()
  }
  x.in (a, k)
  x.vAll()
}

func (x *criticalSection) Leave (k uint, a Any) {
  if k >= x.uint { return }
  x.Locker.Lock()
  x.out (a, k)
  x.vAll()
}
