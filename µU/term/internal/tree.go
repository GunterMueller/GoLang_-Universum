package tree

// (c) Christian Maurer   v. 201113 - license see µU.go

import (
  . "µU/obj"
  "µU/real"
)
const
  M = 80
type
  Type = int; const (
  noType = iota
  realType
  opType
  funcType
  varType
)
type
  tree struct {
              Type
            s []string
              float64
           op real.Operation
  left, right *tree
            f real.Function
          arg *tree
              string
              }
var
  derivationTree[real.NFunctions]*tree


func init() {
  for f := real.Id; f < real.NFunctions; f++ {
    derivationTree[f] = new_().(*tree)
  }
}

func new_() Tree {
  x := new(tree)
  x.Type = noType
  x.s = make([]string, 0)
  x.float64 = real.NaN()
  x.op = real.NoOp
  x.left, x.right = nil, nil
  x.f = real.Undef
  x.arg = nil
  x.string = ""
  return x
}

func (x *tree) imp (Y Any) *tree {
  y, ok := Y.(*tree)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *tree) Set (v []string) {
  n := len(v)
  x.s = make([]string, n)
  for i := 0; i < n; i++ {
    x.s[i] = v[i]
  }
/*/
  if len (v) == 1 {
    x.string = v[0]
  }
/*/
}

func (x *tree) Variables() []string {
  s := x.s
  switch x.Type {
  case opType:
    s = append (s, x.left.Variables()...)
    s = append (s, x.right.Variables()...)
  case funcType:
    s = append (s, x.arg.Variables()...)
  case varType:
    s = append (s, x.string)
  }
  return s
}

func (x *tree) Num() uint {
  return uint(len(x.s))
}

func (x *tree) Empty() bool {
  return x.Type == noType
}

func (x *tree) Clr() {
  x.Type = noType
  x.s = make([]string, 0)
  x.float64 = real.NaN()
  x.op = real.NoOp
  x.left, x.right = nil, nil
  x.f = real.Undef
  x.arg = nil
  x.string = ""
}

func (x *tree) Eq (Y Any) bool {
  y := x.imp(Y)
  if x.Type != y.Type {
    return false
  }
  for i := 0; i < len(x.s); i++ {
    if x.s[i] != y.s[i] {
      return false
    }
  }
  switch x.Type {
  case noType:
    return y.Type == noType
  case realType:
    return real.Eq (x.float64, y.float64)
  case varType:
    return x.string == y.string
  case opType:
    return x.op == y.op && x.left.Eq (y.left) && x.right.Eq (y.right)
  case funcType:
    return x.f == y.f && x.arg.Eq (y.arg)
  }
  panic ("not reachable")
}

func (x *tree) Copy (Y Any) {
  y := x.imp(Y)
  x.Type = y.Type
  switch y.Type {
  case realType:
    x.float64 = y.float64
  case opType:
    x.op = y.op
    x.left.Copy (y.left)
    x.right.Copy (y.right)
  case funcType:
    x.f = y.f
println ("Copy: x ==", x.String(), "   y ==", y.String())
 if x.arg == nil { println ("Copy: x.arg == nil") }
    x.arg.Copy (y.arg)
  case varType:
    x.string = y.string
  }
  n := len(y.s)
  x.s = make([]string, n)
  for i := 0; i < n; i++ {
    x.s[i] = y.s[i]
  }
}

func (x *tree) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *tree) String() string {
  if x == nil { println ("String: x == nil") }
  switch x.Type {
  case realType:
    return real.String(x.float64)
  case opType:
    t := x.left.String() + real.OpString (x.op) + x.right.String()
    if x.op == real.Add || x.op == real.Sub || x.op == real.Pow {
      t = "(" + t + ")"
    }
    return t
  case funcType:
    if x.arg == nil { println ("String: x.arg == nil") }
    return real.FuncString (x.f) + "(" + x.arg.String() + ")" // XXX crash 4
  case varType:
    return x.string
  }
  return "" // noType
}

/*/
func (t *tree) Construct (vs varis.Variables) {
  if x.Empty() { return }
  switch x.Type {
  case realType:
    ;
  case varType:
    vs.Ins (...)
  case opType:
    left.Construct (vs)
    right.Construct (vs)
  case funcType:
    arg.Construct (vs)
  }
}

func (x *tree) Selected (v string, l, c uint) bool {
  x.Construct (x.Variables)
  return Variablen.Selected (x.Variables, v, l, c)
}
/*/

func (t *tree) isSum() bool {
  if t.Type == opType {
    return t.op == real.Add
  }
  return false
}

func (t *tree) isDff() bool {
  if t.Type == opType {
    return t.op == real.Sub
  }
  return false
}

func (t *tree) isQuotient() bool {
  if t.Type == opType {
    return t.op == real.Div
  }
  return false
}

func (t *tree) isProduct() bool {
  if t.Type == opType {
    return t.op == real.Mul
  }
  return false
}

