package main // paint

// (c) murus.org  v. 170814 - license see murus.go

import (
  "murus/env"
//  . "murus/obj"
  "murus/str"
  "murus/kbd"
  "murus/col"
  "murus/scr"
  "murus/box"
  "murus/img"
  "murus/files"
  "murus/sel"
)
type
  figure byte; const (
  line = iota
  rectangle
  circle
  ellipse
  nFigures
)
type
  format byte; const (
  border = iota
  inv
  full
)

func paint (f figure, fo format, x0, y0, x, y int) {
  a := uint(x - x0); if x0 > x { a = uint(x0 - x) }
  b := uint(y - y0); if y0 > x { b = uint(y0 - y) }
  switch f {
  case line:
    switch fo {
    case border:
      scr.Line (x0, y0, x, y)
    case inv:
      scr.LineInv (x0, y0, x, y)
    case full:
      // linewd = Staerke
      // scr.Line (x0, y0, x, y)
  }
  case rectangle:
    switch fo {
    case border:
      scr.Rectangle (x0, y0, x, y)
    case inv:
      scr.RectangleInv (x0, y0, x, y)
    case full:
      scr.RectangleFull (x0, y0, x, y)
    }
  case circle:
    if b > a { a = b }
    switch fo {
    case border:
      scr.Circle (x0, y0, a)
    case inv:
      scr.CircleInv (x0, y0, a)
    case full:
      scr.CircleFull (x0, y0, a)
    }
  case ellipse:
    switch fo {
    case border:
      scr.Ellipse (x0, y0, a, b)
    case inv:
      scr.EllipseInv (x0, y0, a, b)
    case full:
      scr.EllipseFull (x0, y0, a, b)
    }
  }
}

func main() {
  scr.NewMax(); defer scr.Fin()
  files.Cd0()
  symbol := [nFigures]byte { 'l', 'r', 'c', 'e' }
  X, Y := 0, 0
  X1, Y1 := scr.Wd(), scr.Ht()
  colour, paper := col.Black, col.White
  scr.ScrColours (colour, paper)
  scr.Cls()
  paintColour := colour
  scr.ColourF (paintColour)
//  Staerke = 3
  bx := box.New()
  bx.Wd (20)
  bx.Colours (paper, colour)
  name := env.Par(1)
  if str.Empty (name) { name = "temp" }
  scr.Save (0, 0, 20, 1)
  for {
    bx.Edit (&name, 0, 0)
    if ! str.Empty (name) {
      str.OffSpc (&name)
      break
    }
  }
  scr.Restore (0, 0, 20, 1)
  img.Get (name, uint(X), uint(Y))
  scr.MousePointer (true)
  Figur := figure(rectangle)
  var x, y, x0, y0 int
  loop: for {
    scr.ColourF (paintColour)
    scr.MousePointer (true)
    char, cmd, T := kbd.Read()
    switch cmd {
    case kbd.None:
      x, y = scr.MousePosGr()
      scr.Transparence (true)
      scr.Write1Gr (char, x, y - int(scr.Ht1()))
    case kbd.Esc:
      break loop
    case kbd.Back:
      switch T { case 0: x, y = scr.MousePosGr()
        x -= int(scr.Wd1())
        scr.ColourF (paper)
        scr.Write1Gr (' ', x, y - int(scr.Ht1()))
//        scr.RectangleFull (x, y - scr.Ht1(), x + scr.Wd1(), y)
        scr.ColourF (paintColour)
      default:
        scr.Cls()
      }
/*
    case kbd.Ins:
      img.Write (X, Y, X1, Y1 - 16, name)
      box.Edit (Feld, name, scr.Zeilenzahl() - 1, 0)
      img.Get (X, Y, name)
*/
    case kbd.Help:
      paintColour = sel.Colour()
//    case kbd.LookFor:
//      Staerke = Strichstaerken.Staerke()
    case kbd.Enter:
      if T > 0 {
        x0, y0 = scr.MousePosGr()
//        scr.Fill1 (x0, y0)
      }
    case kbd.Print:
      img.Print (uint(X), uint(Y), X1, Y1 - 16)
    case kbd.Tab:
      if T == 0 {
        if Figur + 1 < nFigures { Figur ++ } else { Figur = figure(0) }
      } else {
        if Figur > 0 { Figur -- } else { Figur = figure(nFigures - 1) }
      }
      scr.Colours (col.White, paper)
      scr.Write1 (symbol [Figur], scr.Ht() - 1, 0)
    case kbd.Here:
      x0, y0 = scr.MousePosGr()
      scr.CircleFull (x0, y0, 3 / 2)
    case kbd.Pull:
      x, y = scr.MousePosGr()
      scr.Line (x0, y0, x, y)
      x0, y0 = x, y
    case kbd.Hither:
      x, y = scr.MousePosGr()
      scr.Line (x0, y0, x, y)
    case kbd.There:
      x0, y0 = scr.MousePosGr()
      x, y = x0, y0
    case kbd.Push:
      paint (Figur, inv, x0, y0, x, y)
      x, y = scr.MousePosGr()
      paint (Figur, inv, x0, y0, x, y)
    case kbd.Thither:
      paint (Figur, inv, x0, y0, x, y)
//      scr.Colour (colour)
      x, y = scr.MousePosGr()
      paint (Figur, border, x0, y0, x, y)
      x0, y0 = x, y
    case kbd.This:
      x0, y0 = scr.MousePosGr()
      x, y = x0, y0
    case kbd.Move:
      scr.LineInv (x0, y0, x, y)
      x, y = scr.MousePosGr()
      scr.LineInv (x0, y0, x, y)
    case kbd.Thus:
      scr.LineInv (x0, y0, x, y)
      x, y = scr.MousePosGr()
      scr.Line (x0, y0, x, y)
      x0, y0 = x, y
    }
  }
  scr.Save (0, 0, 20, 1)
  for {
    bx.Edit (&name, 0, 0)
    if ! str.Empty (name) {
      str.OffSpc (&name)
      break
    }
  }
  scr.Restore (0, 0, 20, 1)
  img.Put (name, uint(X), uint(Y), X1, Y1)
}
