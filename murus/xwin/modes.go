package xwin

// (c) Christian Maurer   v. 170818 - license see murus.go

import
  . "murus/mode"

func maxMode() Mode {
  initX()
  return fullScreen
}

func maxRes() (uint, uint) {
  initX()
  return uint(monitorWd), uint(monitorHt)
}

func ok (m Mode) bool {
  return Wd (m) <= uint(monitorWd) && Ht (m) <= uint(monitorHt)
}
