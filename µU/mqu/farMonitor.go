package mqu

// (c) Christian Maurer   v. 170218 - license see µU.go

import (
  . "µU/obj"
  "µU/qu"
  "µU/host"
  "µU/fmon"
)
const (
  ins = uint(iota)
  get
  nFuncs
)
type
  farMonitor struct {
                    qu.Queue
                    fmon.FarMonitor
                    }

func newf (a Any, h host.Host, p uint16, s bool) MQueue {
  x := new(farMonitor)
  x.Queue = qu.New (a)
  ps := func (a Any, i uint) bool {
          if i == get {
            return x.Queue.Num() > 0
          }
          return true // ins
        }
  fs := func (a Any, i uint) Any {
          if i == get {
            return x.Queue.Get()
          }
          x.Queue.Ins (a) // ins
          return a
        }
  x.FarMonitor = fmon.New (x.Queue, nFuncs, fs, ps, h, p, s)
  return x
}

func (x *farMonitor) Fin() {
  x.Fin()
}

func (x *farMonitor) Ins (a Any) {
  x.F(a, ins)
}

func (x *farMonitor) Get() Any {
  var a Any
  return x.F(a, get)
}
