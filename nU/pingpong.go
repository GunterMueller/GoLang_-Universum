package main

// (c) Christian Maurer   v. 171126 - license see nsp.go

import ("time"; "nU/env"; "nU/nchan")

func main() {
  me := uint(env.Par1()) - '0'
  c := nchan.New ("wort", me, 1 - me, "jupiter", 123)
  for i := uint(0); i < 3; i++ {
    if me == 1 {
      println (c.Recv().(string)); time.Sleep(3e9)
      c.Send ("pong"); time.Sleep(3e9)
    } else { // me == 0
      c.Send ("ping"); time.Sleep(3e9)
      println (c.Recv().(string)); time.Sleep(3e9)
    }
  }
}
