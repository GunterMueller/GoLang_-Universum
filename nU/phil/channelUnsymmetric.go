package phil

// (c) Christian Maurer   v. 171229 - license see nU.go

// >>> Unsymmetric solution with synchronous message-passing

type channelUnsymmetric struct {
  lock, lockLeft, unlock, unlockLeft []chan bool
}

func newChU() Philos {
  x := new(channelUnsymmetric)
  x.lock = make([]chan bool, 5)
  x.lockLeft = make([]chan bool, 5)
  x.unlock = make([]chan bool, 5)
  x.unlockLeft = make([]chan bool, 5)
  for p := uint(0); p < 5; p++ {
    x.lock[p] = make(chan bool)
    x.lockLeft[p] = make(chan bool)
    x.unlock[p] = make(chan bool)
    x.unlockLeft[p] = make(chan bool)
  }
  for p := uint(0); p < 5; p++ {
    go func (i uint) {
      for {
        select {
        case <-x.lock[i]:
          <-x.unlock[i]
        case <-x.lockLeft[i]:
          <-x.unlockLeft[i]
        }
      }
    }(p)
  }
  return x
}

func (x *channelUnsymmetric) Lock (p uint) {
  changeStatus (p, hungry)
  if p % 2 == 0 {
    x.lockLeft[left(p)] <- true
    changeStatus (p, hasLeftFork)
    x.lock[p] <- true
  } else {
    x.lock[p] <- true
    changeStatus (p, hasRightFork)
    x.lockLeft[left(p)] <- true
  }
  changeStatus (p, dining)
}

func (x *channelUnsymmetric) Unlock (p uint) {
  changeStatus (p, thinking)
  x.unlockLeft[left(p)] <- true
  x.unlock[p] <- true
}
