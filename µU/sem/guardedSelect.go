package sem

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> Implementation with guarded select

import
  . "µU/obj"
type
  guardedSelect struct {
                  p, v chan any
                       }

func newGS (n uint) Semaphore {
  x := new(guardedSelect)
  x.p, x.v = make(chan any), make(chan any)
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
