package mbbuf

// (c) Christian Maurer   v. 171107 - license see µU.go

import (
  . "µU/obj"
  "µU/bbuf"
  "µU/host"
  "µU/fmon"
)
type
  farMonitor struct {
                    Any "pattern object"
                    bbuf.BoundedBuffer
                    fmon.FarMonitor
                    }

func newfm (a Any, n uint, h host.Host, p uint16, s bool) MBoundedBuffer {
  if a == nil || n == 0 { return nil }
  x := new (farMonitor)
  x.Any = Clone (a)
  x.BoundedBuffer = bbuf.New (a, n)
  c := func (a Any, i uint) bool {
         if i == get {
           return x.BoundedBuffer.Num() > 0
         }
         return true // ins
       }
  f := func (a Any, i uint) Any {
         if i == get {
           return x.BoundedBuffer.Get()
         }
         x.BoundedBuffer.Ins (a)
         return a // ins
       }
  x.FarMonitor = fmon.New (a, nFuncs, f, c, h, p, s)
  return x
}

func (x *farMonitor) Ins (a Any) {
  x.FarMonitor.F(a, ins)
}

func (x *farMonitor) Get() Any {
  return x.FarMonitor.F(x.Any, get)
}
