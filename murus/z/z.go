package z

// (c) murus.org  v. 150121 - license see murus.go

const (
  _b = byte('b'); _B = byte('B')
  _e = byte('e'); _E = byte('E')
  _i = byte('i'); _I = byte('I')
  _o = byte('o'); _O = byte('O')
  _u = byte('u'); _U = byte('U')
  delta = _a - _A
)

func isLatin1 (b byte) bool {
  switch b {
  case
    AE, OE, UE, Ae, Oe, Ue, Sz, Euro, Para, Degree, ToThe2, ToThe3, Mue, Copyright: // ,
    // Registered, Pound, Female, Male, PlusMinus, Times, Division, Negate:
    return true
  }
  return false
}

func string_(b byte) string {
  s:= make ([]byte, 1)
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
  case AE, OE, UE:
    return true
  }
  return false
}

func opensHell (b byte) bool {
  return b == byte(194) ||
         b == byte(195)
}

func devilsDung (s *string) bool {
  n:= len (*s)
  if n == 0 {
    return false
  }
  for i:= 0; i < n; i++ {
    switch (*s)[i] {
    case 194, 195:
      return true
    }
  }
  return false
}

func toHellWithUTF8 (s *string) {
  n:= len (*s)
  if n == 0 { return }
  bs:= []byte(*s)
  i, k:= 0, 0
  var b byte
  for i < n {
    b = bs[i]
    switch b { case 194:
      i++
      b = bs[i]
    case 195:
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
  switch { case a < _A:
    return a == b
  case a <= _Z, _a <= a && a <= _z, a == AE, a == OE, a == UE, a == Ae, a == Oe, a == Ue:
    // see below
  default:
    return a == b
  }
  return a & 31 == b & 31
}

func cap (b byte) byte {
  switch b {
  case Ae:
    return AE
  case Oe:
    return OE
  case Ue:
    return UE
  }
  if _a <= b && b <= _z {
    return b - delta
  }
  return b
}

func lower (b byte) byte {
  switch b {
  case AE:
    return Ae
  case OE:
    return Oe
  case UE:
    return Ue
  }
  if _A <= b && b <= _Z {
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
  case _A, _E, _I, _O, _U, _a, _e, _i, _o, _u, AE, OE, UE, Ae, Oe, Ue:
    return true
  }
  return false
}

func isConsonant (b byte) bool {
  if isVowel (b) {
    return false
  }
  if _B <= b && b <= _Z || _b <= b && b <= _z || b == Sz {
    return true
  }
  return false
}

func postscript (b byte) string {
  switch b {
  case AE:
    return "Adieresis"
  case OE:
    return "Odieresis"
  case UE:
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
    return "Euro"
  case Para:
    return "section"
  case Degree:
    return "degree"
/*
  case ToThe2:
    return ""
  case ToThe3:
    return ""
*/
  case Mue:
    return "mu"
  case Copyright:
    return "copyright"
/*
  case Registered:
    return "registered"
  case Pound:
    return "sterling"
  case Female:
    return ""
  case Male:
    return ""
  case PlusMinus:
    return "plusminus"
  case Times:
    return "multiply"
  case Division:
    return ""
  case Negate:
    return ""
*/
  }
  return ""
}

func init() {
  ord:= []byte(" 0123456789Aa  BbCcDdEeFfGgHhIiJjKkLlMmNnOo  PpQqRrSs TtUu  VvWwXxYyZz")
//                           Ää                            Öö        ß    Üü
//              0         1         2         3         4         5         6
//              0123456789012345678901234567890123456789012345678901234567890123456789
  ord[13] = AE
  ord[14] = Ae
  ord[43] = OE
  ord[44] = Oe
  ord[53] = Sz
  ord[58] = UE
  ord[59] = Ue
  for b:= byte(0); b < byte(len (ord)); b++ {
    nr[b] = b
    in[b] = true
  }
}
