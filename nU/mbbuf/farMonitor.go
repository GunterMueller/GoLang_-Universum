package mbbuf

// (c) Christian Maurer   v. 220702 - license see nU.go

import (. "nU/obj"; "nU/bbuf"; "nU/fmon")

type farMonitor struct {
  any "Musterobjekt"
  bbuf.BoundedBuffer
  fmon.FarMonitor
}

func newfm (a any, n uint, h string, p uint16, s bool) MBoundedBuffer {
  if a == nil || n == 0 { return nil }
  x := new(farMonitor)
  x.any = Clone (a)
  x.BoundedBuffer = bbuf.New (a, n)
  c := func (a any, i uint) bool {
         if i == get {
           return x.BoundedBuffer.Num() > 0
         }
         return true
       }
  f := func (a any, i uint) any {
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

func (x *farMonitor) Ins (a any) {
  x.FarMonitor.F(a, ins)
}

func (x *farMonitor) Get() any {
  return x.FarMonitor.F(x.any, get)
}
