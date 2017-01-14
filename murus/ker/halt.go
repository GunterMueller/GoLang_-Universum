package ker

// (c) murus.org  v. 160101 - license see murus.go

import (
  "os"
  "strconv"
)
var (
  handler = []func(){}
  finished bool
)

func Fin() {
  if finished { return }
  for _, h:= range (handler) {
    h()
  }
  finished = true
}

func Panic (s string) {
  Fin()
  panic (s)
}

func Oops() {
  Panic ("oops")
}

func Panic1 (p string, n uint) {
  Fin()
  panic ("Programm wegen Fehler Nr. " + strconv.Itoa (int(n)) + " im Paket " + p + " abgebrochen")
}

func Stop (p string, n uint) {
  Fin()
  panic ("Programm wegen Fehler Nr. " + strconv.Itoa (int(n)) + " im Paket " + p + " abgebrochen")
}

func StopErr (t string, n uint, e error) {
  Fin()
  s:= ""; if e != nil { s = " => " + e.Error() }
  panic ("Fehler Nr. " + strconv.Itoa (int(n)) + ": " + t + s)
}

func Halt (s int) {
  Fin()
  os.Exit (s)
}

func InstallTerm (h func()) {
  handler = append (handler, h)
}

// func init() { installFin (Fin) } // does not work: attempt to link returns "atexit not defined"
