package smok

// (c) Christian Maurer   v. 221223 - license see µU.go

import (
  "µU/time"
  "µU/env"
  "µU/str"
//  "µU/mode"
  "µU/col"
  "µU/scr"
)
const (
  agent = uint(iota)
  smokerIn
  smokerOut
)
const (
  paper = uint(iota)
  tobacco
  matches
)
const
  raucher = "Raucher mit"
var
  Hölzer = str.Lat1("Hölzer")
var (
  text   = [3]string     {"Papier",         " Tabak",         Hölzer}
  colour = [3]col.Colour {col.LightWhite(), col.LightBrown(), col.SandYellow()}
)
var (
  xm, ym, r0, r1 int
  la, ca, r uint
  lsm [3]uint
  csm [3]int
)

func others (u uint) (uint, uint) {
  return (u + 1) % 3, (u + 2) % 3
}

func init() {
  if env.Call() == "smokers" {
    scr.NewMax() // XXX  scr.New(0, 50, mode.VGA)
    xm, ym, r = int(scr.Wd()) / 2, int(scr.Ht()) / 2, scr.Wd() / 4
    la, ca = scr.NLines() / 2, scr.NColumns() / 2
    r0, r1 = int (r), (866 * int(r)) / 1000
    cr, cc := len(raucher), r1 / int(scr.Wd1()) + 1
    l1 := r / 2 / scr.Ht1()
    lsm = [3]uint { l1 - 1, la + l1 - 1, la + l1 - 1 }
    csm = [3]int { -cr/2, -cc - cr/2, cc - cr/2 }
  }
}

func table() {
  scr.ColourF (col.LightWhite())
  scr.Circle (xm, ym, r)
}

func write (u uint, c col.Colour) {
  scr.Colours (c, col.Black())
  scr.Write (raucher,            lsm[u],     uint(int(ca) + csm[u]))
  scr.Write (text[u] + "vorrat", lsm[u] + 1, uint(int(ca) + csm[u]) - 1)
}

func writeAgent (u uint) {
  scr.Lock()
  table()
  for i := uint(0); i < 3; i++ {
    write (i, colour[i])
  }
  u1, u2 := others(u)
  scr.Colours (colour[u1], col.Black())
  scr.Write (text[u1], la - 1, ca - 2)
  scr.Colours (colour[u2], col.Black())
  scr.Write (text[u2], la, ca - 2)
  scr.Unlock()
  time.Sleep (1)
}

var
  ready chan bool = make (chan bool)

func smoke (u uint, a uint) {
  time.Msleep (a * 200)
  x, y := xm, ym
  switch u {
  case paper:
    y -= r0
  case tobacco:
    x -= r1; y += r0 / 2
  case matches:
    x += r1; y += r0 / 2
  }
  for i := uint(3); true; i++ {
    scr.Lock()
    scr.ColourF (colour[u])
    scr.Circle (x, y, i)
    scr.Unlock()
    time.Msleep (50)
    scr.Lock()
    scr.ColourF (col.Black())
    scr.Circle (x, y, i)
    scr.Unlock()
    select {
    case <-ready:
      return
    default:
    }
  }
}

const
  rings = 10

func start (u uint) {
  scr.Lock()
  write (u, col.Black())
  table()
  scr.Unlock()
  for a := uint(0); a < rings; a++ {
    go func (i uint) { smoke (u, i) } (a)
  }
}

func stop() {
  for a := uint(0); a < rings; a++ {
    ready <- true
  }
}

const
  pause = 2

func writeSmoker (u uint) {
  time.Sleep (pause)
  scr.Lock()
  write (u, colour[u])
  scr.Colours (col.Black(), col.Black())
  scr.Write (text[0], la - 1, ca - 2)
  scr.Write (text[0], la, ca - 2)
  scr.Unlock()
  start (u)
  time.Sleep (2 * pause)
  stop()
}
