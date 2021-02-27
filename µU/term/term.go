package term

// (c) Christian Maurer   v. 210226 - license see µU.go

import (
  "µU/ker"
  "µU/obj"
  "µU/str"
  "µU/col"
  "µU/box"
  "µU/errh"
  "µU/r"
)
type
  errors = int; const (
  none = iota
  noExp
  after
  invalid
  noClose
  noFV
  toThe
  stop
  noArg
  nErrors
)
type
  typ = int; const (
  noType = iota
  realType
  opType
  varType
  funcType
  bracketType
)
const
  min = 10 // minimal width of Box
var (
  derivationTerm [r.NFunctions]*term
  errtext [nErrors]string
)
type
  term struct {
              typ
              string
              box.Box
           wd uint // width of Box
        value,
         real float64
           op r.Operation
  left, right *term
     variable string
    variables []string
            f r.Function
          arg,
     interior *term
              errors
 actual, fehl uint
       cF, cB col.Colour
             }

func init() {
  errtext[none]    = "ok"
  errtext[noExp]   = "expression missing"
  errtext[after]   = "term after '(' missing"
  errtext[invalid] = "invalid term"
  errtext[noClose] = "')' missing"
  errtext[noFV]    = "no func or var"
  errtext[toThe]   = "exponent invalid"
  errtext[stop]    = "truncated"
  errtext[noArg]   = "function argument invalid"
  for f := 1; f < r.NFunctions; f++ {
    derivationTerm[f] = new_(r.Derivation (f)).(*term)
//    println (r.FuncString(f) + "'(x) =", derivationTerm[f].string)
  }
}

func new0() *term {
  x := new(term)
  x.Box = box.New()
  x.Box.Wd (min)
  x.op = r.NoOp
  x.cF, x.cB = col.StartCols()
  x.variables = make([]string, 0)
  x.variable = ""
  x.f = r.Undef
  return x
}

func new_(s string) Term {
  t := new0()
  str.OffSpc (&s)
  if s == "" {
    return nil
  }
  t.string = s
  t = t.Sum().(*term)
  n := uint(len(s))
  if n < min {
    t.Box.Wd (n)
  }
  t.variables = r.Strings (s)
  if len(t.variables) == 0 {
    t.variable = ""
  } else {
    t.variable = t.variables[0]
  }
  return t
}

func (x *term) Type() string {
  switch x.typ {
  case realType:
    return ("realType")
  case opType:
    return ("opType")
  case varType:
    return ("varType")
  case funcType:
    return ("funcType")
  case bracketType:
    return ("bracketType")
  }
  return "noType"
}

func (x *term) imp (Y Term) *term {
  y, ok := Y.(*term)
  if ! ok { obj.TypeNotEqPanic (x, Y) }
  return y
}

func (x *term) Empty() bool {
  return str.Empty (x.string)
}

func (x *term) Clr() {
  x.wd = min
  x.string = str.New (x.wd)
  x.Box.Wd (x.wd)
  x.typ = noType
  x.value = 0
  x.real = 0
  x.op = r.NoOp
  x.left, x.right = nil, nil
  x.variable = ""
  x.variables = make([]string, 0)
  x.f = r.Undef
  x.arg = nil
  x.interior = nil
  x.errors = none
  x.actual, x.fehl = 0, 0
}

func (x *term) String() string {
  return x.string
}

func (x *term) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *term) Wd (w uint) {
  if w < min { w = min }
  x.wd = w
  x.Box.Wd (x.wd)
}

func (x *term) Write (l, c uint) {
  x.Box.Colours (x.cF, x.cB)
  x.Box.Write (x.string, l, c)
}

func (x *term) Edit (l, c uint) {
  x.Write (l, c)
  x.fehl = 0
  for {
    x.Box.Edit (&x.string, l, c)
    str.OffSpc (&x.string)
    str.Norm (&x.string, x.wd)
    if str.Empty (x.string) {
      x.Clr()
      break
    }
    if x.Val() != r.NaN() {
      break
    } else {
      errh.Error0 (errtext[x.errors])
    }
  }
  x.Write (l, c)
}

