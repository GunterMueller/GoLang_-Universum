package char

// (c) Christian Maurer   v. 240311 - license see µU.go

const
  delta = 'a' - 'A'

func init() {
  ord := []byte (" 0123456789Aa  BbCcDdEeFfGgHhIiJjKkLlMmNnOo  PpQqRrSs TtUu  VvWwXxYyZz")
//                             Ää                            Öö        ß    Üü
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
  case Ä, Ö, Ü, Ae, Oe, Ue, Sz, Cent, Pound, Euro, Paragraph, Copyright,
       Not, Registered, Degree, PlusMinus, ToThe2, ToThe3,
       Mu, Pilcrow, Dot, Times, EmptySet, Division:
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
  case Ae, Oe, Ue:
    return true
  }
  return false
}

func isUpperUmlaut (b byte) bool {
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

func devilsDung (s string) bool {
  n := len (s)
  if n == 0 {
    return false
  }
  for i := 0; i < n; i++ {
    switch s[i] {
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
/*/
    case 0xce:
      i++
      b = bs[i] + 11 * 64
/*/
    }
    bs[k] = b
    i++
    k++
  }
  if k == n {
    return
  }
  if k < n {
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

func upper (b byte) byte {
  switch b {
  case Ae, Oe, Ue:
    return b - 32
  }
  if 'a' <= b && b <= 'z' {
    return b - delta
  }
  return b
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

func isUppercaseLetter (b byte) bool {
  return 'A' <= b && b <= 'Z' || isUpperUmlaut(b)
}

func isLowercaseLetter (b byte) bool {
  return 'a' <= b && b <= 'z' || isLowerUmlaut(b) || b == Sz
}

func isLetter (b byte) bool {
  return isUppercaseLetter(b) || IsLowercaseLetter(b)
}

func isDigit (b byte) bool {
  return '0' <= b && b <= '9'
}

func isLetterOrDigit (b byte) bool {
  return isLetter (b) || isDigit (b)
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

func tex (b byte) string {
  switch b {
  case Ä:
    return "\\\"A"
  case Ö:
    return "\\\"O"
  case Ü:
    return "\\\"U"
  case Ae:
    return "\\\"a"
  case Oe:
    return "\\\"o"
  case Ue:
    return "\\\"u"
  case Sz:
    return "\\ss"
  case Cent:
    return ""
  case Pound:
    return "\\it\\S "
  case Euro:
    return ""
  case Paragraph:
    return "\\S "
  case Copyright:
    return "\\copyright "
  case Not:
    return "\\lnot "
  case Registered:
    return "\\textregistered "
  case Degree:
    return "^\\circ "
  case PlusMinus:
    return "\\pm "
  case ToThe2:
    return "^2"
  case ToThe3:
    return "^3"
  case Mu:
    return "\\mu "
  case Pilcrow:
    return "\\P "
  case Dot:
    return "\\cdot "
  case Times:
    return "\\times "
  case EmptySet:
    return "\\emptyset "
  case Division:
    return "\\div "
  }
  return string(b)
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
  case Paragraph:
    return "section"
  case Copyright:
    return "copyright"
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
  case '§': // C2 A7
    return Paragraph
  case '©': // C2 A9
    return Copyright
  case 'ª': // C2 AA
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
  case '×': // C3 97
    return Times
  case 'Ø': // C2 98
    return EmptySet
  case '÷': // C3 B7
    return Division
//  case ' ': // CE BC
//    return Pi
  }
  return 0
}
