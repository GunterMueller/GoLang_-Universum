package macc

// (c) Christian Maurer   v. 170520 - license see murus.go

import (
  . "murus/obj"
  "murus/euro"
  "murus/mon"
)
type
  monitor struct {
                 euro.Euro "balance"
                 mon.Monitor
                 }
var
  zero = euro.New()

func init() {
  zero.SetVal(0)
}

func newm() MAccount {
  x := new (monitor)
  x.Euro = euro.New()
  x.Euro.Set2 (0, 0)
  ps := func (a Any, i uint) bool {
          switch i {
          case draw:
            return ! x.Euro.Less (a.(euro.Euro))
          case deposit:
            return true
          }
          return true
        }
  fs := func (a Any, i uint) Any {
          switch i {
          case deposit:
            x.Euro.Add (a.(euro.Euro))
          case draw:
            x.Euro.Sub (a.(euro.Euro))
          default:
            return zero
          }
          return x.Euro
        }
  x.Monitor = mon.New (nFuncs, fs, ps)
  return x
}

func (x *monitor) Deposit (e euro.Euro) euro.Euro {
  return x.Monitor.F (e, deposit).(euro.Euro)
}

func (x *monitor) Draw (e euro.Euro) euro.Euro {
  return x.Monitor.F (e, draw).(euro.Euro)
}

func (m *monitor) Write (x, y uint) {
  mutex.Lock()
  m.Euro.Write (x, y)
  mutex.Unlock()
}
