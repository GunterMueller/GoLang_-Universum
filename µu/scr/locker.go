package scr

// (c) Christian Maurer   v. 170814 - license see µu.go

import (
  "µu/xwin"
  "µu/cons"
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
