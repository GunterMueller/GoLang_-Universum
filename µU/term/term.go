package term

// (c) Christian Maurer   v. 201113 - license see µU.go

// >>> under development -- not yet to be used, because there are still several errors !

import (
  "µU/obj"
  "µU/z"
  "µU/str"
  "µU/col"
  "µU/box"
  "µU/errh"
  "µU/real"
  "µU/term/internal"
)
const
  M = 40 // scr.Columns() // XXX
type
  errors = int; const (
  none = iota
  stop
  noClose
  noOpen
  noNumber
  noOp
  noFV
  invalid
  nErrors
)
type
  term struct {
              string
       actual,
         fehl uint
              tree.Tree
              errors
       cF, cB col.Colour
              }
var (
  derivationTerm [real.NFunctions]*term
  bx = box.New()
  errtext [nErrors]string
)

func new_() Term {
  x := new(term)
  x.string = str.New(M)
  x.Tree = tree.New()
  x.cF, x.cB = col.StartCols()
  return x
}

func (x *term) imp (Y obj.Any) *term {
  y, ok := Y.(*term)
  if ! ok { obj.TypeNotEqPanic (x, Y) }
  return y
}

func (x *term) Num() uint {
  return x.Tree.Num()
}

func (x *term) Variables() []string {
  return x.Tree.Variables()
}

func (x *term) Empty() bool {
  return x.Tree.Empty()
  return str.Empty (x.string)
}

func (x *term) Clr() {
  x.string = str.New (M)
  x.actual = 0
  x.Tree.Clr()
  x.errors = none
}

func (x *term) Eq (Y obj.Any) bool {
  y := x.imp(Y)
  return x.string == y.string
  return x.string == y.string &&
         x.actual == y.actual &&
         x.Tree.Eq (y.Tree)
}

func (x *term) Copy (Y obj.Any) {
  y := x.imp(Y)
  x.actual = y.actual
  x.Tree.Copy (y.Tree)
  x.string = y.string
  x.errors = none
}

func (x *term) Clone() obj.Any {
  y := new_()
  y.Copy (x)
  return y
}

/*/
func (x *term) Selected (v *string, l, c uint) bool {
  x.Tree.Construct (x.Variables)
  return x.Variables.Selected (*v, l, c)
}
/*/

func (x *term) Val() float64 {
  return x.Tree.Val()
}

func (x *term) Vals (f FuncVal) float64 {
  return x.Tree.Vals (f)
}

func (x *term) Defined (s string) bool {
  str.OffSpc (&s)
  if str.Empty (s) {
    x.Clr()
    return true
  }
  x.string = s
  x.Tree.Set (real.Strings (s))
  x.build()
  if x.errors == none {
    return true
  }
  x.Clr()
  return false
}

func (x *term) String() string {
  x.string = x.Tree.String()
  return x.string
}

func (x *term) SetColours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *term) Write (l, c uint) {
  if c + str.ProperLen (x.string) >= M {
//  serious problem // XXX
  } else { // str.ProperLen (x.string) <= M - c - 1
//    bx.Wd (M - c - 1) // XXX
    bx.Wd (M)
  }
  bx.Colours (x.cF, x.cB)
  bx.Write (x.string, l, c)
}

func (x *term) Edit (l, c uint) {
  x.Write (l, c)
  x.fehl = 0
  for {
    bx.Edit (&x.string, l, c)
    str.OffSpc (&x.string)
    str.Norm (&x.string, M)
    if str.Empty (x.string) {
      x.Clr()
      break
    }
    x.build()
    if x.errors == none {
      break
    }
    errh.Error0 (errtext[x.errors])
  }
  x.Write (l, c)
}

func (x *term) FuncTree (f real.Function, a tree.Tree) tree.Tree {
  return x.Tree.FuncTree (f, a)
}

func (x *term) fail (e errors, f uint) {
  x.errors, x.fehl = e, f
  panic (errtext[x.errors])
}

