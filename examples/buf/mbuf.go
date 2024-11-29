package main

// (c) Christian Maurer   v. 241019 - license see µU.go

// >>> Start the server with "buf 0", the client with "buf 1"

import (
  "µU/kbd"
  "µU/scr"
  "µU/errh"
  "µU/rand"
  "µU/mbuf"
)

func main() {
  scr.NewWH (0, 0, 400, 300); defer scr.Fin()
  buffer := mbuf.New (uint(27))
  for {
    c, _ := kbd.Command()
    switch c {
    case kbd.Help:
      errh.Hint ("in: < key, out: > key, stop: Esc")
    case kbd.Esc:
      return
    case kbd.Left:
      errh.Error ("out", buffer.Get().(uint))
    case kbd.Right:
      n := rand.Natural(10)
      errh.Error ("in", n)
      buffer.Ins (n)
    }
  }
}
