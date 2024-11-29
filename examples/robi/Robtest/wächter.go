package main

// (c) Christian Maurer   v. 240801

// >>> Lösung des Problems von Prof. Dr. J. Nievergelt
//     "Roboter programmieren" - ein Kinderspiel
//     aus Informatik-Spektrum 22 (1999), S. 364-375

import
  . "Robi"

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

func main () {
  Laden ("Stadt")
  zurMauer()
  wachen()
}
