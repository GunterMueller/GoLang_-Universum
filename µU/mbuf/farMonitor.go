package mbuf

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  . "µU/obj"
  "µU/buf"
  "µU/fmon"
)
const (
  ins = uint(iota)
  get
)
type
  farMonitor struct {
                    any "pattern object"
                    buf.Buffer
                    fmon.FarMonitor
                    }

func newfm (a any, h string, p uint16, s bool) MBuffer {
  CheckAtomicOrObject(a)
  x := new(farMonitor)
  x.any = Clone (a)
  x.Buffer = buf.New (a)
  ps := func (a any, i uint) bool {
          if i == get {
            return x.Buffer.Num() > 0
          }
          return true // ins
        }
  fs := func (a any, i uint) any {
          if i == get {
            return x.Buffer.Get()
          }
          x.Buffer.Ins (a) // ins
          return a
        }
  x.FarMonitor = fmon.New (a, 2, fs, ps, h, p, s)
  return x
}

func (x *farMonitor) Fin() {
  x.Fin()
}

func (x *farMonitor) Ins (a any) {
  x.F (a, ins)
}

func (x *farMonitor) Get() any {
  return x.F (x.any, get)
}
