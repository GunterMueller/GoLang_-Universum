package main

// (c) Christian Maurer   v. 241011 - license see µU.go

import (
  . "sync"
  . "µU/lock"
  "µU/time"
  "µU/rand"
  "µU/scr"
  "µU/errh"
  "µU/N"
)
const
  n = 10
var (
  lock Locker
  value uint
  done = make(chan bool)
)

func v() { time.Msleep (rand.Natural (10)) }

func count (p uint) {
  for i := 0; i < 100000; i++ {
    lock.Lock()
    accu := value // "LDA value"
    v()
    accu++        // "INA"
    v()
    value = accu  // "STA value"
    v()
    errh.Hint (N.String(value))
    lock.Unlock()
    v()
  }
  done <- true
}

func main() {
  scr.NewWH (0, 0, 64, 16); defer scr.Fin()
/* choose one of the following lockers:
  lock = NewCAS()
  lock = NewChannel()
  lock = NewDEC()
  lock = NewMorris()
  lock = NewMutex()
  lock = NewTAS()
  lock = NewUdding()
  lock = NewXCHG()
*/
  lock = NewUdding()
  for p := uint(0); p < n; p++ { go count (p) }
  for p := uint(0); p < n; p++ { <-done }
  errh.Error ("", value)
}
