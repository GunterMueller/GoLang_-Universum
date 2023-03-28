package main

// (c) Christian Maurer   v. 230114 - license see µU.go

import (
  "µU/scr"
  "bus/net"
)

func main () {
//  scr.NewWH is called in bus/line/line.go
  for net.StartAndDestinationSelected() {
    net.ShortestPath()
  }
  scr.Fin()
}
