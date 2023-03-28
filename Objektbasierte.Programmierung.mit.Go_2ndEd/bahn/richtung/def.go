package richtung

// (c) Christian Maurer   v. 230107 - license see µU.go

type
  Richtung byte; const (
  Links = Richtung(iota)
  Gerade
  Rechts
  ND
)
var
  Dtext = [ND+1]string {"Links", "Gerade", "Rechts", "ND"}

func Entgegen (d Richtung) Richtung { return entgegen(d) }
