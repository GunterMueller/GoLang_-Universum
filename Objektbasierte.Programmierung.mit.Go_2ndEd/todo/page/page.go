package page

// (c) Christian Maurer   v. 221021 - license see µU.go

import (
  . "µU/obj"
  "µU/font"
  "µU/day"
  "µU/prt"
  "todo/dayattr"
  "todo/appts"
)
type
  page struct {
              day.Calendarday
              day.Period
              appts.Appointments
              }

func New() Page {
  x := new (page)
  x.Calendarday = day.New()
  x.Appointments = appts.New()
  return x
}

func (x *page) imp (Y any) *page {
  y, ok := Y.(*page)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *page) Empty() bool {
  return x.Appointments.Empty()
}

func (x *page) Clr() {
  x.Appointments.Clr()
}

func (x *page) Eq (Y any) bool {
  return x.Calendarday.Eq (x.imp(Y).Calendarday)
}

func (x *page) Copy (Y any) {
  y := x.imp (Y)
  x.Calendarday.Copy (y.Calendarday)
  x.Appointments.Copy (y.Appointments)
}

func (x *page) Clone() any {
  y := New()
  y.Copy (x)
  return y
}

func (x *page) Set (d day.Calendarday) {
  x.Calendarday = d.Clone().(day.Calendarday)
  x.Appointments.Clr()
}

func (x *page) Less (Y any) bool {
  return x.Calendarday.Less ((x.imp(Y)).Calendarday)
}

func (x *page) Leq (Y any) bool {
  return x.Calendarday.Leq ((x.imp(Y)).Calendarday)
}

func (x *page) HasWord() bool {
  return x.Appointments.HasWord()
}

func (x *page) SetFormat (p day.Period) {
  x.Period = p
  x.Appointments.SetFormat (p)
}

func (x *page) Write (l, c uint) {
  if x.Calendarday.IsHoliday() {
    x.Calendarday.Colours (day.HolidayF, day.HolidayB)
  } else {
    x.Calendarday.Colours (day.WeekdayF, day.WeekdayB)
  }
  switch x.Period {
  case day.Daily:
    x.Calendarday.SetFormat (day.WD)
    x.Calendarday.Write (l, c)
    x.Calendarday.SetFormat (day.Dd_mm_yyyy)
    x.Calendarday.Write (l, c + 11)
//    dayattr.WriteAll (x.Calendarday, l, c + 22)
    x.Appointments.Write (l + 2, c)
  case day.Weekly:
    x.Appointments.Write (l + 1, c)
  case day.Monthly:
    x.Appointments.Write (l, c)
  }
}

func (x *page) Edit (l, c uint) {
  x.Write (l, c)
  x.Appointments.Edit (l + 2, c)
}

func (x *page) SetFont (f font.Font) {
// dummy
}

func (x *page) Print (l, c uint) {
  if x.Period == day.Daily {
    x.Calendarday.SetFormat (day.WD)
    x.Calendarday.Print (l, c)
    x.Calendarday.SetFormat (day.Dd_mm_yyyy)
    x.Calendarday.Print (l, c + 11)
    x.Appointments.Print (l + 2, c)
    prt.GoPrint()
  }
}

func (x *page) Codelen() uint {
  return x.Calendarday.Codelen() +
         x.Appointments.Codelen()
}

func (x *page) Encode() Stream {
  b := make (Stream, x.Codelen())
  a := x.Calendarday.Codelen()
  copy (b[:a], x.Calendarday.Encode())
  copy (b[a:], x.Appointments.Encode())
  return b
}

func (x *page) Decode (b Stream) {
  a := x.Calendarday.Codelen()
  x.Calendarday.Decode (b[:a])
  x.Appointments.Decode (b[a:])
}

func (x *page) Day() day.Calendarday {
  return x.Calendarday.Clone().(day.Calendarday)
}

func (x *page) Index() Func {
  return func (a any) any { return Clone (a.(*page).Calendarday) }
}

func (x *page) Fin() {
  dayattr.Fin()
}
