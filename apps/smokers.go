package main

// (c) murus.org  v. 150801 - license see murus.go

import (
  "murus/rand"
  . "murus/menue"
  . "murus/smok/utensil"
  . "murus/smok"
  "murus/smok/scr"
)
var (
  menue = New ("Das Problem der Zigarettenraucher")
  x Smokers
  done chan bool = make (chan bool)
  uLast uint = NUtensils
  round uint
)

func random() uint {
  u:= uint(NUtensils)
  for {
    u = rand.Natural(NUtensils)
    if u != uLast {
      uLast = u
      break
    }
  }
  return u
}

func agent() {
  x.SmokerOut()
  for {
    u:= random()
    x.Agent (u)
    scr.Agent (u)
  }
}

const
  nRounds = 10

func smoker (u uint) {
  for {
    x.SmokerIn (u)
    scr.Smoker (u)
    if round + 1 < nRounds {
      round++
    } else {
      done <- true
      return
    }
    x.SmokerOut()
  }
}

func run() {
  round = 0
  go agent()
  for u:= uint(0); u < NUtensils; u++ {
    go smoker(u)
  }
  <-done
}

func main() {
  sn:= New ("Naive Lösung mit Verklemmungsgefahr")
  sn.Leaf (func() { x = NewNaive(); run() }, false)
  menue.Ins (sn)

  sh:= New ("Lösung mit Helferprozessen nach D. L. Parnas")
  sh.Leaf (func() { x = NewParnas(); run() }, false)
  menue.Ins (sh)

  ss:= New ("Lösung mit einem kritischen Abschnitt")
  ss.Leaf (func() { x = NewCriticalSection(); run() }, false)
  menue.Ins (ss)

  sm:= New ("Monitorlösung")
  sm.Leaf (func() { x = NewMonitor(); run() }, false)
  menue.Ins (sm)

  sc:= New ("Lösung mit einem konditionierten Monitor")
  sc.Leaf (func() { x = NewConditionedMonitor(); run() }, false)
  menue.Ins (sc)
/*
  sp:= New ("Lösung mit Botschaftenaustausch")
  sp.Leaf (func() { x = NewMessage(); run() }, false)
  menue.Ins (sp)
*/
  menue.Exec()
}
