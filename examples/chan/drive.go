package main

// (c) Christian Maurer   v. 241005 - license see µU.go

import (
  "µU/time"
  "µU/rand"
  "µU/col"
  "µU/scr"
  "µU/car"
  "µU/clk"
)

const
  size = 32

func draw (b bool, c col.Colour, t uint, x, y int) {
  car.Draw (b, c, x, (y + 1) * size)
  scr.Flush()
  time.Msleep (t)
  car.Draw (b, scr.ScrColB(), x, (y + 1) * size)
  scr.Flush()
}

func drive (i int, done chan bool) {
//  time.Msleep (rand.Natural (4000))
  time.Msleep (2000)
  colour := col.Rand()
  t := 200 + rand.Natural (400)
  dx := (size / 4 + int(rand.Natural (size))) / 1
  n := int(scr.Wd())
  if rand.Natural (100) % 2 == 1 {
    for x := -size; x <= n; x += dx {
      draw (true, colour, t, x, i)
    }
  } else {
    for x := n; x >= -size; x -= dx {
      draw (false, colour, t, x, i)
    }
  }
  done <- true
}

func main() {
  scr.NewWH (0, 0, 800, 600); defer scr.Fin()
  go clk.Show()
  done := make (chan bool)
  n := int(scr.Ht()) / size - 1
  for i := 0; i < n; i++ { go drive (i, done) }
  for i := 0; i < n; i++ { <-done }
}
