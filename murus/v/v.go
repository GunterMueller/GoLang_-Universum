package v

// (c) murus.org - license see murus.go

import (
  "murus/ker"; "murus/col"; "murus/errh"; "murus/day"
)
const ( // v.
  yy = 17
  mm =  1
  dd = 13
)
var
  v day.Calendarday = day.New()

func init() {
  v.Set(dd, mm, 2000 + yy)
  v.SetFormat(day.Yymmdd)
}

func Colours() (col.Colour, col.Colour, col.Colour) {
  return col.LightWhite, col.WhiteBlue, col.DarkBlue
  return col.FlashYellow, col.WhiteYellow, col.DarkBlue
  return col.MurusF, col.Colour3 (0, 16, 128), col.MurusB
}

func String() string {
  return v.String()
}

func Want(y, m, d uint) {
  wanted:= day.New()
  wanted.SetFormat(day.Yymmdd)
  if wanted.Set(d, m, 2000 + y) {
    if v.Less(wanted) {
      errh.Error0("Your murus " + v.String() + " is outdated. You need " + wanted.String() + ".")
      ker.Halt(-1)
    }
  } else {
    ker.Panic("parameters for v.Want are nonsense")
  }
}
