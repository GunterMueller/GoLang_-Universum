package sel

// (c) Christian Maurer   v. 120909 - license see murus.go

import (
  "murus/font"
  "murus/col"
  "murus/scr"
)

func size (c col.Colour) font.Size {
  cF, cH := c, c
  col.Contrast (&cH)
  n := uint(0)
  z, s:= scr.MousePos()
  Select1 (font.Name, font.NSizes, font.M, &n, z, s, cF, cH)
  if n < font.NSizes {
    return font.Size (n)
  }
  return font.Normal
}
