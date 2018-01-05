package main

import ("nU/col"; "nU/scr")

var l = uint(0)

func w (f col.Colour) {
  scr.ColourF (f)
  scr.Write ("ABCDEFGHIJKLMNOPQR", l, 0)
  l++
}

func main() {
  scr.New(); defer scr.Fin()
  scr.ColourF (col.Red())
  scr.Circle (20, 40, 10)
  scr.ColourF (col.Yellow())
  scr.Circle (40, 40, 20)
  scr.ColourF (col.Cyan())
  scr.Circle (40, 20, 10)
// return
  for i := byte(0); i < 250; i += 10 {
    scr.ColourF (col.New3(127 - i / 2, 127 + i / 2, i))
    scr.Write ("ABCDEFGHIJ", uint(i/10), 0)
  }
  w (col.Black())
  w (col.Brown())
  w (col.Red())
  w (col.LightRed())
  w (col.Yellow())
  w (col.LightGreen())
  w (col.Green())
  w (col.Cyan())
  w (col.LightCyan())
  w (col.LightBlue())
  w (col.Blue())
  w (col.Magenta())
  w (col.LightMagenta())
  w (col.Gray())
  w (col.White())
  w (col.LightWhite())
  scr.ColourF (col.Red())
//  scr.Warp (50, 0)
//  scr.Switch (true)
//  return
  scr.Line (1, 1, 25, 1)
  scr.Write ("*", 1, 1); scr.Write ("*", 25, 1)
  scr.ColourF (col.Yellow())
  scr.Line (1, 1, 1, 50)
  scr.Write ("*", 1, 1); scr.Write ("*", 1, 50)
  scr.ColourF (col.Red())
  scr.Line (1, 1, 10, 60)
  scr.Write ("*", 1, 1); scr.Write ("*", 10, 60)
  scr.ColourF (col.Green())
  scr.Line (25, 1, 1, 30)
  scr.Write ("*", 25, 1); scr.Write ("*", 1, 30)
  scr.ColourF (col.Blue())
  scr.Line (1, 40, 10, 1)
  scr.Write ("*", 1, 40); scr.Write ("*", 10, 1)
  scr.ColourF (col.Magenta())
  scr.Line (25, 20, 1, 10)
  scr.Write ("*", 25, 20); scr.Write ("*", 1, 10)
  scr.ColourF (col.Red())
  scr.Write ("Affe", 0, 20)
  scr.ColourF (col.Blue())
  scr.Write ("Esel", 12, 25)
  scr.ColourF (col.Green())
  scr.Write ("Geier", 15, 40)
//  scr.Warp (30, 0)
}
