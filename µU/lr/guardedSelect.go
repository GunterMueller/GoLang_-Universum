package lr

// (c) Christian Maurer   v. 171015 - license see µU.go

// >>> 1st left/right problem

import
  . "µU/obj"
type
  guardedSelect struct {
        iL, oL, iR, oR chan Any
                  done chan int
                       }

func newGS() LeftRight {
  x:= new(guardedSelect)
  x.iL, x.oL = make(chan Any), make(chan Any)
  x.iR, x.oR = make(chan Any), make(chan Any)
  x.done = make(chan int)
  go func() {
    var nL, nR uint // active lefts, rights
    loop:
    for {
      select {
      case <-x.done:
        break loop
      case <-When (nR == 0, x.iL):
        nL++
      case <-When (nL > 0, x.oL):
        nL--
      case <-When (nL == 0, x.iR):
        nR++
      case <-When (nR > 0, x.oR):
        nR--
      }
    }
  }()
  return x
}

func (x *guardedSelect) LeftIn() {
  x.iL <- 0
}

func (x *guardedSelect) LeftOut() {
  x.oL <- 0
}

func (x *guardedSelect) RightIn() {
  x.iR <- 0
}

func (x *guardedSelect) RightOut() {
  x.oR <- 0
}

func (x *guardedSelect) Fin() {
  x.done <- 0
}
