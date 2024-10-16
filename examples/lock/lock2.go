package main

// (c) Christian Maurer   v. 241011 - license see µU.go

import (
  "runtime"
  . "µU/lock2"
  "µU/time"
  "µU/rand"
  "µU/scr"
  "µU/errh"
  "µU/N"
)
var (
  lock Locker2
  done = make(chan int)
  value uint
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
  done <- 0
}

func main() {
  scr.NewWH (0, 0, 64, 16)
  runtime.GOMAXPROCS (2)
/* choose one of the following lockers:
  lock = NewDekker()
  lock = NewDoranThomas()
  lock = NewKessels()
  lock = NewPeterson()
*/
  lock = NewPeterson()
  go count(0); go count(1)
  <-done; <-done
  errh.Error ("", value)
}
