package ker

// (c) Christian Maurer   v. 220923 - license see µU.go

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
  panic_(s + " " + strconv.Itoa (int(n)))
}

func panic2 (s string, n uint, s1 string, n1 uint) {
  panic_(s + " " + strconv.Itoa (int(n)) + " " + s1 + " " + strconv.Itoa (int(n1)))
}

func panic3 (s string, n uint, s1 string, n1 uint, s2 string, n2 uint) {
  panic_(s + " " + strconv.Itoa (int(n)) + " " + s1 + " " + strconv.Itoa (int(n1)) + " " + s2 + strconv.Itoa (int(n2)))
}

func prePanic() {
  panic_("precondition not met")
}

func oops() {
  panic_("Oops")
}

func shit() {
  panic_("shit happens")
}

func toDo() {
  panic_("ToDo")
}

func stopErr (t string, n uint, e error) {
  s := ""; if e != nil { s = " => " + e.Error() }
  panic_("error nr. " + strconv.Itoa (int(n)) + ": " + t + s)
}

func halt (s int) {
  fin()
  os.Exit (s)
}

func installTerm (h func()) {
  handler = append (handler, h)
}
