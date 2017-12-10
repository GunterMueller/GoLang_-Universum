package phil

// (c) Christian Maurer   v. 171127 - license see µU.go

// >>> Solution with synchronous message-passing
//     Ben-Ari: Principles of Concurrent and Distributed Programming 2nd edition, p. 188
//     modified to be unsymmetric to avoid deadlocks

type
  channel struct {
              ch []chan int
                 }

func newCh() Philos {
  x := new(channel)
  x.ch = make([]chan int, NPhilos)
  for p := uint(0); p < NPhilos; p++ {
    x.ch[p] = make(chan int)
  }
  for p := uint(0); p < NPhilos; p++ {
    go func (i uint) {
      x.fork (i)
    }(p)
  }
  return x
}

func (x *channel) fork (p uint) {
  for {
    x.ch[p] <- 0
    <-x.ch[p]
  }
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
  x.ch[p] <- 0
  x.ch[left (p)] <- 0
  changeStatus (p, satisfied)
}
