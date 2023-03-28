package block

// (c) Christian Maurer   v. 230107 - license see µU.go

import (
  "µU/ker"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  . "bahn/kilo"
  . "bahn/farbe"
)
type
  quintuple struct {
                   uint // number of block
                   Kilometrierung
              x, y int
                   }
var
  all [NK][]*quintuple

func init() {
  for k := Mit; k < NK; k++ {
    all[k] = make([]*quintuple, 0)
  }
}

func newq (n uint, k Kilometrierung, x, y int) *quintuple {
  X := new(quintuple)
  X.uint = n
  X.Kilometrierung = k
  X.x, X.y = x, y
  return X
}

func (X *block) testen (k Kilometrierung, x, y int) {
  return
  X.Ausgeben (Freifarbe)
  var f col.Colour
  if k == Mit {
    f = col.Green()
  } else {
    f = col.Red()
  }
  scr.ColourF (f)
  scr.Circle (x, y, 8)
  kbd.Wait (true)
}

func (X *block) ins (n uint, k Kilometrierung, x, y int) {
  if B[n].Empty() {
    return
  }
  X.testen (k, x, y)
  all[k] = append (all[k], newq (n, k, x, y))
}

func ex (k Kilometrierung, x, y int) (uint, bool) {
  for i := uint(0); i < uint(len(all[k])); i++ {
    a := all[k][i]
    if a.x == x && a.y == y {
      return a.uint, true
    }
  }
  return 0, false
}

func ok() bool {
  for k := Mit; k < NK; k++ {
    for i := uint(0); i < uint(len(all[k])); i++ {
      a := all[k][i]
      _, _, x, y := a.uint, a.Kilometrierung, a.x, a.y
      if _, ok := ex (Gegen, x, y); ! ok {
        return false
      }
    }
  }
  return true
}

func nPaare() uint {
  return uint(len(all[Mit]))
}

func paar (i uint) (uint, uint) {
  a := all[Mit]
  if i >= uint(len(a)) { ker.Shit() }
  s := a[i]
  n := s.uint
  n1, o1 := ex (Gegen, s.x, s.y)
  if o1 {
    return n, n1
  }
  return 0, 0
}
