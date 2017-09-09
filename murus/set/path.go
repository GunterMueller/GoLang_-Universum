package set

// (c) Christian Maurer   v. 140528 - license see murus.go

import (
  . "murus/obj"
  "murus/ker"
)
type
  node struct {
              *tree
         next *node
              }

// x.path is the list of nodes from actual up to root. 
// min/max == true, iff the actual node w.r.t. Less
// is the smallest/largest object in x.
func (x *set) defPath() (bool, bool) {
  t:= x.anchor
  x.path = &node { t, nil }
  min, max:= true, true
//  if e, ok:= x.actual.Any.(Editor) { e.Write (2, 0); errh.Error0("x.actual") }
  for {
    if t == nil { ker.Panic ("piset defPath: t == nil") }
    if t == x.actual {
      if ! Eq (t.Any, x.actual.Any) { ker.Panic ("piset.defPath Eq bug") }
      break
    }
    if Less (x.actual.Any, t.Any) {
      t, max = t.left, false
    } else {
      t, min = t.right, false
    }
    x.path = &node { t, x.path }
  }
  return min && t.left == nil, max && t.right == nil
}

func (x *set) below (f bool) *tree {
  if f {
    return x.path.tree.right
  }
  return x.path.tree.left
}

func (x *set) somethingBelow (f bool) bool {
  return x.below (f) != nil
}

func (x *set) abovePointsToCurrent (f bool) bool {
  t:= x.path.next.tree.right
  if ! f {
    t = x.path.next.tree.left
  }
  return t == x.path.tree
}

func (x *set) up() {
  x.path = x.path.next
}

func (x *set) pointer() *tree {
  return x.path.tree
}
