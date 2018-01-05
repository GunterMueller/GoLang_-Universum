package cr

// (c) Christian Maurer   v. 171230 - license see nU.go

import ("unsafe"; "math"; "nU/cs")

type status struct {
  max []uint // indiziert über die Prozessklassen
  number, class uint
}

type criticalResource struct {
  stat []status // indiziert über die Ressourcen
  nC, nR uint
  cs.CriticalSection
}

func max() uint {
  if unsafe.Sizeof(int(0)) == 32 {
    return math.MaxUint32
  }
  return math.MaxUint64
}

func new_(nc, nr uint) CriticalResource {
  x := new (criticalResource)
  x.nC, x.nR = nc, nr
  x.stat = make ([]status, x.nC)
  for r := uint(0); r < x.nR; r++ {
    x.stat[r].max = make ([]uint, x.nC)
    for c := uint(0); c < x.nC; c++ {
      x.stat[r].max[c] = max()
    }
  }
  c := func (i uint) bool {
        var b bool
        for r := uint(0); r < x.nR; r++ {
          b = b ||
              x.stat[r].number == 0 ||
              x.stat[r].class == i && x.stat[r].number < x.stat[r].max[i]
        }
        return b
      }
  e := func (i uint) uint {
        for r := uint(0); r < x.nR; r++ {
          if x.stat[r].number == 0 || x.stat[r].class == i {
            x.stat[r].class = i
            x.stat[r].number++
            return r
          }
        }
        panic("")
      }
  l := func (i uint) {
        for r := uint(0); r < x.nR; r++ {
          if x.stat[r].class == i && x.stat[r].number > 0 {
            x.stat[r].number--
          }
        }
      }
  x.CriticalSection = cs.New (x.nC, c, e, l)
  return x
}

func (x *criticalResource) Limit (m [][]uint) {
  for c := uint(0); c < x.nC; c++ {
    for r := uint(0); r < x.nR; r++ {
      x.stat[r].max[c] = m[c][r]
    }
  }
}

func (x *criticalResource) Enter (i uint) uint {
  return x.CriticalSection.Enter (i)
}

func (x *criticalResource) Leave (i uint) {
  x.CriticalSection.Leave (i)
}
