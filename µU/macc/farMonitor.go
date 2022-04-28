package macc

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "µU/fmon"
)
type
  farMonitor struct {
                    fmon.FarMonitor
                    }

func newFMon (h string, p uint16, s bool) MAccount {
  balance := uint(0)
  x := new (farMonitor)
  ps := func (a any, i uint) bool {
          if i == deposit {
            return true
          }
          return balance >= a.(uint) // draw
        }
  fs := func (a any, i uint) any {
          if i == deposit {
            balance += a.(uint)
          } else { // draw
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
