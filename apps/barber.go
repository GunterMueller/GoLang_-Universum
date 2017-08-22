package main

// (c) murus.org  v. 140615 - license see murus.go

// >>> Nichtsequentielle Programmierung mit Go 1 kompakt

import (
  "murus/env"
  "murus/kbd"
  "murus/scr"
  "murus/barb"
  . "murus/barb/scr"
)
var (
  x barb.Barber
  stop, done chan bool = make (chan bool), make (chan bool)
)

func customer() {
  TakeSeatInWaitingRoom()
  x.Customer()
}

func barber() {
  x.Barber()
  GetNextCustomer ()
}

func giveHaircuts() {
  for {
    barber ()
    select {
    case <-stop:
      done <- true
      return
    default:
    }
  }
}

func main () {
  switch env.Par1() {
  case 'd':
    x = barb.NewDir()     // Semaphorlösung mit direkter Übergabe des g. A. (S. 69)
  case 'm':
    x = barb.NewMon()     // Monitorlösung (S. 145)
  case 'c':
    x = barb.NewCondMon() // Lösung mit einem konditionierten Monitor (S. 146)
  case 'a':
    x = barb.NewAndrews() // solution with conditions in package "sync" due to
                          // Gregory Andrews: Concurrent Programming, p. 293-294
  default:
    x = barb.NewSem()    // Semaphorlösung (S. 67 ff.)
  }
  go giveHaircuts()
  for {
    c, _:= kbd.Command()
    switch c {
    case kbd.Esc:
      stop <- true
      <-done
      scr.Fin()
      return
    case kbd.Enter:
      go customer()
    }
  }
}
