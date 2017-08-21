package scr

// (c) murus.org  v. 170814 - license see murus.go

import
  "murus/font"

func (X *screen) ActFontsize() font.Size {
  if underX {
    return X.XWindow.ActFontsize()
  }
  return X.Console.ActFontsize()
}

func (X *screen) SetFontsize (f font.Size) {
  if underX {
    X.XWindow.SetFontsize (f)
  } else {
    X.Console.SetFontsize (f)
  }
}
