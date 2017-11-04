package rw

// (c) Christian Maurer   v. 171101 - license see µU.go

// >>> 1st readers/writers problem

type
  channel struct {
       inR, outR,
       inW, outW chan int
            done chan int
                 }

func newCh() ReaderWriter {
  x := new(channel)
  x.inR, x.outR = make(chan int), make(chan int)
  x.inW, x.outW = make(chan int), make(chan int)
  x.done = make(chan int)
  go func() {
    var nR, nW uint // number of active readers/writers
    loop:
    for {
      if _, ok := <-x.done; ok { break loop }
      if nW == 0 {
        if nR == 0 {
          select {
          case <-x.inR:
            nR++
          case <-x.inW:
            nW = 1
          }
        } else { // nR > 0
          select {
          case <-x.inR:
            nR++
          case <-x.outR:
            nR--
          }
        }
      } else { // nW == 1
        select {
        case <-x.outW:
          nW = 0
        }
      }
    }
  }()
  return x
}

func (x *channel) ReaderIn() {
  x.inR <- 0
}

func (x *channel) ReaderOut() {
  x.outR <- 0
}

func (x *channel) WriterIn() {
  x.inW <- 0
}

func (x *channel) WriterOut() {
  x.outW <- 0
}

func (x *channel) Fin() {
  x.done <- 0
}
