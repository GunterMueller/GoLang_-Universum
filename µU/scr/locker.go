package scr

// (c) Christian Maurer   v. 170814 - license see µU.go

import (
  "µU/xwin"
  "µU/cons"
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
