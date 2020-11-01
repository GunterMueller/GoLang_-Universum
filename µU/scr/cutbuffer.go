package scr

// (c) Christian Maurer   v. 190528 - license see µU.go

import
  "µU/str"

func (X *screen) Cut (s *string) {
  n := uint(len(*s))
  if n == 0 { return }
  X.Copy (*s)
  *s = str.New (n)
}

func (X *screen) Copy (s string) {
  if underX {
    X.XWindow.Copy (s)
  } else {
    X.Console.Copy (s)
  }
}

func (X *screen) Paste() string {
  if underX {
    return X.XWindow.Paste()
  }
  return X.Console.Paste()
}
