package macc

// (c) Christian Maurer   v. 241019 - license see µU.go

import
  "µU/fmon"

const (
  deposit = uint(iota)
  draw
  show
)
type
  farMonitor struct {
                    fmon.FarMonitor
                    }

func new_(h string, p uint16, s bool) MAccount {
  balance := uint(0)
  x := new(farMonitor)
  ok := true
  ps := func (a any, i uint) bool {
          switch i {
          case deposit, show:
            return true
          }
          ok = a.(uint) <= balance // draw
          return ok
        }
  fs := func (a any, i uint) any {
          switch i {
          case deposit:
            balance += a.(uint)
          case show:
            // balance unchanged
          case draw:
            if ok {
              balance -= a.(uint)
            } else {
              return 0
            }
          }
          return balance
        }
  x.FarMonitor = fmon.New (balance, 3, fs, ps, h, p, s)
  return x
}

func (x *farMonitor) Deposit (e uint) uint {
  return x.FarMonitor.F (e, deposit).(uint)
}

func (x *farMonitor) Draw (e uint) uint {
  return x.FarMonitor.F (e, draw).(uint)
}

func (x *farMonitor) Show (e uint) uint {
  return x.FarMonitor.F (uint(0), show).(uint)
}
