package rob

// (c) Christian Maurer   v. 220809 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
//  "µU/env"
  "µU/str"
  . "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/errh"
  "µU/bn"
//  "µU/files"
)
type
  Richtung = byte; const (
  Nord = Richtung(iota)
  West
  Süd
  Ost
  nRichtungen
)
type
  roboter struct {
                 uint "R's Nummer"
            x, y uint // R's Position
                 Richtung
                 uint16 "Anzahl Klötze in R's Tasche"
                 aktion
        aktionen []aktion
                 }
var (
  zahl = bn.New (3)
  lfdNummer = uint(1)
  sollProtokolliertWerden bool
  vordergrundfarbe, hintergrundfarbe = col.Black(), col.LightWhite()
  randfarbe, mauerfarbe col.Colour = col.White(), col.CinnabarRed()
  bild [5][z]string
  alle [1+M]*roboter
  schrittweise = true
  sokoban bool
  tx = [nRichtungen+1]string {"Nord", "West", "Süd", "Ost", "ohne"}
  hilfe = []string {
    "               Roboter auf der Stelle drehen: Pfeiltasten              ",
    "         Roboter einen Schritt laufen lassen: Pfeiltasten              ",
    "",
    "einen Klotz auf der Stelle ablegen/aufnehmen: Einfüge-/Entfernungstaste",
    "  Klotz(haufen) einen Schritt weiterschieben: Eingabetaste (Enter)     ",
    "                   Klotz(haufen) wegschießen: Umschalt- + Eingabetaste ",
    "",
    "       Mauer setzen und einen Schritt laufen: Anfangstaste (Pos1)      ",
    "    Mauer vor Roboter abreißen, weiterlaufen: Endetaste                ",
    "",
    "            Markierung auf der Stelle setzen: F5-Taste                 ",
    "         Markierung auf der Stelle entfernen: F6-Taste                 ",
    "                 alle Markierungen entfernen: Umschalt- + F6-Taste     ",
    "",
    "            jeweils letzten Zug zurücknehmen: Rücktaste (<-)           ",
    "alle Züge zurücknehmen, d.h. ganz zum Anfang: Umschalt- + Rücktaste    ",
    "",
    "                      Roboterwelt ausdrucken: Drucktaste               ",
    "",
    "                              Editor beenden: Abbruchtaste (Esc)       "}
  sokoHilfe = []string {
    hilfe[0],
    "laufen: Pfeiltasten         Zug zurücknehmen: <-            fertig: Esc",
    "",
    "           Bedienungshinweise für Sokoban folgen irgendwann            " }
)

func fehler (s string) {
  ker.Panic ("Verstoß gegen die Vor. " + s)
  errh.Error0 ("Verstoß gegen die Vor. " + s)
}

func init() {
  scr.NewWH (0, 0, M * z, M * z + 16) // mars: , M * z)
//  files.Cd (env.Gosrc() + "/robitest")
  zahl.Colours (col.Yellow(), col.Blue())
  for i, h := range hilfe { hilfe[i] = str.Lat1 (h) }
  for i, h := range sokoHilfe { sokoHilfe[i] = str.Lat1 (h) }
}

func freieNummer() uint {
  for i := uint(1); i <= M; i++ {
    if alle[i] == nil {
      return i
    }
  }
  return 0
}

var
  ersterSchritt = true

func schreiten() {
  if ! schrittweise { return }
  if ersterSchritt {
    errh.Hint (errh.ToContinueOrNot)
    ersterSchritt = false
  }
  for {
    switch c, _ := Command(); c {
    case Enter:
      return
    case Esc:
      scr.Fin()
      return
    }
  }
}

func neuerRoboter (x, y uint) Roboter {
  if x >= M || y >= M { ker.PrePanic() }
  r := new(roboter)
  n := freieNummer()
  if n == 0 {
    errh.Error0 ("Die Welt ist bereits voller Roboter !")
    return Roboter(nil)
  }
  r.uint = n
  r.x, r.y = x, y
  r.uint16 = Max
  r.Richtung = Süd
  derRoboter[r.x][r.y] = r
  nummer[r.x][r.y] = n
// println ("neuer Roboter Nr.", n, "bei Pos", r.x, r.y)
  alle[n] = r
  nRoboter++
  r.aktionen = make([]aktion, 0)
  return r
}

