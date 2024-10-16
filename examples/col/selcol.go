package main

// (c) Christian Maurer   v. 241005 - license see µU.go

import (
  "µU/kbd"
  "µU/col"
  "µU/mode"
  "µU/scr"
  "µU/N"
)

func main() {
  scr.New (0, 0, mode.HVGA)
  var r, g, b byte
  for {
    colour := col.New3 (r, g, b)
    scr.ScrColourB (colour)
    scr.Cls()
    scr.Colours (col.FlashWhite(), col.Black())
    scr.Write (N.String (uint(r)), 10, 25)
    scr.Write (N.String (uint(g)), 10, 30)
    scr.Write (N.String (uint(b)), 10, 35)
    switch c, d := kbd.Command(); c {
    case kbd.Esc:
      scr.Fin(); return
    case kbd.Paste:
      if d == 0 {
        if r > 0 && g > 0 && b > 0 { r--; g--; b-- }
      } else {
        if r < 255 && g < 255 && b < 255 { r++; g++; b++ }
      }
    case kbd.Red:
      if d == 0 {
        if r < 255 { r++ }
      } else {
        if r > 0 { r-- }
      }
    case kbd.Green:
      if d == 0 {
        if g < 255 { g++ }
      } else {
        if g > 0 { g-- }
      }
    case kbd.Blue:
      if d == 0 {
        if b < 255 { b++ }
      } else {
        if b > 0 { b-- }
      }
    }
  }
}
