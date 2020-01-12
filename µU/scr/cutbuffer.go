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

func (X *screen) Copy7 (s string, b int) {
  if underX {
    X.XWindow.Copy7 (s, b)
  } else {
    X.Console.Copy7 (s, b)
  }
}

func (X *screen) Paste7 (b int) string {
  if underX {
    return X.XWindow.Paste7 (b)
  }
  return X.Console.Paste7 (b)
}