func (t *tree) isPower() bool {
  if t.Type == opType {
    return t.op == real.Pow
  }
  return false
}

func (t *tree) Val() float64 {
  if len(t.s) == 0 {
    if t.string == "" {
      switch t.Type {
      case realType:
        return t.float64
      case opType:
        return real.OpVal (t.op, t.left.Val(), t.right.Val())
      case funcType:
        return real.FuncVal (t.f, t.arg.Val())
      }
    }
  }
  return real.NaN() // noType, varType
}

func (t *tree) Val1 (x float64) float64 {
  if len(t.s) == 1 {
    switch t.Type {
    case realType:
      return t.float64
    case opType:
      return real.OpVal (t.op, t.left.Val(), t.right.Val())
    case funcType:
      return real.FuncVal (t.f, t.arg.Val())
    case varType:
      return x
    }
  }
  return real.NaN() // noType
}

func (t *tree) Vals (f ValFunc) float64 {
  switch t.Type {
  case realType:
    return t.float64
  case varType:
    return f (t.string)
  case opType:
    return real.OpVal (t.op, t.left.Vals (f), t.right.Vals (f))
  case funcType:
    return real.FuncVal (t.f, t.arg.Vals(f))
  }
  return real.NaN()
}

func (x *tree) IsReal() bool {
  if x.Empty() { return false }
  return x.Type == realType
}

func (x *tree) isVariable() bool {
  if x.Empty() { return false }
  return x.Type == varType
}

func (t *tree) RealTree (x float64) Tree {
  y := new_().(*tree)
  y.Type = realType
  y.float64 = x
  return y
}

func (x *tree) OpTree (op real.Operation, l, r Tree) Tree {
  if l.Empty() || r.Empty() { return new_() }
  t := new_().(*tree)
  t.Type = opType
  t.op = op
//  t.left, t.right = l.Clone().(*tree), r.Clone().(*tree) // XXX
  t.left, t.right = l.(*tree), r.(*tree)
  return t
}

func (x *tree) VarTree (v string) Tree {
  t := new_().(*tree)
  t.Type = varType
  t.string = v
  return t
}

func (x *tree) FuncTree (f real.Function, a Tree) Tree {
  t := new_().(*tree)
  t.Type = funcType
  n := len(x.s)
  t.s = make([]string, n)
  for i := 0; i < n; i++ {
    t.s[i] = x.s[i]
  }
  t.f = f
//  t.arg = a.Clone().(*tree) // XXX
  t.arg = a.(*tree)
  t.string = x.string
  return t
}

func (x *tree) BracketTree (k Tree) Tree { // XXX
  t := new_().(*tree)
  t.Type = k.(*tree).Type
  switch t.Type {
  case noType:
    ;
  case realType:
    t.float64 = k.(*tree).float64
  case opType:
    t.op = k.(*tree).op
    t.left = k.(*tree).left.Clone().(*tree)
    t.right = k.(*tree).right.Clone().(*tree)
  case funcType:
    t.f = k.(*tree).f
    t.arg = k.(*tree).arg.Clone().(*tree)
  case varType:
    t.string = k.(*tree).string
  }
  return t
}

func (t *tree) Less (U Any) bool {
  u := t.imp(U)
  if t.Empty() || u.Empty() { return false }
  if t.IsReal() {
    return ! u.IsReal()
  } else if t.IsReal() {
    return false
  }
  if t.isVariable() {
    if u.isVariable() {
      return false // XXX return t.Variables.Less (u.Variables)
    } else {
//    if u.isPower( {
//      if u.left.isVariable() {
//        return t.Less (u.left)
//      }
//    }
    }
  }
  if t.Type == opType && t.op == real.Pow {
    if t.left.Less (u) {
      return true
    }
  }
  if t.Type == opType && t.op == real.Pow {
    if t.Less (u.left) {
      return true
    }
  }
  return false
}

var ci = 0

func (x *tree) Insert (s string, t Tree) {
ci++; println (">>> ", ci)
println ("Insert: s ==", s, " t ==", t.String())               // 1,2) x x    3,4) x exp(x)
  switch x.Type {
  case opType:
    x.left.Insert (s, t)
    x.right.Insert (s, t)
  case funcType:
println ("Insert funcType: x.arg ==", x.arg.String())          // 1) x   3) x
    x.arg.Insert (s, t)
  case varType:
// println ("Insert: t.arg ==", t.(*tree).arg.String())          //
println ("Insert varType: x.string ==", x.string, "  s ==", s) // 2) x x   4) x x
    if x.string == s {
      x.Type = t.(*tree).Type
      switch t.(*tree).Type {
      case realType:
        x.float64 = t.(*tree).float64
      case opType:
        x.op = t.(*tree).op
        x.left.Copy (t.(*tree).left)
        x.right.Copy (t.(*tree).right)
      case funcType:
        x.f = t.(*tree).f
        if x.arg == nil {
          panic ("ganz große Scheiße")
        }
        x.arg.Copy (t.(*tree).arg) // geht deshalb nicht
      case varType:
        x.string = t.(*tree).string
      }
      n := len(t.(*tree).s)
      x.s = make([]string, n)
      for i := 0; i < n; i++ {
        x.s[i] = t.(*tree).s[i]
      }
    }
  } // noType, realType
}

