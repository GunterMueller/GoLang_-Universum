package mbbuf

// (c) Christian Maurer   v. 171125 - license see nU.go

import (. "nU/obj"; "nU/bbuf"; "nU/fmon")

type farMonitor struct {
  Any "Musterobjekt"
  bbuf.BoundedBuffer
  fmon.FarMonitor
}

func newfm (a Any, n uint, h string, p uint16, s bool) MBoundedBuffer {
  if a == nil || n == 0 { return nil }
  x := new(farMonitor)
  x.Any = Clone (a)
  x.BoundedBuffer = bbuf.New (a, n)
  c := func (a Any, i uint) bool {
         if i == get {
           return x.BoundedBuffer.Num() > 0
         }
         return true
       }
  f := func (a Any, i uint) Any {
         if i == get {
           return x.BoundedBuffer.Get()
         }
         x.BoundedBuffer.Ins (a)
         return a
       }
  x.FarMonitor = fmon.New (a, 2, f, c, h, p, s)
  return x
}

func (x *farMonitor) Fin() {
  x.FarMonitor.Fin()
}

func (x *farMonitor) Ins (a Any) {
  x.FarMonitor.F(a, ins)
}

func (x *farMonitor) Get() Any {
  return x.FarMonitor.F(x.Any, get)
}
