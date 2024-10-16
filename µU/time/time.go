package time

// (c) Christian Maurer   v. 240903 - license see µU.go

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

func actTime() (uint, uint, uint) {
  h, m, s := time.Now().Clock()
  return uint(h), uint(m), uint(s)
}

func actTimeM() (uint, uint, uint, uint) {
  t := time.Now()
  h, m, s := t.Clock()
  ms := t.Nanosecond() / 1000 
  return uint(h), uint(m), uint(s), uint(ms)
}

func actDate() (uint, uint, uint) {
  t := time.Now()
  return uint(t.Day()), uint(t.Month()), uint(t.Year())
}

func secSinceUnix() (s, us uint) {
  t := time.Now()
  return uint(t.Unix()), uint(t.UnixNano())
}