func (r *roboter) Nummer() uint {
  return r.uint
}

func (r *roboter) LinksDrehen() {
  schreiten()
  switch r.Richtung {
  case Ost:
    r.Richtung = Nord
  default:
    r.Richtung++
  }
  r.aktion = linksDrehen
}

func (r *roboter) LinksDrehenZurück() {
  r.RechtsDrehen()
}

func (r *roboter) RechtsDrehen() {
  schreiten()
  switch r.Richtung {
  case Nord:
    r.Richtung = Ost
  default:
    r.Richtung--
  }
  r.aktion = rechtsDrehen
}

func (r *roboter) RechtsDrehenZurück() {
  r.LinksDrehen()
}

func (r *roboter) InLinkerObererEcke() bool {
  return r.x == 0 && r.y == 0
}

func (r *roboter) vorneRandOderMauer() (uint, uint, bool) {
  switch r.Richtung {
  case Nord:
    if r.y == 0 || mauer[r.x][r.y-1] {
      return r.x, M, false
    }
    return r.x, r.y - 1, true
  case West:
    if r.x == 0 || mauer[r.x-1][r.y] {
      return M, r.y, false
    }
    return r.x - 1, r.y, true
  case Süd:
    if r.y + 1 == M || mauer[r.x][r.y+1] {
      return r.x, M, false
    }
    return r.x, r.y + 1, true
  case Ost:
    if r.x + 1 == M || mauer[r.x+1][r.y] {
      return M, r.y, false
    }
    return r.x + 1, r.y, true
  }
  panic ("")
}

// Liefert (x, M, false), wenn R's Richtung Nord oder Süd ist
// und er vor dem Rand steht, wobei x seine x-Position ist;
// liefert (M, y, false), wenn R's Richtung West oder Ost ist
// und er vor dem Rand steht, wobei y seine y-Position ist.
// Liefert (x, y, true), wenn R nicht vor dem Rand steht,
// wobei (x, y) die Position des Platzes vor ihm ist.
func (r *roboter) vorne() (uint, uint, bool) {
  switch r.Richtung {
  case Nord:
    if r.y == 0 {
      return r.x, M, false
    }
    return r.x, r.y - 1, true
  case West:
    if r.x == 0 {
      return M, r.y, false
    }
    return r.x - 1, r.y, true
  case Süd:
    if r.y + 1 == M {
      return r.x, M, false
    }
    return r.x, r.y + 1, true
  case Ost:
    if r.x + 1 == M {
      return M, r.y, false
    }
    return r.x + 1, r.y, true
  }
  panic ("")
}

func (r *roboter) VorRand() bool {
  _, _, ok := r.vorne()
  return ! ok
}

// Spezifikation analog zur Funktion vorne()
func (r *roboter) hinten() (uint, uint, bool) {
  switch r.Richtung {
  case Nord:
    if r.y + 1 == M { // || mauer[r.x][r.y+1] {
      return r.x, M, false
    }
    return r.x, r.y + 1, true
  case West:
    if r.x + 1 == M { // || mauer[r.x+1][r.y] {
      return M, r.y, false
    }
    return r.x + 1, r.y, true
  case Süd:
    if r.y == 0 { // || mauer[r.x][r.y-1] {
      return r.x, M, false
    }
    return r.x, r.y - 1, true
  case Ost:
    if r.x == 0 { // || mauer[r.x-1][r.y] {
      return M, r.y, false
    }
    return r.x - 1, r.y, true
  }
  panic ("")
}

func (r *roboter) freigeben() {
  derRoboter[r.x][r.y] = (*roboter)(nil)
  nummer[r.x][r.y] = 0
  leerAusgeben (r.x, r.y)
  if klötze[r.x][r.y] == 0 {
    if markiert[r.x][r.y] {
      markeAusgeben (r.x, r.y)
    }
  } else {
    klotzAusgeben (r.x, r.y, markiert[r.x][r.y])
  }
  if mauer[r.x][r.y] {
    mauerAusgeben (r.x, r.y)
  }
}

func (r *roboter) besetzen() {
  derRoboter[r.x][r.y] = r
  nummer[r.x][r.y] = r.uint
}

