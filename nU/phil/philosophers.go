package phil

// (c) Christian Maurer   v. 171229 - license see nU.go

import ("sync"; "math"; "nU/col"; "nU/scr")

const (
  lock = iota; unlock
)
const (
  thinking = iota; hungry; starving; hasRightFork; hasLeftFork; dining; nStates
)
var (
  status [5]int
  colour = [5]col.Colour {col.Red(), col.Yellow(), col.Green(),
                          col.LightBlue(), col.Magenta()}
  name = [5]string {"Platon", "Sokrates", "Aristoteles", "Cicero", "Heraklit"}
  text = [nStates]string { "philosophiert   ", "ist hungrig     ", "ist sehr hungrig",
                           "ist sehr hungrig", "ist sehr hungrig", "ist am Speisen  " }
  l0, c0 uint
  ln, cn, lt, ct, lf, cf, cf1, lf1 [5]uint
  mutex sync.Mutex
)

func left (i uint) uint {
  return (i + 5 - 1) % 5
}

func right (i uint) uint {
  return (i + 1) % 5
}

func writePlate (i uint, b bool) {
  scr.Circle (lt[i], ct[i] / 2, 5)
  if b {
    scr.Circle (lt[i], ct[i] / 2, 3)
  }
}

func writeFork (i uint) {
  scr.Line (lf[i], cf[i], lf1[i], cf1[i])
}

func write (i uint) {
  scr.ColourF (colour[i])
  scr.Write (text[status[i]], ln[i] + 1, cn[i])
  if status[i] == thinking {
    scr.ColourF (col.Black())
    writePlate (i, true)
    scr.ColourF (col.Gray())
    writeFork (left (i))
    writeFork (i)
    return
  }
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

func Start() {
  mutex.Lock(); defer mutex.Unlock()
  const pi2 = 2 * math.Pi
  c0, l0 = uint(scr.NColumns()) / 2, uint(scr.NLines()) / 2
  const f = 0.5
  var rr uint
  for i := uint(0); i < 5; i++ {
    w, r := float64(i) / float64(5), 0.7 * float64(l0)
    rr = uint(r)
    lt[i] = uint(math.Trunc (-r * 0.7 * math.Sin (pi2 * w) + f)) + l0
    ct[i] = uint(math.Trunc (2 * r * 0.7 * math.Cos (pi2 * w) + f)) + c0
    ln[i] = uint(math.Trunc (-r * 1.2 * math.Sin (pi2 * w) + f)) + l0
    cn[i] = uint(math.Trunc (2 * r * 1.2 * math.Cos (pi2 * w) + f)) + c0 - 5
    scr.ColourF (colour[i])
    scr.Write (name[i], ln[i], cn[i])
    w += 0.1
    lf[i] = uint(math.Trunc (-r * math.Sin (pi2 * w) + f)) + l0
    cf[i] = uint(math.Trunc (2 * r * math.Cos (pi2 * w) + f)) + c0
    r /= 2
    lf1[i] = uint(math.Trunc (-r * math.Sin (pi2 * w) + f)) + l0
    cf1[i] = uint(math.Trunc (2 * r * math.Cos (pi2 * w) + f)) + c0
    write (i)
  }
  scr.ColourF (col.White())
  scr.Circle (l0, c0 / 2, rr)
}

func changeStatus (i uint, s int) {
  mutex.Lock(); defer mutex.Unlock()
  status[i] = s
  write (i)
  scr.Warp (0, 10)
}
