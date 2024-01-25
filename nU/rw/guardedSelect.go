package rw

// (c) Christian Maurer   v. 220702 - license see nU.go

import
  . "nU/obj"
type
  guardedSelect struct {
             inR, outR,
             inW, outW chan any
                       }

func newGS() ReaderWriter {
  x := new(guardedSelect)
  x.inR, x.outR = make(chan any), make(chan any)
  x.inW, x.outW = make(chan any), make(chan any)
  go func() {
    var nR, nW uint // active readers, writers
    for {
      select {
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
