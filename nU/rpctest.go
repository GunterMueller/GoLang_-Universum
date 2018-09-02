package main

// (c) Christian Maurer   v. 180813 - license see nU.go

import (. "nU/obj"; "nU/ego"; "nU/scr"; "nU/rpc")

func f (a Any) Any {
  p := IntStream{0, 0}
  p = Decode (p, a.(Stream)).(IntStream)
  return p[0] * p[1]
}

func main() {
  me := ego.Me()
  scr.New(); defer scr.Fin()
  input, output := IntStream{7, 8}, 0
  r := rpc.New (input, output, "jupiter", 1234, me == 0, f)
  if me == 0 {
    for { }
  } else {
    output = Decode (output, r.F (input).(Stream)).(int)
    println ("7 * 8 =", output)
  }
}