func (r *roboter) Laufen() {
  if x, y, ok := r.vorne(); ok {
    if mauer[x][y] {
      fehler ("nicht zugemauert")
    }
    schreiten()
    r.freigeben()
/*/
    if klötze[r.x][r.y] > 0 {
      println ("Klotz gefunden bei", uint(r.x), "", uint(r.y))
    }
/*/
    r.x, r.y = x, y
    r.besetzen()
    r.aktion = laufen
  } else {
    fehler ("vor Rand")
    return
  }
}

func (r *roboter) LaufenZurück() {
  if x, y, ok := r.hinten(); ok {
    r.freigeben()
    r.x, r.y = x, y
    r.besetzen()
  } else {
    fehler ("hinten Rand")
    return
  }
}

func (r *roboter) hintenRand() bool {
  _, _, ok := r.hinten()
  return ! ok
}

func (r *roboter) Zurücklaufen() {
  if x, y, ok := r.hinten(); ok {
    if mauer[x][y] { fehler ("nicht zugemauert") }
    schreiten()
    r.freigeben()
    r.x, r.y = x, y
    r.besetzen()
    r.aktion = zurücklaufen
  } else {
    fehler ("hinten Rand")
  }
}

func (r *roboter) Leer() bool {
  return klötze[r.x][r.y] == 0
}

func (r *roboter) NachbarLeer() bool {
  if x, y, ok := r.vorne(); ok {
    return klötze[x][y] == 0
  }
  return false
}

func (r *roboter) HatKlötze() bool {
  return r.uint16 > 0
}

func (r *roboter) AnzahlKlötze() uint {
  return uint(r.uint16)
}

func (r *roboter) Ablegen() {
  if r.uint16 == 0 {
    fehler ("Tasche nicht leer")
    return
  }
  schreiten()
  klötze[r.x][r.y]++
  klotzAusgeben (r.x, r.y, false)
  r.uint16--
  r.aktion = ablegen
}

func (r *roboter) Aufnehmen() {
  if klötze[r.x][r.y] == 0 {
    fehler ("kein Klotz da")
    return
  }
  schreiten()
  klötze[r.x][r.y]--
  r.uint16++
  r.aktion = aufnehmen
}

func (r *roboter) Geschoben() bool {
  if r.NachbarLeer() {
    return false
  }
  ok := false
  x0, y0 := r.x, r.y
  if x1, y1, o1 := r.vorne(); o1 {
    r.x, r.y = x1, y1
    n := klötze[x1][y1]
    if x2, y2, o2 := r.vorne(); o2 {
      klötze[x1][y1] = 0
      k := klötze[x2][y2]
      if k == 0 {
        ok = true
        klötze[x2][y2] += n
        klotzAusgeben (x2, y2, markiert[x2][y2])
      }
    }
  }
  r.x, r.y = x0, y0
  if ok {
    r.Laufen()
    r.aktion = geschoben
    return true
  }
  return false
}

func (r *roboter) geschoben() bool {
  if r.VorRand() || r.VorMauer() || r.NachbarLeer() {
    return false
  }
  x0, y0 := r.x, r.y
  x1, y1, _ := r.vorne()
  if derRoboter[x1][y1] != nil || klötze[x1][y1] != 1 {
    return false
  }
  r.x, r.y = x1, y1
  if r.VorRand() || r.VorMauer() || ! r.NachbarLeer() {
    r.x, r.y = x0, y0
    return false
  }
  klötze[x1][y1] = 0
  x1, y1, _ = r.vorne()
  klötze[x1][y1] = 1
  return true
}

func (r *roboter) Schießen() {
  x0, y0 := r.x, r.y
  for {
    if ! r.geschoben() {
      break
    }
  }
  r.x, r.y = x0, y0
  weltAusgeben()
}

func (r *roboter) Markieren() {
// schreiten()
  if r == nil { ker.Panic ("r == nil") }
  markiert[r.x][r.y] = true
  markenBesitzer[r.x][r.y] = r.uint
  markeAusgeben (r.x, r.y)
  if klötze[r.x][r.y] > 0 {
    klotzAusgeben (r.x, r.y, true)
  }
  r.aktion = markieren
}

func (r *roboter) Entmarkieren() {
  schreiten()
  markiert[r.x][r.y] = false
  markenBesitzer[r.x][r.y] = 0
  if klötze[r.x][r.y] > 0 {
    klotzAusgeben (r.x, r.y, true)
  }
  r.aktion = entmarkieren
}

