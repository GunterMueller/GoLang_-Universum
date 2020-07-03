package main

// This program must be started on a tty-console or in fullscreen-mode in a GUI !

import ("math/rand"; "nU/scr"; . "nU/smok")

var (
  smokers Smokers
  done = make(chan bool)
  uLast = uint32(3)
)

func random() uint {
  u := uint32(3)
  for {
    u = rand.Uint32() % 3
    if u != uLast {
      uLast = u
      break
    }
  }
  return uint(u)
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
  scr.New(); defer scr.Fin()
  Init()
/*
  smokers = NewNaive()
  smokers = NewParnas()
  smokers = NewCriticalSection()
  smokers = NewMonitor()
  smokers = NewConditionedMonitor()
  smokers = NewChannel()
*/
  smokers = NewMonitor()
  go agent()
  for u := uint(0); u < 3; u++ {
    go smoker(u)
  }
  <-done
}
