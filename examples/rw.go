package main

// (c) Christian Maurer   v. 230105 - license see µU.go

// >>> Start the first competitor with "rw 0", the second with "rw 1", and so on

import (
  "µU/time"
  "µU/rand"
  "µU/env"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  . "µU/rw"
)
const (
  ht = 50; wd = 20
  s = "####################"
)
var (
  rw ReaderWriter
  b, w, y = col.Black(), col.White(), col.LightYellow()
  r, g = col.Red(), col.Green()
  ch = make(chan int, 1)
)

func write (s string, f, b col.Colour, n, c uint) {
  <-ch
  scr.Colours (f, b); scr.Write (s, n, c)
  ch <- 0
}

func readOrWrite() {
  const t = 5; time.Msleep (1000 * (t + rand.Natural(t)))
}

func goR (n uint) {
  write (s, w, b, n, 0)
  rw.ReaderIn()
  write (s, g, b, n, 0)
  readOrWrite()
  rw.ReaderOut()
  write (s, b, b, n, 0)
}

func goW (n uint) {
  write (s, w, b, n, 0)
  rw.WriterIn()
  write (s, r, b, n, 0)
  readOrWrite()
  rw.WriterOut()
  write (s, b, b, n, 0)
}

func main() {
  x := uint(0) + env.N(1) * (wd * 8 + 8)
  scr.NewWH (x, 0, wd * 8, ht * 16); defer scr.Fin()
// choose one of the following implementations (see µU/rw/def.go):
/*/
  rw = New1()
  rw = New2()
  rw = NewSemaphore()
  rw = NewAddS(4)
  rw = NewGo()
  rw = NewCriticalSection1()
  rw = NewCriticalSection2()
  rw = NewCriticalSectionBounded(3)
  rw = NewCriticalSectionFair()
  rw = NewCriticalResource (3)
  rw = NewMonitor1()
  rw = NewMonitor2()
  rw = NewConditionedMonitor()
  rw = NewConditionedMonitorBounded(3)
  rw = NewChannel()
  rw = NewGuardedSelect()
  rw = NewKangLee()
  rw = NewFarMonitor ("s", 5000, env.N(1) == 0) // s = name of the server
  rw = NewFarMonitorBounded (3, "s", 5000, env.N(1) == 0)
/*/
  rw = New1()

  ch <- 0
  for n := uint(0);; {
    c, _ := kbd.Command()
    switch c {
    case kbd.Esc:
      return
    case kbd.Left:
      go goR (n)
    case kbd.Right:
      go goW (n)
    }
    if n + 1 < ht {
      n++
    } else {
      n = 0
    }
  }
}
