package main

// (c) Christian Maurer   v. 241025

import
  . "Robi"

func umkehren() {
  LinksDrehen()
  LinksDrehen()
}

func vorKlotzLaufen() {
  if Leer() && NachbarLeer() {
    Laufen()
    vorKlotzLaufen()
  }
}

func zumRandLaufen() {
  if ! VorRand() {
    Laufen()
    zumRandLaufen()
  }
}

func leeren() {
  if ! Leer() {
    Aufnehmen()
    leeren()
  }
}

func MauerSchieben() {
  Entmauern()
  Laufen()
  umkehren()
  Zumauern()
  umkehren()
}

func zumRandLaufenUndLeeren() {
  Aufnehmen()
  if ! VorRand() {
    Laufen()
    zumRandLaufenUndLeeren()
  }
}

func zumAnfangLaufen() {
  if VorRand() {
    LinksDrehen()
    zumAnfangLaufen()
  } else {
    if InLinkerObererEcke() {
      RechtsDrehen()
      if VorRand() {
        LinksDrehen()
      }
    } else {
      Laufen()
      zumAnfangLaufen()
    }
  }
}

func pingpongLaufen() {
  LinksDrehen()
  if VorRand() {
    RechtsDrehen()
    if VorRand () {
      LinksDrehen()
      LinksDrehen()
    } else {
      Laufen()
      RechtsDrehen()
      if ! VorRand () {
        Laufen()
      }
    }
  } else {
    Laufen()
    RechtsDrehen()
    if VorRand() {
      LinksDrehen()
      if VorRand() {
        LinksDrehen()
        if ! VorRand() {
          Laufen()
        }
      } else {
        Laufen()
      }
    } else {
      Laufen()
    }
  }
  pingpongLaufen()
}

func WeltLeeren() {
  zumRandLaufenUndLeeren()
  LinksDrehen()
  if ! VorRand() {
    Laufen()
    LinksDrehen()
    zumRandLaufenUndLeeren()
    RechtsDrehen()
    if ! VorRand() {
      Laufen()
      RechtsDrehen()
      WeltLeeren()
    }
  }
}

// Auf allen aufeinanderfolgenden leeren Plätzen in Robis Richtung ab
// seinem Nachbarplatz liegt jetzt genau ein Klotz und Robi steht jetzt
// in seiner Richtung auf dem letzten dieser vorher freien Plätze.
// In Robis Tasche sind entsprechend weniger Klötze.
func laufenUndLegen() {
  Ablegen()
  if NachbarLeer() {
    Laufen()
    laufenUndLegen()
  }
}

func zurMitteLaufenUndLegen1() {
  if NachbarLeer() {
    laufenUndLegen()
    LinksDrehen()
    if NachbarLeer() {
      Laufen()
      zurMitteLaufenUndLegen1()
    }
  }
}

func laufen (n uint) {
  if n > 0 {
    Laufen()
    laufen (n - 1)
  }
}

func legen (n uint) {
  if n > 0 && HatKlötze() {
    Ablegen()
    legen (n - 1)
  }
}

// Liefert die Anzahl der Klötze, die vorher auf Robis Platz lagen.
// Diese Klötze liegen jetzt nicht mehr auf seinem Platz,
// sondern sind zusätzlich in seiner Tasche.
func AnzahlAufgenommen() uint {
  if Leer() {
    return 0
  }
  Aufnehmen()
  return 1 + AnzahlAufgenommen()
}

func anzahl (n uint) uint {
  if Leer() {
    legen (n)
    return 0
  }
  Aufnehmen()
  return 1 + anzahl (n + 1)
}

// Liefert die Anzahl der Klötze, die auf Robis Platz liegen.
// Genau diese Zahl an Klötzen liegt da immer noch.
func Anzahl() uint {
  return anzahl (0)
}

// Vor.: Zwischen Robis Platz und dem Rand in seiner Richtung ist kein Platz zugemauert.
// Liefert die Anzahl der Plätze in Robis Richtung von Robis Platz vorher bis zum Rand. 
// Robi steht jetzt in seiner Richtung am Rand.
func entfernung (n uint) uint {
  if VorRand() {
    umkehren()
    laufen (n)
    umkehren()
    return 0
  }
  Laufen()
  return 1 + entfernung (n + 1)
}

// Vor.: s. entfernung.
// Liefert die Anzahl der Plätze in Robis Richtung von Robis Platz vorher bis zum Rand. 
func Entfernung() uint {
  return entfernung (0)
}

// Vor.: s. entfernung.
// Liefert die Anzahl der Plätze in Robis Richtung von Robis Platz vorher bis zum Rand. 
// Robi steht jetzt in seiner Richtung am Rand.
func EntfernungGelaufen() uint {
  if VorRand() {
    return 0
  }
  Laufen()
  return 1 + EntfernungGelaufen ()
}

func zurücklaufen (n uint) {
  if n > 0 {
    Zurücklaufen()
    zurücklaufen (n - 1)
  }
}

func AN (a uint) uint {
  Zurücklaufen()
  return a
}

func anzahlNachbar() uint {
  if VorRand() {
    return 0
  }
  Laufen()
  return AN (Anzahl())
}

func anzahlNachbarn (a, n uint) uint {
  if n == 0 {
    return a
  }
  LinksDrehen()
  return anzahlNachbarn (a + anzahlNachbar(), n - 1)
}

// Liefert die Anzahl von Robis Nachbarn in allen 4 Richtungen.
func AnzahlNachbarn() uint {
  return anzahlNachbarn (0, 4)
}

func anzahlGelaufen() uint {
  if Leer() {
    if ! VorRand() { Laufen() }
    return 0
  }
  Aufnehmen()
  return 1 + anzahlGelaufen()
}

func AnzahlBisRandGelaufen() uint {
  if VorRand() {
    return Anzahl()
  }
  return anzahlGelaufen() + AnzahlBisRandGelaufen()
}

func anzahlBisRand (a, n uint) uint {
  if VorRand() {
    zurücklaufen (n)
    return Anzahl()
  }
  Laufen()
  return anzahlBisRand (a + anzahlGelaufen(), n + 1)
}

func AnzahlBisRand() uint {
  return anzahlBisRand (0, 0)
}

func anzahlWelt (a uint, links bool) uint {
  if VorRand() {
    LinksDrehen()
    if VorRand() {
      return a
    }
    Laufen()
    LinksDrehen()
    return anzahlWelt (a + anzahlGelaufen(), ! links)
  }
  return anzahlWelt (a + anzahlGelaufen(), links)
}

// Vor.: Kein Platz ist zugemauert.
// Liefert die Anzahl der Klötze in der Welt.
func AnzahlWelt() uint {
  zumAnfangLaufen()
  return anzahlWelt (0, true)
}

func main() {
  Laden()
  Fertig()
}
