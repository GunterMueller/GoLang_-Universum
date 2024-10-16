package main

// (c) Christian Maurer   v. 241005 - license see µU.go

import (
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/errh"
)
var
  xt, yt, x1, y1 int

func setc (F col.Colour) {
  scr.ColourF (F)
  scr.Line (0, yt, x1, y1)
  yt += 16
}

func w (F col.Colour, y int) {
  scr.ColourF (F)
  scr.Line (0, yt, x1, y)
}

func main() {
  scr.NewWH (0, 0, 1600, 1152); defer scr.Fin()
  errh.Error2 ("Zeilen =", scr.NLines (), "/ Spalten =", scr.NColumns ())
  errh.Error2 ("Graphikzeilen =", scr.Ht(), "/ Graphikspalten =", scr.Wd())
  xt, yt = 0, 0
  x1, y1 = int(scr.Wd() - 1), int(scr.Ht() - 1)

  setc (col.Red())
  setc (col.LightRed())
  setc (col.Orange())
  setc (col.Yellow())
  setc (col.LightGreen())
  setc (col.DarkGreen())
  setc (col.Blue())
  setc (col.LightBlue())
  setc (col.Magenta())
  setc (col.Cyan())

  X, Y := make ([]int, 5), make ([]int, 5)
  X[0], Y[0] = 400,   0
  X[1], Y[1] = 700, 300
  X[2], Y[2] =   0, int(y1)
  X[3], Y[3] =  x1, 400
  X[4], Y[4] = 200, 250
  scr.Polygon (X, Y)

  scr.ColourF (col.Pink())
  for r := scr.Wd() / 4; r <= scr.Wd() / 4 + 15; r++ {
    scr.Circle (int(scr.Wd() / 2), int(scr.Ht() / 2), r)
  }
  scr.ColourF (col.LightOrange())
  for r := scr.Wd() / 4 + 16; r <= scr.Wd() / 4 + 31; r++ {
    scr.Circle (int(scr.Wd() / 2), int(scr.Ht() / 2), r)
  }

  for yt := 400; yt <= 415; yt++ { w (col.Red(), yt) }
  for yt := 416; yt <= 431; yt++ { w (col.Green(), yt) }
  for yt := 432; yt <= 447; yt++ { w (col.Blue(), yt) }

  bx:= box.New ()
  bx.Wd (10)
  bx.Colours (col.FlashWhite(), col.Blue())
  T := "          "
  errh.Hint ("Graphik und Text sind verträglich: geben Sie ein Wort ein !")
  bx.Edit (&T, scr.NLines () - 3, 0)
  errh.Error0 (T)
}