func (t *term) next() byte {
  if t.string == "" {
    return 0
  }
  b := t.string[0]
  t.string = t.string[1:]
  t.actual++
  return b
}

func (t *term) back (b byte) {
  t.string = string(b) + t.string
  t.actual--
}

func (t *term) rem (n uint) {
  t.string = t.string[n:]
  t.actual += n
}

func (x *term) Num() uint {
  return uint(len(x.variables))
}

func (x *term) Variables() []string {
  return x.variables
}

// func (x *term) Selected (v *string, l, c uint) bool {
//   x.tree.Construct (x.Variables)
//   return x.Variables.Selected (*v, l, c)
// }

func realTerm (x float64) Term {
  t := new0()
  t.typ = realType
  t.real = x
  t.string = r.String (t.real)
	return t
}

func opTerm (op r.Operation, left, right Term) Term {
 if left.(*term) == nil { panic ("left.(*term) == nil") }
  t := new0()
  t.typ = opType
  t.op = op
  t.left, t.right = left.(*term), right.(*term)
  t.variables = make([]string, 0)
 if t.left == nil { panic ("Mist") }
  n := len (t.left.variables)
  for i := 0; i < n; i++ {
    t.variables = append (t.variables, t.left.variables[i])
  }
  n = len (t.right.variables)
  for i := 0; i < n; i++ {
    t.variables = append (t.variables, t.right.variables[i])
  }
  t.variable = t.left.variable
  rv := t.right.variable
  for i := 0; i < len(rv); i++ {
    if _, ok := str.Pos (t.left.variable, rv[i]); ! ok {
      t.variable += string(rv[i])
    }
  }
  t.string = t.left.string + r.OpString (t.op) + t.right.string
  return t
}

func varTerm (v string) Term {
  t := new0()
  t.typ = varType
  t.variable = v
  t.string = v
	return t
}

func funcTerm (f r.Function, a Term) Term {
  t := new0()
  t.typ = funcType
  t.f = f
  t.arg = a.(*term)
  t.string = r.FuncString (t.f) + "(" + t.arg.string + ")"
  n := len(t.arg.variables)
  t.variables = make([]string, 0)
  for i := 0; i < n; i++ {
    t.variables = append (t.variables, t.arg.variables[i])
  }
  t.variable = t.arg.variable
	return t
}

func bracketTerm (t Term) Term {
  b := new0()
  b.typ = bracketType
  b.string = "(" + t.String() + ")"
  n := len(t.(*term).variables)
  b.variables = make([]string, 0)
  for i := 0; i < n; i++ {
    b.variables = append (b.variables, t.(*term).variables[i])
  }
  b.variable = t.(*term).variable
  b.interior = t.(*term)
  return b
}

func (x *term) Val() float64 {
  var val float64
  switch x.typ {
  case realType:
    val = x.real
  case opType:
    val = r.OpVal (x.op, x.left.Val(), x.right.Val())
  case funcType:
    val = r.FuncVal (x.f, x.arg.Val())
  case bracketType:
    val = x.interior.Val()
  default:
    val = r.NaN()
  }
  return val
}

func (x *term) Val1 (a float64) float64 {
  if len(x.variables) != 1 {
    return r.NaN()
  }
  var val float64
  switch x.typ {
  case noType:
    val = r.NaN()
  case realType:
    val = x.real
  case opType:
    val = r.OpVal (x.op, x.left.Val(), x.right.Val())
  case varType:
    val = a
  case funcType:
    val = r.FuncVal (x.f, x.arg.Val())
  case bracketType:
    val = x.interior.Val()
  }
  return val
}

