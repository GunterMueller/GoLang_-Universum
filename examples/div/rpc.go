package main

// (c) Christian Maurer   v. 241012 - license see µU.go

// start server with "rpc 0" and then client with "rpc 1"

import (
  . "µU/obj"
  "µU/ego"
  "µU/scr"
  "µU/errh"
  "µU/rpc"
)

func f (a any, i uint) any {
  p := IntStream {0, 0}
  p = Decode (p, a.(Stream)).(IntStream)
  return p[0] * p[1]
}

func main() {
  me := ego.Me()
  scr.NewWH (0, 0, 96, 16); defer scr.Fin()
  input, output := IntStream {7, 8}, 0
  r := rpc.New (input, output, 1, "terra", 1234, me == 0, f)
  if me == 0 { // rpc-server is called
    for { }
  } else { // rpc-client
    output = Decode (output, r.F (input, 0).(Stream)).(int)
    errh.Error ("7 * 8 =", uint(output))
  }
}
