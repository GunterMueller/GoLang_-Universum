package scr

// (c) murus.org  v. 150124 - license see murus.go

func (X *screen) Cut (s string) {
  if len (s) == 0 { return }
  if underX {
    X.Window.Cut (s)
  } else {
    X.Console.Cut (s)
  }
}

func (X *screen) Paste() string {
  if underX {
    return X.Window.Paste()
  }
  return X.Console.Paste()
}
