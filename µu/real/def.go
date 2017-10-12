package real

// (c) Christian Maurer   v. 140214 - license see µu.go

// >>> alpha-version !

import (
  "math"
  "µu/col"
)

/* float64 = real numbers representable in the 64-bit floating point
             format due to IEEE 754 (including Inf, -Inf and NaN)
             with external representation in the "normal format"
               ["-"] c1 [ "." c ]
             or in the "scientific format"
               ["-"] c1 "." { c } "e" ["-"] c,
               where c1 = "1" | "2" | ... | "9", z = "0" | z1,
                     s1 = c1 { c }, s = c { c } . */
//const ( // Format 
//  sc = iota; ... // TODO
//)
const (
  Epsilon = 10e-15 // Eps = math.SmallestNonzeroFloat64
  Pi = math.Pi
  E = math.E
  Max = math.MaxFloat64
)
type
  Operation byte; const (
  NoOp = iota
  Add
  Sub
  Mul
  Div
  Pow
  NOperations )
type
  Function byte; const (
  Undef = iota                   // undefined function
  Floor; Fract                   // integer / fractional part
  Chs                            // change sign
  Rec                            // recipro value
  Sqr; Sqrt                      // square, square root
  Exp; Exp10; Exp2               // exponentials
  Log; Lg; Ld                    // logarithms
  Sin; Cos; Tan; Cot             // arithmetic functions
  Arcsin; Arccos; Arctan; Arccot // and their inverses
  Sinh; Cosh; Tanh; Coth         // hyperbolic functions
  Arsinh; Arcosh; Artanh; Arcoth // and their inverses
  Gamma                          // Gamma-function
  F; F1; F2; G; G1; G2           // for symbolic derivation
  H; H1; H2                      // F1 = F', F2 = F" etc.
  NFunctions
)

// Returns, what the name says.
func NaN () float64 { return math.NaN() }
func Inf () float64 { return math.Inf(1) }
func MinusInf () float64 { return math.Inf(-1) }
func IsZero (x float64) bool { return x - 0 < Epsilon }
func IsNaN (x float64) bool { return math.IsNaN(x) }
func IsInf (x float64) bool { return math.IsInf(x,1) }
func IsMinusInf (x float64) bool { return math.IsInf(x,-1) }
func IsAbsInf (x float64) bool { return math.IsInf(x,1)||math.IsInf(x,-1) }

// Returns true, iff x is neither infinite nor a NaN.
func Valid (x float64) bool { return valid(x) }

// Returns the number, iff s defines one; returns otherwise NaN.
func Number (s string) float64 { return number(s) }

func Defined (s string) (float64, bool) { return defined(s) } // TODO
func String (x float64) string { return string_(x) }

func Colours (f, b col.Colour) { colours(f,b) }

// TODO
// func SetFormat (f Format) { setFormat (f) }

// Pre: (l, c) < (scr.NLines, scr.NColumns).
// x is written to the screen starting at position (l, c).
func Write (x float64, l, c uint) { write(x,l,c) }

// Pre: (l, c) < (scr.NLines, scr.NColumns).
// x is the number, that was edited at screen position (l, c).
func Edit (x *float64, l, c uint) { edit(x,l,c) }
