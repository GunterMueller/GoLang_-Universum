package main

// (c) Christian Maurer   v. 220606

import
  . "robi"

func runToStart() {
  if InUpperLeftCorner() {
    if AtEdge() {
      TurnLeft()
      runToStart()
    } else {
      TurnLeft()
      if AtEdge() {
        TurnRight()
      }
      TurnRight()
    }
  } else {
    if AtEdge() {
      TurnLeft()
    } else {
      Run()
    }
    runToStart()
  }
}

func runToEdge() {
  if ! AtEdge() {
    Run()
    runToEdge()
  }
}

func search() {
  if AtEdge() {
    TurnLeft()
    TurnLeft()
    runToEdge()
    TurnRight()
    Run()
    TurnRight()
    search()
  } else {
    if ! InFrontOfWall() {
      Run()
      search()
    }
  }
}

func aussenRechtsherumRun() {
  if ! InFrontOfWall() {
    TurnRight()
    if InFrontOfWall() {
      TurnLeft()
    }
    Run()
    aussenRechtsherumRun()
  }
}

func innenLinksherumRun() {
  TurnRight()
  if InFrontOfWall() {
    TurnLeft()
    if InFrontOfWall() {
      TurnLeft()
    }
    Run()
    innenLinksherumRun()
  }
}

func laufen () {
  runToStart()
  search()
  TurnLeft()
  aussenRechtsherumRun()
  innenLinksherumRun()
  runToStart()
}

func main() {
  Load ("house")
  Run()
  Ready()
}
