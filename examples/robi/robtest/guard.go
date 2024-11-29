package main

// (c) Christian Maurer   v. 240803

// >>> Solution of the problem of Prof. Dr. J. Nievergelt
//     "Roboter programmieren" - ein Kinderspiel
//     in Informatik-Spektrum 22 (1999), S. 364-375

import
  . "robi"

func toTheWall() {
  if ! InFrontOfWall() {
    Run()
    toTheWall()
  }
}

func guard() {
  if InFrontOfWall() {
    TurnLeft()
  } else {
    Run()
    TurnRight()
    if ! InFrontOfWall() {
      Run()
    }
  }
  guard()
}

/*/
func rechtsMauer() bool {
  TurnRight()
  if InFrontOfWall() {
    TurnLeft()
    return true
  }
  TurnLeft()
  return false
}
/*/

func main () {
  Load ("city")
  toTheWall()
  guard()
}
