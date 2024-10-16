package main

// (c) Christian Maurer   v. 241005 - license see µU.go

import (
  "µU/kbd"
  . "µU/col"
  "µU/mode"
  "µU/scr"
)
const (
  m = 7
  mm = 5 * m
)
var (
  i, a, b, c int
  x, y [13]int
  cl Colour
)

func wait () { kbd.Wait (true) }

func line (a, b uint) { scr.Line (x[a], y[a], x[b], y[b]) }

func quadrangle (a, b, c, d uint) { line (a, b); line (b, c); line (c, d); line (d, a) }

func scheren4 (a1, a2, q1, q2, z1, z2 uint) {
/* Idee:
  for t:= 0.0; t < 1.0; t += 0.01 {
    scr.Colour (cl)
    quadrangle (a1, a2, q1 + t * (z1 - q1), q2 + t * (z2 - q2))
    wait
    scr.Colour (Black())
    quadrangle (a1, a2, q1 + t * (z1 - q1), q2 + t * (z2 - q2))
  }
  scr.Colour (cl); quadrangle (a1, a2, z1, z2)
*/
}

func main() {
  scr.New (0, 0, mode.XGA); defer scr.Fin()
  i = 1
  a = 5 * mm
  b = 4 * mm
  c = 3 * mm
  cl = LightBlue()
  scr.ColourF (cl)
  x[1] = 150;           y[1] = 450
  x[2] = x[1] +  a;     y[2] = y[1]
  x[3] = x[2];          y[3] = y[1] - a
  x[4] = x[1];          y[4] = y[3]
  x[5] = x[4] +  9 * m; y[5] = y[4] - 12 * m
  x[6] = x[3] + 12 * m; y[6] = y[3] - 16 * m
  x[7] = x[5] + 12 * m; y[7] = y[5] - 16 * m
  x[8] = x[3];          y[8] = y[3] - a
  x[9] = x[5];          y[9] = y[5] - a
  x[10] = x[1] + 9 * m; y[10] = y[1]
  x[11] = x[10];        y[11] = y[3]
  x[12] = x[10];        y[12] = y[11] - a

  quadrangle (1, 2, 3, 4); line (4, 5); quadrangle (3, 5, 7, 6)
  wait()
  scr.ColourF (LightRed()); line (3, 8); line (5, 9); line (7, 9)
  wait()
  scr.ColourF (Blue()); line (3, 6); line (6, 8); line (5, 7)
  scr.ColourF (LightRed()); line (7, 8)
  wait()
  scr.ColourF (LightGreen()); line (8, 12); line (12, 11); line (11, 3)
  wait()
  scr.ColourF (Black()); line (7, 9); line (9, 12)
  scr.ColourF (Blue()); line (8, 7)
  scr.ColourF (Green()); quadrangle (8, 12, 11, 3)
  wait()
  scr.ColourF (Yellow()); line (11, 10); line (10, 2); line (2, 3)
  wait()
  scr.ColourF (Black()); line (8, 12); line (12, 5); line (3, 8)
  scr.ColourF (Gray()); line (11, 5)
  scr.ColourF (Blue()); line (7, 5); line (7, 8)
  scr.ColourF (Yellow()); line (11, 3)
  wait()
/*
  for i:= 150; i <= 350; i ++ {
    scr.ColourF (cl); line (1, 3); line (2, 3)
    wait()
    scr.ColourF (Black()); line (1, 3); line (2, 3)
  }
*/
}
