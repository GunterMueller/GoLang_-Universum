package macc

// (c) Christian Maurer   v. 171125 - license see µU.go

import (. "µU/obj"; "µU/mon")

type monitor struct {
  mon.Monitor
}

func newM() MAccount {
  x := new(monitor)
  balance := uint(0)
  f := func (a Any, i uint) Any {
         if i == deposit {
           x.Monitor.Signal (deposit)
           balance += a.(uint)
         } else { // draw
           if balance < a.(uint) {
             x.Monitor.Wait (deposit)
           }
           balance -= a.(uint)
         }
         return balance
       }
  x.Monitor = mon.New (2, f)
  return x
}

func (x *monitor) Deposit (a uint) uint {
  return x.Monitor.F (a, deposit).(uint)
}

func (x *monitor) Draw (a uint) uint {
  return x.Monitor.F (a, draw).(uint)
}
