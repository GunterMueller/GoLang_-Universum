package phil

// (c) Christian Maurer   v. 171228 - license see nU.go

type
  channel struct {
              ch []chan int
                 }

func newCh() Philos {
  x := new(channel)
  x.ch = make([]chan int, 5)
  for p := uint(0); p < 5; p++ {
    x.ch[p] = make(chan int)
  }
  for p := uint(0); p < 5; p++ {
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
  changeStatus (p, thinking)
}
