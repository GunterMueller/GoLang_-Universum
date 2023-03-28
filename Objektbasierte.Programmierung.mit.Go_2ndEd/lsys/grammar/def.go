package grammar

// (c) Christian Maurer   v. 230303 - license see µU.go

import
  "µU/col"
const (
  MaxL =  4
  MaxR = 80
  Comment      = '%'
  Step         = 'F'
  YetiStep     = 'f'
  TurnLeft     = '+' // around z-axis
  TurnRight    = '-'
  Invert       = '|'
  TiltDown     = '_' // around x-axis
  TiltUp       = '^'
  RollLeft     = byte(92) // '\' // around y-axis
  RollRight    = '/'
  BranchStart  = '['
  BranchEnd    = ']'
  PolygonStart = '{'
  PolygonEnd   = '}'
)
type
  Symbol = byte
var (
  StartColour col.Colour
  Startword string
  Startangle, Turnangle float64
  NumIterations uint
  colours = []col.Colour {col.Brown(),      // n
                          col.Red(),        // r
                          col.LightRed(),   // l
                          col.Orange(),     // o
                          col.Green(),      // g
                          col.DarkGreen(),  // d
                          col.Cyan(),       // c
                          col.LightBlue(),  // e
                          col.Blue(),       // b
                          col.Magenta(),    // m
                          col.Black(),      // k
                          col.Gray(),       // y
                          col.White(),      // w
                          col.LightWhite(), // z
                         }
)

func Initialize (s string) { initialize(s) }

func IsColour (s Symbol) (col.Colour, bool) { return isColour(s) }

// Pre: s is not empty.
// Returns true, iff there is a rule with a left side starting with s.
func ExRule (s string) bool { return exRule(s) }

// Pre: There is at most one rule with s as left side.
// Returns the right side of the rule with left side s,
// if such a rule exists; otherwise "".
func Derivation (s string) string { return derivation(s) }
