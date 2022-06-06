package col

// (c) Christian Maurer   v. 211127 - license see µU.go

import
 . "µU/obj"
const
  NColours = 1 << 24
type
  Colour interface {

  Object // empty colour is black

// TODO Spec
  EncodeInv() Stream

// String returns the name of x, defined by the name given with New3.
  Stringer

// String1 returns "rrggbb", where "rr", "gg" and "bb" are the rgb-values
// in sedecimal basis (with uppercase letters).
  String1() string

// Defined returns true, iff s is a string of 3 values in sedecimal basis
// (with uppercase letters). In that case, c is the colour with
// the corresponding rgb-values; otherwise, nothing has happened.
  Defined1 (s string) bool

// Return the values of red/green/blue intensity of x.
  R() byte; G() byte; B() byte

// x is the colour defined by the values of r, g and b.
  Set (r, g, b byte)
  SetR (b byte); SetG (b byte); SetB (b byte)

// Liefert x.R() + 256 * x.G() + x.B().
  Code() uint

// Returns true, if c is, what the name of the func says.
  IsBlack() bool
  IsWhite() bool
  IsLightWhite() bool

// Returns the rgb-values of x scaled to the range from 0 to 1.
  Float32() (float32, float32, float32)
  Float64() (float64, float64, float64)

// c is changed in a manner suggested by the name of the method.
  Invert()
  Contrast()
}

// Returns the colour White.
func New() Colour { return new_() }

// Returns the colour defined by (r, g, b).
func New3 (r, g, b byte) Colour { return new3(r,g,b) }

// Returns the colour defined by (r, g, b) with name n.
func New3n (n string, r, g, b byte) Colour { return new3n(n,r,g,b) }

// Returns a random colour.
func Rand() Colour { return random() }

// Returns the fore- and backgroundcolours at the start of the system
// for unmarked and marked objects.
func StartCols() (Colour, Colour) { return startCols() }
func StartColsA() (Colour, Colour) { return startColsA() }

func AllColours() []Colour { return allColours() }

// simple colours see cols.go
// RAL-Farben see ral.go
// X-Colours see X.go
