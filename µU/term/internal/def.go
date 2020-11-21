package tree

// (c) Christian Maurer   v. 201111 - license see µU.go

import (
  . "µU/obj"
  "µU/real"
)
type
  ValFunc func (string) float64
type
  Tree interface {

  Clearer
  Equaler

  String() string

// Returns the number of variables in x.
  Num() uint

// The v's are the variables of x.
  Set (v []string)

// Returns all variables of x.
  Variables() []string

// Returns true, if x is a tree with only one node of type float64.
  IsReal() bool

  RealTree (x float64) Tree
  OpTree (op real.Operation, l, r Tree) Tree
  VarTree (v string) Tree
  FuncTree (f real.Function, a Tree) Tree
  BracketTree (t Tree) Tree

// x is a bit simpified.
  Simplify()

// Pre: i < Num().
// Every instance of s in x is replaced by t.
  Insert (s string, t Tree)

// Returns the value of x, if x does not contain any variable, otherwise NaN.
  Val() float64

// Returns the value of x at the a's, if it is defined for them; otherwise NaN.
  Vals (f ValFunc) float64

//  Construct (vs varis.Variables) {

  Build (f real.Function, t Tree)

// Pre: i < Num.
// Returns the derivation of x to the i-th variable of x.
  Derivation (s string) Tree // TODO

// Liefert genau dann true, wenn entweder x genau eine Variable enthält
// oder wenn x mehrere Variablen enthält und die/der Benutzer/in
// an der Bildschirmposition (l, c) eine Variable eingegeben hat, die in x vorkommt.
//  Selected (l, c uint) (bool, vari.Variable)
//  Selected (v vari.Variable, l, c uint) bool
}

func New() Tree { return new_() }
