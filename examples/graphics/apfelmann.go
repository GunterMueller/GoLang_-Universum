package main

// (c) Christian Maurer   v. 241005 - license see µU.go

import (
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/errh"
)

func f (x, y, a, b float64) (float64, float64) {
  return x * x - y * y + a, 2 * x * y + b
}

func main() {
  scr.NewMax(); defer scr.Fin()
  no := "2"
  errh.HintPos ("eine Nummer von 1 bis 5 eingeben:", 0, 0)
  bx := box.New ()
  bx.Wd (1)
  bx.Edit (&no, 0, 34)
  errh.DelHint()
  scr.Cls()
  var xMid, yMid, enlargement float64
  switch no {
  case "1": // just some examples:
    xMid, yMid, enlargement = -0.8, 0.0, 0.1
  case "3":
    xMid, yMid, enlargement = -0.747, 0.110, 250.
  case "4":
    xMid, yMid, enlargement = -0.745345, 0.11307, 6000.
  case "5":
    xMid, yMid, enlargement = -0.7454285, 0.1130090, 200000.
  default: // case "2":
    xMid, yMid, enlargement = -0.8, 0.0, 0.4
  }
  enlargement = 5./8. / enlargement
  xMin, xMax:= xMid - enlargement, xMid + enlargement
  sf := 3./4.
  yMin, yMax:= yMid - sf * enlargement, yMid + sf * enlargement
  limit := 50.0
  dx, dy:= (xMax - xMin) / float64(scr.Wd()), (yMax - yMin) / float64(scr.Ht())
  c := col.New()
  for y := uint(0); y < scr.Ht() - 1; y++ {
    for x := uint(0); x < scr.Wd() - 1; x++ {
      u, v := 0.0, 0.0
      a, b := xMin + float64(x) * dx, yMin + float64(scr.Ht() - y) * dy
      n:= uint8(0)
      for {
        u, v = f (u, v, a, b)
        n++
        if u * u + v * v > limit {
          c = col.New3 (8 * n % 255, 2 * n % 255, 4 * n % 255)
          break
        } else if n == 255 {
          c = col.Black()
          break
        }
      }
      if c.B() != 0 {
        scr.ColourF (c)
        scr.Point (int(x), int(y))
      }
    }
  }
  kbd.Wait (false)
}
