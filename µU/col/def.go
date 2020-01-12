package col

// (c) Christian Maurer   v. 191107 - license see µU.go

import
 . "µU/obj"
const
  P6 = 3
type
  Colour interface {

  Object // empty colour is black

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

  SetR (b byte); SetG (b byte); SetB (b byte)

// TODO Spec
  Cc() Stream

// TODO Spec
  Code() uint

// Returns true, if c is, what the name of the func says.
  IsBlack() bool
  IsWhite() bool
  IsLightWhite() bool

// Returns the rgb-values of x scaled to the range from 0 to 1.
  Float32() (float32, float32, float32)
  Float64() (float64, float64, float64)

// c is changed in a manner suggested by the name of the func.
  Invert()
  Contrast()
}

// Returns the colour White.
func New() Colour { return new_() }

// Returns the colour defined by (r, g, b) with name n.
func New3 (n string, r, g, b byte) Colour { return new3(n,r,g,b) }

// Returns a random colour.
func Rand() Colour { return random() }

func StartCols() (Colour, Colour) { return startCols() }
func StartColsA() (Colour, Colour) { return startColsA() }

// Pre: b is one of 4, 8, 15, 16, 24 or 32.
// depth() Colour { return = (b + 4) / 8.
func SetDepth (b uint) { setDepth(b) }

func Depth() uint { return depth } // in bytes - must not be altered after the call to SetDepth !

// Returns the number of available colours, depending on depth.
func NCols() uint { return nCols() }

func P6Encode (a, p Stream) { p6Encode(a,p) }
func P6Colour (a Stream) Colour { return p6Colour(a) }

func AllColours() []Colour { return allColours() }

// simple colours see cols.go
// RAL-Farben see ral.go
// X-Colours see X.go
