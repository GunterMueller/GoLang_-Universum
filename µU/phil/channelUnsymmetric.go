package phil

// (c) Christian Maurer   v. 170627 - license see µU.go

// >>> Unsymmetric solution with synchronous message-passing

import
  . "µU/lockn"
type
  channelUnsymmetric struct {
           cl, cll, cu, cul []chan bool
                            }

func newChU() LockerN {
  x := new(channelUnsymmetric)
  x.cl  = make([]chan bool, NPhilos)
  x.cll = make([]chan bool, NPhilos)
  x.cu  = make([]chan bool, NPhilos)
  x.cul = make([]chan bool, NPhilos)
  for p := uint(0); p < NPhilos; p++ {
    x.cl [p] = make(chan bool)
    x.cll[p] = make(chan bool)
    x.cu [p] = make(chan bool)
    x.cul[p] = make(chan bool)
  }
  for p := uint(0); p < NPhilos; p++ {
    go func (i uint) {
         for {
           select {
           case <-x.cl[i]:
             <-x.cu[i]
           case <-x.cll[i]:
             <-x.cul[i]
           }
         }
       }(p)
  }
  return x
}

func (x *channelUnsymmetric) Lock (p uint) {
  changeStatus (p, hungry)
  if p % 2 == 0 {
    x.cll[left(p)] <- true
    changeStatus (p, hasLeftFork)
    x.cl[p] <- true
  } else {
    x.cl[p] <- true
    changeStatus (p, hasRightFork)
    x.cll[left(p)] <- true
  }
  changeStatus (p, dining)
}

func (x *channelUnsymmetric) Unlock (p uint) {
  changeStatus (p, satisfied)
  x.cul[left(p)] <- true
  x.cu[p] <- true
}
