package barb

// (c) Christian Maurer   v. 171019 - license see µU.go

import
  "µU/cmon"
type
  conditionedMonitor struct {
                            cmon.Monitor
                            }

func newCM() Barber {
  var n uint
  c := func (i uint) bool {
         if i == customer {
           return true
         }
         return n > 0
       }
  f := func (i uint) uint {
         if i == customer {
           n++
         } else {
           n--
         }
         return n
       }
  x := new(conditionedMonitor)
  x.Monitor = cmon.New (2, f, c)
  return x
}

func (x *conditionedMonitor) Customer() {
  x.F (customer)
}

func (x *conditionedMonitor) Barber() {
  x.F (barber)
}
