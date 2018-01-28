package time

// (c) Christian Maurer   v. 171217 - license see µU.go

import
  "time"

func sleep (s uint) {
  time.Sleep (time.Duration(s) * 1e9)
}

func msleep (s uint) {
  time.Sleep (time.Duration(s) * 1e6)
}

func usleep (s uint) {
  time.Sleep (time.Duration(s) * 1e3)
}

/*
var
  start, now uint

func mess0() {
  println()
  start = Secµsec()
}

func mess (s string) {
  now = Secµsec()
//  if now < start { now += 1000000 }
  println (s, now - start)
  start = now
}
*/

func secmsec() uint {
  t := time.Now()
  return uint(t.Second()) * 1e3 + uint(t.Nanosecond() + 5e5) / 1e6
}

func secµsec() uint {
  t := time.Now()
  return uint(t.Second()) * 1e6 + uint(t.Nanosecond() + 5e2) / 1e3
}

func secnsec() uint {
  t := time.Now()
  return uint(t.Second()) * 1e9 + uint(t.Nanosecond())
}

func uTime() (uint, uint, uint) {
  h, m, s := time.Now().Clock()
  return uint(h), uint(m), uint(s)
}

func uDate() (uint, uint, uint) {
  t := time.Now()
  return uint(t.Day()), uint(t.Month()), uint(t.Year())
}

func sSU() (s uint, us uint64) {
  t := time.Now()
  s, us = uint(t.Unix()), uint64(t.UnixNano())
  return
}
