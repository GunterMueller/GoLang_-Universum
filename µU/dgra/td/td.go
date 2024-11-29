package td

// (c) Christian Maurer   v. 241015 - license see µU.go

import (
  . "µU/obj"
  "µU/fmon"
)
type
  farMonitor struct {
                    fmon.FarMonitor
                    }

func new_(h string, port uint16, s bool) TD {
  const R = 8 // number of involved processes
  var uints [R]uint
  x := new(farMonitor)
  p := AllTrueSp
  f := func (a any, i uint) any {
         n := a.(uint) / 100 // n-th process
         if i == register {
           uints[n] = a.(uint) % 100
         }
         return uints[n]
       }
  x.FarMonitor = fmon.New (uint(0), 2, f, p, h, port, s)
  return x
}

func (x *farMonitor) Register (a uint) {
  x.FarMonitor.F (a, register)
}

func (x *farMonitor) Answer (a uint) uint {
  return x.FarMonitor.F (a, answer).(uint)
}

func (x *farMonitor) Fin() {
  x.FarMonitor.Fin()
}
