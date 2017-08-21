package dgra

// (c) murus.org  v. 170810 - license see murus.go

// XXX uncompleted - different hosts missing, load missing, embedding missing !

import (
  "murus/kbd"
  "murus/col"
  "murus/mode"
  "murus/scr"
  "murus/box"
  "murus/nat"
  "murus/ego"
  "murus/bnat"
  "murus/vtx"
  "murus/edg"
  "murus/host"
  "murus/nchan"
  "murus/gra"
)

func construct() DistributedGraph {
  const max = uint(16)
  cf, ca, cb := col.Blue, col.Red, col.LightWhite
  scr.New(900, 100, mode.HQVGA); defer scr.Fin()
  scr.ScrColours (cf, cb); scr.Cls()
  scr.Colours (cf, cb)
  scr.Write ("number of vertices (4..16)", 9, 0)
  b := bnat.New(2)
  b.Colours (cf, cb)
  var n uint
  for {
    b.Clr(); b.Edit (9, 27)
    n = b.Val(); if n > 4 && n <= 16 { break }
  }
  scr.Write ("click the vertices in position", 9, 0)
  wd := nat.Wd(n)
  g := gra.New (false, vtx.New(bnat.New(wd), wd, 1), edg.New(false, uint16(nchan.Port0)))
  g.SetWrite (vtx.W, edg.W)
  v := make([]vtx.Vertex, n)
  x, y := make([]int, n), make([]int, n)
  for i := uint(0); i < n; i++ {
    for {
      c, _ := kbd.Command()
      if c == kbd.Here {
        x[i], y[i] = scr.MousePosGr()
        break
      }
    }
    b0 := bnat.New(wd)
    b0.SetVal(i)
    v[i] = vtx.New (b0, wd, 1)
    v[i].Set (x[i], y[i])
    v[i].Colours (cf, cb); v[i].ColoursA (ca, cb)
    g.Ins(v[i])
    g.Write()
  }
  scr.Write("edge   --   ? (yes = Enter)    ", 9, 0)
  for i := uint(0); i < n; i++ {
    for j := i + 1; j < n; j++ {
      scr.Write(nat.String(i), 9, 5)
      scr.Write(nat.String(j), 9, 9)
      loop: for {
        switch c, _ := kbd.Command(); c {
        case kbd.Enter:
          e := edg.New(false, nchan.Port(n, i, j, 1))
          e.Label(false)
          e.SetPos0 (v[i].Pos()); e.SetPos1 (v[j].Pos())
          e.Colours(cf, cb); e.ColoursA(ca, cb)
          g.Edge(e)
          g.Write()
          break loop
        case kbd.Back, kbd.Esc:
          break loop
        }
      }
    }
  }
  var m uint
  scr.Write("diameter of the graph:     ", 9, 0)
  for {
    b.Clr(); b.Edit (9, 23)
    m = b.Val(); if m > 0 && m < n { break }
  }
  scr.Write("name of graph:                ", 9, 0)
  bx := box.New()
  bx.Wd(14)
  name := ""
  bx.Edit(&name, 9, 15)
  g.Name(name)
  h := make([]host.Host, n)
  for i := uint(0); i < n; i++ { h[i] = host.Localhost() } // XXX
  g.Ex (v[ego.Me()])
  d := New (g)
  n = g.Num() - 1
  d.SetHosts (h)
  d.SetDiameter (m)
  g.Fin()
  return d
}
