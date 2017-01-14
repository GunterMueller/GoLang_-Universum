package scr

// (c) murus.org  v. 140527 - license see murus.go

import (
  "murus/xker"; "murus/cons"
)

func lock() {
  if underX {
    xker.Lock()
  } else {
    cons.Lock()
  }
}

func unlock() {
  if underX {
    xker.Unlock()
  } else {
    cons.Unlock()
  }
}
