package scr

// (c) Christian Maurer   v. 170814 - license see murus.go

func (X *screen) Cut (s string) {
  if len (s) == 0 { return }
  if underX {
    X.XWindow.Cut (s)
  } else {
    X.Console.Cut (s)
  }
}

func (X *screen) Paste() string {
  if underX {
    return X.XWindow.Paste()
  }
  return X.Console.Paste()
}
