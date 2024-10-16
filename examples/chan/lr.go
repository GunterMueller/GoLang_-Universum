package main

// (c) Christian Maurer   v. 241007 - license see µU.go

import (
  "µU/time"
  "µU/rand"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  . "µU/lr"
)
const (
  ht = 60
  wd = 20
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
  const t = 2
  time.Msleep (1000 * (t + rand.Natural(t)))
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
  scr.NewWH (0, 0, wd * 8, ht * 16); defer scr.Fin()
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
  lr = NewChannel()
  lr = NewChannelBounded (2, 3)
  lr = NewGuardedSelect()
  lr = NewFarMonitor ("s", 5000, env.N(1) == 0) // s = name of the server
/*/
  lr = NewMutex()
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
