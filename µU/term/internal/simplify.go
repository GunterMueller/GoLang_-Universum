package tree

// (c) Christian Maurer   v. 201113 - license see µU.go

import (
//  . "µU/obj"
  "µU/real"
  "µU/errh"
)

const Test = false

func Fall (s string) {
  if Test { errh.Error0 (s + " erkannt") }
}

// This function handles only some trivial cases.
//
// Missing is a systematic approach with a "normal form" roughly as follows:
// production only to left descending *-chains with 1 as left leaf,
// then sorting of the elements in this chain (powers e.g. to the front),
// simplifying by multiplying the powers in this chain
// (replacing quotients x / y by x * y^(-1) ?) to order polynoms by descending powers.
//
// ==>  still a really large amount of work TODO :-(
//
func (t *tree) Simplify() {
  if t.Empty() { return }
  switch t.Type {
  case realType:
    // nothing todo
  case opType:
    t.left.Simplify()
    t.right.Simplify()
    switch t.op {
    case real.Add:
      if t.left.Val() == 0 && t.right.Val() == 0 {
        t.Copy (t.RealTree (0))
      } else if t.left.Val() == 0 {
        t.Copy (t.right)
      } else if t.right.Val() == 0 {
        t.Copy (t.left)
      } else if t.right.IsReal() && t.left.IsReal() {
        t.Copy (t.RealTree (real.OpVal (real.Add, t.left.float64, t.right.float64)))
      } else if t.left.isProduct() && t.right.isProduct() {
        if t.left.right.Eq (t.right.right) { // y * x + z * x -> (y + z) * x
          t0 := t.OpTree (real.Add, t.left.left, t.right.left)
          t.Copy (t.OpTree (real.Mul, t0, t.left.right))
        } else if t.left.left.Eq (t.right.left) { // x * y + x * z -> x * (y + z)
          t0 := t.OpTree (real.Add, t.left.right, t.right.right)
          t.Copy (t.OpTree (real.Mul, t.left.left, t0))
        }
      }
// x + x * y == x * (y + 1)
// y + x * y == (x + 1) * y
    case real.Sub:
      if t.left.Val() == 0 && t.right.Val() == 0 {
        t.Copy (t.RealTree (0))
      } else if t.left.Val() == 0 {
        t.Copy (t.FuncTree (real.Chs, t.right))
      } else if t.right.Val() == 0 {
        t.Copy (t.left)
      } else if t.right.IsReal() && t.left.IsReal() {
        t.Copy (t.RealTree (real.OpVal (real.Sub, t.left.float64, t.right.float64)))
      } // TODO more as above
    case real.Mul:
      if t.left.Val() == 0 || t.right.Val() == 0 {
        t.Copy (t.RealTree (0))
      } else if t.left.Val() == 1 {
        t.Copy (t.right)
      } else if t.right.Val() == 1 {
        t.Copy (t.left)
      } else if t.right.IsReal() && t.left.IsReal() {
        t.Copy (t.RealTree (real.OpVal (real.Mul, t.left.float64, t.right.float64)))
      } else if t.left.Eq (t.right) {
        t.Copy (t.OpTree (real.Mul, t.RealTree(2), t.left))
      } else if t.left.isProduct() {
        if t.right.Eq (t.left.left) {
          t0 := t.FuncTree (real.Sqr, t.right)
          t.Copy (t.OpTree (real.Mul, t0, t.left.right))
        }
      } // etc.
// (x * y) * y = x * y^2
// x * (x * y) = x^2 * y
// y * (x * y) = y^2 * x
// x * y + x == x * (y + 1)
// x * y + y == (x + 1) * y
// x^y * x^z = x^(y+z)
// x^y * x = x^(y+1)
// x * x^y = x^(y+1)
// x * x = x^2
// (1 / x) * y = y / x
// x * (1 / y) = x / y
    case real.Div:
      if t.left.Val() == 0 {
        t.Copy (t.RealTree (0))
      } else if t.right.Val() == 1 {
        t.Copy (t.left)
      } else if t.right.IsReal() && t.left.IsReal() {
        t.Copy (t.RealTree (real.OpVal (real.Div, t.left.float64, t.right.float64)))
      } else if t.right.IsReal() {
        t.Copy (t.OpTree (real.Mul, t.RealTree (1 / t.right.Val()), t.right))
      } else if t.left.Eq (t.right) {
        t.Copy (t.RealTree (1))
      }
    case real.Pow:
      if t.right.Val() == 1 {
        t.Copy (t.left)
      } else if t.left.Val() == 0 {
        t.Copy (t.RealTree (0))
      } else if t.left.Val() == 1 {
        t.Copy (t.RealTree (1))
      } else if t.right.Val() == 0 {
        t.Copy (t.RealTree (1))
      } else if t.right.Val() == 1 {
        t.Copy (t.left)
      } else if t.left.isPower() {
        t0 := t.OpTree (real.Mul, t.left.right, t.right)
        t.Copy (t.OpTree (real.Pow, t.left.left, t0))
      } else if t.left.isProduct() {
        t0 := t.OpTree (real.Pow, t.left.left, t.right)
        t1 := t.OpTree (real.Pow, t.left.right, t.right)
        t.Copy (t.OpTree (real.Mul, t0, t1))
      }
// sqrt(x)^2 = x
// sqrt(x^y) = x^(y/2)
    }
  case funcType:
    t.arg.Simplify()
    if t.arg.IsReal() {
      t.Copy (t.RealTree (real.FuncVal (t.f, t.arg.float64)))
    }
  case varType:
    // nothing todo
  }
/*/
   Ausklammern und Zusammenfassen fehlt 
   e.g. x + 2 + 3 = x + 5              
        x + 3 + x = 2 * x + 3          
        x + x^2 = x^2 + x              
        etc.                           

   if right < left
     left <-> right)
        if isProduct (left) {
          if BaumKleiner (right, left.right) {
   (x * z) * y == (x * y) * z, if y < z
            left.right <-> right
            left.Simplify()
          if isMulDivTree (right) {
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
  
// (x^y * z) * x = x^(y +1) * z
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
}
