package mon

// (c) Christian Maurer   v. 170629 - license see µu.go

// >>> Synchronization with Locker - what happens behind the scene, depends
//     on the choice of the implementation of Locker by the call of the
//     constructor New... (Mutex, Channel, CAS, TAS, XCHG, ... see µu/lock)
//     One might of course get around this step of indirection and replace
//     lock.Locker e.g. by sync.Mutex and then directy call Lock/Unlock.

//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 154 ff.

import (
  "sync"
  . "µu/obj"
  "µu/perm"
)
type
  monitor struct {
                 uint "number of monitor functions"
                 sync.Mutex "monitor entry queue"
               s []sync.Mutex "condition variable queues"
        nBlocked []uint "numbers of goroutines blocked on s"
          urgent sync.Mutex "urgent queue"
              nU uint "number of goroutines blocked on urgent"
                 FuncSpectrum "monitor functions"
                 PredSpectrum "conditions"
                 bool "true, iff monitor is conditioned"
                 perm.Permutation "indeterminism"
                 }

func new_(n uint, f FuncSpectrum, p PredSpectrum) Monitor {
  if n == 0 { return nil }
  x := new(monitor)
  x.uint = n
  x.s = make ([]sync.Mutex, x.uint)
  for i := uint(0); i < x.uint; i++ {
    x.s[i].Lock()
  }
  x.nBlocked = make ([]uint, x.uint)
  x.urgent.Lock()
  x.FuncSpectrum, x.PredSpectrum = f, AllTrueSp
  x.bool = p != nil
  if x.bool {
    x.PredSpectrum = p
  }
  x.Permutation = perm.New (x.uint)
  return x
}

func (x *monitor) Wait (i uint) {
  if i >= x.uint { WrongUintParameterPanic ("mon.Wait", x, i) }
  x.nBlocked[i]++
  if x.nU > 0 {
    x.urgent.Unlock()
  } else {
    x.Mutex.Unlock()
  }
  x.s[i].Lock()
  x.nBlocked[i]--
}

func (x *monitor) Awaited (i uint) bool {
  if i >= x.uint { WrongUintParameterPanic ("mon.Awaited", x, i) }
  return x.nBlocked[i] > 0
}

func (x *monitor) Signal (i uint) {
  if i >= x.uint { WrongUintParameterPanic ("mon.Signal", x, i) }
  if x.nBlocked[i] > 0 {
    x.nU++
    x.s[i].Unlock()
    x.urgent.Lock()
    x.nU--
  }
}

func (x *monitor) SignalAll (i uint) {
  if i >= x.uint { return }
  if i >= x.uint { WrongUintParameterPanic ("mon.SignalAll", x, i) }
  for {
    if x.nBlocked[i] == 0 { break }
    x.nU++
    x.s[i].Unlock()
    x.urgent.Lock()
    x.nU--
  }
}

func (x *monitor) F (a Any, i uint) Any {
  if i >= x.uint { WrongUintParameterPanic ("mon.F", x, i) }
  x.Mutex.Lock()
  if x.bool {
    for ! x.PredSpectrum (a, i) {
      x.Wait (i)
    }
  }
  b := x.FuncSpectrum (a, i)
  if x.bool {
    x.Permutation.Permute()
    for j := uint(0); j < x.uint; j++ {
      x.Signal (x.Permutation.F (j))
    }
  }
  if x.nU > 0 {
    x.urgent.Unlock()
  } else {
    x.Mutex.Unlock()
  }
  return b
}

// experimental
func (x *monitor) S (a Any, i uint, c chan Any) {
  c <- x.F (a, i)
}
