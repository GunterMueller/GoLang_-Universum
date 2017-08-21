package sem

// (c) murus.org  v. 170121 - license see murus.go

// >>> Implementation with guarded select
//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 185

import
  . "murus/obj"
type
  guardedSelect struct {
                  p, v chan Any
                       }

func newGSel (n uint) Semaphore {
  x:= new (guardedSelect)
  x.p, x.v = make (chan Any), make (chan Any)
  go func() {
    val:= n
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
