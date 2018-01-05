package mbuf

// (c) Christian Maurer   v. 171127 - license see nU.go

import (. "nU/obj"; "nU/buf"; "nU/fmon")

const (ins = uint(iota); get)

type farMonitor struct {
  Any "Musterobjekt"
  buf.Buffer
  fmon.FarMonitor
}

func newfm (a Any, h string, p uint16, s bool) MBuffer {
  x := new(farMonitor)
  x.Any = Clone (a)
  x.Buffer = buf.New (a)
  c := func (a Any, i uint) bool {
         if i == get {
           return x.Buffer.Num() > 0
         }
         return true
       }
  f := func (a Any, i uint) Any {
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

func (x *farMonitor) Ins (a Any) {
  x.F (a, ins)
}

func (x *farMonitor) Get() Any {
  return x.F (x.Any, get)
}
