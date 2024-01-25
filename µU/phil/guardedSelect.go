package phil

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> Implementation with a universal guardedSelect
//
// >>> N O T   C O R R E C T

import (
  "µU/ker"
  . "µU/obj"
)
type
  guardedSelect struct {
          lock, unlock []chan any
                       }

func newGS() Philos {
  x := new(guardedSelect)
  nForks := make([]uint, NPhilos)
  x.lock = make([]chan any, NPhilos)
  x.unlock = make([]chan any, NPhilos)
  for p := uint(0); p < NPhilos; p++ {
    nForks[p] = 2
    x.lock[p] = make(chan any)
    x.unlock[p] = make(chan any)
  }
  for p := uint(0); p < NPhilos; p++ {
    go func (i uint) {
      for {
for j := uint(0); j < NPhilos; j++ { print (nForks[j], " ") }; println("\n")
        when := When (nForks[i] == 2, x.lock[i])
// println (i, "will wait"); ker.Sleep(10); println (i, "waited")
        select {
        case any, ok := <-when:
          if ok {
j := any.(uint); if x.lock[i] == nil || j != i { ker.Oops() }
println ("lock ok", i)
            nForks[left(i)]--
            nForks[right(i)]--
          } else {
println ("lock not ok", i)
          }
println(i, "has", nForks[i])
//          nForks[left(i)]--
//          nForks[right(i)]--
        case any, ok := <-x.unlock[i]:
          if ok {
j := any.(uint); if x.lock[i] == nil || j != i { ker.Oops() }
println ("unlock ok", i)
            nForks[left(i)]++
            nForks[right(i)]++
          }
        }
      }
    }(p)
  }
  return x
}

func (x *guardedSelect) Lock (p uint) {
  changeStatus (p, hungry)
  x.lock[p] <- p
  changeStatus (p, dining)
}

func (x *guardedSelect) Unlock (p uint) {
  changeStatus (p, satisfied)
  x.unlock[p] <- p
}