func (r *roboter) Markiert() bool {
  return markiert[r.x][r.y]
}

func (r *roboter) NachbarMarkiert() bool {
  if x, y, ok := r.vorne(); ok {
    return markiert[x][y]
  }
  return false
}

func (r *roboter) VorMauer() bool {
  if r.VorRand() { return false }
  if x, y, ok := r.vorne(); ok {
    return mauer[x][y]
  }
  return false
}

func (r *roboter) Zumauern() {
  if r.VorRand() {
    return
    fehler ("am Rand")
  }
  if x, y, ok := r.vorne(); ok {
    n := klötze[x][y]
    klötze[x][y] = 0
    r.uint16 += n
    schreiten()
    markiert[x][y] = false
    r.freigeben()
    mauer[r.x][r.y] = true
    mauerAusgeben (r.x, r.y)
    r.x, r.y = x, y
    r.besetzen()
    r.aktion = zumauern
  }
}

func (r *roboter) Entmauern() {
  if r.VorRand() {
    fehler ("vor Rand")
    return
  }
  x, y, _ := r.vorne()
  schreiten()
  mauer[x][y] = false
  leerAusgeben (x, y)
  r.freigeben()
  r.x, r.y = x, y
  r.besetzen()
  r.aktion = entmauern
}

func (r *roboter) Ausgeben() {
  x0, y0 := z * int(r.x), z * int(r.y)
  m := markiert[r.x][r.y]
  if m {
    markeAusgeben (r.x, r.y)
  }
  switch r.Richtung {
  case Nord:
    for y := 0; y < z; y++ {
      for x := 0; x < z; x++ {
        scr.ColourF (r.farbe (bild[1][y][x], m))
        scr.Point (x0 + x, y0 + y)
      }
    }
  case West:
    for y := 0; y < z; y++ {
      for x := 0; x < z; x++ {
        scr.ColourF (r.farbe (bild[1][x][z-1-y], m))
        scr.Point (x0 + x, y0 + y)
      }
    }
  case Süd:
    for y := z - 1; y > 0; y-- {
      for x := 0; x < z; x++ {
        scr.ColourF (r.farbe (bild[1][z-1-y][z-1-x], m))
        scr.Point (x0 + x, y0 + y)
      }
    }
  case Ost:
    for y := 0; y < z; y++ {
      for x := 0; x < z; x++ {
        scr.ColourF (r.farbe (bild[1][z-1-x][y], m))
        scr.Point (x0 + x, y0 + y)
      }
    }
  default:
    ker.Panic ("Roboter hat keine Richtung")
  }
}

var
  erster = true

