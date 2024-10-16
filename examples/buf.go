package main

// (c) Christian Maurer   v. 241005 - license see µU.go

// >>> Start the server with "buf 0", the client with "buf 1"

import (
  "µU/kbd"
  "µU/scr"
  "µU/rand"
  "µU/buf"
)

func main() {
  scr.NewWH (0, 0, 400, 300); defer scr.Fin()
  buffer := buf.New(uint(27))
  for {
    c, _ := kbd.Command()
    switch c {
    case kbd.Esc:
      return
    case kbd.Left:
      println ("raus", buffer.Get().(uint))
    case kbd.Right:
      n := rand.Natural(10)
      println ("rein", n)
      buffer.Ins (n)
    }
  }
}
