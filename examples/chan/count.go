package main

// (c) Christian Maurer   v. 241005 - license see µU.go

import (
  "sync"
  "µU/clk"
  "µU/time"
  "µU/rand"
  "µU/col"
  "µU/scr"
  "µU/N"
)
const
  Z1 = 1
var
  mutex sync.Mutex

func count (n uint, c chan bool) {
  const X = 5
  var n0, n1 uint
  for i := 0; i < 2 * X; i++ {
    mutex.Lock()
    N.Colours (col.Yellow(), col.Red())
    N.Write (uint(i), n, 2)
    scr.Flush()
    mutex.Unlock()
    time.Msleep (rand.Natural (50))
    if i == X {
      n0 = n + 1
      if n0 >= scr.NLines() { n0 = Z1 }
      go count (n0, make (chan bool))
    }
    n1 = 2 * X + 1
  }
  mutex.Lock()
  N.Colours (col.Yellow(), col.Black())
  N.Write (n1, n, 2)
  scr.Flush()
  mutex.Unlock()
  c <- true
}

func main() {
  scr.NewWH (0, 0, 64, 400); defer scr.Fin()
  go clk.Show()
  done := make (chan bool)
  go count (Z1, make (chan bool))
  <-done
}
