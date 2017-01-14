package scr

// (c) murus.org  v. 140527 - license see murus.go

import
  "murus/font"

func (X *screen) ActFontsize() font.Size {
  if underX {
    return X.Window.ActFontsize()
  }
  return X.Console.ActFontsize()
}

func (X *screen) SetFontsize (f font.Size) {
  if underX {
    X.Window.SetFontsize (f)
  } else {
    X.Console.SetFontsize (f)
  }
}
