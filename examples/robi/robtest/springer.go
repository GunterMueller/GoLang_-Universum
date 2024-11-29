package main

// (c) Christian Maurer   v. 240801

import
  . "robi"

func jump (far, right bool) {
  if AtEdge() { return }
  Run()
  if far {
    if AtEdge() {
      RunBack()
      return
    } else {
      Run()
    }
  }
  if right {
    TurnRight()
  } else {
    TurnLeft()
  }
  if AtEdge() {
    if right {
      TurnLeft()
    } else {
      TurnRight()
    }
    RunBack()
    if far {
      RunBack()
    }
    return
  } else {
    Run()
  }
  if ! far {
    if AtEdge() {
    } else {
      Run()
    }
  }
  PutDown()
  jump (! far, ! right)
}

func main() {
  Load ("Springer")
  jump (false, false)
  jump (true, false)
  jump (true, true)
  jump (false, true)
  Ready()
}
