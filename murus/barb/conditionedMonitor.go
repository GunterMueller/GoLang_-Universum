package barb

// (c) Christian Maurer   v. 170731 - license see murus.go

import (
  . "murus/obj"
  "murus/mon"
)
type
  barberCMon struct {
                    mon.Monitor
                    }

func newCM() Barber {
  var n uint
  fs := func (a Any, i uint) Any {
          if i == customer {
            n++
          } else { // i == barber
            n--
          }
          return 0
        }
  ps := func (a Any, i uint) bool {
          if i == customer {
            return true
          }
          return n > 0 // i == barber
        }
  return &barberCMon { mon.New (2, fs, ps) }
}

func (x *barberCMon) Customer() {
  x.F (nil, customer)
}

func (x *barberCMon) Barber() {
  x.F (nil, barber)
}
