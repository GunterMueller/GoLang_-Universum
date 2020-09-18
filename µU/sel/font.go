package sel

// (c) Christian Maurer   v. 200902 - license see µU.go

import (
  "µU/font"
  "µU/col"
  "µU/scr"
)

func size (c col.Colour) font.Size {
  cF, cH := c, c
  cH.Contrast()
  n := uint(0)
  z, s := scr.MousePos()
  Select1 (font.Name, uint(font.NSizes), font.M, &n, z, s, cF, cH)
  if n < uint(font.NSizes) {
    return font.Size (n)
  }
  return font.Normal
}
