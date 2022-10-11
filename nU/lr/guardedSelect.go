package lr

// (c) Christian Maurer   v. 220702 - license see nU.go

import . "nU/obj"

type guardedSelect struct {
  inL, outL, inR, outR chan any
}

func newGS() LeftRight {
  x := new(guardedSelect)
  x.inL, x.outL = make(chan any), make(chan any)
  x.inR, x.outR = make(chan any), make(chan any)
  go func() {
    var nL, nR uint
    for {
      select {
      case <-When (nR == 0, x.inL):
        nL++
      case <-When (nL > 0, x.outL):
        nL--
      case <-When (nL == 0, x.inR):
        nR++
      case <-When (nR > 0, x.outR):
        nR--
      }
    }
  }()
  return x
}

func (x *guardedSelect) LeftIn() {
  x.inL <- 0
}

func (x *guardedSelect) LeftOut() {
  x.outL <- 0
}

func (x *guardedSelect) RightIn() {
  x.inR <- 0
}

func (x *guardedSelect) RightOut() {
  x.outR <- 0
}
