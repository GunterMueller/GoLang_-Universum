package rob

// (c) Christian Maurer   v. 230309 - license see µU.go

import (
  "µU/ker"
  "µU/env"
  . "µU/obj"
  . "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/errh"
  "µU/pseq"
  "µU/bn"
  "µU/day"
  "µU/clk"
  "µU/gra"
)
const
  z = 32 // Pixelgröße eines Platzes
type
  aktion = byte; const (
  linksDrehen = aktion(iota)
  rechtsDrehen
  laufen
  zurücklaufen
  ablegen
  aufnehmen
  geschoben
  schießen
  markieren
  entmarkieren
  zumauern
  entmauern
  nAktionen
)
var (
  nRoboter uint
  derRoboter [M][M]*roboter
  nummer [M][M]uint
  klötze [M][M]uint16
  markiert [M][M]bool
  markenBesitzer [M][M]uint
  mauer [M][M]bool
  aktionstext = []string {"LinksDrehen",
                          "RechtsDrehen",
                          "Laufen",
                          "Zurücklaufen",
                          "Ablegen",
                          "Aufnehmen",
                          "Geschoben",
                          "Schießen",
                          "Markieren",
                          "Entmarkieren",
                          "Zumauern",
                          "Entmauern"}
  programmdatei = pseq.New (byte(0))
  datei = pseq.New (byte(0))
  name string
  graph gra.Graph
)

func init() {
  n := bn.New (5)
  graph = gra.New (false, n, nil)
  for y := uint(0); y < M; y++ {
    for x := uint(0); x < M; x++ {
      nummer[x][y] = 0
      klötze[x][y] = 0
      markiert[x][y] = false
      markenBesitzer[x][y] = 0
      mauer[x][y] = false
    }
  }
}

func weltAusgeben() {
  for y := uint(0); y < M; y++ {
    for x := uint(0); x < M; x++ {
      n := nummer[x][y]
      if n > 0 {
        r := alle[n]
        if r == nil { ker.Panic ("ausgeben r == nil") }
        r.Ausgeben()
      } else if mauer[x][y] {
        mauerAusgeben (x, y)
      } else if klötze[x][y] == 0 {
        if markiert[x][y] {
          markeAusgeben (x, y)
        } else {
          leerAusgeben (x, y)
        }
      } else {
        klotzAusgeben (x, y, markiert[x][y])
      }
    }
  }
}

func codelen() uint {
  return M * M * (1 + 2 + 1 + 1 + 1 + 1) + uint(nRoboter) * 6
}

func encode() Stream {
  s := make(Stream, codelen())
  i := uint(0)
  for x := uint(0); x < M; x++ {
    for y := uint(0); y < M; y++ {
      s[i] = uint8(nummer[x][y])
      n := s[i]
      i++
      if n > 0 {
        copy (s[i:i+6], derRoboter[x][y].Encode())
        i += 6
      }
      copy (s[i:i+2], Encode (klötze[x][y]))
      i += 2
      s[i] = 0
      i++
      s[i] = 0; if markiert[x][y] { s[i] = 1 }
      i++
      s[i] = uint8(markenBesitzer[x][y])
      i++
      s[i] = 0; if mauer[x][y] { s[i] = 1 }
      i++
    }
  }
  return s
}

func decode (s Stream) {
  i := uint(0)
  nRoboter = 0
  for x := uint(0); x < M; x++ {
    for y := uint(0); y < M; y++ {
      n := uint(s[i])
      i++
      nummer[x][y] = n
      if n > 0 {
        derRoboter[x][y] = neuerRoboter (uint(x), uint(y)).(*roboter)
        derRoboter[x][y].Decode (s[i:i+6])
        alle[n] = derRoboter[x][y]
        nRoboter++
        i += 6
      }
      klötze[x][y] = Decode (uint16(0), s[i:i+2]).(uint16)
      i += 2
      i++
      markiert[x][y] = false; if s[i] == 1 { markiert[x][y] = true }
      i++
      markenBesitzer[x][y] = uint(s[i])
      i++
      mauer[x][y] = false; if s[i] == 1 { mauer[x][y] = true }
      i++
    }
  }
}

var
  geladen = false

func speichern() {
  if geladen {
    geladen = false
  } else {
    ker.Panic ("Welt nicht geladen")
  }
  datei.Name (name)
  s := encode()
  for i := uint(0); i < uint(len(s)); i++ {
    datei.Seek (i)
    datei.Put (s[i])
  }
}

