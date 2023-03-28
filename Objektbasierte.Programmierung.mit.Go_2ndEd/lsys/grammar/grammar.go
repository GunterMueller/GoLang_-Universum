package grammar

// (c) Christian Maurer   v. 230301 - license see µU.go

import (
  "µU/ker"
  "µU/char"
  "µU/str"
  "µU/col"
  "µU/N"
  "µU/pseq"
  "µU/br"
)
const (
  max = 26 // maximal number of rules, that start with the same symbol
  separator = "->"
  suffix = ".ls"
)
const
  ab = 123 // 'A'..'Z', 'a'..'z' = 65..90, 97..122
var (
  leftSide, rightSide [ab][max]string
  nRules [ab]uint
)

func stop (f int) {
  var s string
  switch f {
  case  0:
    s = "Startwort fehlt"
  case  1:
    s = "Winkel muss < 360 sein"
  case  2:
    s = "keine Regel angegeben"
  case  3:
    s = "Abzweigungswinkel fehlt "
  case  4:
    s = "Regel ohne Symbol am Anfang"
  case  5:
    s = "höchstens 26 Regeln dürfen mit dem gleichen Symbol anfangen"
  case  6:
    s = "Ableitungssymbol '->' fehlt"
  case  7:
    s = "linke Seite einer Regel darf höchstens 4 Zeichen lang sein"
  case  8:
    s = "linke Seite einer Regel darf nur Buchstaben enthalten"
  case  9:
    s = "rechte Seite einer Regel muss mindestens 2 Zeichen lang sein"
  case 10:
    s = "rechte Seite einer Regel enthält falsche Zeichen"
  case 11:
    s = "Anzahl der Iterationen fehlt"
  case 12:
    s = "Fehler bei der Angabe eines Winkels"
  case 13:
    s = "Fehler bei der Anzahl der Iterationen"
  case 14:
    s = "Anzahl der Iterationen muss > 0 sein"
  case 15:
    s = "unverwertbare Informationen am Ende"
  }
println (s)
//  ker.Panic (s)
}

func initialize (ls string) {
  filename := ls + ".ls"
  m := pseq.Length (filename)
  word := str.New (m)
  file := pseq.New (word)
  file.Name (filename)
  text := file.Get().(string)
  str.OffSpc (&text)
  for {
    Startword = str.SplitLine (&text)
    if Startword[0] != Comment {
      break
    }
  }
  if _, b := N.Natural (Startword); b {
    word = Startword
    Startword = str.SplitLine (&text)
  } else {
    word = "90"
  }
  if str.Empty (Startword) { stop(0) }
  var bc bool
  if StartColour, bc = isColour (Startword[0]); bc {
    str.Rem (&Startword, 0, 1)
  } else {
    StartColour = col.LightWhite()
    StartColour = col.Black()
  }
  k, t := N.Natural (word)
  if t { if k >= 360 { stop(1) } }
  Startangle = float64(k)
  if ! okR (Startword) { stop(2) }
// production rules:
  for {
    if str.Empty (text) { stop(3) }
    b := text[0]
    if char.IsDigit (b) { break }
    if ! char.IsLetter (b) { stop(4) }
    k = nRules[b]; if k == max { stop(5) }
    for {
      word = str.SplitLine (&text)
      if word[0] != Comment {
        break
      }
    }
    if i, b := str.Sub ("%", word); b {
      word = str.Part (word, 0, i)
    }
    i, t := str.Sub (separator, word); if ! t { stop(6) }
    if i >= MaxL { stop(7) }
    leftSide[b][k] = str.Part (word, 0, i)
    if ! okL (leftSide[b][k]) { stop(8) }
    i += uint(len(separator))
    rightSide[b][k] = str.Part (word, i, uint(len(word)) - i)
    if len(rightSide[b]) < 2 { stop(9) }
    if ! okR (rightSide[b][k]) { } // stop(10) }
    nRules[b]++
  }
  word = str.SplitLine (&text)
  if str.Empty (word) { stop(11) }
  n := br.New(3)
  if ! n.Defined (word) { stop(12) }
  if str.Empty (text) {
    Turnangle = 90.
    NumIterations = uint(n.RealVal())
  } else {
    Turnangle = n.RealVal()
    word = str.SplitLine (&text)
    NumIterations, t = N.Natural (word); if ! t { stop(13) }
//    if NumIterations == 0 { stop(14) }
    if ! str.Empty (text) { stop(15) }
  }
}

