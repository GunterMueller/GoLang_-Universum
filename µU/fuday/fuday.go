package fuday

// (c) Christian Maurer   v. 161216 - license see µU.go

import (
  "µU/nat"
  "µU/day"
)
type
  fuDay struct {
               day.Calendarday
               }
var
  tmp, tmp1, tmp2 day.Calendarday = day.New(), day.New(), day.New()

func new_() FUDay {
  x := new (fuDay)
  x.Calendarday = day.New()
  x.Update()
  return x
}

func (x *fuDay) Set (d day.Calendarday) {
  x.Calendarday.Copy (d)
}

func (x *fuDay) summer() bool {
  y := x.Year()
  tmp.Set (1, 4, y)
  if x.Calendarday.Less (tmp) {
    return false
  }
  tmp.Set (1, 10, y)
  if x.Calendarday.Less (tmp) {
    return true
  }
  return false
}

func (x *fuDay) Semester (b, e day.Calendarday) {
  y := x.Year() % 100
  tmp.Set (1, 4, y)
  if x.Calendarday.Less (tmp) {
    b.Set (1, 10, y); b.Dec (day.Yearly)
    e.Set (1,  4, y)
  } else {
    tmp.Set (1, 10, y)
    if x.Calendarday.Less (tmp) {
      b.Set (1,  4, y)
      e.Set (1, 10, y)
    } else {
      b.Set (1, 10, y)
      e.Set (1,  4, y); e.Inc (day.Yearly)
    }
  }
  e.Dec (day.Daily)
}

func (x *fuDay) Lectures (b, e day.Calendarday) {
  x.Semester (b, e)
  y := x.Year() % 100
  w := uint(14)
  if x.summer() {
    b.Set (14, 4, y)
  } else {
    b.Set (18, 10, y)
    w += 2 + 2 // Weihnachtsferien
  }
  for ! b.IsBeginning (day.Weekly) {
    b.Dec (day.Daily)
  }
  e.Copy (b)
  for i := uint(0); i < w; i++ {
    e.Inc (day.Weekly)
  }
  e.Dec (day.Daily)
  e.Dec (day.Daily) // Saturday
}

func (x *fuDay) String() string {
  y := x.Year() % 100
  tmp.Set (1, 4, y)
  s := nat.StringFmt (y, 2, true)
  if x.Calendarday.Less (tmp) {
    tmp.Dec (day.Yearly)
    return "WS " + nat.StringFmt (tmp.Year() % 100, 2, true) + "/" + s
  }
  tmp.Set (1, 10, y)
  if x.Calendarday.Less (tmp) {
    return "SS " + s
  }
  tmp.Inc (day.Yearly)
  return "WS" + s + "/" + nat.StringFmt (tmp.Year() % 100, 2, true)
}

func (x *fuDay) LectureDay (d day.Calendarday) bool {
  summer := x.summer()
  x.Lectures (tmp1, tmp2)
  if d.Less (tmp1) || tmp2.Less (d) || d.IsHoliday() {
    return false
  }
  if ! summer {
    tmp.Copy (tmp1) // Beginn Akademische Ferien:
    for i := uint(0); i < 10; i++ { tmp.Inc (day.Weekly) }
// tmp.Write (10, 0)
    tmp2.Copy (tmp) // Vorlesungsbeginn Januar:
    for i := uint(0); i < 2; i++ { tmp2.Inc (day.Weekly) }
// tmp2.Write (10, 0)
    if tmp.Eq(d) || tmp.Less (d) && d.Less (tmp2) {
      return false
    }
  }
  return true
}

func (x *fuDay) NumWeeks() uint {
  if x.summer() {
    return 14
  }
  return 16
}

func (x *fuDay) Monday (d day.Calendarday, n uint) {
  w := x.NumWeeks()
  if n == 0 || n > w {
    d.Clr()
    return
  }
  x.Lectures (d, tmp)
  for i := uint(0); i + 1 < n; i++ {
    d.Inc (day.Weekly)
  }
  if ! x.summer() && n > 10 { // Akademische Ferien
    d.Inc (day.Weekly)
    d.Inc (day.Weekly)
  }
}
