package main

// (c) Christian Maurer   v. 240801

import
  . "Robi"

func zumAnfangLaufen() {
  if InLinkerObererEcke() {
    if VorRand() {
      LinksDrehen()
      zumAnfangLaufen()
    } else {
      LinksDrehen()
      if VorRand() {
        RechtsDrehen()
      }
      RechtsDrehen()
    }
  } else {
    if VorRand() {
      LinksDrehen()
    } else {
      Laufen()
    }
    zumAnfangLaufen()
  }
}

func zumRandLaufen() {
  if ! VorRand() {
    Laufen()
    zumRandLaufen()
  }
}

func suchen() {
  if VorRand() {
    LinksDrehen()
    LinksDrehen()
    zumRandLaufen()
    RechtsDrehen()
    Laufen()
    RechtsDrehen()
    suchen()
  } else {
    if ! VorMauer() {
      Laufen()
      suchen()
    }
  }
}

func aussenRechtsherumLaufen() {
  if ! VorMauer() {
    RechtsDrehen()
    if VorMauer() {
      LinksDrehen()
    }
    Laufen()
    aussenRechtsherumLaufen()
  }
}

func innenLinksherumLaufen() {
  RechtsDrehen()
  if VorMauer() {
    LinksDrehen()
    if VorMauer() {
      LinksDrehen()
    }
    Laufen()
    innenLinksherumLaufen()
  }
}

func laufen () {
  zumAnfangLaufen()
  suchen()
  LinksDrehen()
  aussenRechtsherumLaufen()
  innenLinksherumLaufen()
  zumAnfangLaufen()
}

func main() {
  Laden ("Haus")
  laufen()
  Fertig()
}
