package dgra

// (c) Christian Maurer   v. 171011 - license see µu.go

import
  "time"

func Time() string {
  t := time.Now()
  m, s, r := t.Minute(), t.Second(), t.Nanosecond() / 1e5
  a, b := r % 100, r / 100
  return string('0' + m / 10) + string('0' + m % 10) +
         string('0' + s / 10) + string('0' + s % 10) +
         string('0' + a / 10) + string('0' + a % 10) +
         string('0' + b / 10) + string('0' + b % 10)
}
