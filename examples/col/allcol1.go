package main

// (c) Christian Maurer   v. 241003 - license see µU.go

import (
  "µU/kbd"
  . "µU/col"
  "µU/scr"
)

func main() {
  scr.NewMax(); defer scr.Fin()
  ac := AllColours()
  const m = 4
  cs := make([][]Colour, m)
  cs[0] = ac[  0: 64]
  cs[1] = ac[ 64:128]
  cs[2] = ac[128:192]
  cs[2] = ac[192:300]
  const a = 8
  wd, ht := int(scr.Wd()), int(scr.Ht())
  dx, dy := wd / a, ht / a
  draw := true
  i := 0
  for {
    if draw {
      scr.Cls()
      for j, f := range cs[i] {
        x, y := (j * dx) % (a * dx), (j / a) * dy
        scr.ColourF (f)
        scr.RectangleFull (x, y, x + dx - 1, y + dy - 1)
        if f != Black() {
          scr.Colours (FlashWhite(), Black())
          scr.WriteGr (f.String(), x + 40, y + 80)
        }
      }
    }
    draw = true
    switch k, c := kbd.Command(); k {
    case kbd.Esc:
      return
    case kbd.Up:
      if c == 0 {
        if i > 0 {
          i--
        } else {
          draw = false
        }
      } else {
        i = 0
      }
    case kbd.Down:
      if c == 0 {
        if i + 1 < m {
          i++
        } else {
          draw = false
        }
      } else {
        i = m - 1
      }
    default:
      draw = false
    }
  }
}
