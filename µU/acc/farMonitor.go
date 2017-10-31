package acc

// (c) Christian Maurer   v. 171020 - license see µU.go

import (
  . "µU/obj"
  "µU/host"
  "µU/fmon"
)
type
  farMonitor struct {
                    fmon.FarMonitor
                    }

func newFM (h host.Host, p uint16, s bool) Account {
  balance := uint(0)
  x := new (farMonitor)
  ps := func (a Any, i uint) bool {
          if i == draw {
            return balance >= a.(uint)
          }
          return true // deposit
        }
  fs := func (a Any, i uint) Any {
          switch i {
          case deposit:
            balance += a.(uint)
          case draw:
            balance -= a.(uint)
          }
          return balance
        }
  x.FarMonitor = fmon.New (balance, nFuncs, fs, ps, h, p, s)
  return x
}

func (x *farMonitor) Deposit (a uint) uint {
  return x.FarMonitor.F (a, deposit).(uint)
}

func (x *farMonitor) Draw (a uint) uint {
  return x.FarMonitor.F (a, draw).(uint)
}
