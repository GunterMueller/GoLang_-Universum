package phil

// (c) Christian Maurer   v. 221213 - license see µU.go

import (
  "sync"
  "math"
  "µU/env"
  "µU/col"
  "µU/mode"
  "µU/scr"
  "µU/N"
  "µU/ppm"
)
const (
  min =  5 // minimal and
  max = 13 // maximal number of philosophers
)
const (
  lock = iota
  unlock
)
const (
  satisfied = iota
  hungry
  starving
  hasRightFork
  hasLeftFork
  dining
  nStates
)
var (
  status [max]int
  colour = [max]col.Colour { col.Red(), col.Yellow(), col.Magenta(), col.Cyan(),
                             col.Green(), col.Orange(), col.LightBlue(), col.LightRed(),
                             col.LightYellow(), col.LightMagenta(), col.LightCyan(),
                             col.LightOrange(), col.LightGreen() }
  thinkColour col.Colour
  notUsedColour = col.DarkGray()
  name = [max+1]string { "Platon", "Sokrates", "Aristoteles", "Cicero", "Heraklit",
                         "Pythagoras", "Diogenes", "Thales", "Epikur", "Anaxagoras",
                         "Protagoras", "Demokrit", "Theophrast", "Nemo" }
  image [max+1]string
  text = [nStates]string { "    satt    ", "  hungrig   ", "sehr hungrig",
                           "sehr hungrig", "sehr hungrig", "  speisend  " }
  x0, y0 int
  x, y, x1, y1, xt, yt [max]int
  mutex, writeMutex sync.Mutex
)

func init() {
  scr.New (0, 0, mode.SVGA)
  thinkColour = scr.ScrColB()
  s := env.Arg(1)
  if k, ok := N.Natural (s); ok {
    if k < min { k = min }
    if k > max { k = max }
    NPhilos = k
  }
}

func left (i uint) uint {
  return (i + NPhilos - 1) % NPhilos
}

func right (i uint) uint {
  return (i + 1) % NPhilos
}

func writePlate (i uint, b bool) {
  r := uint(y0) / 4
  if NPhilos > 7 {
    r = uint(y0) / 6
  }
  scr.Circle (xt[i], yt[i], r)
  if b {
    scr.Circle (xt[i], yt[i], (2 * r) / 3)
  }
}

func writeFork (i uint) {
  scr.Line (x[i], y[i], x1[i], y1[i])
}

func _write (i uint) {
  writeMutex.Lock(); defer writeMutex.Unlock()
  if status[i] == satisfied {
    scr.Colours (thinkColour, thinkColour)
    scr.WriteGr (text[status[i]], xt[i] - 12 * 4, yt[i] - 8)
    writePlate (i, false)
    writePlate (i, true)
    scr.Colours (notUsedColour, thinkColour)
    writeFork (left (i))
    writeFork (i)
    im := ppm.New()
    im.Load (image[i])
    scr.WriteImage (im.Colours(), xt[i] - 30, yt[i] - 40)
//    ppm.Get (image[i], uint(xt[i]) - 30, uint(yt[i]) - 40)
    return
  }
//  ppm.Get (image[max], uint(xt[i] - 30), uint(yt[i]) - 40)
  scr.Colours (colour[i], thinkColour)
  scr.WriteGr (text[status[i]], xt[i] - 12 * 4, yt[i] - 8)
  writePlate (i, false)
  switch status[i] {
  case hasRightFork:
    writeFork (i)
  case hasLeftFork:
    writeFork (left (i))
  case dining:
    writeFork (i)
    writeFork (left (i))
    writePlate (i, true)
  }
}

func SitDownAtTable() (uint, bool) {
  mutex.Lock(); defer mutex.Unlock()
  for i := uint(0); i < NPhilos; i++ {
    if scr.UnderMouseGr (xt[i], yt[i], xt[i], yt[i], 50) && status[i] == satisfied {
      return i, true
    }
  }
  return max, false
}

func Start() {
  const twoPi = 2 * math.Pi
  for i := uint(0); i <= max; i++ {
    image[i] = env.Val ("GOSRC") + "/µU/phil/pics/" + name[i]
  }
  x0, y0 = int(scr.Wd()) / 2, int(scr.Ht()) / 2
  const f = 0.5
  for i := uint(0); i < NPhilos; i++ {
// Attention: the mathematical positiv sense is inverted on the screen,
// because lines count upwards !
// middlepoint of plates:
    w, r := float64(i) / float64(NPhilos), 0.75 * float64(y0)
    xt[i] = int(math.Trunc (r * math.Cos (twoPi * w) + f)) + x0
    yt[i] = int(math.Trunc (- r * math.Sin (twoPi * w) + f)) + y0
// endpoints of the fork to the right (with the same number):
    w, r = w + 0.5 / float64(NPhilos), float64(y0)
    x[i] = int(math.Trunc (r * math.Cos (twoPi * w) + f)) + x0
    y[i] = int(math.Trunc (- r * math.Sin (twoPi * w) + f)) + y0
    r = r / 2
    x1[i] = int(math.Trunc (r * math.Cos (twoPi * w) + f)) + x0
    y1[i] = int(math.Trunc (- r * math.Sin (twoPi * w) + f)) + y0
    _write (i)
  }
}

func changeStatus (p uint, s int) {
  mutex.Lock(); defer mutex.Unlock()
  status[p] = s
  _write (p)
}
