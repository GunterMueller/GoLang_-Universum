package main

// (c) Christian Maurer   v. 171126 - license see nsp.go

import ("time"; "nU/env"; "nU/nchan")

const (h0 = "venus"; h1 = "mars")

func main() {
  mars := env.Localhost() == h1
  partner := h1; if mars { partner = h0 }
  c := nchan.NewD ("wort", partner, 123)
  for i := uint(0); i < 3; i++ {
    if mars {
      println (c.Recv().(string)); time.Sleep(3e9)
      c.Send ("pong"); time.Sleep(3e9)
    } else {
      c.Send ("ping"); time.Sleep(3e9)
      println (c.Recv().(string)); time.Sleep(3e9)
    }
  }
}
