package mbuf

// (c) murus.org  v. 161226 - license see murus.go

import (
  . "murus/obj"
  "murus/buf"
  "murus/host"
  "murus/fmon"
)
type
  farMonitor struct {
                    buf.Buffer
                    fmon.FarMonitor
                    }

func NewFarMonitor (a Any, n uint, h host.Host, p uint16, s bool) MBuffer {
  if a == nil || n == 0 { return nil } // TODO panic
  x:= new (farMonitor)
  x.Buffer = buf.New (a, n)
  c:= func (a Any, i uint) bool {
        if i == get {
          return x.Buffer.Num() > 0
        }
        return true // ins
      }
  f:= func (a Any, i uint) Any {
        if i == get {
          return x.Buffer.Get()
        }
        x.Buffer.Ins (a)
        return a // ins
      }
  x.FarMonitor = fmon.New (a, nFuncs, f, c, h, p, s)
  return x
}

func (x *farMonitor) Ins (a Any) {
  x.FarMonitor.F (a, ins)
}

func (x *farMonitor) Get() Any {
  var dummy Any
  return x.FarMonitor.F (dummy, get)
}
