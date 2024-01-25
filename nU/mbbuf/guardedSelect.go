package mbbuf

// (c) Christian Maurer   v. 220702 - license see nU.go

// >>> implementation with synchronous message passing and guarded selective waiting

import
  . "nU/obj"
type
  guardedSelect struct {
            cIns, cGet chan any
                       }

func newgs (a any, n uint) MBoundedBuffer {
  if a == nil || n == 0 { return nil }
  x := new (guardedSelect)
  x.cIns, x.cGet = make (chan any), make (chan any)
  go func() {
    buffer := make ([]any, n)
    var in, out, num uint
    for {
      select {
      case buffer [in] = <-When (num < n, x.cIns):
        in = (in + 1) % n
        num++
      case When (num > 0, x.cGet) <- buffer [out]:
        out = (out + 1) % n
        num--
      }
    }
  }()
  return x
}

func (x *guardedSelect) Ins (a any) {
  x.cIns <- a
}

func (x *guardedSelect) Get() any {
  return Clone (<-x.cGet)
}
