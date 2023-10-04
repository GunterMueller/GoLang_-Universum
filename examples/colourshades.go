package main

// (c) Christian Maurer   v. 230924 - license see µU.go

import (
  "µU/kbd"
  "µU/col"
  "µU/scr"
)

func main() {
  scr.NewMax(); defer scr.Fin()
  const b = 256
  nx, ny := int(scr.Wd()), int(scr.Ht())
  z := ny / 3
  F := col.Black()
  for i := 0; i < nx; i++ {
    F.SetR (uint8(((i * b) / nx) % b))
    scr.ColourF (F)
    scr.Line (i, 0, i, z - 1)
  }
  F = col.Black()
  for i := 0; i < nx; i++ {
    F.SetG (uint8(((i * b) / nx) % b))
    scr.ColourF (F)
    scr.Line (i, z, i, 2 * z - 1)
  }
  F = col.Black()
  for i := 0; i < nx; i++ {
    F.SetB (uint8(((i * b) / nx) % b))
    scr.ColourF (F)
    scr.Line (i, 2 * z, i, 3 * z - 1)
  }
  kbd.Wait (false)
}