func (x *term) tree (s string) tree.Tree {
  if s == "" {
    x.fail (stop, 0)
  }
// println ("get", s)
  if s[0] == '(' {
    s = s[1:]
    if s == "" { x.fail (stop, 0) }
    l := len(s)
    if s[l-1] != ')' {
      x.fail (noClose, 0)
    }
    s = s[:l-1]
    p := l - 1
    a := x.tree (s) // tree.BracketTree (x.tree (s))
    s = s[p:]
    if s == "" {
      return a
    }
    if op, ok := real.IsOp (s[0]); ok {
      s = s[1:]
      if s == "" { x.fail (stop, 0) }
      return x.Tree.OpTree (op, a, x.tree (s))
    }
    return a
  } else if s[0] == '-' {
    s = s[1:]
    return x.Tree.FuncTree (real.Chs, x.tree(s))
  } else if z.IsDigit(s[0]) {
    if r, p, ok := real.RealStarted (s); ok {
      a := x.Tree.RealTree (r)
      if p == uint(len(s)) {
        return a
      }
      s = s[p:]
      if op, ok := real.IsOp (s[0]); ok {
        s = s[1:]
        if s == "" { x.fail (stop, 0) }
        return x.Tree.OpTree (op, a, x.tree(s))
      } else {
        x.fail (noOp, 0)
      }
      if s == "" {
        panic ("s empty")
      } else {
        return a
      }
    }
  } else if z.IsLetter (s[0]) {
    if f, p, ok := real.FuncStarted (s); ok {
      s = s[p:]
      if s[0] != '(' {
        x.fail (noOpen, 0)
      }
      p := str.RightBr (s)
      if s[p] == ')' {
        ft := x.Tree.FuncTree (f, x.tree (s[1:p]))
        s = s[p+1:]
        if s == "" { return ft }
        if op, ok := real.IsOp (s[0]); ok {
          if s == "" {
            x.fail (invalid, 0)
          }
          s = s[1:]
          ot := x.OpTree (op, ft, x.tree (s))
          return ot
        } else {
          x.fail (invalid, 0)
        }
        return ft
      } else {
        x.fail (noClose, 0)
      }
    } else {
      if v, p, ok := str.StartsWithVar (s); ok {
        vt := x.Tree.VarTree (v)
        if p == uint(len(s)) {
          return vt
        } else {
          s = s[1:]
          if op, ok := real.IsOp (s[0]); ok {
            s = s[1:]
            if s == "" {
              x.fail (invalid, 0)
            }
            ot := x.OpTree (op, vt, x.tree (s))
            return ot
          } else {
            x.fail (invalid, 0) // bluse
          }
        }
      } else {
      }
      x.fail (noFV, 0)
    }
  } else {
    x.fail (invalid, 0)
  }
  panic ("evil end")
}

func (x *term) build() {
  x.actual = 0
  x.errors = none
  x.fehl = 0
  x.Tree = x.tree (x.string) // bluse
  if x.errors == none {
    x.Simplify()
  }
}

func (x *term) Simplify() {
  x.Tree.Simplify()
  x.string = x.Tree.String()
}

func (x *term) Insert (s string, t Term) {
  tree := t.(*term).tree (x.string)
  x.Tree.Insert (s, tree)
  x.Simplify()
}

func (x *term) Derivation (s string) Term {
  t := new_().(*term)
  t.Tree = x.Tree.Derivation (s)
  t.Simplify()
  return t
}

func init() {
  errtext[none]     = "null problemo"
  errtext[stop]     = "truncated"
  errtext[noClose]  = "')' missing"
  errtext[noOpen]   = "'(' missing"
  errtext[noNumber] = "no number"
  errtext[noOp]     = "no op"
  errtext[noFV]     = "no func or var"
  errtext[invalid]  = "invalid character"
  for f := 1; f < real.NFunctions; f++ {
    derivationTerm[f] = new_().(*term)
    t := real.Derivation (f)
    derivationTerm[f].string = t
    derivationTerm[f].build()
    d := derivationTerm[f].tree (t)
    derivationTerm[f].Tree.Build (f, d)
    fs := real.FuncString (f)
    if derivationTerm[f].errors == none {
//      println ("derivation of", fs, "at", s, "is", derivationTerm[f].String())
    } else {
      panic ("derivation faulty for " + fs)
    }
  }
}
