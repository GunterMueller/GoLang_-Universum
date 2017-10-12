package scr

// (c) Christian Maurer   v. 170814 - license see µu.go

import
  "µu/font"

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
