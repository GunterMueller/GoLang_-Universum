package dayattr

// (c) Christian Maurer   v. 211215 - license see µU.go

import (
  "µU/env"
  "µU/font"
  "µU/str"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/day"
  "µU/pseq"
  "µU/files"
  "todo/pdays"
)
const
  length = 8 // of names
type
  attribute = uint
var (
  attrs = pseq.New (byte(0))
  nAttrs uint
  name = make([]string, 0)
  set = make([]pdays.PersistentDays, 0)
  actual attribute
  bx = box.New()
  workdayAF, workdayAB, holidayAF, holidayAB =
    col.LightWhite(), col.Blue(), col.LightWhite(), col.Red()
)

func init() {
  bx.Wd (length)
  bx.Colours (workdayAF, workdayAB)
  files.Cd (env.Gosrc() + "/todo/")
  attrs.Name ("Tagesattribute.kfg")
  s := ""
  for i := uint(0); i < attrs.Num(); i++ {
    attrs.Seek (i)
    s += string(attrs.Get().(byte))
  }
  name, nAttrs = str.SplitByte (s, byte(10))
  for a := uint(0); a < nAttrs; a++ {
    str.Norm (&name[a], length)
    d := pdays.New()
    set = append (set, d)
    set[a].Name (name[a])
  }
}

func normalize() {
  actual = attribute (0)
}

func change (w bool) {
  if w {
    if actual + 1 < nAttrs {
      actual++
    } else {
      actual = attribute (0)
    }
  } else if actual > 0 {
    actual--
  } else {
    actual = attribute (nAttrs - 1)
  }
}

func write (a attribute, visible bool, l, c uint) {
  if visible {
    bx.Write (name[a], l, c)
  } else {
    // bx.Clr (l, c)
    scr.Clr (l, c, length, 1)
  }
}

func writeActual (l, c uint) {
  write (actual, true, l, c)
}

func actualize (d day.Calendarday, b bool) {
  if true /* actual > 0 */ {
    if b {
      set[actual].Ins (d)
    } else {
      set[actual].Del (d)
    }
  }
}

func clr() {
  set[0].Clr()
}

func attrib (d day.Calendarday) {
  if set[actual].Ex (d) {
    if d.IsHoliday() {
      d.Colours (holidayAF, holidayAB)
      d.SetFont (font.Bold)
    } else {
      d.Colours (workdayAF, workdayAB)
      d.SetFont (font.Bold)
    }
  } else if d.IsHoliday() {
    d.Colours (day.HolidayF, day.HolidayB)
    d.SetFont (font.Bold)
  } else {
    d.Colours (day.WeekdayF, day.WeekdayB)
    d.SetFont (font.Roman)
  }
}

func fin() {
  for a := uint(0); a < nAttrs; a++ {
    set[a].Fin()
  }
}
