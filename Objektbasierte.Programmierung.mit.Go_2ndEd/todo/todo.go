package main

// (c) Christian Maurer   v. 230112 - license see µU.go

import (
  "µU/ker"
  "µU/kbd"
  "µU/fontsize"
  "µU/prt"
  "µU/col"
  "µU/mode"
  "µU/scr"
  "µU/errh"
  "µU/day"
  . "todo/help"
  "todo/dayattr"
  "todo/cal"
)
const (
  l1 = 24; c1 = 35
)
var (
  actualDay = day.New()
  period = day.Yearly
  l0, c0, dc uint
)

func setWeekdayColours (d day.Calendarday) {
  d.Colours (day.WeekdayNameF, day.WeekdayNameB)
}

func pos (d day.Calendarday) (uint, uint) {
  switch period {
  case day.Yearly:
    return d.PosInYear()
  case day.Monthly:
    return d.PosInMonth (true, 1, 3, dc)
  case day.Weekly:
    return d.PosInWeek (false, dc)
  }
  return 0, 0
}

func writeAll (d day.Calendarday) {
  switch period {
  case day.Monthly, day.Weekly:
    if d.Empty() {
      return
    }
  default:
    return
  }
  d1 := day.New()
  d1.Copy (d)
  d1.SetBeginning (period)
  cal.SetFormat (period)
  for d.Equiv (d1, period) {
    l, c := pos (d1)
    cal.Seek (d1)
    cal.SetFormat (period) // weil Seek über Define <- Clone das Format mitkopiert
    cal.WriteDay (l0 + l, c0 + c)
    d1.Inc (day.Daily)
  }
}

func write() {
  switch period {
  case day.Yearly:
    actualDay.WriteYear (l0, c0)
  case day.Monthly:
    actualDay.SetFormat (day.Dd_mm_)
    actualDay.WriteMonth (true, 1, 3, dc, l0, c0)
  case day.Weekly:
    actualDay.WriteWeek (false, dc, l0, c0)
  }
}

func edited() bool {
  scr.Cls()
  l0, c0, dc = 0, 0, 0
  switch period {
  case day.Decadic:
    actualDay.SetFormat (day.Yyyy)
    actualDay.Colours (day.YearnumberF, day.YearnumberB)
  case day.Yearly:
    ;
  case day.HalfYearly, day.Quarterly:
    ker.ToDo()
  case day.Monthly:
    l0, c0 = 3, 5; dc = 12 // 11 dayattr + 1
    actualDay.SetAttribute (setWeekdayColours)
    actualDay.SetFormat (day.Wd)
    actualDay.WriteWeek (true, 3, l0, 2)
    actualDay.WriteWeek (true, 3, l0, 2 + 6 * dc + 3)
  case day.Weekly:
    l0, c0 = 2, 2; dc = 11 // 7 x 11 == 77 < 80
    actualDay.SetAttribute (setWeekdayColours)
    actualDay.SetFormat (day.Wd)
    actualDay.WriteWeek (false, dc, l0, c0)
  case day.Daily:
    cal.Edit (actualDay, l0, c0)
  }
  switch period {
  case day.Weekly, day.Monthly, day.Quarterly, day.Yearly:
    dayattr.WriteActual (l1, c1)
  }
  startDate, original := day.New(), day.New()
  for {
    switch period {
    case day.Yearly:
      actualDay.SetFormat (day.Dd)
    case day.Quarterly, day.HalfYearly:
      errh.Error ("nicht erreichbarer Punkt", 3)
    case day.Weekly, day.Monthly:
      actualDay.SetFormat (day.Dd_mm_)
    }
    if ! actualDay.Equiv (startDate, period) {
      startDate.Copy (actualDay)
      switch period {
      case day.Yearly:
        dayattr.WriteActual (l1, c1)
        actualDay.SetFormat (day.Dd)
      case day.Monthly, day.Weekly:
        if period == day.Monthly {
          actualDay.Colours (day.MonthF, day.MonthB)
        } else {
          actualDay.Colours (day.WeekdayNameF, day.WeekdayNameB)
        }
        actualDay.SetFormat (day.Yyyy)
        actualDay.Colours (day.YearnumberF, day.YearnumberB)
        actualDay.Write (0, 0)
        actualDay.Write (0, 80 - 4)
        if period == day.Monthly {
          actualDay.SetFormat (day.M)
        } else {
          actualDay.SetFormat (day.WN)
        }
        actualDay.Colours (day.MonthF, day.MonthB)
        actualDay.Write (0, 30)
        actualDay.SetFormat (day.Dd_mm_)
      }
      actualDay.SetAttribute (dayattr.Attrib)
      write()
      writeAll (actualDay)
    }
    l, c := pos (actualDay)
    dayattr.Attrib (actualDay)
    original.Copy (actualDay)
    switch period {
    case day.Daily:
      ;
    case day.Weekly:
      actualDay.Edit (l0 + l, c0 + c)
    case day.Monthly, day.Quarterly, day.HalfYearly, day.Yearly:
      actualDay.Edit (l0 + l, c0 + c)
    case day.Decadic:
      actualDay.Edit (0, 0)
    }
    if actualDay.Empty() {
      actualDay.Copy (original)
    }
    C, d := kbd.LastCommand()
    actualDay.Write (l0 + l, c0 + c)
    scr.MousePointer (true)
    switch C {
    case kbd.Enter:
      switch period {
      case day.Decadic, day.Monthly, day.Weekly:
        period--
      case day.Yearly, day.HalfYearly, day.Quarterly:
        period = day.Monthly
      }
      return false
    case kbd.Esc, kbd.Back:
      switch period {
      case day.Daily, day.Weekly, day.Yearly:
        period++
      case day.Monthly, day.Quarterly, day.HalfYearly:
        period = day.Yearly
      case day.Decadic:
        return true
      }
      return false
    case kbd.Tab:
      dayattr.Change (d == 0)
      write()
      dayattr.WriteActual (l1, c1)
      C = kbd.Enter // see above
    case kbd.Help:
      Help()
    case kbd.Search:
      dayattr.Normalize()
      dayattr.WriteActual (l1, c1)
      cal.EditWord (l1, c1 + 1 + 8)
      write()
      if period == day.Weekly {
        writeAll (actualDay)
      }
      C = kbd.Enter // so "actualDay" is not influenced
    case kbd.Mark:
      dayattr.Actualize (actualDay, true)
      dayattr.Attrib (actualDay)
      actualDay.Write (l0 + l, c0 + c)
      actualDay.Change (kbd.Down, 0)
    case kbd.Unmark:
      dayattr.Actualize (actualDay, false)
      dayattr.Attrib (actualDay)
      actualDay.Write (l0 + l, c0 + c)
      actualDay.Change (kbd.Down, 0)
    case kbd.Print:
      if period == day.Yearly {
        actualDay.PrintYear (0, 0)
        prt.GoPrint()
      }
    default:
      actualDay.Change (C, d)
    }
  }
  return true
}

func main() {
  if scr.Ok (mode.WXGA) {
    scr.New (0, 0, mode.WXGA) // 80 * 16 = 1280, 25 * 32 = 800
    scr.SetFontsize (fontsize.Huge)
  } else {
    scr.New (0, 0, mode.TXT)
  }
  defer scr.Fin()
  cF, cB := col.Black(), col.LightWhite()
  scr.ScrColours (cF, cB)
  actualDay.Update()
  day.WeekdayF, day.WeekdayB = cF, cB
  day.HolidayB = cB
  day.WeekdayNameF, day.WeekdayNameB = col.Magenta(), cB
  for ! edited() { }
  cal.Fin()
}
