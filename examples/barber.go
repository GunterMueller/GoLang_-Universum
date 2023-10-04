package main

// (c) Christian Maurer   v. 230924 - license see µU.go

// The problem of the sleeping barber

import (
  "µU/kbd"
  "µU/scr"
  "µU/barb"
  . "µU/barb/scr"
)
var (
  b barb.Barber
  stop, done chan bool = make (chan bool), make (chan bool)
)

func customer() {
  TakeSeatInWaitingRoom()
  b.Customer()
}

func barber() {
  b.Barber()
  GetNextCustomer()
}

func giveHaircuts() {
  for {
    barber()
    select {
    case <-stop:
      done <- true
      return
    default:
    }
  }
}

func main() {
// choose one of the following implementations (see µU/barb/def.go):
/*/
  b = barb.NewSem()
  b = barb.NewDir()
  b = barb.NewCS()
  b = barb.NewMon()
  b = barb.NewCondMon()
  b = barb.NewAndrews()
/*/
  b = barb.NewSem()

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