func (x *term) Vals (f FuncVal) float64 {
  var val float64
  switch x.typ {
  case noType:
    val = r.NaN()
  case realType:
    val = x.real
  case opType:
    val = r.OpVal (x.op, x.left.Vals(f), x.right.Vals(f))
  case varType:
    val = f (x.variable)
  case funcType:
    val = r.FuncVal (x.f, x.arg.Vals(f))
  case bracketType:
    val = x.interior.Vals(f)
  }
  return val
}

func Null() Term {
  return realTerm (0)
}

func (t *term) fail (e errors) {
  t.errors = e
  t.fehl = t.actual
  ker.Panic (errtext[e])
}

func (t *term) Expression() Term {
  b := t.next()
  if b == 0 {
    t.fail (noExp); return nil
  }
  if b == '(' {
    if t.string == "" {
      t.fail (after); return nil
    }
    A := t.Sum()
    if A == nil {
      t.fail (invalid); return nil
    }
    b = t.next()
    if b != ')' {
      t.fail (noClose); return nil
    }
    return A
  }
  t.back (b)
  if x, n, ok := r.RealStarted (t.string); ok {
    t.rem (n)
    return realTerm (x)
  }
  if f, n, ok := r.FuncStarted (t.string); ok {
    t.rem (n)
    A := t.Expression()
    if A == nil {
      t.fail (noArg); return nil
    }
    return funcTerm (f, A)
  }
  if v, n, ok := str.StartsWithVar (t.string); ok {
    V := varTerm (v)
    t.rem (n)
    return V
  }
  t.fail (noFV); return nil
}

func (t *term) Power() Term {
  P := t.Expression()
  if P == nil {
    t.fail (noExp); return nil
  }
  for {
    b := t.next()
    if b == 0 {
      break
    }
    if op, ok := r.PowOp (b); ok {
      P1 := t.Power()
      if P1 == nil {
        t.fail (toThe); return nil
      }
      P = opTerm (op, P, P1)
      if P == nil {
        t.fail (invalid); return nil
      }
    } else {
      t.back (b)
      break
    }
  }
  return P
}

func (t *term) Product() Term {
  P := t.Power()
  if P == nil {
    t.fail (invalid); return nil
  }
  for {
    b := t.next()
    if b == 0 {
      break
    }
    if op, ok := r.MulDivOp (b); ok {
      P1 := t.Power()
      if P1 == nil {
        t.fail (invalid); return nil
      }
      P = opTerm (op, P, P1)
      if P == nil {
        t.fail (invalid); return nil
      }
    } else {
      t.back (b)
      break
    }
  }
  return P
}

func (t *term) Sum() Term {
  var S Term
  b := t.next()
  if b == 0 {
    t.fail (stop); return nil
  }
  if b == '-' {
    S = t.Product()
    if S == nil {
      t.fail (invalid); return nil
    }
    S = opTerm (r.Sub, Null(), S)
    if S == nil {
      t.fail (invalid); return nil
    }
  } else {
    t.back (b)
    S = t.Product()
    if S == nil {
      t.fail (invalid); return nil
    }
  }
  for {
    b = t.next()
    if b == 0 {
      break
    }
    if op, ok := r.AddSubOp (b); ok {
      S = opTerm (op, S, t.Product())
      if S == nil {
        t.fail (invalid); return nil
      }
    } else {
      t.back (b)
      break
    }
  }
  return S
}

func (t *term) isProduct() bool {
  return t.typ == opType && t.op == r.Mul
}

func (t *term) isPower() bool {
  return t.typ == opType && t.op == r.Pow
}

func (x *term) Eq (Y Term) bool {
  y := x.imp(Y)
  if x.typ != y.typ {
    return false
  }
  return x.string == y.string // &&
//         x.float64 == y.float64
}

var zähler int

