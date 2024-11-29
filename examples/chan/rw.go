package main

// (c) Christian Maurer   v. 231020 - license see µU.go

import (
  "µU/ego"
  "µU/time"
  "µU/rand"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/errh"
  . "µU/rw"
)
const (
  ht = 60; wd = 8
  re = " Reader "
  wr = " Writer "
)
var (
  rw ReaderWriter
  b, cw = col.Black(), col.White()
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

func Read (n uint) {
  write (re, cw, b, n, 0)
  rw.ReaderIn()
  write (re, g, b, n, 0)
  readOrWrite()
  rw.ReaderOut()
  write (re, b, b, n, 0)
}

func Write (n uint) {
  write (wr, cw, b, n, 0)
  rw.WriterIn()
  write (wr, r, b, n, 0)
  readOrWrite()
  rw.WriterOut()
  write (wr, b, b, n, 0)
}

func main() {
  me := ego.Me()
  w := uint(wd * 8)
  scr.NewWH ((w + 8) * me, 0, w, ht * 16); defer scr.Fin()
// choose one of the following implementations (see µU/rw/def.go):
/*/
  rw = New1()
  rw = New2()
  rw = NewSemaphore()
  rw = NewAddS(4)
  rw = NewBaton()
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
  rw = NewFarMonitor ("terra", 5000, me == 0) // replace "terra" by the name of your computer
  rw = NewFarMonitorBounded (3, "terra", 5000, me == 0)
/*/
  rw = NewBaton()
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
      b, c, _ := kbd.Read()
      if c == kbd.Esc {
        return
      }
      if b == 'r' || c == kbd.Left {
        go Read (n)
      }
      if b == 'w' || c == kbd.Right {
        go Write (n)
      }
      if n + 1 < ht {
        n++
      } else {
        n = 0
      }
    }
  }
}
