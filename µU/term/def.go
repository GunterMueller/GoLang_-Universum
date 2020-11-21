package term

// (c) Christian Maurer   v. 201111 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/real"
  "µU/term/internal"
)
type
  FuncVal = func (string) float64
type
  Term interface {

  Clearer
  Equaler
  Stringer
  Editor

// Returns the number of variables in x.
  Num() uint

// Returns the variables of x.
  Variables() []string

// f/b is the fore-/backgroundcolour of x.
  SetColours (f, b col.Colour)

// x is a bit simplified.
  Simplify()

// Every occurence of the variable s of x is replaced by t.
  Insert (s string, t Term)

//
  VarTree (v string) tree.Tree
//
  FuncTree (f real.Function, a tree.Tree) tree.Tree

// Pre: s occurs in x.
// Returns the derivation of x to the variable s.
  Derivation (s string) Term

// Liefert den Wert von x, falls x keine Variable enthält, andernfalls NaN.
  Val() float64

// Liefert den Wert von x an den Stelle a für die Variablen von x,
// falls er für sie definert ist, andernfalls NaN.
  Vals (f FuncVal) float64

// Liefert genau dann true, wenn entweder x genau eine Variable enthält
// oder wenn x mehrere Variablen enthält und die/der Benutzer/in
// an der Bildschirmposition (l, c) eine Variable eingegeben hat, die in x vorkommt.
//  Selected (l, c uint) (bool, string)
}

func New() Term { return new_() }