func (t *term) Simplification() Term {
//
// >>> under development
//
// This function handles only some trivial cases.
//
// Missing is a systematic approach with a "normal form" roughly as follows:
// production only to left descending *-chains with 1 as left leaf,
// then sorting of the elements in this chain (powers e.g. to the front),
// simplifying by multiplying the powers in this chain
// (replacing quotients x / y by x * y^(-1) ?) to order polynoms by descending powers.
// ==>  still a large amount of work TODO :-(
//
  y := new_(t.string)
// println ("Simplification called with t ==", t.string)
  if t.Empty() { return y }
  switch y.(*term).typ {
  case noType:
    panic ("noType")
    return nil
  case realType:
    return y
  case opType:
    switch y.(*term).op {
    case r.Add:
      if y.(*term).left.typ == realType {
        if y.(*term).left.Val() == 0 {                     // 0 + x = 0
          return y.(*term).right
        }
        if y.(*term).right.typ == realType {
          return realTerm (y.(*term).left.Val() + y.(*term).right.Val())
        } else {
          return opTerm (r.Add, y.(*term).left, y.(*term).right)
        }
      }
      if y.(*term).right.typ == realType &&
         y.(*term).right.Val() == 0 {                      // x + 0 = 0 
        return y.(*term).left
      }
/*/
      if t.left.isProduct() 
//      if t.left.typ == opType && t.left.op == r.Mul {
        if t.left.left.string == t.right.string {
          t0 := opTerm (r.Add, t.left.right, realTerm (1))
          return opTerm (r.Mul, t.left, t0)                // x * y + x = x * (y + 1)
        }
        if t.left.right.string == t.right.string {
          t0 := opTerm (r.Add, t.left.left, realTerm (1))
          return opTerm (r.Mul, t0, t.left.right)          // x * y + y = (x + 1) * y
        }
//        if t.right.Type == opType && t.right.op == r.Mul {
        if t.right.isProduct() {
          if t.left.left.Eq (t.right.left) {
            t0 := opTerm (r.Add, t.left.right, t.right.right)
            return opTerm (r.Mul, t.left.left, t0)         // x * y + x * z = x * (y + z)
          }
          if t.left.right.string == t.right.right.string {
            t0 := opTerm (r.Add, t.left.left, t.right.left)
            return opTerm (r.Mul, t0, t.left.right)        // x * z + y * z = (x + y) * z
          }
          if t.left.string == t.right.left.string {
            t0 := opTerm (r.Add, t.right.right, realTerm (1))
            return opTerm (r.Mul, t.left, t0)              // x + x * y = x * (y + 1)
          }
          if t.left.string == t.right.right.string {
            t0 := opTerm (r.Add, t.left, realTerm (1))
            return opTerm (r.Mul, t0, t.right.left)        // x + y * x = (x + 1) * y
          }
        }
      }
/*/
      return y
    case r.Sub:
      if y.(*term).right.typ == realType {
        if y.(*term).right.Val() == 0 {                    // x - 0 = x
          return y.(*term).left
        }
        if y.(*term).left.typ == realType {
          return realTerm (r.OpVal (r.Sub, y.(*term).left.Val(), y.(*term).right.Val()))
        }
      }
/*/
      if t.left.isProduct() && t.right.isProduct() {
        if t.left.right.Eq (t.right.right) {
          t0 := opTerm (r.Sub, t.left.left, t.right.left)
          return opTerm (r.Mul, t0, t.left.right)          // x * z - y * z = (x - y) * z
        }
        if t.left.left.Eq (t.right.left) {
          t0 := opTerm (r.Sub, t.left.right, t.right.right)
          return opTerm (r.Mul, t.left.left, t0)           // x * y - x * z = x * (y - z)
        }
      }
/*/
// x - x * y = (x - 1) * y
// x * y - x = (x - 1) * y
      return y
    case r.Mul:
// println ("r.Mul", y.(*term).string)
      if y.(*term).right.typ == realType {
// println ("aa   ", y.(*term).left.String())
        if y.(*term).left.typ == opType {
          if y.(*term).left.left.typ == realType {             // (a * x) * b = ab * x
            p := realTerm (y.(*term).left.left.Val() * y.(*term).right.Val())
            return opTerm (r.Mul, p, y.(*term).left.right)
          }
        }
      }
      if y.(*term).right.typ == realType {
// println ("bb   ", y.(*term).left.String())
        return opTerm (r.Mul, y.(*term).right, y.(*term).left)     // x * a = a * x
      }
      if y.(*term).left.typ == realType {
        switch y.(*term).left.Val() {
        case 0:                                                    // 0 * x = 0
// println ("cc   ", y.(*term).left.String())
          return realTerm (0)
        case 1:                                                    // 1 * x = x
// println ("dd   ", y.(*term).left.String())
          return y.(*term).right
        }
        if y.(*term).right.typ == realType {                   // a * b = ab
          return realTerm (y.(*term).left.Val() * y.(*term).right.Val())
        }
        return opTerm (r.Mul, y.(*term).left, y.(*term).right)
      }
      if y.(*term).right.typ == realType {
// println ("ee   ", y.(*term).left.String())
        switch y.(*term).right.Val() {
        case 0:                                                    // x * 0 = 0
          return realTerm (0)
        case 1:                                                    // x * 1 = x
// println ("ff   ", y.(*term).left.String())
          return y.(*term).left
        }
      }
      if y.(*term).left.string == y.(*term).right.string {         // x * x = x^2
        return funcTerm (r.Sqr, y.(*term).left)
      }
/*/
// println ("r.Mul left", y.(*term).left.String())                           // cos(x)*1/sin(x)
   if y.(*term).left.left != nil {
     println ("r.Mul left.right", y.(*term).left.right.String())             // sin(x)
     println ("r.Mul left.left", y.(*term).left.left.String(),
                                 r.OpText (y.(*term).left.left.op))          // cos(x)*1  Mul
     if y.(*term).left.left.left != nil {
       println ("r.Mul left.left.left", y.(*term).left.left.left.String())   // cos(x)
     }
     if y.(*term).left.left.right != nil {
       println ("r.Mul left.left.right", y.(*term).left.left.right.String()) // 1
     }
   }
/*/
// println ("r.Mul right", y.(*term).right.String()) // cos(x)
      if y.(*term).left.isProduct() {
        if y.(*term).left.typ == opType {
// println ("A op ==", r.OpText(y.(*term).left.op))    // Div
        }
        if y.(*term).left.typ == opType && y.(*term).left.op == r.Mul {
// println ("A t0 ==")
          if y.(*term).left.left.string == y.(*term).right.string {  // (x * y) * x
            t0 := opTerm (r.Pow, y.(*term).right, realTerm(2))       // x^2
// println ("B t0 ==", t0.String())
            return opTerm (r.Mul, y.(*term).left.right, t0)          // y * x^2
          }
          if y.(*term).right.string == y.(*term).left.right.string { // (x * y) * y
            t0 := opTerm (r.Pow, y.(*term).right, realTerm (2))      // y^2
// println ("C t0 ==", t0.String())
            return opTerm (r.Mul, y.(*term).left.left, t0)           // x * y^2
          }
        }
        if y.(*term).left.isProduct() {
// x * (x * y) = x^2 * y
// x * (y * x) = x^2 * y
      }
      return y
    }
// x * y + x == x * (y + 1)
// x * y + y == (x + 1) * y
//
// x^y * x^z = x^(y+z)
// x^y * x = x^(y+1)
// x * x^y = x^(y+1)
//
// (1 / x) * y = y / x
// x * (1 / y) = x / y
    case r.Div:
      if y.(*term).left.string == y.(*term).right.string {
        return realTerm (1)
      }
      if y.(*term).left.typ == realType {
        if y.(*term).left.Val() == 0 {                       // 0 / x = 0
          return realTerm (0)
        }
        if y.(*term).right.typ == realType {             // a / b = a/b
          return realTerm (r.OpVal (r.Div, t.left.Val(), t.right.Val()))
        }
      }
      if y.(*term).right.typ == realType {
        if y.(*term).right.Val() == 1 {                      // x / 1 = x
          return y.(*term).left
        }
      }
      if y.(*term).right.typ == realType {               // x / a
        t0 := realTerm (1 / y.(*term).right.Val())           // 1/a
        return opTerm (r.Mul, t0, y.(*term).left)            // 1/a * x
      }
      return y
    case r.Pow:
      if y.(*term).left.typ == realType {
        a := y.(*term).left.Val()
        switch a {
        case 0:                                              // 0^x = 0
          return realTerm (0)
        case 1:                                              // 1^x = 1
          return realTerm (1)
        }
      }
      if y.(*term).right.typ == realType {
        switch y.(*term).right.Val() {
        case 0:                                              // x^0 = 1
          return realTerm (0)
        case 1:                                              // x^1 = x
          return t.left.Simplification()
        default:
          a := y.(*term).right.Val()                         // x^a
          if r.Integer(a) {
            if a > 0 {
              power := r.Power (y.(*term).left.Val(), uint(a))
              return realTerm (power)
            }
            a = -a                                           // a > 0
            power := r.Power (y.(*term).left.Val(), uint(a)) // x^(-a) = 1/x^a for a > 0
            return realTerm (1/power)
          }
          b := a * r.FuncVal (r.Log, y.(*term).left.Val())
          return realTerm (r.FuncVal (r.Exp, b))             // exp(a * log(x))
        }
      }
      if y.(*term).right.isPower() {
// println ("y.right", y.(*term).right.String())
        yz := y.(*term).right.Simplification()
// println ("yz", yz.(*term).string)
        return opTerm (r.Pow, y.(*term).left, yz)
      }
//                                                           // x^(y^z)
      switch y.(*term).left.op {
      case r.Pow:                                            // (x^y)^z
        t0 := y.(*term).left.right                           // y
        t1 := opTerm (r.Mul, t0, y.(*term).right)            // y * z
        return opTerm (r.Pow, y.(*term).left.left, t1)       // x^(y * z)
      case r.Mul:                                            // (x*y)^z
        t0 := y.(*term).right
        yl := y.(*term).left                                 // x
        yl = yl.left                                         // nil
if yl == nil { panic ("murx") }
//  XXX
// println ("yl ==", yl.string)
        t1 := opTerm (r.Pow, yl, t0)                         // x^z
        t2 := opTerm (r.Pow, y.(*term).left.right, t0)       // y^z
        return opTerm (r.Mul, t1, t2).Simplification()       // x^z * y^z
      }
// sqrt(x)^2 = x
// sqrt(x^y) = x^(y/2)
      return y
    }
  case funcType:                                             // f(x)
    if y.(*term).arg.typ == realType {                   // f(a)
      return realTerm (r.FuncVal (y.(*term).f, y.(*term).arg.Val()))
    }
    return y
  case varType:
    return y
  case bracketType:
    return y
  }
/*/
  x + a + b = x + a+b
  x + x^2 = x^2 + x

   if right < left
     left <-> right)
        if isProduct (left) {
          if right.Less (left.right) {
   (x * z) * y == (x * y) * z, if y < z
            left.right <-> right
            left.Simplify()
          if isMulDivTerm (right) {
            if IsReal (right.left) {
   x * (a . y) = a * (x . y)", 0)
              left <-> right.left
              right.Simplify()
   x * (a . y)
              return
            }
          }

   normalize products: only produce *-chains descending to the left
// x * (y * z) = (x * y) * z
  
// (x^y * z) * x = x^(y + 1) * z
// (z * x^y) * x = z * x^(y + 1)
// (x * z) * x^y = x^(y + 1) * z
// (z * x) * x^y = z * x^(y + 1)
// (xz * t) * xy = x * (z + y) * t
// (t * xy) * xz = t * x * (y + z)
  
  
// sqrt(x) * sqrt(y)) == sqrt(x * y)
  
// (x / x^a) = x^(1-a)
// x^a / x^b = x^(a-b)
// x^a / x == x^(a-1)
  
// exp(log(x)) == x
  
// sqrt (x^y) == x^(y/2)
// log(x^y) == y * log(x)
// f(g(x)) = x, if g is inverse of f
/*/
  return y
}

