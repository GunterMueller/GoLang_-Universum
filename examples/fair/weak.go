package main

// (c) Christian Maurer   v. 241005 - license see µU.go

import (
  "µU/ker"
  "µU/time"
  "µU/rand"
  "µU/col"
  "µU/scr"
  "µU/errh"
  "µU/N"
)
var (
  aheadW = true
  haltW = false
  doneW = make(chan bool)
)

func verifyW() {
  m := uint (0)
  for aheadW {
    time.Usleep (rand.Natural (3 * 1e6))
    m++
    N.Write (m, 0, 7)
    haltW = true
  }
}

func falsifyW() {
  m := uint (0)
  for aheadW {
    m++
    N.Write (m, 1, 7)
    haltW = false
  }
  doneW <- true
}

func stopW() {
  m := uint (0)
  for ! haltW {
    m++
    N.Write (m, 2, 7)
  }
  aheadW = false
}

func main() {
  scr.NewWH (0, 0, 800, 600); defer scr.Fin()
  N.Colours (col.Yellow(), col.Black())
  go verifyW()
  go falsifyW()
  go stopW()
  _ = <-doneW
  errh.Error ("program terminated", 0)
  ker.Fin()
}
