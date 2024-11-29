package main

// (c) Christian Maurer   v. 241005 - license see µU.go

// >>> Start the first program "pingpong 0", the second with "pingpong 1"

import (
  "time"
  "µU/env"
  "µU/nchan"
)

func main() {
  me := uint(env.Arg1()) - '0'
  c := nchan.New ("wort", me, 1 - me, "terra", 123)
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
