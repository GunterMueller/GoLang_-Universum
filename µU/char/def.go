package char

// (c) Christian Maurer   v. 240311 - license see µU.go

const (
  Ä                = byte(0xc4) // 'Ä'  196
  Ö                = byte(0xd6) // 'Ö'  214
  Ü                = byte(0xdc) // 'Ü'  220
  Ae               = byte(0xe4) // 'ä'  228
  Oe               = byte(0xf6) // 'ö'  246
  Ue               = byte(0xfc) // 'ü'  252
  Sz               = byte(0xdf) // 'ß'  223
  Cent             = byte(0xa2) // '¢'  162
  Pound            = byte(0xa3) // '£'  163
  Euro             = byte(0xa4) // '€'  164
  Paragraph        = byte(0xa7) // '§'  167
  Copyright        = byte(0xa9) // '©'  169
  Not              = byte(0xac) // '¬'  172
  Registered       = byte(0xae) // '®'  174
  Degree           = byte(0xb0) // '°'  176
  PlusMinus        = byte(0xb1) // '±'  177
  ToThe2           = byte(0xb2) // '²'  178
  ToThe3           = byte(0xb3) // '³'  179
  Mu               = byte(0xb5) // 'µ'  181
  Pilcrow          = byte(0xb6) // '¶'  182
  Dot              = byte(0xb7) // '·'  183
  Times            = byte(0xd7) // '×'  215
  EmptySet         = byte(0xd8) // 'Ø'  216
  Division         = byte(0xf7) // '÷'  247
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

// Returns true, iff b is a letter or a digit.
func IsLetterOrDigit (b byte) bool { return isLetterOrDigit(b) }

// Returns the postscript name of b.
func TeX (b byte) string { return tex(b) }

// Returns the postscript name of b.
func Postscript (b byte) string { return postscript(b) }
