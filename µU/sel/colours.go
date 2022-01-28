package sel

// (c) Christian Maurer   v. 220124 - license see µU.go

import (
  "µU/col"
  "µU/scr"
  "µU/kbd"
)

var (
  pattern = []col.Colour {col.Black(), col.DarkRed(), col.Red(), col.FlashRed(), col.LightRed(),
                          col.FlashOrange(), col.DarkYellow(), col.Yellow(),
                          col.FlashGreen(), col.Green(), col.DarkGreen(),
                          col.DarkCyan(), col.Cyan(), col.FlashCyan(),
                          col.LightBlue(), col.FlashBlue(), col.Blue(),
                          col.Magenta(), col.FlashMagenta(), col.LightMagenta(),
                          col.White(), col.LightGray(), col.Gray(), col.DarkGray(),
                          col.DarkBrown(), col.Brown(), col.LightBrown(), col.LightWhite()}
  n = len(pattern)
)

func colour (x0, y0 int) (col.Colour, bool) {
  wd, ht := int(scr.Wd()), int(scr.Ht())
  w := wd / n
  if w > 24 { w = 24 }
  if x0 + n * int(w) >= wd { x0 = int(wd) - int(n * w) }
  if y0 + w >= ht { y0 = ht - w }
  scr.SaveGr (x0, y0, uint(n * w), uint(y0 + w))
  c := scr.ScrColF()
  for i := 0; i < n; i++ {
    scr.ColourF (pattern[i])
    scr.RectangleFull (x0 + i * w, y0, x0 + (i + 1) * w, y0 + w)
  }
  scr.MousePointer (true)
  loop:
  for {
    xm, ym := scr.MousePosGr()
    switch comm, _ := kbd.Command(); comm {
    case kbd.Esc, kbd.This:
      break loop
    case kbd.Here, kbd.Drag:
      if x0 <= xm && xm < x0 + n * w && y0 <= ym && ym <= y0 + w {
        i := (xm - x0) / w
        c = pattern[i]
        scr.ColourF (pattern[i])
      } else {
        break loop
      }
      scr.RestoreGr (x0, y0, uint(n * w), uint(w))
      return c, true
    }
  }
  scr.RestoreGr (x0, y0, uint(n * w), uint(w))
  return col.Black(), false
}
