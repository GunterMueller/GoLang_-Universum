package main

// (c) Christian Maurer   v. 241025

import (
  "sync"
  "µU/env"
  "µU/ker"
  "µU/perm"
  "µU/scr"
  "µU/clk"
  "rob"
)
var (
  r []rob.Robot
  time0, time1 []clk.Clocktime
  m sync.Mutex
  done chan bool
)

func rennen (i uint) {
  for n := uint(0); n < 2; n++ {
    for ! r[i].AtEdge() {
      m.Lock()
      r[i].Run()
      rob.WriteWorld()
      m.Unlock()
    }
    m.Lock()
    if n == 0 {
      r[i].TurnLeft()
      r[i].TurnLeft()
    }
    m.Unlock()
  }
  time1[i].Update()
  time1[i].Dec (time0[i])
  done <- true
}

func less (i, k int) bool {
  return time1[i].Less (time1[k])
}

func main() {
  rob.Load ("Rennwelt")
  if rob.NumberRobots() != 0 { ker.Panic ("The racing world must not contain any robot.") }
  var max uint
  m := env.N(1)
  if m == 0 || m > rob.M {
    max = rob.M
  } else if m == 1 {
    max = 2
  } else {
    max = m
  }
  p := perm.New (max)
  r = make([]rob.Robot, max)
  time0, time1 = make([]clk.Clocktime, max), make([]clk.Clocktime, max)
  for i := uint(0); i < max; i++ {
    r[i] = rob.NewRobot (i, 0)
    time0[i] = clk.New()
    time1[i] = clk.New()
  }
  done = make (chan bool)
  rob.WriteWorld()
  for i := uint(0); i < max; i++ {
    j := p.F (i)
    time1[j].Update()
    go rennen (j)
  }
  for i := uint(0); i < max; i++ { <-done }
  n := uint(0)
  z := time1[0]
  for i := uint(1); i < max; i++ {
    if time1[i].Less (z) {
      z = time1[i]
      n = i
    }
  }
  r[n].TurnLeft()
  r[n].TurnLeft()
  r[n].Run()
  r[n].TurnLeft()
  r[n].TurnLeft()
  rob.WriteWorld()
  rob.ReportError ("winner is nr.", n + 1)
  scr.Fin()
}
