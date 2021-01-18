package cons

// (c) Christian Maurer   v. 210105 - license see µU.go

import (
  "µU/ker"
  . "µU/mode"
)
var (
  fullScreen, modus Mode
)

func init() {
  visible = true // TODO -> only in initConsole()
}

func maxMode() Mode {
  width, height = maxResConsole()
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
