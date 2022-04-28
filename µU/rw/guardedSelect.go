package rw

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> 1st readers/writers problem with guarded select

import
  . "µU/obj"
type
  guardedSelect struct {
  inR, outR, inW, outW chan any
                  done chan int
                       }

func newGS() ReaderWriter {
  x := new(guardedSelect)
  x.inR, x.outR = make(chan any), make(chan any)
  x.inW, x.outW = make(chan any), make(chan any)
  x.done = make(chan int)
  go func() {
    var nR, nW uint // active readers, writers
    loop:
    for {
      select {
      case <-x.done:
        break loop
      case <-When (nW == 0, x.inR):
        nR++
      case <-When (nR > 0, x.outR):
        nR--
      case <-When (nR == 0 && nW == 0, x.inW):
        nW = 1
      case <-When (nW == 1, x.outW):
        nW = 0
      }
    }
  }()
  return x
}

func (x *guardedSelect) ReaderIn() {
  x.inR <- 0
}

func (x *guardedSelect) ReaderOut() {
  x.outR <- 0
}

func (x *guardedSelect) WriterIn() {
  x.inW <- 0
}

func (x *guardedSelect) WriterOut() {
  x.outW <- 0
}

func (x *guardedSelect) Fin() {
  x.done <- 0
}
