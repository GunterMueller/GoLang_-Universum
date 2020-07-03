package z

// (c) Christian Maurer   v. 200403 - license see µU.go

const
  delta = 'a' - 'A'

func init() {
  ord := []byte(" 0123456789Aa  BbCcDdEeFfGgHhIiJjKkLlMmNnOo  PpQqRrSs TtUu  VvWwXxYyZz")
//                            Ää                            Öö        ß    Üü
//               0         1         2         3         4         5         6
//               0123456789012345678901234567890123456789012345678901234567890123456789
  ord[13] = Ä
  ord[14] = Ae
  ord[43] = Ö
  ord[44] = Oe
  ord[53] = Sz
  ord[58] = Ü
  ord[59] = Ue
  for b := byte(0); b < byte(len (ord)); b++ {
    nr[b] = b
    in[b] = true
  }
}

func isLatin1 (b byte) bool {
  switch b {
  case Ä, Ö, Ü, Ae, Oe, Ue, Sz, Cent, Pound, Euro, Yen, BrokenBar, Paragraph, Copyright, Female,
       LeftDoubleAngle, Not, Registered, Degree, PlusMinus, ToThe2, ToThe3, Mu, Pilcrow, Dot,
       ToThe1, Male, RightDoubleAngle, Quarter, Half, ThreeQuarters, Times, EmptySet, Division:
    return true
  }
  return false
}

func str (b byte) string {
  s := make ([]byte, 1)
  s[0] = b
  return string(s)
}

func isLowerUmlaut (b byte) bool {
  switch b {
  case Ae, Oe, Ue, Sz:
    return true
  }
  return false
}

func isCapUmlaut (b byte) bool {
  switch b {
  case Ä, Ö, Ü:
    return true
  }
  return false
}

func opensHell (b byte) bool {
  return b == byte(0xc2) ||
         b == byte(0xc3)
}

func devilsDung (s *string) bool {
  n := len (*s)
  if n == 0 {
    return false
  }
  for i := 0; i < n; i++ {
    switch (*s)[i] {
    case 0xc2, 0xc3:
      return true
    }
  }
  return false
}

func toHellWithUTF8 (s *string) {
  n := len (*s)
  if n == 0 { return }
  bs := []byte(*s)
  i, k := 0, 0
  var b byte
  for i < n {
    b = bs[i]
    switch b {
    case 0xc2:
      i++
      b = bs[i]
    case 0xc3:
      i++
      b = bs[i] + 64
    }
    bs[k] = b
    i++
    k++
  }
  if k == n {
    return
  } else if k < n {
    *s = string(bs[:k])
  }
}

func Equiv (a, b byte) bool {
  switch {
  case a < 'A':
    return a == b
  case a <= 'Z', 'a' <= a && a <= 'z', a == Ä, a == Ö, a == Ü, a == Ae, a == Oe, a == Ue:
    // see below
  default:
    return a == b
  }
  return a & 31 == b & 31
}

func cap (b byte) byte {
  switch b {
  case Ae, Oe, Ue:
    return b - 32
  }
  if 'a' <= b && b <= 'z' {
    return b - delta
  }
  return b
}

func isCapLetter (b byte) bool {
  return 'A' <= b && b <= 'Z' || isCapUmlaut(b)
}

func isLowerLetter (b byte) bool {
  return 'a' <= b && b <= 'z' || isLowerUmlaut(b)
}

func isLetter (b byte) bool {
  return isCapLetter(b) || IsLowerLetter(b)
}

func isDigit (b byte) bool {
  return '0' <= b && b <= '9'
}

func lower (b byte) byte {
  switch b {
  case Ä, Ö, Ü:
    return b + 32
  }
  if 'A' <= b && b <= 'Z' {
    return b + delta
  }
  return b
}

var (
  nr [256]byte
  in [256]bool // in = make (map[byte] bool)
)

