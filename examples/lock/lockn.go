package main

// (c) Christian Maurer   v. 241011 - license see µU.go

import (
  "runtime"
  . "µU/lockn"
  "µU/time"
  "µU/rand"
  "µU/scr"
  "µU/errh"
  "µU/N"
)
const
  n = 10
var (
  lock LockerN
  value uint
  done = make(chan bool)
)

func v() { time.Msleep (rand.Natural (10)) }

func count (p uint) {
  for i := uint(0); i < 100000; i++ {
    lock.Lock (p)
    accu := value // "LDA value"
    v()
    accu++        // "INA"
    v()
    value = accu  // "STA value"
    v()
    errh.Hint (N.String(value))
    lock.Unlock (p)
    v()
  }
  done <- true
}

func main() {
  scr.NewWH (0, 0, 64, 16); defer scr.Fin()
  runtime.GOMAXPROCS (6)
/* choose one of the following lockers:
  lock = NewDijkstra (n)
  lock = NewDijkstraGoto (n)
  lock = NewHabermann (n)
  lock = NewBakery (n)
  lock = NewBakery1 (n)
  lock = NewTicket (n)
  lock = NewTiebreaker (n)
  lock = NewFast (n)
  lock = NewKessels (n)
  lock = NewSzymanski (n)
  lock = NewKnuth (n)
  lock = NewDeBruijn (n)
  lock = NewEisenbergMcGuire (n)
  lock = NewChannel (n)
  lock = NewGuardedSelect (n)
*/
  lock = NewDijkstra (n)
  for p := uint(0); p < n; p++ { go count (p) }
  for p := uint(0); p < n; p++ { <-done }
  errh.Error ("", value)
}