func isComment (s Symbol) bool {
  return s == Comment
}

func isStep (s Symbol) bool {
  switch s {
  case Step, YetiStep:
    return true
  }
  return false
}

func isTurn (s Symbol) bool {
  switch s {
  case TurnLeft, TurnRight, Invert:
    return true
  }
  return false
}

func isTilt (s Symbol) bool {
  switch s {
  case TiltDown, TiltUp:
    return true
  }
  return false
}

func isRoll (s Symbol) bool {
  switch s {
  case RollLeft, RollRight:
    return true
  }
  return false
}

func isBranch (s Symbol) bool {
  switch s {
  case BranchStart, BranchEnd:
    return true
  }
  return false
}

func isPolygon (s Symbol) bool {
  switch s {
  case PolygonStart, PolygonEnd:
    return true
  }
  return false
}

func isCol (s Symbol) bool {
  return char.IsLowercaseLetter (s) && s != 'f'
}

func isSymbol (b byte) bool {
  return isComment (b) ||
         isStep (b) ||
         isTurn (b) ||
         isTilt (b) ||
         isRoll (b) ||
         isBranch (b) ||
         isPolygon (b) ||
         isCol (b) ||
         char.IsUppercaseLetter (b)
}

func isColour (s Symbol) (col.Colour, bool) {
  switch s {
  case 'n':
    return colours[0], true // Brown
  case 'r':
    return colours[1], true // Red
  case 'l':
    return colours[2], true // LightRed
  case 'o':
    return colours[3], true // Orange
  case 'g':
    return colours[4], true // Green
  case 'd':
    return colours[5], true // DarkGreen
  case 'c':
    return colours[6], true // Cyan
  case 'e':
    return colours[7], true // LightBlue
  case 'b':
    return colours[8], true // Blue
  case 'm':
    return colours[9], true // Magenta
  case 'k':
    return colours[10], true // Black
  case 'y':
    return colours[11], true // Gray
  case 'w':
    return colours[12], true // White
  case 'z':
    return colours[13], true // LightWhite
  }
  return col.LightWhite(), false
}

func okL (s string) bool {
  for i := 0; i < len(s); i++ {
    if ! isSymbol (s[i]) {
      return false
    }
  }
  return true
}

func okR (s string) bool {
  k, p := uint(0), uint(0)
  for i := 0; i < len(s); i++ {
    if ! isSymbol (s[i]) {
// println ("kein Symbol: " + string(s[i]))
ker.Panic ("affe" + string(s[i]) + "esel")
      return false
    }
    switch s[i] {
    case BranchStart:
      k++
    case BranchEnd:
      if k > 0 {
        k--
      } else {
        return false
      }
    case PolygonStart:
      p++
    case PolygonEnd:
      if p > 0 {
        p--
      } else {
        return false
      }
    }
  }
  return k + p == 0
}

func exRule (s string) bool {
  if ! char.IsLetter (s[0]) { return false }
  b := s[0]
  if nRules[b] == 0 {
    return false
  }
  for i := uint(0); i < nRules[b]; i++ {
    p, t := str.Sub (s, leftSide[b][i])
    if t && p == 0 {
      return true
    }
  }
  return false
}

func derivation (s string) string {
  if ! str.Empty(s) {
    b := s[0]
    if char.IsLetter (b) && nRules[b] > 0 {
      for i := uint(0); i < nRules[b]; i++ {
        if s == leftSide[b][i] {
          return rightSide[b][i]
        }
      }
    }
  }
  return ""
}
