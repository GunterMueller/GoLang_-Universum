package main

// (c) Christian Maurer   v. 241005 - license see µU.go

import (
  "µU/kbd"
  . "µU/col"
  "µU/fontsize"
  "µU/scr"
)

func main() {
  scr.NewMax(); defer scr.Fin()
  ac := AllColours()
  n, i := len(ac), 0
  wd, ht := int(scr.Wd()), int(scr.Ht())
  draw := true
  for {
    if draw {
      scr.ColourF (ac[i])
      scr.RectangleFull (0, 0, wd, ht)
      scr.Colours (FlashWhite(), Black())
      scr.SetFontsize (fontsize.Huge)
      scr.WriteGr (ac[i].String(), wd/2 - 64, ht/2)
    }
    draw = true
    switch c, d := kbd.Command(); c {
    case kbd.Esc:
      return
    case kbd.Up:
      if d == 0 {
        if i > 0 {
          i--
        } else {
          draw = false
        }
      } else {
        if i >= 10 {
          i -= 10
        } else {
          draw = false
        }
      }
    case kbd.Down:
      if d == 0 {
        if i + 1 < n {
          i++
        } else {
          draw = false
        }
      } else {
        if i + 10 < n {
          i += 10
        } else {
          draw = false
        }
      }
    case kbd.Pos1:
      i = 0
    case kbd.End:
      i = n - 1
    default:
      draw = false
    }
  }
}
