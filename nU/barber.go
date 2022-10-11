package main

import ("nU/term"; "nU/scr"; "nU/barb"; . "nU/barb/barbscr")

var (
  x barb.Barber
  stop, done chan int = make(chan int), make(chan int)
)

func customer() {
  TakeSeatInWaitingRoom()
  x.Customer()
}

func barber() {
  x.Barber()
  GetNextCustomer()
}

func giveHaircuts() {
  for {
    barber()
    select {
    case <-stop:
      done <- 0
      return
    default:
    }
  }
}

func main () {
  scr.New(); defer scr.Fin()
  term.New(); defer term.Fin()
  Init()
/*/
  x = barb.NewCS()
  x = barb.NewDir()
  x = barb.NewMon()
  x = barb.NewCondMon()
  x = barb.NewAndrews()
/*/
  x = barb.NewSem()
  go giveHaircuts()
  for {
//    scr.Warp (5, 0)
//    scr.Switch (false)
    b := term.Read()
    switch b {
    case term.Esc:
      stop <- 0
      <-done
      return
    case term.Enter:
      go customer()
    }
  }
}
