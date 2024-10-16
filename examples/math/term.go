package main

// (c) Christian Maurer   v. 241011 - license see  mU.go

import (
  "µU/scr"
  "µU/str"
  "µU/errh"
  "µU/R"
  "µU/term"
)

func main() {
  scr.NewWH (0, 0, 1024, 32); defer scr.Fin()
  R.SetFormat (0)
/* choose one of the following strings:
  a := "x + y"
  a := "2 * x + 3 * x"
  a := "2 * x"
  a := "2"
  a := "tan(x)"
  a := "2 * x^3 + 3 * x^2 + 5 * x + 7"
  a := "3 * sin(x)"
  a := "sin(x) * 3"
  a := "2 * x + 3 * sin(x)"
  a := "x+y"
  a := "3 * x^2 + 5 * x"
  a := "2 * 3 * x"
  a := "2 * x * 3"
  a := "x * 3 * x"
  a := "sin(x)^2"
  a := "sin(exp(x))"
  a := "x^3"
  a := "2 * x^3 + 3 * x^2 + 5 * x + 7"
  a := "x^7^2"
  a := "exp(sin(x))"
  a := "sin(x)^cos(x)"
  a := "x * 2 * 4"
  a := "sin(x) + sin(y)"
  a := "sin(x)"
  a := "sin(cos(sin(x)))"
  a := "sin(x + tan(y)) * exp(2 * x + 2 * y) + 7"
  a := "sin(x + tan(x)) * exp(2 * x)"
*/
  a := "sin(exp(x))"
  str.OffSpc (&a)
  t := term.New (a)
  if t == nil { errh.Error0 (a + " is no valid term"); return }
//  for i := uint(0); i < t.Num(); i++ { errh.Error (">" + t.Variables()[i] + "<", i) }
//  errh.Error0 ("New t == " + t.String() + " of " + t.Type())
//  t = t.Simplification(); errh.Error0 ("simplified t == " + t.String()); return
/* choose to which variable should be derived
  v := "x"
  v := "y"
*/
  v := "x"
  d := t.Derivation (v)
  errh.Error0 ("derivation to " + v + " == " + d.String())
/*/
  d = d.Simplification() // is buggy
  errh.Error0 ("derivation to " + v + " == " + d.String())
/*/
}
