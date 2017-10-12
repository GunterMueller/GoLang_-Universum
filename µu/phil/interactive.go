package phil

// (c) Christian Maurer   v. 170918 - license see µu.go

import (
  "sync"; "math"
  "µu/env"
  "µu/col"
  "µu/mode"
  "µu/scr"
  "µu/img"
)
const (
  min =  5 // minimal and
  max = 12 // maximal number of philosophers
)
var (
  colour = [max]col.Colour { col.Red(), col.DarkYellow(), col.Cyan(), col.Magenta(), col.LightGreen(),
                             col.Orange(), col.Green(), col.Brown(), col.Blue(), col.Pink(),
                             col.Gray(), col.LightOrange() }
  thinkColour col.Colour
  notUsedColour = col.DarkGray()
  name = [max+1]string { "Thales", "Pythagoras", "Heraklit", "Anaxagoras", "Diogenes",
                         "Protagoras", "Sokrates", "Demokrit", "Platon", "Aristoteles",
                         "Epikur", "Cicero", "Nemo" }
  image [max+1]string
  status1 [max]state // < nStates
  text = [nStates]string { "    satt    ", "  hungrig   ", "sehr hungrig",
                           "sehr hungrig", "sehr hungrig", "  speisend  " }
  writeMutex sync.Mutex
  x0, y0 int
  x, y, x1, y1, xt, yt [max]int
)

func init() {
  if env.Call() == "philos" {
    scr.New (0, 0, mode.XGA)
    thinkColour = scr.ScrColB()
  }
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
  writeMutex.Lock()
  defer writeMutex.Unlock()
  if status[i] == satisfied {
    scr.Colours (thinkColour, thinkColour)
    scr.WriteGr (text[status[i]], xt[i] - 12 * 4, yt[i] - 8)
    writePlate (i, false)
    writePlate (i, true)
    scr.Colours (notUsedColour, thinkColour)
    writeFork (left (i))
    writeFork (i)
    img.Get (image[i], uint(xt[i]) - 50, uint(yt[i]) - 50)
    return
  }
  img.Get (image[max], uint(xt[i] - 50), uint(yt[i]) - 50)
  scr.Colours (colour[i], thinkColour)
  scr.WriteGr (text[status1[i]], xt[i] - 12 * 4, yt[i] - 8)
  writePlate (i, false)
  switch status1[i] {
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

func changeStatus1 (i uint, s state) {
  mutex.Lock()
  status1[i] = s
  _write (i)
  mutex.Unlock()
}

func WannaEat() (uint, bool) {
  mutex.Lock()
  defer mutex.Unlock()
  for i := uint(0); i < NPhilos; i++ {
    if scr.UnderMouseGr (xt[i], yt[i], xt[i], yt[i], 50) && status1[i] == satisfied {
       return i, true
    }
  }
  return max, false
}

func Start() {
  const twoPi = 2 * math.Pi
  for i := uint(0); i <= max; i++ {
    image[i] = env.Val ("GOSRC") + "/µu/phil/pics/" + name[i]
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
