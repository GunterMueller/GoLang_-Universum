package mbuf

// (c) Christian Maurer   v. 170218 - license see µu.go

import (
  "µu/ker"
  . "µu/obj"
  "µu/buf"
  "µu/host"
  "µu/fmon"
)
type
  farMonitor struct {
                    buf.Buffer
                    fmon.FarMonitor
                    }

func newfm (a Any, n uint, h host.Host, p uint16, s bool) MBuffer {
  if a == nil || n == 0 {
    ker.Panic("mbuf.NewFM with 1st param nil or 2nd param 0")
  }
  x := new (farMonitor)
  x.Buffer = buf.New (a, n)
  c := func (a Any, i uint) bool {
         if i == get {
           return x.Buffer.Num() > 0
         }
         return true // ins
       }
  f := func (a Any, i uint) Any {
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
  x.FarMonitor.F(a, ins)
}

func (x *farMonitor) Get() Any {
  var a Any
  return x.FarMonitor.F(a, get)
}
