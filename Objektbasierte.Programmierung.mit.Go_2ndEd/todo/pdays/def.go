package pdays

// (c) Christian Maurer   v. 210723 - license see µU.go

import (
  . "µU/obj"
  "µU/day"
)
type
  PersistentDays interface { // persistent sets of calendardays

  Clearer
  Persistor

// Returns true, iff d is contained in x.
  Ex (d day.Calendarday) bool

// d is contained in x.
  Ins (d day.Calendarday)

// d is not contained in x.
  Del (d day.Calendarday)

// Returns the number of days in x.
  Num() uint
}
