package macc

// (c) Christian Maurer   v. 220702 - license see nU.go

import
   "nU/mon"
type
  monitor struct {
                 mon.Monitor
                 }

func newM() MAccount {
  balance := uint(0)
  x := new(monitor)
  f := func (a any, i uint) any {
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
