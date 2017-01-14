package cs

// (c) murus.org  v. 161216 - license see murus.go

// >>> Synchronization with Locker - what happens behind the scene, depends
//     on the choice of the implementation of Locker by the call of the
//     constructor New... (Mutex, Channel, CAS, TAS, XCHG, ... see murus/lock)
//     One might of course get around this step of indirection and replace
//     lock.Locker e.g. by sync.Mutex and then directy call Lock/Unlock.

//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 87 ff.
import (
  . "murus/obj"; "murus/perm"; "murus/lock"
)
type
  criticalSection struct {
                      nP uint "number of process classes"
                         lock.Locker "the baton"
                       s []lock.Locker "on which goroutines are blocked, if ! CondSpectrum"
                      nB []uint "/ numbers of goroutines, that are blocked on these semaphores"
                         CondSpectrum "conditions to enter the critical section"
                 in, out OpSpectrum "operations in the entry and exit protocols"
                         perm.Permutation "random permutation"
                         }

func newCS (n uint, c CondSpectrum, e, l OpSpectrum) CriticalSection {
  if n == 0 { return nil }
  x := new (criticalSection)
  x.nP = n
  x.Locker = lock.NewMutex ()
  x.s = make ([]lock.Locker, x.nP)
  x.nB = make ([]uint, x.nP)
  for k := uint(0); k < x.nP; k++ {
    x.s[k] = lock.NewMutex ()
    x.s[k].Lock()
  }
  x.CondSpectrum, x.in, x.out = c, e, l
  x.Permutation = perm.New (x.nP)
  return x
}

func (x *criticalSection) vall() {
  x.Permutation.Permute()
  for i := uint(0); i < x.nP; i++ {
    k := x.Permutation.F (i)
    if x.CondSpectrum (k) && x.nB[k] > 0 {
      x.nB[k] --
      x.s[k].Unlock()
      return
    }
  }
  x.Locker.Unlock()
}

func (x *criticalSection) Blocked (k uint) bool {
  if k >= x.nP { return false }
  return x.nB[k] > 0
}

func (x *criticalSection) Enter (k uint, a Any) {
  if k >= x.nP { return }
  x.Locker.Lock()
  if ! x.CondSpectrum (k) {
    x.nB[k] ++
    x.Locker.Unlock()
    x.s[k].Lock()
  }
  x.in (a, k)
  x.vall()
}

func (x *criticalSection) Leave (k uint, a Any) {
  if k >= x.nP { return }
  x.Locker.Lock()
  x.out (a, k)
  x.vall()
}
