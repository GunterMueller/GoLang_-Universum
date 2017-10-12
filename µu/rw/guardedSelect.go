package rw

// (c) Christian Maurer   v. 170121 - license see µu.go

//     readers/writers problem: solution with "guarded selective waiting"
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 185

import
  . "µu/obj"
type
  guardedSelect struct {
        iR, oR, iW, oW chan Any
                  done chan int
                       }

func newGSel() ReaderWriter {
  x := new (guardedSelect)
  x.iR, x.oR = make (chan Any), make (chan Any)
  x.iW, x.oW = make (chan Any), make (chan Any)
  x.done = make (chan int)
  go func() {
    var nR, nW uint // active readers, writers
//    loop:
    for {
      select {
//      case <-x.done: break loop
      case <-When (nW == 0, x.iR):
        nR++
      case <-When (nR > 0, x.oR):
        nR--
      case <-When (nR == 0 && nW == 0, x.iW):
        nW = 1
      case <-When (nW == 1, x.oW):
        nW = 0
      }
    }
  }()
  return x
}

func (x *guardedSelect) ReaderIn() {
  x.iR <- 0
}

func (x *guardedSelect) ReaderOut() {
  x.oR <- 0
}

func (x *guardedSelect) WriterIn() {
  x.iW <- 0
}

func (x *guardedSelect) WriterOut() {
  x.oW <- 0
}

func (x *guardedSelect) Fin() {
  x.done <- 0
}
