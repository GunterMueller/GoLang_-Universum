package mbuf

// (c) murus.org  v. 140330 - license see murus.go

// >>> Implementation with a conditioned monitor

import (
  . "murus/obj"
  "murus/mon"; "murus/buf"
)
type
  monitor struct {
                 buf.Buffer
                 mon.Monitor
                 }

func NewMonitor (a Any, n uint) MBuffer {
  if a == nil || n == 0 { return nil } // TODO panic
  x:= new (monitor)
  x.Buffer = buf.New (a, n)
  p:= func (a Any, i uint) bool {
        if i == ins {
          return ! x.Buffer.Full()
        }
        return x.Buffer.Num() > 0 // get
      }
  f:= func (a Any, i uint) Any {
        if i == ins {
          x.Buffer.Ins (a)
          return a
        }
        return x.Buffer.Get() // get
      }
  x.Monitor = mon.New (nFuncs, f, p)
  return x
}

func (x *monitor) Ins (a Any) {
  x.F (a, ins)
}

func (x *monitor) Get() Any {
  var dummy Any
  return x.F (dummy, get)
}
