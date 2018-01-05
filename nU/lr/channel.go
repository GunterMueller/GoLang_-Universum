package lr

// (c) Christian Maurer   v. 171125 - license see nU.go

type channel struct {
  inL, outL, inR, outR chan int
}

func newCh() LeftRight {
  x := new(channel)
  x.inL, x.outL = make(chan int), make(chan int)
  x.inR, x.outR = make(chan int), make(chan int)
  go func() {
    var nL, nR uint
    for {
      if nL == 0 {
        if nR == 0 {
          select {
          case <-x.inL:
            nL++
          case <-x.inR:
            nR++
          }
        } else { // nL == 0 && nR > 0
          select {
          case <-x.inR:
            nR++
          case <-x.outR:
            nR--
          }
        }
      } else { // nL > 0
        select {
        case <-x.inL:
          nL++
        case <-x.outL:
          nL--
        }
      }
    }
  }()
  return x
}

func (x *channel) LeftIn() {
  x.inL <- 0
}

func (x *channel) LeftOut() {
  x.outL <- 0
}

func (x *channel) RightIn() {
  x.inR <- 0
}

func (x *channel) RightOut() {
  x.outR <- 0
}
