package mbuf

// (c) murus.org  v. 140330 - license see murus.go

// >>> implementation with synchronous message passing and guarded selective waiting

import
  . "murus/obj"
type
  guardedSelect struct {
                cI, cG chan Any
                       }

func NewGuardedSelect (a Any, n uint) MBuffer {
  if n == 0 { return nil } // panic
  x:= new (guardedSelect)
  x.cI, x.cG = make (chan Any), make (chan Any)
  go func() {
    buffer:= make ([]Any, n)
    var in, out, num uint
    for {
      select {
      case buffer [in] = <-When (num < n, x.cI):
        in = (in + 1) % n
        num ++
      case When (num > 0, x.cG) <- buffer [out]:
        out = (out + 1) % n
        num --
      }
    }
  }()
  return x
}

func (x *guardedSelect) Ins (a Any) {
  x.cI <- a
}

func (x *guardedSelect) Get() Any {
  return Clone (<-x.cG)
}
