package main

// (c) Christian Maurer   v. 241005 - license see µU.go

import (
  "µU/time"
  "µU/rand"
  "µU/col"
  "µU/scr"
  "µU/errh"
  "µU/menue"
  . "µU/barr"
)
const (
  nBarrs = 15 // number of involved processes
  K = 11 // number of tasks
)
var (
  T menue.Menue
  b Barrier = New (nBarrs)
  ready chan int = make (chan int)
)

func point (l, c uint, C col.Colour) {
  scr.Lock()
  scr.Colours (C, scr.ScrColB())
  scr.Write1 ('*', l, 2 * c)
  scr.Unlock()
}

func process (n uint) {
  for x := uint(0); x < K; x++ {
    point (n, x, col.LightRed())
    time.Sleep (1 + rand.Natural (4))
// >>> if the following 2 code lines are commented out, the program runs
//     with no synchronisation - good for demonstration purposes !
    point (n, x, col.Yellow())
    b.Wait()
    time.Sleep (1)
    point (n, x, col.Green())
  }
  ready <- 0
}

func main() {
  scr.NewWH (0, 0, 2 * K * 8, 2 * nBarrs * 8 + 16); defer scr.Fin()
  for n := uint(0); n < nBarrs; n++ {
    go process (n)
  }
  for n := uint(0); n < nBarrs; n++ {
    _ = <-ready
  }
  errh.Error0 ("alle Prozesse fertig")
}
