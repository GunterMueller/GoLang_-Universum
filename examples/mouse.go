package main

// (c) Christian Maurer   v. 230924 - license see µU.go

import (
  "µU/kbd"
  "µU/col"
  "µU/scr"
)

func main() {
  scr.NewWH (0, 0, 300, 300); defer scr.Fin()
  scr.ScrColourB (col.Blue())
  scr.Cls()
  const (x = 100; x1 = 200; y = 100; y1 = 200)
  var (
    c [2]col.Colour
    a1, a uint
  )
  c[0], c[1] = col.Black(), col.LightRed()
  scr.MousePointer (true)
  if scr.UnderMouseGr (x, y, x1, y1, 0) {
    a1 = 1
  } else {
    a1 = 0
  }
  scr.ColourF (c[a1])
  scr.RectangleFull (x, y, x1, y1)
  for {
    scr.MousePointer (true)
    switch co, _ := kbd.Command(); co {
    case kbd.Esc:
      return
    case kbd.Go:
      if scr.UnderMouseGr (x, y, x1, y1, 0) {
        a = 1
      } else {
        a = 0
      }
      if a1 != a {
        a1 = 1 - a1
        scr.ColourF (c[a1])
        scr.RectangleFull (x, y, x1, y1)
      }
    }
  }
}
