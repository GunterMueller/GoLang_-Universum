package mqu

// (c) murus.org  v. 161226 - license see murus.go

import (
  . "murus/obj"
  "murus/qu"
  "murus/host"
  "murus/fmon"
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

func NewFarMonitor (a Any, h host.Host, p uint16, s bool) MQueue {
  x:= new (farMonitor)
  x.Queue = qu.New (a)
  c:= func (a Any, i uint) bool {
        if i == get {
          return x.Queue.Num() > 0
        }
        return true // ins
      }
  f:= func (a Any, i uint) Any {
        if i == get {
          return x.Queue.Get()
        }
        x.Queue.Ins (a) // ins
        return a
      }
  x.FarMonitor = fmon.New (x.Queue, nFuncs, f, c, h, p, s)
  return x
}

func (x *farMonitor) Fin() {
  x.Fin()
}

func (x *farMonitor) Ins (a Any) {
  x.F (a, ins)
}

func (x *farMonitor) Get() Any {
  var dummy Any
  return x.F (dummy, get)
}