func (x *tree) Build (f real.Function, t Tree) {
//  fs := real.FuncString (f)
//  ts := t.String()
// println ("Build called with", fs, ts) // "sin" "cos(x)"
//  derivationTree[f].Copy (t) // crash
  vs := t.Variables()
  if len(vs) > 1 { panic ("len(vs) > 1") }
  if len(vs) == 0 {
// println ("variables: none")
    derivationTree[f].Type = realType
    derivationTree[f].float64 = t.(*tree).float64
  } else {
    if f == real.Sin { println ("variables:", vs[0]) }
    derivationTree[f].Type = funcType
    derivationTree[f].f = t.(*tree).f
    derivationTree[f].arg = x.VarTree (vs[0]).(*tree)
  }
//  if f == real.Sin { println ("derivationTree[Sin] ==", derivationTree[f].String()) }
}

func (x *tree) Derivation (s string) Tree {
println ("tree Derivation of x to", s, "==", x.String())
  if x.Empty() {
println ("tree empty")
    return new_()
  }
  switch x.Type {
  case realType:
    return x.RealTree(0)
  case varType:
println ("nasse Bluse", x.string, s, ".")
    if x.string == s {
      return x.RealTree(1)
    }
    return x.RealTree(0)
  case opType:
    switch x.op {
    case real.Add, real.Sub:
// (f +|- g)' = f' +|- g'
// println (x.left.String(), x.right.String())
// println (x.left.Derivation(s).(*tree).String(), x.right.Derivation(s).(*tree).String())
      t := x.OpTree (x.op, x.left.Derivation(s), x.right.Derivation(s))
      t.Simplify()
      return t
    case real.Mul:
// (f * g)' = f' * g + f * g'
 println ("left", x.left.Derivation(s).String(), x.right.String())
 println ("right", x.left.String(), x.right.Derivation(s).String())
      t0 := x.OpTree (real.Mul, x.left.Derivation (s), x.right)
 println ("t0", t0.String())
      t1 := x.OpTree (real.Mul, x.left, x.right.Derivation (s))
 println ("t1", t1.String())
      t := x.OpTree (real.Add, t0, t1)
 println ("t", t.String())
      t.Simplify()
      return t
    case real.Div:
// (f / g)' = (f' * g - f * g') / g^2
      t0 := x.OpTree (real.Mul, x.left.Derivation (s), x.right)
      t1 := x.OpTree (real.Mul, x.left, x.right.Derivation (s))
      t2 := x.OpTree (real.Sub, t0, t1)
      t3 := x.OpTree (real.Mul, x.right, x.right)
      t := x.OpTree (real.Div, t2, t3)
      t.Simplify()
      return t
    case real.Pow:
// (f^a)' == a * f(a-1) * f', if a is a number
// (f^g)' == (f^g) * (g * log f)' in the general case
      if x.right.IsReal() {
        t0 := x.OpTree (real.Pow, x.left, x.RealTree (x.right.float64 - 1))
        t1 := x.OpTree (real.Mul, x.RealTree (x.right.float64), t0)
        t := x.OpTree (real.Mul, t1, x.left.Derivation (s))
        t.Simplify()
        return t
      }
      t0 := x.OpTree (real.Pow, x.left, x.right).(*tree)
      t1 := x.OpTree (real.Mul, x.right, x.FuncTree (real.Log, x.left))
      t := x.OpTree (real.Mul, t0, t1.Derivation (s))
      t.Simplify()
      return t
    }
  case funcType:
    fs := real.FuncString (x.f)
    t := derivationTree[x.f]
    t.Insert ("x", x.VarTree (s))
    if x.arg.Type == varType {
println ("affe", x.String(), fs)                              // exp(x) exp
println ("esel", t.String())                                  // exp(x)
println ("geier", x.arg.String())                             // x
      t.Insert (s, x.arg)
println ("return t >" + t.String() + "<")                     // exp(x)
      return t
    } else { // (f(g(x)))' = f'(g(x)) * g'(x)
println ("kamel  t == >" + t.String() + "<", fs)              // cos(x) sin
println ("ratte  x.arg == >" + x.arg.String() + "<")          // exp(x)
      ad := x.arg.Derivation (s)
println ("tiger  x.arg.Derivation == >" + ad.String() + "<")  // exp(x)
      t.Insert (s, x.arg) // t == f'(g(x))                    // XXX crash
println ("zebra  t == >" + t.String() + "<")
      return x.OpTree (real.Mul, t, ad)
    }
  }
  panic ("not reachable")
}
