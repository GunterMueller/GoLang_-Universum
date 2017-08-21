package nelect

// (c) murus.org  v. 160101

import (
//  "murus/ker"
  "murus/kbd"
  "murus/col"
  "murus/scr"
  "murus/nat"
  "murus/errh"
  "murus/node"
  "murus/gra"
)
/*
func write (s string, l, c uint) {
  scr.Lock()
  scr.Write (s, l, c)
  scr.Unlock()
}
*/

func (e *netElection) log (i ...int) {
// this is suboptimal TODO
  if ! e.bool { return }
  for k:= 0; k < len(i); k++ {
    if i[k] < 0 { i[k] = -i[k] }
  }
  f, b := col.Blue, col.LightWhite
  scr.Colours (f, b)
  nat.Colours (f, b)
  switch len(i) {
  case 1:
    scr.Write (" tid                        ", 8, 0)
    nat.Write (uint(i[0]), 8,  5)
  case 2:
    scr.Write (" tid      ntid              ", 8, 0)
    nat.Write (uint(i[0]), 8,  5)
    nat.Write (uint(i[1]), 8, 15)
  case 3:
    scr.Write (" tid      ntid      nntid   ", 8, 0)
//                        1         2
//              012345678901234567890123456789
    nat.Write (uint(i[0]), 8,  5)
    nat.Write (uint(i[1]), 8, 15)
    nat.Write (uint(i[2]), 8, 26)
  }
  errh.Hint ("weiter mit Enter")
  kbd.Wait (true)
  errh.DelHint()
}

/*
func error (s string, i int) {
  errh.Error (s, uint(i))
}

func hint (s string) {
  errh.Hint (s)
}

func showScr (i uint) {
  errh.Error ("elected is", i)
}
*/

func (e *netElection) logm (g gra.Graph) {
  g.Trav3Cond (node.O, node.O3)
//  errh.Error ("log", e.me)
}
