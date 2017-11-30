package rw

// (c) Christian Maurer   v. 171125 - license see nU.go

type channel struct {
  rI, rO, wI, wO chan int
}

func newCh() ReaderWriter {
  x := new(channel)
  x.rI, x.rO = make(chan int), make(chan int)
  x.wI, x.wO = make(chan int), make(chan int)
  go func() {
    var nR, nW uint
    for {
      if nW == 0 {
        if nR == 0 {
          select {
          case <-x.rI:
            nR++
          case <-x.wI:
            nW = 1
          }
        } else { // nR > 0
          select {
          case <-x.rI:
            nR++
          case <-x.rO:
            nR--
          }
        }
      } else { // nW == 1
        select {
        case <-x.wO:
          nW = 0
        }
      }
    }
  }()
  return x
}

func (x *channel) ReaderIn() {
  x.rI <- 0
}

func (x *channel) ReaderOut() {
  x.rO <- 0
}

func (x *channel) WriterIn() {
  x.wI <- 0
}

func (x *channel) WriterOut() {
  x.wO <- 0
}
