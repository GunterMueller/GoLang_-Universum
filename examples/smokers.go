package main

// (c) Christian Maurer   v. 230105 - license see µU.go

// >>> The problem of the cigarette smokers

import (
  "µU/rand"
  . "µU/smok"
)
var (
  smokers Smokers
  done = make(chan bool)
  uLast uint = 3
)

func random() uint {
  u := uint(3)
  for {
    u = rand.Natural(3)
    if u != uLast {
      uLast = u
      break
    }
  }
  return u
}

func agent() {
  for {
    u := random()
    smokers.Agent (u)
    WriteAgent (u)
  }
}

func smoker (u uint) {
  for i := 0; i < 10; i++ {
    smokers.SmokerIn (u)
    WriteSmoker (u)
    smokers.SmokerOut()
  }
  done <- true
}

func main() {
// choose one of the following implementions (see µU/go/smok/def.go):
/*/
  smokers = NewNaive()
  smokers = NewParnas()
  smokers = NewCriticalSection()
  smokers = NewMonitor()
  smokers = NewConditionedMonitor()
  smokers = NewChannel()
/*/
  smokers = NewNaive()

  go agent()
  for u := uint(0); u < 3; u++ {
    go smoker(u)
  }
  <-done
}
