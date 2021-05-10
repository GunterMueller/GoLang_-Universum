package term

// (c) Christian Maurer   v. 210506 - license see µU.go
//
// >>> under development

import (
  . "µU/obj"
  "µU/col"
)
type
  FuncVal = func (string) float64
type
  Term interface {

  Clearer

  Eq (Y Term) bool

  Write (l, c uint)
  Edit (l, c uint)
  col.Colourer

  Colours (f, b col.Colour)
  Wd (w uint)

  String() string

  Type() string

// Returns the number of variables in x.
  Num() uint

// Returns the variables of x.
  Variables() []string

// Returns x, but a bit simplified.
  Simplification() Term

// Returns a term ... TODO
// Every occurence of the variable v of x is replaced by t.
// >>> implementation under development
  Insert (v string, t Term) Term

// Pre: x does not contain any variables.
// Returns the value of x.
  Val() float64

// Pre: x contains exactly one variable.
// Returns the value of x for with value a for that variable.
  Val1 (a float64) float64

// Pre: f defines values for all variables in x.
// Returns the value of x.
  Vals (f FuncVal) float64

// Pre: v occurs in x.
// Returns the derivation of x to the variable v.
// >>> implementation under development
  Derivation (v string) Term

// Returns true, if either x contains exactly one variable or
// iff x contains more variables and the user at the screen position (l, c)
// has edited a variable, that is contained in x.
//  Selected (l, c uint) (bool, string)
}

// letter = 'A' | 'B' | ... | 'Z' | 'a' | 'b' | ... | 'z'
// digit = '0' | '1' | ... | '9'
// variable = letter { letter | digit }
// function = 'sqr' | 'exp' | 'sin' etc. (see type Function in µU/r/def.go)
// expression = int | uint | float64 | variable | function '('expression')' | '('sum')'
// power = expression { '^' power }
// product = power { '*'|'/' power }
// sum = ['-'] product { '+'|'-' product }
//
// Pre: s describes a term due to the above EBNF.
// Returns that term.
func New (s string) Term { return new_(s) }
