package ker

// (c) Christian Maurer   v. 171005 - license see µU.go

import
  "time"

func Sleep (s uint) {
  time.Sleep (time.Duration(s) * 1e9)
}

func Msleep (s uint) {
  time.Sleep (time.Duration(s) * 1e6)
}

func Usleep (s uint) {
  time.Sleep (time.Duration(s) * 1e3)
}

var
  start, now uint

func Mess0() {
  println()
  start = Secµsec()
}

func Mess (s string) {
  now = Secµsec()
//  if now < start { now += 1000000 }
  println (s, now - start)
  start = now
}

func Secmsec() uint {
  t := time.Now()
  return uint(t.Second()) * 1e3 + uint(t.Nanosecond() + 5e5) / 1e6
}

func Secµsec() uint {
  t := time.Now()
  return uint(t.Second()) * 1e6 + uint(t.Nanosecond() + 5e2) / 1e3
}

func Secnsec() uint {
  t := time.Now()
  return uint(t.Second()) * 1e9 + uint(t.Nanosecond())
}

func UpdateTime() (uint, uint, uint) {
  h, m, s := time.Now().Clock()
  return uint(h), uint(m), uint(s)
}

func UpdateDate() (uint, uint, uint) {
  t := time.Now()
  return uint(t.Day()), uint(t.Month()), uint(t.Year())
}

func SecondsSinceUnix() (s uint, us uint64) {
  t := time.Now()
  s, us = uint(t.Unix()), uint64(t.UnixNano())
  return
}
