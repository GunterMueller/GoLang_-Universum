package ker

// (c) Christian Maurer   v. 201226 - license see µU.go

import (
  "os"
  "strconv"
)
var (
  handler = []func(){}
  finished bool
)

func fin() {
  if finished { return }
  for _, h:= range handler {
    h()
  }
  finished = true
}

func panic_(s string) {
  fin()
  panic (s)
}

func panic1 (s string, n uint) {
  fin()
  panic (s + strconv.Itoa (int(n)))
}

func panic2 (s string, n uint, s1 string, n1 uint) {
  fin()
  panic (s + " " + strconv.Itoa (int(n)) + " " + s1 + " " + strconv.Itoa (int(n1)))
}

func prePanic() {
  fin()
  panic ("precondition not met")
}

func shit() {
  fin()
  panic ("shit happens")
}

func stopErr (t string, n uint, e error) {
  fin()
  s := ""; if e != nil { s = " => " + e.Error() }
  panic ("error nr. " + strconv.Itoa (int(n)) + ": " + t + s)
}

func halt (s int) {
  fin()
  os.Exit (s)
}

func installTerm (h func()) {
  handler = append (handler, h)
}
