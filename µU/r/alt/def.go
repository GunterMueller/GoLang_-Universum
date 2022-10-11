package r

// (c) Christian Maurer   v. 220827 - license see µU.go

// >>> For all functions without specification ;-) the source is the spec.

import (
  "math"
  "µU/obj"
  "µU/col"
)
const (
  E = math.E
  Pi = math.Pi
  Epsilon = math.SmallestNonzeroFloat64
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
type
  Func64 = func (float64) float64

// Returns, what the name says.
func Inf() float64 { return math.Inf(1) }
func MinusInf() float64 { return math.Inf(-1) }
func IsZero (x float64) bool { return isZero(x) }
func IsInf (x float64) bool { return math.IsInf(x,1) }
func IsMinusInf (x float64) bool { return math.IsInf(x,-1) }
func IsAbsInf (x float64) bool { return math.IsInf(x,1)||math.IsInf(x,-1) }

// Returns true, iff x is integer (and not NaN or Inf).
func Integer (x float64) bool { return integer(x) }

// Returns true, iff x is neither infinite nor a NaN.
func Finite (x float64) bool { return finite(x) }

// Returns true, iff IsZero (x - y).
func Eq (x, y float64) bool { return eq(x,y) }

// Returns (f, true), iff s represents the real number f in the range of float64;
// panics otherwise.
func Real (s string) (float64, uint, bool) { return float(s) }

func String (x float64) string { return string_(x) }

func Colours (f, b col.Colour) { colours(f,b) }

// The number of digits after "." in Write and Edit is n.
// Initially it is 2.
func SetFormat (n uint) { setFormat(n) }

// n is sufficiently greater than the number of digits given by SetFormat.
// n is the width for writing end editing.
func Wd (n uint) { wd(n) }

// Pre: (l, c) < (scr.NLines, scr.NColumns).
// x is written to the screen starting at position (l, c).
func Write (x float64, l, c uint) { write(x,l,c) }

// Pre: (l, c) < (scr.NLines, scr.NColumns).
// x is the number, that was edited at screen position (l, c).
func Edit (x *float64, l, c uint) { edit(x,l,c) }

func FuncDefined (s string) (Function, bool) { return funcDefined(s) }

func RealStarted (s string) (float64, uint, bool) { return realStarted(s) }
func ConstStarted (s string) (float64, uint, bool) { return constStarted(s) }
func IsOp (b byte) (Operation, bool) { return isOp(b) }
func FuncStarted (s string) (Function, uint, bool) { return funcStarted(s) }

func OpString (op Operation) string { return opString(op) }
func OpText (op Operation) string { return opText(op) }
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

func PowOp (b byte) (Operation, bool) { return powOp(b) }
func AddSubOp (b byte) (Operation, bool) { return addSubOp(b) }
func MulDivOp (b byte) (Operation, bool) { return mulDivOp(b) }

func Strings (s string) []string { return strings(s) }
func Power (x float64, k uint) float64 { return power(x,k) }