func constructGraph() {
  var n uint
  k, k1, k2, k3 := bn.New (5), bn.New (5), bn.New (5), bn.New (5)
  for y := uint8(0); y < M; y++ {
    for x := uint8(0); x < M; x++ {
      if ! mauer[x][y] {
        n = uint(x) * M + uint(y)
        if klötze[x][y] > 0 { n += M * M * M }
        k.SetVal (n)
        graph.Ins (k)
      }
    }
  }
  for y := uint8(0); y < M; y++ {
    for x := uint8(0); x < M; x++ {
      n = uint(x) * M + uint(y)
      k.SetVal (n)
      k1.SetVal (uint(x + 1) * M + uint(y))
      k2.SetVal (n + M * M * M)
      k3.SetVal (uint(x + 1) * M + uint(y) + M * M * M)
      if x + 1 < M {
        if graph.Ex2 (k, k1) || graph.Ex2 (k2, k1) || graph.Ex2 (k, k3) || graph.Ex2 (k2, k3) {
          graph.Edge (nil)
        }
      }
      if y + 1 < M {
        k1.SetVal (uint(x) * M + uint(y + 1))
        k3.SetVal (uint(x) * M + uint(y + 1) + M * M * M)
        if graph.Ex2 (k, k1) || graph.Ex2 (k2, k1) || graph.Ex2 (k, k3) || graph.Ex2 (k2, k3) {
          graph.Edge (nil)
        }
      }
    }
  }
}

func laden (s ...string) {
  if len(s) == 0 {
    name = env.Arg (1)
  } else {
    name = s[0]
  }
  if name == "" { name = "Welt" }
  name += ".rob"
  datei.Name (name)
  if ! datei.Empty() {
    s := make (Stream, datei.Num())
    for i := uint(0); i < datei.Num(); i++ {
      datei.Seek (i)
      s[i] = datei.Get().(byte)
    }
    decode (s)
  }
  geladen = true
  constructGraph()
  weltAusgeben()
}

func (r *roboter) traversieren (op Op) {
  r = alle[1] // 5
  if r == nil { ker.Panic ("r == nil") }
  for _, a := range(r.aktionen) {
    op (a)
  }
  r.aktionen = make([]aktion, 0)
}

func prog (s string) {
  for i := 0; i < len(s); i++ {
    programmdatei.Ins (byte(s[i]))
  }
  programmdatei.Ins (byte('\n'))
}

func programmErzeugen() {
  heute := day.New()
  heute.SetFormat (day.Dd)
  heute.Update()
  zeit := clk.New()
  zeit.SetFormat (clk.Hh_mm)
  zeit.Update()
  s := zeit.String()
  s = s[0:2] + s[3:5]
  programmdatei.Name (env.User() + "-" + heute.String() + "-" + s + ".go")
  programmdatei.Clr()
  scr.Cls()
  scr.Colours (col.Yellow(), col.DarkBlue())
  prog ("package main\n\nimport . \"robi\"\n")
  prog ("func main() {")
  alle[0].traversieren (func (a any) { prog ("  " + aktionstext[a.(aktion)] + "()") })
  prog ("  Fertig()\n}")
  programmdatei.Fin()
}

func protokollSchalten (ein bool) {
  sollProtokolliertWerden = ein
  scr.Colours (col.ErrorF(), col.ErrorB())
  if sollProtokolliertWerden {
    scr.Write ("Protokoll eingeschaltet", 0, scr.NColumns() - M - 1)
  } else {
    scr.Write ("                       ", 0, scr.NColumns() - M - 1)
  }
}

func sokobanSchalten (ein bool) {
  sokoban = ein
}

func ausgabe (x uint) {
  zahl.SetVal (x)
  zahl.Write (scr.NLines() - 2, 9)
  Wait (true)
}

func eingabe() uint {
  zahl.Edit (scr.NLines() - 2, 9)
  return zahl.Val()
}

func fehlerMelden (s string, n uint) {
  errh.Error (s, n)
}

func hinweisAusgeben (s string, n uint) {
  errh.Hint1 (s, n)
}

func fertig() {
  errh.Error0 ("Programm beendet")
  scr.Fin()
}

func anzahl() uint {
  a := uint(0)
  for y := uint8(0); y < M; y++ {
    for x := uint8(0); x < M; x++ {
      if nummer[x][y] > 0 {
        a++
      }
    }
  }
  return a
}
