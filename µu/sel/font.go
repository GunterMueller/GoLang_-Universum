package sel

// (c) Christian Maurer   v. 170918 - license see µu.go

import (
  "µu/font"
  "µu/col"
  "µu/scr"
)

func size (c col.Colour) font.Size {
  cF, cH := c, c
  cH.Contrast()
  n := uint(0)
  z, s:= scr.MousePos()
  Select1 (font.Name, font.NSizes, font.M, &n, z, s, cF, cH)
  if n < font.NSizes {
    return font.Size (n)
  }
  return font.Normal
}
