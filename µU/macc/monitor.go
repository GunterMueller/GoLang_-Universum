package macc

// (c) Christian Maurer   v. 171019 - license see µU.go

import (
  . "µU/obj"
  "µU/euro"
  "µU/mon"
)
type
  monitor struct {
                 mon.Monitor
                 }

func newM() MAccount {
  x := new(monitor)
  cent := uint(0)
  f := func (a Any, i uint) Any {
         e := a.(uint)
         if i == deposit {
           x.Monitor.Signal (deposit)
           cent += e
         } else { // draw
           if cent < e {
             x.Monitor.Wait (deposit)
           }
           cent -= e
         }
         return cent
       }
  x.Monitor = mon.New (nFuncs, f)
  return x
}

func (x *monitor) Deposit (e euro.Euro) euro.Euro {
  e.SetVal (x.Monitor.F (e.Val(), deposit).(uint))
  return e
}

func (x *monitor) Draw (e euro.Euro) euro.Euro {
  e.SetVal (x.Monitor.F (e.Val(), draw).(uint))
  return e
}
