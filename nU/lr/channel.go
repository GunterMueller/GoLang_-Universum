package lr

// (c) Christian Maurer   v. 171125 - license see nU.go

type channel struct {
  lI, lO, rI, rO chan int
}

func newCh() LeftRight {
  x := new(channel)
  x.lI, x.lO = make(chan int), make(chan int)
  x.rI, x.rO = make(chan int), make(chan int)
  go func() {
    var nL, nR uint
    for {
      if nL == 0 {
        if nR == 0 {
          select {
          case <-x.lI:
            nL++
          case <-x.rI:
            nR++
          }
        } else { // nR > 0
          select {
          case <-x.rI:
            nR++
          case <-x.rO:
            nR--
          }
        }
      } else { // nL > 0
        select {
        case <-x.lI:
          nL++
        case <-x.lO:
          nL--
        }
      }
    }
  }()
  return x
}

func (x *channel) LeftIn() {
  x.lI <- 0
}

func (x *channel) LeftOut() {
  x.lO <- 0
}

func (x *channel) RightIn() {
  x.rI <- 0
}

func (x *channel) RightOut() {
  x.rO <- 0
}
