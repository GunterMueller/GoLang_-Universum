package scr

// (c) Christian Maurer   v. 170919 - license see µu.go

import (
  "µu/ker"
  "µu/mode"
  "µu/col"
  "µu/scr"
  . "µu/smok/utensil"
)
const (
  R = "Raucher mit"; V = "vorrat"
)
var (
  xm, ym, r0, r1 int
  la, ca, r uint
  lsm [NUtensils]uint
  csm [NUtensils]int
)

func init() {
  scr.New(0, 50, mode.VGA)
  xm, ym, r = int(scr.Wd()) / 2, int(scr.Ht()) / 2, scr.Wd() / 4
  la, ca = scr.NLines() / 2, scr.NColumns() / 2
  r0, r1 = int (r), (866 * int(r)) / 1000
  cr, cc:= len(R), r1 / int(scr.Wd1()) + 1
  l1:= r / 2 / scr.Ht1()
  lsm = [NUtensils]uint { l1 - 1, la + l1 - 1, la + l1 - 1 }
  csm = [NUtensils]int { -cr/2, -cc - cr/2, cc - cr/2 }
}

func table() {
  scr.ColourF (col.LightWhite())
  scr.Circle (xm, ym, r)
}

func agent (u uint) {
  scr.Lock()
  table()
  for u:= uint(0); u < NUtensils; u++ {
    scr.Colours (Colour(u), col.Black())
    scr.Write (R,             lsm[u],     uint(int(ca) + csm[u]))
    scr.Write (String(u) + V, lsm[u] + 1, uint(int(ca) + csm[u] - 1))
  }
  u1, u2:= Others(u)
  scr.Colours (Colour(u1), col.Black())
  scr.Write (String(u1), la - 1, ca - 2)
  scr.Colours (Colour(u2), col.Black())
  scr.Write (String(u2), la, ca - 2)
  scr.Unlock()
  ker.Sleep (1)
}

var
  ready chan bool = make (chan bool)

func smoke (u uint, a uint) {
  ker.Msleep (a * 200)
  x, y:= xm, ym
  switch u {
  case Papier:
    y -= r0
  case Tabak:
    x -= r1; y += r0 / 2
  case Hölzer:
    x += r1; y += r0 / 2
  }
  for i:= uint(3); true; i++ {
    scr.Lock()
    scr.ColourF (Colour(u))
    scr.Circle (x, y, i)
    scr.Unlock()
    ker.Msleep (50)
    scr.Lock()
    scr.ColourF (col.Black())
    scr.Circle (x, y, i)
    scr.Unlock()
    select { case <-ready: return; default: }
  }
}

const
  rings = 10

func start (u uint) {
  scr.Lock()
  scr.Colours (col.Black(), col.Black())
  scr.Write (R,             lsm[u],     uint(int(ca) + csm[u]))
  scr.Write (String(u) + V, lsm[u] + 1, uint(int(ca) + csm[u] - 1))
  table()
  scr.Unlock()
  for a:= uint(0); a < rings; a++ {
    go func (i uint) { smoke (u, i) } (a)
  }
}

func stop() {
  for a:= uint(0); a < rings; a++ {
    ready <- true
  }
}

const
  pause = 2

func smoker (u uint) {
  ker.Sleep (pause)
  scr.Lock()
  scr.Colours (col.Black(), col.Black())
  scr.Write (String(0), la - 1, ca - 2)
  scr.Write (String(0), la, ca - 2)
  scr.Unlock()
  start (u)
  ker.Sleep (2 * pause)
  stop()
}
