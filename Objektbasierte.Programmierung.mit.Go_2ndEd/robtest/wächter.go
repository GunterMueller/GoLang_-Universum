package main

// (c) Christian Maurer   v. 210508

// >>> Lösung des Problems von Prof. Dr. J. Nievergelt
//         "Roboter programmieren" - ein Kinderspiel
//     aus Informatik-Spektrum 22 (1999), S. 364-375

import
  . "robi"

func zurMauer() {
  if ! VorMauer() {
    Laufen()
    zurMauer()
  }
}

func wachen() {
  if VorMauer() {
    LinksDrehen()
  } else {
    Laufen()
    RechtsDrehen()
    if ! VorMauer() {
      Laufen()
    }
  }
  wachen()
}

func rechtsMauer() bool {
  RechtsDrehen()
  if VorMauer() {
    LinksDrehen()
    return true
  }
  LinksDrehen()
  return false
}

/*/
func seek() {
  if VorMauer() {
    LinksDrehen()
  } else if ! rechtsMauer() {
    Laufen()
    seek()
  }
}

func track() {
  if rechtsMauer() {
    if VorMauer() {
      LinksDrehen()
    } else {
      Laufen()
    }
  } else {
    RechtsDrehen()
    Laufen()
  }
  track()
}
/*/

func main () {
  Laden ("Stadt")
  zurMauer()
  wachen()
/*
  seek()
  track()
*/
}
