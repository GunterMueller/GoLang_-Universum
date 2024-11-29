package main

// (c) Christian Maurer   v. 241020 - license see µU.go

import (
  "µU/ego"
  "µU/time"
  "µU/rand"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/errh"
  . "µU/lr"
)
const (
  ht, wd = 60, 20
	li = "<<<<<<<<<<<<<<<<<<<<"
  re = ">>>>>>>>>>>>>>>>>>>>"
)
var (
  lr LeftRight
  b, w = col.Black(), col.White()
  l, r = col.LightGreen(), col.LightRed()
  ch = make (chan int, 1)
)

func write (s string, f, b col.Colour, l, c uint) {
  <-ch
  scr.Colours (f, b); scr.Write (s, l, c)
  ch <- 0
}

func drive() {
  const t = 2; time.Msleep (1000 * (t + rand.Natural(t)))
}

func left (n uint) {
  write (li, w, b, n, 0)
  lr.LeftIn()
  write (li, l, b, n, 0)
  drive()
  lr.LeftOut()
  write (li, b, b, n, 0)
}

func right (n uint) {
  write (re, w, b, n, 0)
  lr.RightIn()
  write (re, r, b, n, 0)
  drive()
  lr.RightOut()
  write (re, b, b, n, 0)
}

func main() {
  me := ego.Me()
  w := uint(wd * 8)
  scr.NewWH ((w + 8) * me, 0, w, ht * 16); defer scr.Fin()
// choose one of the following implementations (see µU/lr/def.go):
/*/
  lr = NewMutex()
  lr = NewSemaphore()
  lr = NewBaton()
  lr = NewCriticalSection1()
  lr = NewCriticalSection2()
  lr = NewCriticalSectionBounded (2, 3)
  lr = NewCriticalResource (2, 3)
  lr = NewMonitor1()
  lr = NewMonitor2()
  lr = NewMonitorBounded (2, 3)
  lr = NewConditionedMonitor1()
  lr = NewConditionedMonitor2()
  lr = NewConditionedMonitorBounded (2, 3)
  lr = NewFarMonitor ("terra", 5000, me == 0) // replace "terra" by the name of your computer
  lr = NewChannel()
  lr = NewChannelBounded (2, 3)
  lr = NewGuardedSelect()
/*/
  lr = NewBaton()
  if me == 0 {
    errh.Hint ("I am the server")
    for {
      if c, _ := kbd.Command(); c == kbd.Esc {
        return
      }
    }
  } else {
    ch <- 0
    for n := uint(0);; {
      switch c, _ := kbd.Command(); c {
      case kbd.Esc:
        return
      case kbd.Left:
        go left (n)
      case kbd.Right:
        go right (n)
      }
      if n + 1 < ht {
        n++
      } else {
        n = 0
      }
    }
  }
}
