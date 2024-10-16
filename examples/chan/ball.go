package main

// (c) Christian Maurer   v. 241005 - license see µU.go

import (
  "µU/clk"
  "µU/time"
  "µU/rand"
  "µU/col"
  "µU/mode"
  "µU/scr"
)
var
  stopped bool

func start (d chan bool) {
  r := int(scr.Ht() / 100)
  XX, YY := int(scr.Wd()), int(scr.Ht())
  x := r + int(rand.Natural (uint(XX - 2 * r)))
  y := r + int(rand.Natural (uint(YY - 2 * r)))
  dx, dy := 0, 0
  for dx == 0 || dy == 0 {
    dx = rand.Integer (1 + r / 2)
    dy = rand.Integer (1 + r / 2)
  }
  colour := col.Rand()
  change := false
  for ! stopped {
    if (x > r && x < XX - r && y > r && y < YY - r) || change {
      change = false
      scr.Lock()
      scr.ColourF (scr.ScrColB())
      scr.CircleFull (x, y, uint(r))
      x += dx
      y += dy
      scr.ColourF (colour)
      scr.CircleFull (x, y, uint(r))
      scr.Unlock()
    }
    if x <= r || x >= XX - r {
      dx = -dx
      change = true
    }
    if y <= r || y >= YY - r {
      dy = -dy
      change = true
    }
    time.Msleep (8)
  }
  d <- true
}

func stop() {
  stopped = true
}

func main() {
  scr.New (0, 0, mode.SVGA); defer scr.Fin()
  go clk.Show()
  done := make (chan bool)
  const n = 10
  for i := 0; i < n; i++ { go start (done) }
  for i := 0; i < n; i++ { <-done }
  stop()
}
