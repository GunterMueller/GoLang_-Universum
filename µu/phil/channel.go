package phil

// (c) Christian Maurer   v. 170627 - license see µu.go

// >>> Solution with message-passing
//     Ben-Ari: Principles of Concurrent and Distributed Programming 2nd edition, p. 188
//     modified to be unsymmetric to avoid deadlocks
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 187

import
  . "µu/lockp"
type
  channel struct {
              ch []chan bool
                 }

func newCh() LockerP {
  x := new (channel)
  x.ch = make ([]chan bool, NPhilos)
  for p := uint(0); p < NPhilos; p++ {
    x.ch[p] = make (chan bool)
  }
  for p := uint(0); p < NPhilos; p++ {
    go func (i uint) {
         for {
           x.ch[i] <- true
           <-x.ch[i]
         }
       }(p)
  }
  return x
}

func (x *channel) Lock (p uint) {
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

func (x *channel) Unlock (p uint) {
  x.ch[p] <- true
  x.ch[left (p)] <- true
  changeStatus (p, satisfied)
}
