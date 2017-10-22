package barb

// (c) Christian Maurer   v. 171019 - license see µU.go

import (
  . "µU/obj"
  "µU/mon"
)
type
  monitor struct {
                 mon.Monitor
                 }

func newM() Barber {
  var n uint
  var m mon.Monitor
//  barberFree := true
  f := func (a Any, i uint) Any {
         if i == customer {
//           if ! barberFree {
//             m.Wait (barber)
//           }
//           barberFree = false
           n++
           m.Signal (customer)
         } else {
//           barberFree = true
//           m.Signal (barber)
           for n == 0 {
             m.Wait (customer)
           }
           n--
         }
         return n
       }
  m = mon.New (2, f)
  x := new(monitor)
  x.Monitor = m
  return x
}

func (x *monitor) Customer() {
  x.F (nil, customer)
}

func (x *monitor) Barber() {
  x.F (nil, barber)
}
