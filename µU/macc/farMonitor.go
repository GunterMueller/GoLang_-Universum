package macc

// (c) Christian Maurer   v. 170520 - license see µU.go

import (
  . "µU/obj"
  "µU/euro"
  "µU/host"
  "µU/fmon"
)
type
  farMonitor struct {
                    fmon.FarMonitor
                    }

func newf (h host.Host, p uint16, s bool) MAccount {
  balance := euro.New()
  balance.SetVal (0)
  x := new (farMonitor)
  ps := func (a Any, i uint) bool {
          if i == draw {
            return ! balance.Less (a.(euro.Euro))
          }
          return true // deposit
        }
  fs := func (a Any, i uint) Any {
          switch i {
          case deposit:
            balance.Add (a.(euro.Euro))
          case draw:
            balance.Sub (a.(euro.Euro))
            if balance.Empty() { panic("should not happen") }
          }
          return balance
        }
  x.FarMonitor = fmon.New (balance, nFuncs, fs, ps, h, p, s)
  return x
}

func (x *farMonitor) Deposit (e euro.Euro) euro.Euro {
  return x.FarMonitor.F (e, deposit).(euro.Euro)
}

func (x *farMonitor) Draw (e euro.Euro) euro.Euro {
  return x.FarMonitor.F (e, draw).(euro.Euro)
}