func (r *roboter) manipulieren (c Comm, d uint) {
  s := schrittweise
  schrittweise = false
  neu := true
  if erster {
    erster = false
    r.Ausgeben()
  }
  switch c {
  case Esc:
    return
  case Enter:
    if d == 0 {
      r.Geschoben()
    } else {
      r.Schießen()
    }
  case Left:
    if sokoban {
      switch r.Richtung {
      case Nord:
        r.LinksDrehen()
      case West:
        ;
      case Süd:
        r.RechtsDrehen()
      case Ost:
        r.LinksDrehen()
        r.LinksDrehen()
      }
      if r.NachbarLeer() {
        r.Laufen()
      } else {
        r.Geschoben()
      }
    } else {
      switch r.Richtung {
      case Nord:
        r.LinksDrehen()
      case West:
        if ! r.VorRand() && ! r.VorMauer() { r.Laufen() }
      case Süd:
        r.RechtsDrehen()
      case Ost:
        r.LinksDrehen()
        r.LinksDrehen()
      }
    }
  case Right:
    if sokoban {
      switch r.Richtung {
      case Nord:
        r.RechtsDrehen()
      case West:
        r.LinksDrehen()
        r.LinksDrehen()
      case Süd:
        r.LinksDrehen()
      case Ost:
        ;
      }
      if r.NachbarLeer() {
        r.Laufen()
      } else {
        r.Geschoben()
      }
    } else {
      switch r.Richtung {
      case Nord:
        r.RechtsDrehen()
      case West:
        r.LinksDrehen()
        r.LinksDrehen()
      case Süd:
        r.LinksDrehen()
      case Ost:
        if ! r.VorRand() && ! r.VorMauer() { r.Laufen() }
      }
    }
  case Up:
    if sokoban {
      switch r.Richtung {
      case Nord:
        ;
      case West:
        r.RechtsDrehen()
      case Süd:
        r.LinksDrehen()
        r.LinksDrehen()
      case Ost:
        r.LinksDrehen()
      }
      if r.NachbarLeer() {
        r.Laufen()
      } else {
        r.Geschoben()
      }
    } else {
      switch r.Richtung {
      case Nord:
        if ! r.VorRand() && ! r.VorMauer() { r.Laufen() }
      case West:
        r.RechtsDrehen()
      case Süd:
        r.LinksDrehen()
        r.LinksDrehen()
      case Ost:
        r.LinksDrehen()
      }
    }
  case Down:
    if sokoban {
      switch r.Richtung {
      case Nord:
        r.LinksDrehen()
        r.LinksDrehen()
      case West:
        r.LinksDrehen()
      case Süd:
        ;
      case Ost:
        r.RechtsDrehen()
      }
      if r.NachbarLeer() {
        r.Laufen()
      } else {
        r.Geschoben()
      }
    } else {
      switch r.Richtung {
      case Nord:
        r.LinksDrehen()
        r.LinksDrehen()
      case West:
        r.LinksDrehen()
      case Süd:
        if ! r.VorRand() && ! r.VorMauer() { r.Laufen() }
      case Ost:
        r.RechtsDrehen()
      }
    }
  case Ins:
    if ! sokoban {
      r.Ablegen()
    }
  case Del:
    if ! sokoban {
      r.Aufnehmen()
    }
  case Pos1:
    if ! sokoban {
      r.Zumauern()
    }
  case End:
    if ! sokoban {
      r.Entmauern()
    }
  case Tab:
    // Roboter wechseln
  case Back:
    na := len(r.aktionen)
    if na > 0 {
      a := r.aktionen[na-1]
      if na > 1 {
        r.aktionen = r.aktionen[:na-1]
      } else {
        r.aktionen = make ([]aktion, 0)
      }
      switch a {
      case linksDrehen:
        r.LinksDrehenZurück()
      case rechtsDrehen:
        r.RechtsDrehenZurück()
      case laufen:
        r.LaufenZurück()
      case zurücklaufen:
//        r.zurücklaufenZurück() // TODO
      case ablegen:
        r.Aufnehmen()
      case aufnehmen:
        r.Ablegen()
      case geschoben:
//        r.GeschobenZurück() // TODO
      case schießen:
//        r.SchießenZurück() // TODO
      case markieren:
        r.Entmarkieren()
      case entmarkieren:
        r.Markieren()
      case zumauern:
        r.Entmauern() // XXX
      case entmauern:
        r.Zumauern() // XXX
      }
    }
  case Help:
    if sokoban {
      errh.Help (sokoHilfe)
    } else {
      errh.Help (hilfe)
    }
//  case Search: // ?
//    // neuen Roboter initialisieren
  case Mark:
    if ! sokoban {
      r.Markieren()
    }
  case Unmark:
    if ! sokoban {
      r.Entmarkieren()
    }
//  case Print:
//    img.Print (0, 0, scr.Wd(), scr.Ht() - scr.Ht1())
  default:
    neu = false
  }
  schrittweise = s
  if neu {
    if c != Back {
      r.aktionen = append (r.aktionen, r.aktion)
    }
    r.Ausgeben()
  }
}

func (r *roboter) Codelen() uint {
  return 6
}

func (r *roboter) Encode() Stream {
  s := make(Stream, 6)
  s[0] = uint8(r.uint)
  s[1] = uint8(r.x)
  s[2] = uint8(r.y)
  s[3] = uint8(r.Richtung)
  copy (s[4:6], Encode(r.uint16))
  return s
}

func (r *roboter) Decode (s Stream) {
  r.uint = uint(s[0])
  r.x = uint(s[1])
  r.y = uint(s[2])
  r.Richtung = Richtung(s[3])
  r.uint16 = Decode(uint16(0), s[4:6]).(uint16)
  r.aktion = nAktionen
  r.aktionen = make([]aktion, 0)
}

func (r *roboter) Pos() (uint, uint) {
  return r.x, r.y
}

