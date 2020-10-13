package z

// (c) Christian Maurer   v. 200403 - license see µU.go

const (
  Ä                = byte(196) // 'Ä'
  ÄÄ = 'Ä'
  Ö                = byte(214) // 'Ö'
  Ü                = byte(220) // 'Ü'
  Ae               = byte(228) // 'ä'
  Oe               = byte(246) // 'ö'
  Ue               = byte(252) // 'ü'
  Sz               = byte(223) // 'ß'
//NoBreakSpace     = byte(160)
//InvExclamation   = byte(161) // '¡'
  Cent             = byte(162) // '¢'
  Pound            = byte(163) // '£'
  Euro             = byte(164) // '€'
  Yen              = byte(165) // '¥'
  BrokenBar        = byte(166) // '¦'
  Paragraph        = byte(167) // '§'
//Diaeresis        = byte(168) // '¨'
  Copyright        = byte(169) // '©'
  Female           = byte(170) // 'ª'
  LeftDoubleAngle  = byte(171) // '«'
  Not              = byte(172) // '¬'
//SoftHyphen       = byte(173)
  Registered       = byte(174) // '®'
//Macron           = byte(175) // '¯'
  Degree           = byte(176) // '°'
  PlusMinus        = byte(177) // '±'
  ToThe2           = byte(178) // '²'
  ToThe3           = byte(179) // '³'
//Acute            = byte(180) // '´'
  Mu               = byte(181) // 'µ'
  Pilcrow          = byte(182) // '¶'
  Dot              = byte(183) // '·`
//Cedilla          = byte(184) // '·'
  ToThe1           = byte(185) // '¹'
  Male             = byte(186) // 'º'
  RightDoubleAngle = byte(187) // '»'
  Quarter          = byte(188) // '¼'
  Half             = byte(189) // '½'
  ThreeQuarters    = byte(190) // '¾'
//InvQuestionMark  = byte(191) // '¿'
  Times            = byte(215) // '×'
  EmptySet         = byte(216) // 'Ø'
  Division         = byte(247) // '÷'
)

// Returns true, if b is one of the constants that are defined
// internally. Eventually they are shown here in the spec.
func IsLatin1 (b byte) bool { return isLatin1(b) }

// Returns the correspondings string of len 1.
func String (b byte) string { return str(b) }

// Returns true, if b is a small german Umlaut or 'ß'.
func IsLowerUmlaut (b byte) bool { return isLowerUmlaut(b) }

// Returns true, if b is a capital german umlaut.
func IsCapUmlaut (b byte) bool { return isCapUmlaut(b) }

// Returns true, if b is 194 or 195.
func OpensHell (b byte) bool { return opensHell(b) }

// Returns true, iff s contains one of the bytes that open hell.
func DevilsDung (s string) bool { return devilsDung(s) }

// All UTF8-runes in s starting with one of the bytes, that open hell,
// are converted to the corresponding latin1-bytes.
func ToHellWithUTF8 (s *string) { toHellWithUTF8(s) }

// Returns b transformed into the corresponding capital.
// Beware: Cap('ß') = 'ß' !
func Cap (b byte) byte { return cap(b) }

// Returns b transformed into the corresponding small letter.
func Lower (b byte) byte { return lower(b) }

// Returns true, iff b equals its corresponding capital letter.
func IsCap (b byte) bool { return b == cap(b) }

// Returns true, iff b is a capital letter.
func IsCapLetter (b byte) bool { return isCapLetter(b) }

// Returns true, iff b is a small letter.
func IsLowerLetter (b byte) bool { return isLowerLetter(b) }

// Returns true, iff b is a letter.
func IsLetter (b byte) bool { return isCapLetter(b) || isLowerLetter(b) }

// Returns true, iff b is a vowel or a german Umlaut.
func IsVowel (b byte) bool { return isVowel (b) }

// Returns true, iff b is a consonant.
func IsConsonant (b byte) bool { return isConsonant (b) }

// Returns true, iff b is a digit.
func IsDigit (b byte) bool { return isDigit(b) }

// Returns the postscript name of b.
func Postscript (b byte) string { return postscript(b) }
