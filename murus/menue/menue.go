package menue

// (c) murus.org  v. 161216 - license see murus.go

import (
  . "murus/obj"; "murus/str"
  "murus/col"; "murus/scr"; "murus/box"; "murus/errh"
  "murus/sel"
)
type
  menue struct {
               string
        isMenu,
     withTitle bool
     nextLevel []*menue
               Stmt "to Stmt"
               uint "last position"
          next *menue
               }
var (
  bx = box.New()
  depth uint
)

func newMenue (s string) Menue {
  x := new (menue)
  x.string = str.Lat1 (s) + " "
  str.Norm (&x.string, scr.NColumns())
  x.isMenu = true
  x.Stmt = Null
  return x
}

func (x *menue) imp (Y Menue) *menue {
  y, ok := Y.(*menue)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *menue) Leaf (s Stmt, w bool) {
  if x.nextLevel != nil { return }
  x.withTitle = w
  x.isMenu = false
  x.Stmt = s
}

func (x *menue) Ins (Y Menue) {
  y := x.imp (Y)
  if y == nil || ! x.isMenu { return }
  n := uint(len (x.nextLevel))
  if n >= scr.NLines() - 2 { return }
  x.nextLevel = append (x.nextLevel, y)
}

func (x *menue) selected (l, c uint) bool {
  n := uint(len (x.nextLevel))
  if n == 0 || ! x.isMenu { return false }
  if n == 1 { return true }
  errh.Head (x.string)
  errh.Hint (errh.ToSelect)
  i := x.uint
  sel.Select (func (p, l, c uint, f, b col.Colour) {
                bx.Colours (f, b); bx.Write (x.nextLevel[p].string, l, c)
              }, n, scr.NLines() - 2, scr.NColumns(), &i, 2, 0, col.MenuF, col.MenuB)
  if i < n {
    x.uint = i
    x.next = x.nextLevel[i]
  }
  errh.DelHead()
  errh.DelHint()
  return i < n
}

func (x *menue) Exec() {
  depth ++
  if x.isMenu {
    for {
      if x.selected (0, 0) {
        x.next.Exec()
      } else {
        break
      }
    }
  } else { // is leaf
    scr.Cls()
    if x.withTitle {
      errh.Head (x.string)
    }
    x.Stmt()
    scr.Cls()
  }
  depth --
  if depth == 0 { scr.Fin() }
}
