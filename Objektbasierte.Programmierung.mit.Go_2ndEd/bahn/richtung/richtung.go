package richtung

// (c) Christian Maurer   v. 230107 - license see µU.go

func entgegen (d Richtung) Richtung {
  switch d {
  case Links:
    return Rechts
  case Gerade:
    return Gerade
  case Rechts:
    return Links
  }
  return ND
}