func Less (a, b byte) bool {
  if a == b {
    return false
  }
  if in[a] {
    if in[b] {
      return nr[a] < nr[b]
    } else {
      return true // Sonderzeichen hinter Buchstaben
    }
  } else {
    if in[b] {
      return false // s. o.
    }
  }
  return a < b // nach ASCII
}

func isVowel (b byte) bool {
  switch b {
  case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u', Ä, Ö, Ü, Ae, Oe, Ue:
    return true
  }
  return false
}

func isConsonant (b byte) bool {
  if isVowel (b) {
    return false
  }
  if 'B' <= b && b <= 'Z' || 'b' <= b && b <= 'z' || b == Sz {
    return true
  }
  return false
}

func postscript (b byte) string {
  switch b {
  case Ä:
    return "Adieresis"
  case Ö:
    return "Odieresis"
  case Ü:
    return "Udieresis"
  case Ae:
    return "adieresis"
  case Oe:
    return "odieresis"
  case Ue:
    return "udieresis"
  case Sz:
    return "germandbls"
  case Cent:
    return "cent"
  case Pound:
    return "sterling"
  case Euro:
    return "euro"
  case Yen:
    return "yen"
  case BrokenBar:
    return "brokenbar"
  case Paragraph:
    return "section"
  case Copyright:
    return "copyright"
  case Female:
    return "ordfeminine"
  case LeftDoubleAngle:
    return "quotedblleft"
  case Not:
    return "logicalnot"
  case Registered:
    return "registered"
  case Degree:
    return "deg"
  case PlusMinus:
    return "plusminus"
  case ToThe2:
    return "twosuperior"
  case ToThe3:
    return "threesuperior"
  case Mu:
    return "mu"
  case Pilcrow:
    return "paragraph"
  case Dot:
    return "periodcentered"
  case ToThe1:
    return "onesuperior"
  case Male:
    return "ordmasculine"
  case RightDoubleAngle:
    return "quotedblright"
  case Quarter:
    return "onequarter"
  case Half:
    return "onehalf"
  case ThreeQuarters:
    return "threequarters"
  case Times:
    return "multiply"
  case EmptySet:
    return "emptyset"
  case Division:
    return "divisionslash"
  }
  return ""
}

func Latin1Byte (r rune) byte {
  switch r {
  case 'A': // C3 84
    return Ä
  case 'Ö': // C3 96
    return Ö
  case 'Ü': // C3 9C
    return Ü
  case 'ä': // C3 A4
    return Ae
  case 'ö': // C3 B6
    return Oe
  case 'ü': // C3 BC
    return Ue
  case 'ß': // C3 9F
    return Sz
  case '¢': // C2 A2
    return Cent
  case '£': // C2 A3
    return Pound
  case '€': // E2 82 AC
    return Euro
  case '¥': // C2 A5
    return Yen
  case '¦': // C2 A6
    return BrokenBar
  case '§': // C2 A7
    return Paragraph
  case '©': // C2 A9
    return Copyright
  case 'ª': // C2 AA
    return Female
  case '«': // C2 AB
    return LeftDoubleAngle
  case '¬': // C2 AC
    return Not
  case '®': // C2 AE
    return Registered
  case '°': // C2 B0
    return Degree
  case '±': // C2 B1
    return PlusMinus
  case '²': // C2 B3
    return ToThe2
  case '³': // C2 B2
    return ToThe3
  case 'µ': // C2 B5
    return Mu
  case '¶': // C2 B6
    return Pilcrow
  case '·': // C2 B7
    return Dot
  case '¹': // C2 B9
    return ToThe1
  case 'º': // C2 BA
    return Male
  case '»': // C2 BB
    return RightDoubleAngle
  case '¼': // C2 BC
    return Quarter
  case '½': // C2 BD
    return Half
  case '¾': // C2 BE
    return ThreeQuarters
  case '×': // C3 97
    return Times
  case 'Ø': // C2 98
    return EmptySet
  case '÷': // C3 B7
    return Division
  }
  return 0
}
