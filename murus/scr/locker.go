package scr

// (c) murus.org  v. 170814 - license see murus.go

import (
  "murus/xwin"
  "murus/cons"
)

func lock() {
  if underX {
    xwin.Lock()
  } else {
    cons.Lock()
  }
}

func unlock() {
  if underX {
    xwin.Unlock()
  } else {
    cons.Unlock()
  }
}
