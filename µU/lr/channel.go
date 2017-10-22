package lr

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> 1st left/right problem

type
  channel struct {
  lI, lO, rI, rO,
            done chan int
                 }

func newCh() LeftRight {
  x := new(channel)
  x.lI, x.lO = make(chan int), make(chan int)
  x.rI, x.rO = make(chan int), make(chan int)
  x.done = make(chan int)
  go func() {
    var nL, nR uint
    for {
//      if _, ok:= <-x.done; ok { break }
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

func (x *channel) Fin() {
  x.done <- 0
}
