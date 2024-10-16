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
  . "µU/sem"
)
var (
  haltS Semaphore
  aheadS = true
  doneS = make(chan bool)
)

func verifyS() {
  m := uint (0)
  for aheadS {
    time.Usleep (rand.Natural (3 * 1e6))
    m++
    N.Write (m, 0, 7)
    haltS.V()
  }
}

func falsifyS() {
  m := uint (0)
  for aheadS {
    m++
    N.Write (m, 1, 7)
    haltS.P()
  }
}

func stopS() {
  m := uint (0)
  haltS.P()
  m++
  N.Write (m, 2, 7)
  aheadS = false
  haltS.V()
  doneS <- true
}

func main() {
  scr.NewWH (0, 0, 800, 600); defer scr.Fin()
  haltS = New (0)
  N.Colours (col.Yellow(), col.Black())
  go verifyS()
  go falsifyS()
  go stopS()
  <-doneS
  errh.Error ("program terminated", 0)
  ker.Fin()
}