func (x *term) Insert (v string, t Term) Term {
  switch x.typ {
  case noType:
    panic ("cannot insert in term of noType")
  case realType:
    return realTerm (x.Val())
  case opType:
    return opTerm (x.op, x.left.Insert (v, t), x.right.Insert (v, t))
  case varType:
    if x.variable == v {
      return t
    } else {
      return varTerm (x.variable)
    }
  case funcType:
    return funcTerm (x.f, x.arg.Insert (v, t))
  case bracketType:
    return x.interior.Insert (v, t)
  }
  return nil
}

func (x *term) Derivation (v string) Term {
  if x.Empty() {
    return nil
  }
  switch x.typ {
  case realType:
    return realTerm(0)
  case varType:
    if x.variable == v {
      return realTerm(1)
    }
    return realTerm(0)
  case opType:
    f, g := x.left, x.right
    f1 := f.Derivation (v).Simplification()
    g1 := g.Derivation (v).Simplification()
    f1g := opTerm (r.Mul, f1, g).Simplification()
    fg1 := opTerm (r.Mul, f, g1).Simplification()
    switch x.op {
    case r.NoOp:
      panic ("NoOp")
      return nil
    case r.Add, r.Sub:                                   // (f +|- g)' = f' +|- g'
      return opTerm (x.op, f1, g1).Simplification()
    case r.Mul:                                          // (f * g)' = f' * g + f * g'
      return opTerm (r.Add, f1g, fg1).Simplification()
    case r.Div:                                          // (f / g)' = (f' * g - f * g') / g^2
      t0 := opTerm (r.Sub, f1g, fg1).Simplification()
      t1 := opTerm (r.Mul, g, g).Simplification()
      return opTerm (r.Div, t0, t1).Simplification()
    case r.Pow:
                                                         // (t^a)' == a * t^(a-1) * t' for real a
      if x.right.typ == realType {
        a := x.right.Val()                               // a
        t0 := opTerm (r.Pow, x.left, realTerm (a - 1))   // t^(a-1)
        t1 := opTerm (r.Mul, realTerm (a), t0)           // a*t^(a-1)
          t := opTerm (r.Mul, t1, x.left.Derivation (v)) // a*t^(a-1) * t'
        return t.Simplification()                        // a*t^(a-1) * t'
      }
                                                         // (f^g)' == (f^g) * (g * log f)'
                                                         // (f^g)' == (f^g) * (g' * log f + g / f)
      t0 := opTerm (r.Pow, x.left, x.right)              // f^g
      t1 := opTerm (r.Mul, x.right, funcTerm (r.Log, x.left)) // g*log(f)
      t2 := t1.Derivation (v)                            // (g*log(f))' = g / f + g' * log(f)
      t3 := bracketTerm (t2)
      return opTerm (r.Mul, t0, t3) // .Simplification()
    }
  case funcType: // f(g(x))' = f'(g(x)) * g'(x)
    t0 := derivationTerm[x.f]
    t0.Insert ("x", varTerm (v))
    if x.arg.typ == varType {
      t0.Insert (v, x.arg)
      return t0
    }
    f1 := t0.Insert (v, x.arg)                           // f'(g(x))
    if f1 == nil { panic ("Kacke") }
    g1 := x.arg.Derivation (v).Simplification()          // g'(x)
    return opTerm (r.Mul, f1, g1)
  case bracketType:                                      // (f)' = f'
// println (bracketTerm (x.interior.Derivation (v)).String())
    return bracketTerm (x.interior.Derivation (v))
  }
  panic ("not reachable")
}
