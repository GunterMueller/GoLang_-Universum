package macc

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

func newFM (h host.Host, p uint16, s bool) MAccount {
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
            return balance
          case draw:
            balance -= a.(uint)
          }
          return a
        }
  x.FarMonitor = fmon.New (balance, 2, fs, ps, h, p, s)
  return x
}

func (x *farMonitor) Deposit (a uint) uint {
  return x.FarMonitor.F (a, deposit).(uint)
}

func (x *farMonitor) Draw (a uint) uint {
  return x.FarMonitor.F (a, draw).(uint)
}
