package mon

// (c) murus.org  v. 140525 - license see murus.go

// >>> Synchronization with Locker - what happens behind the scene, depends
//     on the choice of the implementation of Locker by the call of the
//     constructor New... (Mutex, Channel, CAS, TAS, XCHG, ... see murus/lock)
//     One might of course get around this step of indirection and replace
//     lock.Locker e.g. by sync.Mutex and then directy call Lock/Unlock.

//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 154 ff.

import (
  . "murus/obj"; "murus/lock"; "murus/perm"
)
type
  monitor struct {
            nFns uint "number of monitor functions"
                 lock.Locker "monitor entry queue"
               s []lock.Locker "condition variable queues"
              nB []uint "numbers of goroutines blocked on s"
          urgent lock.Locker "urgent queue"
              nU uint "number of goroutines blocked on urgent"
                 FuncSpectrum "monitor functions"
                 PredSpectrum "conditions"
            cond bool // true, iff monitor is conditioned
                 perm.Permutation "indeterminism"
                 }

func New (n uint, f FuncSpectrum, p PredSpectrum) Monitor {
  if n == 0 { return nil }
  x:= new (monitor)
  x.nFns = n
  x.Locker = lock.NewMutex()
  x.s = make ([]lock.Locker, x.nFns)
  for i:= uint(0); i < x.nFns; i++ {
    x.s[i] = lock.NewMutex()
    x.s[i].Lock()
  }
  x.nB = make ([]uint, x.nFns)
  x.urgent = lock.NewMutex()
  x.urgent.Lock()
  x.FuncSpectrum = f
  x.PredSpectrum = AllTrueSp
  x.cond = p != nil
  if x.cond {
    x.PredSpectrum = p
  }
  x.Permutation = perm.New (x.nFns)
  return x
}

func (x *monitor) Wait (i uint) {
  if i >= x.nFns { WrongUintParameterPanic ("mon.Wait", x, i) }
  x.nB[i] ++
  if x.nU > 0 {
    x.urgent.Unlock()
  } else {
    x.Locker.Unlock()
  }
  x.s[i].Lock()
  x.nB[i] --
}

func (x *monitor) Awaited (i uint) bool {
  if i >= x.nFns { WrongUintParameterPanic ("mon.Awaited", x, i) }
  return x.nB[i] > 0
}

func (x *monitor) Signal (i uint) {
  if i >= x.nFns { WrongUintParameterPanic ("mon.Signal", x, i) }
  if x.nB[i] > 0 {
    x.nU ++
    x.s[i].Unlock()
    x.urgent.Lock()
    x.nU --
  }
}

func (x *monitor) SignalAll (i uint) {
  if i >= x.nFns { return }
  if i >= x.nFns { WrongUintParameterPanic ("mon.SignalAll", x, i) }
  for {
    if x.nB[i] == 0 { break }
    x.nU ++
    x.s[i].Unlock()
    x.urgent.Lock()
    x.nU --
  }
}

func (x *monitor) F (a Any, i uint) Any {
  if i >= x.nFns { WrongUintParameterPanic ("mon.F", x, i) }
  x.Locker.Lock()
  if x.cond {
    for ! x.PredSpectrum (a, i) {
      x.Wait (i)
    }
  }
  b:= x.FuncSpectrum (a, i)
  if x.cond {
    x.Permutation.Permute()
    for j:= uint(0); j < x.nFns; j++ {
      x.Signal (x.Permutation.F (j))
    }
  }
  if x.nU > 0 {
    x.urgent.Unlock()
  } else {
    x.Locker.Unlock()
  }
  return b
}

// experimental
func (x *monitor) S (a Any, i uint, c chan Any) {
  c <- x.F (a, i)
}
