package main

// (c) Christian Maurer   v. 230308

import (
  "sync"
  "µU/ker"
  "µU/env"
  "µU/perm"
  "µU/scr"
  "µU/clk"
//  "µU/files"
  "rob"
)
const
  dt  =  1
var (
  r = make([]rob.Roboter, max)
  zeit0, zeit1 = make([]clk.Clocktime, max), make([]clk.Clocktime, max)
  m sync.Mutex
  done chan bool
  max = uint(24)
  p = perm.New (max)
)

func rennen (i uint) {
  for n := uint(0); n < 2; n++ {
    for ! r[i].VorRand() {
      m.Lock()
      r[i].Laufen()
      rob.WeltAusgeben()
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
  rob.Laden ("Rennwelt")
  if rob.AnzahlRoboter() != 0 { ker.Panic ("Die Rennwelt darf keinen Roboter enthalten") }
  m := env.N(1)
  if m == 0 || m > rob.M {
    max = rob.M
  } else if m == 1 {
    max = 2
  } else {
    max = m
  }
  for i := uint(0); i < max; i++ {
    r[i] = rob.NeuerRoboter (i, 0)
    zeit0[i] = clk.New()
    zeit1[i] = clk.New()
  }
  done = make (chan bool)
  rob.WeltAusgeben()
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
  rob.WeltAusgeben()
  rob.FehlerMelden ("Sieger ist Nr.", n + 1)
  scr.Fin()
}
