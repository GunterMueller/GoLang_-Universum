package mbbuf

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  . "µU/obj"
  "µU/bbuf"
  "µU/fmon"
)
type
  farMonitor struct {
                    any "pattern object"
                    bbuf.BoundedBuffer
                    fmon.FarMonitor
                    }

func newfm (a any, n uint, h string, p uint16, s bool) MBoundedBuffer {
  if a == nil || n == 0 { return nil }
  x := new (farMonitor)
  x.any = Clone (a)
  x.BoundedBuffer = bbuf.New (a, n)
  c := func (a any, i uint) bool {
         if i == get {
           return x.BoundedBuffer.Num() > 0
         }
         return true // ins
       }
  f := func (a any, i uint) any {
         if i == get {
           return x.BoundedBuffer.Get()
         }
         x.BoundedBuffer.Ins (a)
         return a // ins
       }
  x.FarMonitor = fmon.New (a, nFuncs, f, c, h, p, s)
  return x
}

func (x *farMonitor) Ins (a any) {
  x.FarMonitor.F(a, ins)
}

func (x *farMonitor) Get() any {
  return x.FarMonitor.F(x.any, get)
}
