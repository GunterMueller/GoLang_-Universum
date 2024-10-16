package main

// (c) Christian Maurer   v. 241005 - license see µU.go

import (
  "µU/time"
  "µU/env"
  "µU/rand"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  . "µU/lr"
  "µU/ego"
)
var (
  lr LeftRight
  w, b = col.White(), col.Black()
  ch = make(chan int, 1)
  s = ""
  s0 = "                    "
)

func drive() {
  t := uint(5); time.Sleep (t + rand.Natural(t))
}

func write() {
  for {
    scr.Lock()
    scr.Colours (w, b)
    scr.Write (s0, 0, 0)
    scr.Write (s, 0, 0)
    scr.Unlock()
    time.Msleep (10)
  }
}

func plusL() {
  scr.Lock()
  s += "<"
  scr.Unlock()
}

func minusL() {
  scr.Lock()
  s = s[1:]
  scr.Unlock()
}

func plusR() {
  scr.Lock()
  s += ">"
  scr.Unlock()
}

func minusR() {
  scr.Lock()
  s = s[1:]
  scr.Unlock()
}

func left() {
  lr.LeftIn()
  plusL()
  drive()
  minusL()
  lr.LeftOut()
}

func right() {
  lr.RightIn()
  plusR()
  drive()
  minusR()
  lr.RightOut()
}

func main() {
  me := ego.Me()
  if env.UnderX() {
    scr.NewWH (264 * me, 100, 30 * 8, 10 * 16)
  } else {
    scr.NewMax()
  }
  scr.Colours (w, b)
  defer scr.Fin()
  lr = NewFarMonitor (env.Localhost(), 50000, me == 0)
  if me == 0 { // server
    for {
      switch c, _ := kbd.Command(); c {
        case kbd.Pause:
        return
      }
    }
  } else { // client
    go write()
    for {
      switch c, _ := kbd.Command(); c {
      case kbd.Esc:
        return
      case kbd.Left:
        go left()
      case kbd.Right:
        go right()
      }
    }
  }
}
