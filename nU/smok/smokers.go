package smok

// (c) Christian Maurer   v. 230326 - license see nU.go

import ("time"; "nU/col"; "nU/scr")

const raucher = "Raucher mit"
var (
  text   = [3]string     {"Papier",         " Tabak",         "Hölzer" }
  colour = [3]col.Colour {col.FlashWhite(), col.LightBrown(), col.Yellow()}
  cm, lm, r0, r1 uint
  la, ca, r uint
  lsm [3]uint
  csm [3]int
  ready = make (chan bool)
)

func init_() {
  cm, lm, r = scr.NColumns() / 4, scr.NLines() / 2, 3 * scr.NLines() / 8
  la, ca = scr.NLines() / 2, scr.NColumns() / 2
  r0, r1 = r, (866 * r) / 1000
  cr, cc := len(raucher), 2 * (int(r1) + 1)
  l1 := r / 2
  lsm = [3]uint { l1 - 2, la + l1 - 1, la + l1 - 1 }
  csm = [3]int { -cr/2, -cc - cr/2, cc - cr/2 }
}

func table() {
  scr.ColourF (col.FlashWhite())
  scr.Circle (lm, cm, r)
}

func write (u uint, f col.Colour) {
  scr.Colours (f, col.Black())
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
  scr.ColourF (colour[u1])
  scr.Write (text[u1], la - 1, ca - 2)
  scr.ColourF (colour[u2])
  scr.Write (text[u2], la, ca - 2)
  scr.Unlock()
  time.Sleep (1e9)
}

func smoke (u uint, a uint) {
  time.Sleep (time.Duration(a) * 200 * 1e6)
  c, l := cm, lm
  switch u {
  case paper:
    l -= r0
  case tobacco:
    c -= r1; l += r0 / 2
  case matches:
    c += r1; l += r0 / 2
  }
  for i := float64(3); true; i += 0.1 {
    r := uint(i)
    scr.Lock()
    scr.ColourF (colour[u])
    scr.Circle (l, c, r)
    scr.Unlock()
    time.Sleep (50 * 1e6)
    scr.Lock()
    scr.ColourF (col.Black())
    scr.Circle (l, c, r)
    scr.Unlock()
    select {
    case <-ready:
      return
    default:
    }
  }
}

const rings = 10

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

const pause = 2

func writeSmoker (u uint) {
  time.Sleep (pause * 1e9)
  scr.Lock()
  write (u, colour[u])
  scr.Colours (col.Black(), col.Black())
  scr.Write (text[0], la - 1, ca - 2)
  scr.Write (text[0], la, ca - 2)
  scr.Unlock()
  start (u)
  time.Sleep (2 * pause * 1e9)
  stop()
}
