package phil

// (c) Christian Maurer   v. 171104 - license see µU.go

// >>> mit asynchronem Bta noch eleganter ?

import
  . "µU/lockn"
type
  channel1 struct {
               ch []chan int
                  }

func newCh1() LockerN {
  x := new(channel1)
  x.ch = make([]chan int, NPhilos)
  for p := uint(0); p < NPhilos; p++ {
    x.ch[p] = make(chan int, 1) // wenn die ch's Kapazität 1 hätten,
  }
// könnte man die Goroutinen mit der for-Schleife weglassen
/*
  for p := uint(0); p < NPhilos; p++ {
    go func (i uint) {
         for {
           x.ch[i] <- 0
           <-x.ch[i]
         }
       }(p)
  }
*/
  return x
}


func (x *channel1) Lock (p uint) {
  changeStatus (p, hungry)
  if p % 2 == 0 {
    <-x.ch[left (p)]
    changeStatus (p, hasLeftFork)
    <-x.ch[p]
  } else {
    <-x.ch[p]
    changeStatus (p, hasRightFork)
    <-x.ch[left (p)]
  }
  changeStatus (p, dining)
}

func (x *channel1) Unlock (p uint) {
  x.ch[p] <- 0
  x.ch[left (p)] <- 0
  changeStatus (p, satisfied)
}
