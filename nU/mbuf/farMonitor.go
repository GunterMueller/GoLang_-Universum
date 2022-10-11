package mbuf

// (c) Christian Maurer   v. 220702 - license see nU.go

import (. "nU/obj"; "nU/buf"; "nU/fmon")

const (ins = uint(iota); get)

type farMonitor struct {
  any "Musterobjekt"
  buf.Buffer
  fmon.FarMonitor
}

func newfm (a any, h string, p uint16, s bool) MBuffer {
  x := new(farMonitor)
  x.any = Clone (a)
  x.Buffer = buf.New (a)
  c := func (a any, i uint) bool {
         if i == get {
           return x.Buffer.Num() > 0
         }
         return true
       }
  f := func (a any, i uint) any {
         if i == get {
           return x.Buffer.Get()
         }
         x.Buffer.Ins (a)
         return a
       }
  x.FarMonitor = fmon.New (a, 2, f, c, h, p, s)
  return x
}

func (x *farMonitor) Fin() {
  x.FarMonitor.Fin()
}

func (x *farMonitor) Ins (a any) {
  x.F (a, ins)
}

func (x *farMonitor) Get() any {
  return x.F (x.any, get)
}
