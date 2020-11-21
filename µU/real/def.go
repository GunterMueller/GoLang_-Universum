package real

// (c) Christian Maurer   v. 201112 - license see µU.go

// >>> For all functions without specification ;-) the source is the spec.

import (
  "math"
  "µU/obj"
  "µU/col"
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
  Operation = int; const (
  NoOp = iota
  Add
  Sub
  Mul
  Div
  Pow
  NOperations
)
type
  Function = int; const (
  Undef = iota                   // undefined function
  Id                             // identity function
  Chs                            // change sign
  Rec                            // reciproc value
  Sqr; Sqrt                      // square, square root
  Exp; Exp10; Exp2               // exponentials
  Log; Lg; Ld                    // logarithms
  Sin; Cos; Tan; Cot             // arithmetic functions
  Arcsin; Arccos; Arctan; Arccot // and their inverses
  Sinh; Cosh; Tanh; Coth         // hyperbolic functions
  Arsinh; Arcosh; Artanh; Arcoth // and their inverses
  F; F1; F2                      // for symbolic derivation
  G; G1; G2                      // F1 = F', F2 = F"
  H; H1; H2
  NFunctions
)

// Returns, what the name says.
func NaN() float64 { return math.NaN() }
func Inf() float64 { return math.Inf(1) }
func MinusInf() float64 { return math.Inf(-1) }
func IsZero (x float64) bool { return x - 0 < Epsilon }
func IsNaN (x float64) bool { return math.IsNaN(x) }
func IsInf (x float64) bool { return math.IsInf(x,1) }
func IsMinusInf (x float64) bool { return math.IsInf(x,-1) }
func IsAbsInf (x float64) bool { return math.IsInf(x,1)||math.IsInf(x,-1) }

// Returns true, iff x is neither infinite nor a NaN.
func Finite (x float64) bool { return finite(x) }

// Returns true, iff | x - y | < Epsilon.
func Eq (x, y float64) bool { return eq(x,y) }

// If s defines a number of type float64, that number is returned; returns otherwise NaN.
func Val (s string) float64 { return val(s) }

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

func FuncDefined (s string) (Function, bool) { return funcDefined(s) }

func RealStarted (s string) (float64, uint, bool) { return realStarted(s) }
func IsOp (b byte) (Operation, bool) { return isOp(b) }
func FuncStarted (s string) (Function, uint, bool) { return funcStarted(s) }

func OpString (op Operation) string { return opString(op) }
func FuncString (f Function) string { return funcString(f) }

func OpVal (op Operation, x, y float64) float64 { return opVal(op,x,y) }
func FuncVal (f Function, x float64) float64 { return funcVal(f,x) }

func OpCodelen() uint { return opCodelen() }
func OpEncode (op Operation) obj.Stream { return opEncode(op) }
func OpDecode (s obj.Stream) Operation { return opDecode(s) }

func FuncCodelen() uint { return funcCodelen() }
func FuncEncode (f Function) obj.Stream { return funcEncode(f) }
func FuncDecode (s obj.Stream) Function { return funcDecode(s) }

func Inverse (f Function) Function { return inverse(f) }

func Derivation (f Function) string { return derivation(f) }

func PowOp (op *Operation, b byte) bool { return powOp(op,b) }
func AddSubOp (op *Operation, b byte) bool { return addSubOp(op,b) }
func MulDivOp (op *Operation, b byte) bool { return mulDivOp(op,b) }

func Strings (s string) []string { return strings(s) }
