package cons

// (c) murus.org  v. 170818 - license see murus.go

import (
  "murus/ker"
  . "murus/mode"
)
var (
  fullScreen, modus Mode
)

func init() {
  visible = true // TODO -> only in initConsole()
}

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

func (x *console) Fin() {
  ker.Fin()
}
