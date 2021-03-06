package set

// (c) Christian Maurer   v. 201004 - license see µU.go

import (
  . "µU/obj"
  "µU/ker"
)
type
  node struct {
              *tree
         next *node
              }

// x.path is the list of nodes from actual up to root. 
// min/max == true, iff the actual node w.r.t. Less is the smallest/largest object in x.
func (x *set) defPath() (bool, bool) {
  t := x.anchor
  x.path = &node { t, nil }
  min, max := true, true
  for {
    if t == x.actual {
      if ! Eq (t.Any, x.actual.Any) { ker.Panic ("set.defPath Eq bug") }
      break
    }
    if t == nil { ker.Panic ("set defPath: t == nil") }
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
  t := x.path.next.tree.right
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
