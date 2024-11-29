package main

// (c) Christian Maurer   v. 240803 - license see µU.go
//
// >>> depth first search

import
  . "robi"
var
  nBlocks uint

func ahead() bool {
  return ! AtEdge() && ! InFrontOfWall()
}

func search() {
  var leftOk, ok, rightOk bool
  if Marked() {
    return
  }
  Mark()
  if ! Empty() {
    nBlocks++
    Hint ("number of blocks =", nBlocks)
  }
  TurnLeft(); leftOk = ok; TurnRight()
  ok = ahead()
  TurnRight(); rightOk = ok; TurnLeft()
  x, y := Pos()
  if leftOk {
    TurnLeft()
    Run()
    search()
    TurnRight()
  }
  Set (x, y)
  if ok {
    Run()
    search()
  }
  Set (x, y)
  if rightOk {
    TurnRight()
    Run()
    search()
    TurnLeft()
  }
}

func main() {
  Load ("maze")
  search()
  ReportError ("number of blocks =", nBlocks)
  Ready()
}
