package main

import (
  "µU/scr"
  "µU/col"
  "µU/kbd"
)

func main() {
  const (d = 200; a = 8)
  scr.NewWH (0, 0, 1600, 224); defer scr.Fin()
  all := col.AllColours()
  n := len(all)
  i := 0
  for {
    for j := 0; j < a; j++ {
      c := all[i + j]
      scr.ColourF (c)
      scr.CircleFull (j * d + d/2, d/2, d/2)
      scr.Colours (col.FlashWhite(), col.Black())
      scr.WriteGr ("                        ", j * d, d)
      scr.WriteGr (c.String(), j * d, d)
    }
    switch k, _ := kbd.Command(); k {
    case kbd.Esc:
      return
    case kbd.Right:
      if i + a < n { i++ }
    case kbd.Left:
      if i > 0 { i-- }
    case kbd.Pos1:
      i = 0
    case kbd.End:
      i = n - a
    }
  }
}
