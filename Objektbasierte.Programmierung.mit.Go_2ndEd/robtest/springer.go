package main

// (c) Christian Maurer   v. 220606

import
  . "robi"

func springen (weit, rechts bool) {
  if VorRand() { return }
  Laufen()
  if weit {
    if VorRand() {
      Zurücklaufen()
      return
    } else {
      Laufen()
    }
  }
  if rechts {
    RechtsDrehen()
  } else {
    LinksDrehen()
  }
  if VorRand() {
    if rechts {
      LinksDrehen()
    } else {
      RechtsDrehen()
    }
    Zurücklaufen()
    if weit {
      Zurücklaufen()
    }
    return
  } else {
    Laufen()
  }
  if ! weit {
    if VorRand() {
    } else {
      Laufen()
    }
  }
  Ablegen()
  springen (! weit, ! rechts)
}

func main() {
  Laden ("Springer")
  springen (false, false)
  springen (true, false)
  springen (true, true)
  springen (false, true)
  Fertig()
}
