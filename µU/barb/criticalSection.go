package barb

// (c) Christian Maurer   v. 171019 - license see µU.go

import (
  . "µU/obj"
  "µU/cs"
)
type
  barberCS struct {
                  cs.CriticalSection
                  }

func newCS() Barber {
  var n uint
  c := func (i uint) bool {
         if i == customer {
           return true
         }
         return n > 0
       }
  e := func (i uint) uint {
         if i == customer {
           n++
         } else {
           n--
         }
         return n
       }
  x := new(barberCS)
  x.CriticalSection = cs.New (2, c, e, NothingSp)
  return x
}

func (x *barberCS) Customer() {
  x.Enter (customer)
}

func (x *barberCS) Barber() {
  x.Enter (barber)
}
