package sem

// (c) Christian Maurer   v. 170121 - license see nU.go

import . "nU/obj"

type guardedSelect struct {
  p, v chan Any
}

func newGS (n uint) Semaphore {
  x := new(guardedSelect)
  x.p, x.v = make(chan Any), make(chan Any)
  go func() {
    val := n
    for {
      select {
      case <-When (val > 0, x.p):
        val--
      case <-x.v:
        val++
      }
    }
  }()
  return x
}

func (x *guardedSelect) P() {
  x.p <- 0
}

func (x *guardedSelect) V() {
  x.v <- 0
}