func (r *roboter) Set (x, y uint) {
  leerAusgeben (x, y)
  m := markiert[x][y]
  if m {
    r.Markieren()
  }
  if klötze[x][y] > 0 {
    klotzAusgeben (x, y, m)
  }
  r.x, r.y = x, y
}

func editieren() {
  var r *roboter
  weltAusgeben()
  if nRoboter == 0 {
    errh.Hint ("Roboter erzeugen: linke Maustaste")
    scr.MousePointer (true)
    loop0:
    for {
      switch c, _ := Command(); c {
      case Esc:
        break loop0
      case Here:
        ym, xm := scr.MousePos()
        r = neuerRoboter (xm / 4, ym / 2).(*roboter)
        r.Ausgeben()
        break loop0
      }
    }
    errh.DelHint()
  }
  if nRoboter > 0 {
    r = alle[1]
    loop:
    for {
      switch c, d := Command(); c {
      case Esc:
        break loop
      default:
        if c < Go {
          r.Ausgeben()
          r.manipulieren (c, d)
        }
      }
    }
    if sollProtokolliertWerden {
      programmErzeugen()
    }
  }
  speichern()
}

func (r *roboter) farbe (x byte, m bool) col.Colour {
  var v, h col.Colour
  switch r.uint % 16 {
  case 0:
    v, h = col.LightBrown(), col.LightGreen()
  case 1:
    v, h = col.Red(), col.Yellow()
  case 2:
    v, h = col.Green(), col.LightGreen()
  case 3:
    v, h = col.Yellow(), col.DarkCyan()
  case 4:
    v, h = col.Blue(), col.White()
  case 5:
    v, h = col.Gray(), col.LightOrange()
  case 6:
    v, h = col.LightRed(), col.Blue()
  case 7:
    v, h = col.LightBlue(), col.DarkYellow()
  case 8:
    v, h = col.Orange(), col.DarkRed()
  case 9:
    v, h = col.Brown(), col.LightBlue()
  case 10:
    v, h = col.Cyan(), col.Orange()
  case 11:
    v, h = col.Magenta(), col.LightRed()
  case 12:
    v, h = col.DarkYellow(), col.LightCyan()
  case 13:
    v, h = col.LightGreen(), col.DarkBlue()
  case 14:
    v, h = col.Pink(), col.Blue()
  case 15:
    v, h = col.DarkRed(), col.LightGray()
  }
  f := h
  switch x {
  case 'x':
    f = randfarbe
  case 'o':
    f = v
  case 'm':
    if m {
      f = v
    } else {
      f = h
    }
  case ' ':
    f = hintergrundfarbe
  case '+':
    f = vordergrundfarbe
  }
  return f
}

func markeAusgeben (x, y uint) {
  x0, y0 := z * int(x), z * int(y)
  for dy := 0; dy < z; dy++ {
    for dx := 0; dx < z; dx++ {
      n := markenBesitzer[x][y]
      if n == 0 { ker.Panic ("kein Markenbesitzer") }
      r := alle[n]
      scr.ColourF (r.farbe (bild[4][dy][dx], true))
      scr.Point (x0 + dx, y0 + dy)
    }
  }
}

func klotzAusgeben (x, y uint, m bool) {
  r := alle[1]
  if r == nil { ker.Panic ("r == nil") }
  x0, y0 := z * int(x), z * int(y)
  for dy := 0; dy < z; dy++ {
    for dx := 0; dx < z; dx++ {
      scr.ColourF (r.farbe (bild[3][dy][dx], m))
      scr.Point (x0 + dx, y0 + dy)
    }
  }
}

func mauerAusgeben (x, y uint) {
  r := alle[1]
  if r == nil { ker.Panic ("r == nil") }
  x0, y0 := z * int(x), z * int(y)
  for dy := 0; dy < z; dy++ {
    for dx := 0; dx < z; dx++ {
      scr.ColourF (r.farbe (bild[2][dy][dx], false))
      scr.Point (x0 + dx, y0 + dy)
    }
  }
}

func leerAusgeben (x, y uint) {
  x0, y0 := z * int(x), z * int(y)
  scr.ColourF (hintergrundfarbe)
  scr.RectangleFull (x0 + 1, y0 + 1, x0 + z - 1, y0 + z - 1)
  scr.ColourF (randfarbe)
  scr.Rectangle (x0, y0, x0 + z - 1, y0 + z - 1)
}
