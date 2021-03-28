package errh

// (c) Christian Maurer   v. 210315 - license see µU.go

import
  "µU/col"
var (
  ToWait, ToContinue, ToContinueOrNot, ToCancel, ToScroll,
  ToSelect, ToChange, ToSwitch, ToSelectWithPrint, ToPrint string
)

// s is written to the first line of the screen
func Head (s string) { head(s) }
// The head is deleted, the former content of the screen is restored.
func DelHead() { delHead() }

// s is written to the last line of the screen
// resp. to the screen starting at position (line, column) == (l, c).
func Hint (s string) { hint(s) }
func HintPos (s string, l, c uint) { hintPos(s,l,c) }
func Hint1 (s string, n uint) { hint1(s,n) }
func Hint2 (s string, n uint, s1 string, n1 uint) { hint2(s,n,s1,n1) }

// The hints are deleted, the former content of the screen is restored.
func DelHint() { delHint() }
func DelHintPos (s string, l, c uint) { delHintPos(s,l,c) }

// s (and n resp.) is written to the last line of the screen.
// The calling process is blocked, until Enter or left mouse button is pressed;
// then the former content of the last line of the screen is restored.
// func Proceed0 (s string) { proceed0(s) }
// func Proceed (s string, n uint) { proceed(s,n) }

// s (and n resp.) is written to the last line of the screen.
// The calling process is blocked, until Escape or Backspace is pressed;
// then the former content of the last line of the screen is restored.
func Error0 (s string) { error0(s) }
func Error (s string, n uint) { error(s,n) }
func Error2 (s string, n uint, s1 string, n1 uint) { error2(s,n,s1,n1) }
func Error3 (s string, n uint, s1 string, n1 uint, s2 string, n2 uint) {
  error3(s,n,s1,n1,s2,n2)
}
func Error4 (s string, n uint, s1 string, n1 uint, s2 string, n2 uint, s3 string, n3 uint) {
  error4(s,n,s1,n1,s2,n2,s3,n3)
}
func ErrorF (s string, f float64) { errorF(s,f) }
func ErrorZ (s string, z int) { errorZ(s,z) }

// s is written to the screen, starting at position (line, column) == (l, c).
// The calling process is blocked, until Escape or Backspace is pressed;
// then the former content of the screen, starting at (l, c), is restored.
func Error0Pos (s string, l, c uint) { error0Pos(s,l,c) }
func ErrorPos (s string, n, l, c uint) { errorPos(s,n,l,c) }
func Error2Pos (s string, n uint, s1 string, n1 uint, l, c uint) { error2Pos(s,n,s1,n1,l,c) }

// The calling process is blocked, until the user has confirmed by some action.
func Confirmed () bool { return confirmed() }

// TODO Spec
func WriteLicense (p, v, a string, f, l, b col.Colour, g []string, t *string) {
  writeLicense(p,v,a,f,l,b,g,t)
}
func MuLicense (p, v, a string, f, l, b col.Colour) { µULicense(p,v,a,f,l,b) }
func Headline (p, v, a string, f, b col.Colour) { headline(p,v,a,f,b) }

// h is written to the center of the screen.
// The calling process is blocked, until Enter, Esc, Back or a mouse button is pressed;
// then the former content of the screen is restored.
func Help (h []string) { help(h) }

// Like Help, but only a short hint for F1 is given.
func Help1() { help1() }
