package main

// (c) Christian Maurer   v. 230924 - license see µU.go

import (
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/errh"
  "µU/N"
)

func main() {
  scr.NewWH (0, 0, 400, 300); defer scr.Fin()
  errh.Hint ("Ende mit Alt+Pause")
  W, B := col.FlashWhite(), col.Black()
  l := scr.NLines()
  loop:
  for {
    b, c, d := kbd.Read()
    scr.Write ("       ", 0, 0)
    scr.MousePointer (c >= kbd.Go)
    if c == kbd.None {
      scr.Write1 (b, 0, 0)
    } else {
      switch c {
      case kbd.Pause:
        if d > 1 { break loop }
      case kbd.Go:
        scr.ColourF (col.Green())
      case kbd.Here, kbd.This, kbd.That:
        scr.ColourF (col.LightGreen())
      case kbd.Drag, kbd.Drop, kbd.Move:
        scr.ColourF (col.Orange())
      case kbd.To, kbd.There, kbd.Thither:
        scr.ColourF (col.Yellow())
      case kbd.Roll:
        scr.ColourF (col.Pink())
      }
      scr.Colours (W, B)
      scr.Write (c.String(), 0, 0)
      if c >= kbd.Go {
        xm, ym := scr.MousePosGr()
        scr.Colours (W, B)
        scr.Write ("    ", l - 2, 0); scr.WriteNat (uint(xm), l - 2, 0)
        scr.Write ("    ", l - 2, 4); scr.WriteNat (uint(ym), l - 2, 4)
        if c == kbd.Go {
          scr.MousePointer (true)
        } else {
          scr.MousePointer (false)
          scr.Colours (W, B)
          scr.Point (xm, ym)
        }
      }
    }
    if d == 0 {
      scr.Write1 (' ', 0, 8)
    } else {
      N.Write (d, 0, 8)
      if c == kbd.Del && d > 0 {
        scr.Cls()
      }
    }
  }
}
