package v

// (c) Christian Maurer - license see µu.go

import (
  "µu/ker"
  "µu/col"
  "µu/errh"
  "µu/day"
)
const ( // v.
  yy = 17
  mm = 10
  dd = 12
)
var
  v day.Calendarday = day.New()

func init() {
  v.Set(dd, mm, 2000 + yy)
  v.SetFormat(day.Yymmdd)
}

func colours() (col.Colour, col.Colour, col.Colour) {
  return col.LightWhite(), col.LightGreen(), col.BlackGreen()
  return col.LightWhite(), col.WhiteBlue(), col.DarkBlue()
  return col.FlashYellow(), col.WhiteYellow(), col.DarkBlue()
  return col.MuF(), col.New3 (0, 16, 128), col.MuB()
}

func want (y, m, d uint) {
  wanted := day.New()
  wanted.SetFormat(day.Yymmdd)
  if wanted.Set(d, m, 2000 + y) {
    if v.Less(wanted) {
      errh.Error0("Your " + ker.Mu + " " + v.String() + " is outdated. You need " + wanted.String() + ".")
      ker.Halt(-1)
    }
  } else {
    ker.Panic("parameters for v.Want are nonsense")
  }
}
