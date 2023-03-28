package robi

// (c) Christian Maurer   v. 230308 - license see µU.go

import (
  "µU/scr"
  "µU/pseq"
  "rob"
)
var (
  r rob.Roboter
  programmdatei = pseq.New (byte(0))
  sollProtokolliertWerden = false
)

func m() uint { return rob.M }

func linksDrehen() { r.LinksDrehen(); r.Ausgeben() }
func rechtsDrehen() { r.RechtsDrehen(); r.Ausgeben() }
func inLinkerObererEcke() bool { return r.InLinkerObererEcke() }
func vorRand() bool { return r.VorRand() }
func laufen() { r.Laufen(); r.Ausgeben() }
func zurücklaufen() { r.Zurücklaufen(); r.Ausgeben() }
func leer() bool { return r.Leer() }
func nachbarLeer() bool { return r.NachbarLeer() }
func hatKlötze() bool { return r.HatKlötze() }
func anzahlKlötze() uint { return r.AnzahlKlötze() }
func ablegen() { r.Ablegen(); r.Ausgeben() }
func aufnehmen() { r.Aufnehmen(); r.Ausgeben() }
func geschoben() bool { defer r.Ausgeben(); return r.Geschoben() }
func schießen() { r.Schießen(); r.Ausgeben() }
func markieren() { r.Markieren(); r.Ausgeben() }
func entmarkieren() { r.Entmarkieren(); r.Ausgeben() }
func markiert() bool { return r.Markiert() }
func nachbarMarkiert() bool { return r.NachbarMarkiert() }
func vorMauer() bool { return r.VorMauer() }
func zumauern() { r.Zumauern(); r.Ausgeben() }
func entmauern() { r.Entmauern(); r.Ausgeben() }

func prog (Zeile string) {
  for i := 0; i < len(Zeile); i++ {
    programmdatei.Ins (byte(Zeile[i]))
  }
  programmdatei.Ins ('\n')
}

func programmErzeugen() {
  rob.ProgrammErzeugen()
}

func laden (s ...string) {
  rob.Laden (s...)
  r = rob.Nr(1)
}

func editieren() {
  rob.Editieren()
  if sollProtokolliertWerden {
    programmErzeugen()
  }
  scr.Fin()
}

func protokollSchalten (ein bool) {
  rob.ProtokollSchalten (ein)
}

func sokobanSchalten (ein bool) {
  rob.SokobanSchalten (ein)
}

func ausgeben (n uint) {
  rob.Ausgabe (n)
}

func eingabe() uint {
  return rob.Eingabe()
}

func fehlerMelden (s string, n uint) {
  rob.FehlerMelden (s, n)
}

func hinweisAusgeben (s string, n uint) {
  rob.HinweisAusgeben (s, n)
}

func fertig() {
  rob.Fertig()
  scr.Fin()
}

func pos() (uint, uint) {
  return r.Pos()
}

func set(x, y uint) {
  r.Set (x, y)
}
