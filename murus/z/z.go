package z

// (c) murus.org  v. 170204 - license see murus.go

const
  delta = 'a' - 'A'

func init() {
  ord := []byte(" 0123456789Aa  BbCcDdEeFfGgHhIiJjKkLlMmNnOo  PpQqRrSs TtUu  VvWwXxYyZz")
//                           ÄAe                            ÖOe        ß    ÜUe
//              0         1         2         3         4         5         6
//              0123456789012345678901234567890123456789012345678901234567890123456789
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
  case
    Ä, Ö, Ü, Ae, Oe, Ue, Sz, Euro, Cent, Pound, Paragraph, Degree, Copyright, Registered,
    Mue, PlusMinus, Times, Division, Dot, Negate, ToThe1, ToThe2, ToThe3, Female, Male:
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
  case Euro:
    return "euro"
/*
  case toThe2:
    return ""
  case toThe3:
    return ""
*/
  case Mue:
    return "mu"
  case Copyright:
    return "copyright"
/*
  case registered:
    return "registered"
  case pound:
    return "sterling"
  case female:
    return ""
  case male:
    return ""
  case plusMinus:
    return "plusminus"
  case times:
    return "multiply"
  case division:
    return ""
  case negate:
    return ""
*/
  }
  return ""
}
