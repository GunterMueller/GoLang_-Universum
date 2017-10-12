package macc

// (c) Christian Maurer   v. 170520 - license see µu.go

import (
  "sync"
  . "µu/obj"
  "µu/euro"
  "µu/host"
  "µu/fmon"
)
type
  farMonitor struct {
                    euro.Euro "balance"
                    fmon.FarMonitor
                    }
var
  mutex sync.Mutex

func newf (h host.Host, p uint16, s bool) MAccount {
  x := new (farMonitor)
  x.Euro = euro.New()
  x.Euro.SetVal (0)
  ps := func (a Any, i uint) bool {
          if i == draw {
            return ! x.Euro.Less (a.(euro.Euro))
          }
          return true // deposit
        }
  fs := func (a Any, i uint) Any {
          switch i {
          case deposit:
            x.Euro.Add (a.(euro.Euro))
          case draw:
            x.Euro.Sub (a.(euro.Euro))
            if x.Euro.Empty() { panic("should not happen") }
          }
          return x.Euro
        }
  x.FarMonitor = fmon.New (x.Euro, nFuncs, fs, ps, h, p, s)
  return x
}

func (x *farMonitor) Deposit (e euro.Euro) euro.Euro {
  return x.FarMonitor.F (e, deposit).(euro.Euro)
}

func (x *farMonitor) Draw (e euro.Euro) euro.Euro {
  return x.FarMonitor.F (e, draw).(euro.Euro)
}

func (m *farMonitor) Write (x, y uint) {
  mutex.Lock()
  m.Euro.Write (x, y)
  mutex.Unlock()
}
