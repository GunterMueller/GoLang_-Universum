package cons

// (c) murus.org  v. 140615 - license see murus.go

import (
  "murus/ker"
  . "murus/mode"
)
var (
  fullScreen, modus Mode
)

func maxMode() Mode {
  width, height = resMaxConsole()
  return ModeOf (width, height)
}

func maxRes() (uint, uint) {
  return Res (maxMode())
}

func ok (m Mode) bool {
  fullScreen = maxMode()
  return Wd(m) <= Wd(fullScreen) && Ht(m) <= Ht(fullScreen)
}

func (x *console) X() uint{
  return uint(x.x)
}

func (x *console) Y() uint{
  return uint(x.y)
}

func (x *console) Wd() uint{
  return x.wd
}

func (x *console) Ht() uint{
  return x.ht
}

func (x *console) Wd1() uint{
  return x.wd1
}

func (x *console) Ht1() uint{
  return x.ht1
}

func (x *console) NLines() uint{
  return x.nLines
}

func (x *console) NColumns() uint{
  return x.nColumns
}

func (x *console) Fin() {
  ker.Fin()
}

func init() {
  visible = true // TODO -> only in initConsole()
}
