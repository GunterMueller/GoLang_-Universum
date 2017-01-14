package z

// (c) murus.org  v. 160428 - license see murus.go

const (
  AE = byte(196); OE = byte(214); UE = byte(220)
  Ae = byte(228); Oe = byte(246); Ue = byte(252); Sz = byte(223)
  Euro = byte(164); Para = byte(167); Degree = byte(176)
  ToThe2 = byte(178); ToThe3 = byte(179); Mue = byte(181)
  Copyright = byte(169); Registered = byte(174); Pound = byte(163)
  Female = byte(170); Male = byte(186); PlusMinus = byte(177)
  Times = byte(151); Division = byte(183); Negate = byte(172)
  _a = byte('a'); _A = byte('A')
  _z = byte('z'); _Z = byte('Z')
)

// Returns true, if b is one of the above constants.
func IsLatin1 (b byte) bool { return isLatin1(b) }

// Returns the correspondings string of len 1.
func String (b byte) string { return string_(b) }

// Returns true, if b is one of Ae, Oe, Ue or Sz.
func IsLowerUmlaut (b byte) bool { return isLowerUmlaut(b) }

// Returns true, if b is one of AE, OE or UE.
func IsCapUmlaut (b byte) bool { return isCapUmlaut(b) }

// Returns true, if b is 194 or 195.
func OpensHell (b byte) bool { return opensHell(b) }

// Returns true, iff s contains one of the bytes that open hell.
func DevilsDung (s *string) bool { return devilsDung(s) }

// All UTF8-runes in s starting with one of the bytes, that open hell,
// are converted to the corresponding latin1-bytes (one of the above constants)
func ToHellWithUTF8 (s *string) { toHellWithUTF8(s) }

// Returns b transformed into the corresponding capital, e.g.
// Cap('a') == Cap('A') == 'A', Cap('9') == '9', Cap('.') == '.'),
// Cap(Ae) == Cap(AE) == AE etc. Beware: Cap(Sz) = Sz !
func Cap (b byte) byte { return cap(b) }

// Returns b transformed into the corresponding small letter, e.g.
// Lower('A') == Lower('a') == 'a'), Lower('9') == '9', Lower('.') == '.'),
// Loser(AE) == Lower(Ae) == Ae etc.
func Lower (b byte) byte { return lower(b) }

func IsCap (b byte) bool { return b == cap(b) }

func IsCapLetter (b byte) bool { return _A <= b && b <= _Z || isCapUmlaut(b) }

func IsLowerLetter (b byte) bool { return _a <= b && b <= _z || isLowerUmlaut(b) }

func IsLetter (b byte) bool { return IsCapLetter(b) || IsLowerLetter(b) }

func IsVowel (b byte) bool { return isVowel (b) }

func IsConsonant (b byte) bool { return isConsonant (b) }

func IsDigit (b byte) bool { return byte('0') <= b && b <= byte('9') }

func Postscript (b byte) string { return postscript(b) }
