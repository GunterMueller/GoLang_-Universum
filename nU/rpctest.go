package main

import (. "nU/obj"; "nU/ego"; "nU/scr"; "nU/rpc")

func f (a any, i uint) any {
  p := IntStream{0, 0}
  p = Decode (p, a.(Stream)).(IntStream)
  return p[0] * p[1]
}

func main() {
  me := ego.Me()
  serving := me == 0
  scr.New(); defer scr.Fin()
  input, output := IntStream{7, 8}, 0
  r := rpc.New (input, output, 1, "jupiter", 1234, me == 0, f)
  if serving { // rpc-server is called
    for { }
  } else { // rpc-client
    output = Decode (output, r.F (input, 0).(Stream)).(int)
    println ("7 * 8 =", output)
  }
}
