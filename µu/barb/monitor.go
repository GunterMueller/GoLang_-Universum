package barb

// (c) Christian Maurer   v. 170731 - license see µu.go

import (
  . "µu/obj"
  "µu/mon"
)
type
  barberMon struct {
                   mon.Monitor
                   }

func newM() Barber {
  var x mon.Monitor
  var n uint
//  barberFree:= true
  do := func (a Any, i uint) Any {
          if i == customer {
//            for ! barberFree {
//              x.Wait (barber)
//            }
//            barberFree = false
            n++
            x.Signal (customer)
          } else { // i == barber
//            barberFree = true
//            x.Signal (barber)
            for n == 0 {
              x.Wait (customer)
            }
            n--
          }
          return 0
        }
  x = mon.New (2, do, nil)
  return &barberMon { x }
}

func (x *barberMon) Customer() {
  x.F (nil, customer)
}

func (x *barberMon) Barber() {
  x.F (nil, barber)
}
