package main

// (c) Christian Maurer   v. 241025

import (
  "sync"
  "µU/env"
  "µU/ker"
  "µU/perm"
  "µU/scr"
  "µU/clk"
  "Rob"
)
var (
  r []Rob.Roboter
  zeit0, zeit1 []clk.Clocktime
  m sync.Mutex
  done chan bool
)

func rennen (i uint) {
  for n := uint(0); n < 2; n++ {
    for ! r[i].VorRand() {
      m.Lock()
      r[i].Laufen()
      Rob.WeltAusgeben()
      m.Unlock()
    }
    m.Lock()
    if n == 0 {
      r[i].LinksDrehen()
      r[i].LinksDrehen()
    }
    m.Unlock()
  }
  zeit1[i].Update()
  zeit1[i].Dec (zeit0[i])
  done <- true
}

func less (i, k int) bool {
  return zeit1[i].Less (zeit1[k])
}

func main() {
  Rob.Laden ("Rennwelt")
  if Rob.AnzahlRoboter() != 0 { ker.Panic ("The racing world must not contain any robot.") }
  var max uint
  m := env.N(1)
  if m == 0 || m > Rob.M {
    max = Rob.M
  } else if m == 1 {
    max = 2
  } else {
    max = m
  }
  p := perm.New (max)
  r = make([]Rob.Roboter, max)
  zeit0, zeit1 = make([]clk.Clocktime, max), make([]clk.Clocktime, max)
  for i := uint(0); i < max; i++ {
    r[i] = Rob.NeuerRoboter (i, 0)
    zeit0[i] = clk.New()
    zeit1[i] = clk.New()
  }
  done = make (chan bool)
  Rob.WeltAusgeben()
  for i := uint(0); i < max; i++ {
    j := p.F (i)
    zeit1[j].Update()
    go rennen (j)
  }
  for i := uint(0); i < max; i++ { <-done }
  n := uint(0)
  z := zeit1[0]
  for i := uint(1); i < max; i++ {
    if zeit1[i].Less (z) {
      z = zeit1[i]
      n = i
    }
  }
  r[n].LinksDrehen()
  r[n].LinksDrehen()
  r[n].Laufen()
  r[n].LinksDrehen()
  r[n].LinksDrehen()
  Rob.WeltAusgeben()
  if env.E() {
    Rob.FehlerMelden ("winner is Nr.", n + 1)
  } else {
    Rob.FehlerMelden ("Sieger ist Nr.", n + 1)
  }
  scr.Fin()
}
