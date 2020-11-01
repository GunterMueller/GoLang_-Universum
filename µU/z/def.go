package z

// (c) Christian Maurer   v. 201013 - license see µU.go

const (
  Ä                = byte(0xc4) // 'Ä'
  Ö                = byte(0xd6) // 'Ö'
  Ü                = byte(0xdc) // 'Ü'
  Ae               = byte(0xe4) // 'ä'
  Oe               = byte(0xf6) // 'ö'
  Ue               = byte(0xfc) // 'ü'
  Sz               = byte(0xdf) // 'ß'
//NoBreakSpace     = byte(0xa0) // ' '
//InvExclamation   = byte(0xa1) // '¡'
  Cent             = byte(0xa2) // '¢'
  Pound            = byte(0xa3) // '£'
  Euro             = byte(0xa4) // '€'
  Yen              = byte(0xa5) // '¥'
  BrokenBar        = byte(0xa6) // '¦'
  Paragraph        = byte(0xa7) // '§'
//Diaeresis        = byte(0xa8) // '¨'
  Copyright        = byte(0xa9) // '©'
  Female           = byte(0xaa) // 'ª'
  LeftDoubleAngle  = byte(0xab) // '«'
  Not              = byte(0xac) // '¬'
//SoftHyphen       = byte(0xad)
  Registered       = byte(0xae) // '®'
//Macron           = byte(0xaf) // '¯'
  Degree           = byte(0xb0) // '°'
  PlusMinus        = byte(0xb1) // '±'
  ToThe2           = byte(0xb2) // '²'
  ToThe3           = byte(0xb3) // '³'
//Acute            = byte(0xb4) // '´'
  Mu               = byte(0xb5) // 'µ'
  Pilcrow          = byte(0xb6) // '¶'
  Dot              = byte(0xb7) // '·`
//Cedilla          = byte(0xb8) // '·'
  ToThe1           = byte(0xb9) // '¹'
  Male             = byte(0xba) // 'º'
  RightDoubleAngle = byte(0xbb) // '»'
  Quarter          = byte(0xbc) // '¼'
  Half             = byte(0xbd) // '½'
  ThreeQuarters    = byte(0xbe) // '¾'
//InvQuestionMark  = byte(0xbf) // '¿'
  Times            = byte(0xd7) // '×'
  EmptySet         = byte(0xd8) // 'Ø'
  Division         = byte(0xf7) // '÷'
)

// Returns true, if b is one of the constants that are defined
// internally. Eventually they are shown here in the spec.
func IsLatin1 (b byte) bool { return isLatin1(b) }

// Returns the correspondings string of len 1.
func String (b byte) string { return str(b) }

// Returns true, if b is a small german Umlaut or 'ß'.
func IsLowerUmlaut (b byte) bool { return isLowerUmlaut(b) }

// Returns true, if b is a capital german umlaut.
func IsUpperUmlaut (b byte) bool { return isUpperUmlaut(b) }

// Returns true, if b is 194 or 195.
func OpensHell (b byte) bool { return opensHell(b) }

// Returns true, iff s contains one of the bytes that open hell.
func DevilsDung (s string) bool { return devilsDung(s) }

// All UTF8-runes in s starting with one of the bytes, that open hell,
// are converted to the corresponding latin1-bytes.
func ToHellWithUTF8 (s *string) { toHellWithUTF8(s) }

// Returns b transformed into the corresponding upper-case letter.
// Beware: Cap('ß') = 'ß' !
func Upper (b byte) byte { return upper(b) }

// Returns b transformed into the corresponding lower-case letter.
func Lower (b byte) byte { return lower(b) }

// Returns true, iff b equals its corresponding upper-case letter.
func IsUpper (b byte) bool { return b == upper(b) }

// Returns true, iff b equals its corresponding lower-case letter.
func IsLower (b byte) bool { return b == lower(b) }

// Returns true, iff b is an upper-case letter.
func IsUppercaseLetter (b byte) bool { return isUppercaseLetter(b) }

// Returns true, iff b is an lower-case letter.
func IsLowercaseLetter (b byte) bool { return isLowercaseLetter(b) }

// Returns true, iff b is a letter.
func IsLetter (b byte) bool { return isUppercaseLetter(b) || isLowercaseLetter(b) }

// Returns true, iff b is a vowel or a german Umlaut.
func IsVowel (b byte) bool { return isVowel (b) }

// Returns true, iff b is a consonant.
func IsConsonant (b byte) bool { return isConsonant (b) }

// Returns true, iff b is a digit.
func IsDigit (b byte) bool { return isDigit(b) }

// Returns the postscript name of b.
func Postscript (b byte) string { return postscript(b) }
