package date

// (c) Christian Maurer   v. 220804 - license see µU.go

import (
  . "µU/obj"
  "µU/str"
  "µU/font"
  "µU/prt"
  "µU/col"
  "µU/scr"
  "µU/day"
  "µU/clk"
)
type
  daytime struct {
                 day.Calendarday
                 clk.Clocktime
                 }
const
  separator = ','
var
  one, two, three, lastTime clk.Clocktime = clk.New(), clk.New(), clk.New(), clk.New()

func init() {
  one.Set (1, 0, 0)
  two.Set (2, 0, 0)
  three.Set (3, 0, 0)
  lastTime.Set (0, 0, 0)
}

func new_() DayTime {
  return &daytime { day.New(), clk.New() }
}

func (x *daytime) imp (Y any) *daytime {
  y, ok := Y.(*daytime)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *daytime) Set (d day.Calendarday, t clk.Clocktime) {
  x.Calendarday = d.Clone().(day.Calendarday)
  x.Clocktime = t.Clone().(clk.Clocktime)
}

func (x *daytime) Day() day.Calendarday {
  return x.Calendarday.Clone().(day.Calendarday)
}

func (x *daytime) Time() clk.Clocktime {
  return x.Clocktime.Clone().(clk.Clocktime)
}

func (x *daytime) Update() {
  x.Calendarday.Update()
  x.Clocktime.Update()
}

func (x *daytime) Normal() bool {
  if x.Calendarday.Normal() {
    if x.Calendarday.Normal1() { // last Sunday in October
      return ! x.Clocktime.Less (two)
    }
    return true
  }
  if x.Calendarday.Normal1() { // last Sunday in March
    return x.Clocktime.Less (two)
  }
  return false
}

func (x *daytime) Empty() bool {
  return x.Calendarday.Empty() || x.Clocktime.Empty()
}

func (x *daytime) Clr() {
  x.Calendarday.Clr()
  x.Clocktime.Clr()
}

func (x *daytime) Eq (Y any) bool {
  y := x.imp (Y)
  return x.Calendarday.Eq (y.Calendarday) && x.Clocktime.Eq (y.Clocktime)
}

func (x *daytime) Less (Y any) bool {
  return false
}

func (x *daytime) Colours (f, b col.Colour) {
  x.Calendarday.Colours (f, b)
  x.Clocktime.Colours (f, b)
}

func (x *daytime) Cols() (col.Colour, col.Colour) {
  return x.Calendarday.Cols()
}

func (x *daytime) SetFormat (d, c Format) {
  if d < day.NFormats && c < clk.NFormats {
    x.Calendarday.SetFormat (d)
    x.Clocktime.SetFormat (c)
  }
}

func (x *daytime) writeMask (l, c uint) {
  scr.Write1 (separator, l, c + 10) // TODO depends on Format
}

func (x *daytime) Write (l, c uint) {
  x.writeMask (l, c)
  x.Calendarday.Write (l, c)
  x.Clocktime.Write (l, c + 10 + 1 + 1) // TODO depends on Format
}

func (x *daytime) Edit (l, c uint) {
  x.Calendarday.Edit (l, c)
  x.Clocktime.Write (l, c + 10 + 1 + 1) // TODO depends on Format
}

func (x *daytime) SetFont (f font.Font) {
  x.Calendarday.SetFont (f)
  x.Clocktime.SetFont (f)
}

func (x *daytime) printMask (l, c uint) {
  prt.SetFont (font.Roman)
  prt.SetFontsize (font.Normal)
  prt.Print1 (separator, l, c + 10) // TODO depends on Format
}

func (x *daytime) Print (l, c uint) {
  x.printMask (l, c)
  x.Calendarday.Print (l, c)
  x.Clocktime.Print (l, c + 12) // TODO depends on Format
}

func (x *daytime) String() string {
  return x.Calendarday.String() + string(separator) + " " + x.Clocktime.String()
}

func (x *daytime) Defined (s string) bool {
  if p, ok := str.Pos (s, separator); ok {
    return x.Calendarday.Defined (s[:p]) && x.Clocktime.Defined (s[p+1:])
  }
  x.Clr()
  return false
}

func (x *daytime) Copy (Y any) {
  y := x.imp (Y)
  x.Calendarday.Copy (y.Calendarday)
  x.Clocktime.Copy (y.Clocktime)
}

func (x *daytime) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *daytime) Actualize (Y DayTime) {
  y := Y.(*daytime)
  if x.Empty() || y.Empty() { return }
  y.Copy (x)
  if x.Normal() {
    return
  }
  y.Clocktime.Inc (one)
  if y.Clocktime.Less (one) {
    y.Calendarday.Inc (day.Daily)
  }
}

func (x *daytime) Normalize() {
  if x.Empty() { return }
  if x.Calendarday.Normal() {
    if x.Calendarday.Normal1() { // TODO get rid of Leq
/*
      if two.Leq (x.Clocktime) && x.Clocktime.Less (three) {
      // kritischer Fall 2 <= Zeit < 3 (Stunden 2A und 2B): wenn
      // wenn Zeit Leq Zeit beim vorigen Aufruf war, ist es 2B,
      // d.h. es wird nicht mehr eine Stunde zurückgestellt
        if x.Clocktime.Leq (lastTime) {
          return
        }
        lastTime.Copy (x.Clocktime)
      } else {
     // für eventuelle weitere Aufrufe im gleichen Programmlauf:
        lastTime.Set (0, 0, 0)
        if three.Leq (x.Clocktime) {
          return
        }
      }
*/
    } else {
      return
    }
  } else if x.Calendarday.Normal1() && x.Clocktime.Less (three) {
    return
  }
  if x.Clocktime.Less (one) {
    x.Calendarday.Dec (day.Daily)
  }
  x.Clocktime.Dec (one)
}

func (x *daytime) Inc (dt clk.Clocktime) {
  x.Clocktime.Inc (dt)
  if x.Clocktime.Less (dt) {
    x.Calendarday.Inc (day.Daily)
  }
}

func (x *daytime) Dec (dt clk.Clocktime) {
  if x.Clocktime.Less (dt) {
    x.Calendarday.Dec (day.Daily)
  }
  x.Clocktime.Dec (dt)
}

func (x *daytime) Codelen() uint {
  return x.Calendarday.Codelen() + x.Clocktime.Codelen()
}

func (x *daytime) Encode() Stream {
  b := make (Stream, x.Codelen())
  a := x.Calendarday.Codelen()
  copy (b[:a], x.Calendarday.Encode())
  copy (b[a:], x.Clocktime.Encode())
  return b
}

func (x *daytime) Decode (b Stream) {
  a := x.Calendarday.Codelen()
  x.Calendarday.Decode (b[:a])
  x.Clocktime.Decode (b[a:])
}
