package sel

// (c) Christian Maurer   v. 220110 - license see µU.go

import (
  "µU/font"
  "µU/col"
  "µU/scr"
)

func fontsize (f, b col.Colour) font.Size {
  n := uint(0)
  scr.MousePointer (true)
  z, s := scr.MousePos()
  Select1 (font.Name, uint(font.NSizes), font.M, &n, z, s, f, b)
  if n < uint(font.NSizes) {
    return font.Size (n)
  }
  return font.Normal
}
