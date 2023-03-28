package world

// (c) Christian Maurer   v. 230311 - license see µU.go

import (
  . "µU/obj"
  "µU/kbd"
  "µU/mode"
  "µU/scr"
  "µU/errh"
  "µU/pseq"
  "life/species"
)
const
  y0 = 1
type
  world struct {
               string "name of the world"
     spec, old []species.Species
  line, column uint16
               }
var (
  nx, ny uint
  suffix string
  theMode = mode.PAL
  system species.System
)

func init() {
  nx, ny = mode.Wd (theMode) / 8 / 2, mode.Ht (theMode) / 16 - 2
}

func m() mode.Mode {
  return theMode
}

func sys (s species.System) {
  species.Sys (s)
  suffix = species.Suffix
  system = s
}

func new_() World {
  w := new(world)
  w.spec = make([]species.Species, ny * nx)
  w.old = make([]species.Species, ny * nx)
  for y := uint(0); y < ny; y++ {
    for x := uint(0); x < nx; x++ {
      w.spec[nx * y + x] = species.New()
      w.old[nx * y + x] = species.New()
    }
  }
  w.line, w.column = uint16(ny / 2), uint16(nx / 2)
  return w
}

func (x *world) imp (Y any) *world {
  y, ok := Y.(*world)
  if ! ok {
    TypeNotEqPanic(x, Y)
  }
  return y
}

func (w *world) Name (name string) {
  w.string = name
  file := pseq.New (New().String())
  file.Name (w.string + "." + suffix)
  if ! file.Empty() {
    w.Defined (file.Get().(string))
  }
  w.Write()
  file.Fin()
}

func (w *world) Fin() {
  s := w.String()
  file := pseq.New (s)
  file.Name (w.string + "." + suffix)
  file.Put (s)
  file.Fin()
}

func (w *world) Rename (name string) {
  file := pseq.New (w)
  file.Name (w.string + "." + suffix)
  w.string = name
  file.Rename (w.string + "." + suffix)
  file.Fin()
}

func (w *world) Eq (Y any) bool {
  v := w.imp(Y)
  for y := uint(0); y < ny; y++ {
    for x := uint(0); x < nx; x++ {
      if ! w.spec[nx * y + x].Eq (v.spec[nx * y + x]) {
        return false
      }
    }
  }
  return true
}

func (w *world) Copy (Y any) {
  v := w.imp(Y)
  for y := uint(0); y < ny; y++ {
    for x := uint(0); x < nx; x++ {
      w.spec[nx*y+x].Copy (v.spec[nx*y+x])
    }
  }
  w.line, w.column = v.line, v.column
}

func (w *world) Clone() any {
  y := New()
  y.Copy(w)
  return y
}

func (w *world) Write() {
  for y := uint(0); y < ny; y++ {
    for x := uint(0); x < nx; x++ {
      w.spec[nx*y+x].Write (y0 + uint(y), 2 * uint(x))
    }
  }
  if system == species.Life {
    errh.Hint ("weiter: Enter        Zelle einsetzen/entfernen: linker/rechter Mausknopf        Ende: Esc")
  } else {
    errh.Hint ("weiter: Enter  Fuchs/Hase/Pflanze: linker/rechter Mausknopf (Fuchs mit Umschalttaste)  Ende: Esc")
  }
}

func (w *world) modify() {
  for y := uint(0); y < ny; y++ {
    for x := uint(0); x < nx; x++ {
      w.old[nx * y + x].Copy (w.spec[nx * y + x])
    }
  }
  for y := uint(0); y < ny; y++ {
    for x := uint(0); x < nx; x++ {
      w.spec[nx*y+x].Modify (func (s species.Species) uint {
        var n uint
        if x + 1 < nx {
          if w.old[nx * y + x + 1].Eq (s) {
            n++
          }
        }
        if y > 0 {
          if w.old[nx * (y - 1) + x].Eq (s) {
            n++
          }
        }
        if x > 0 {
          if w.old[nx * y + x - 1].Eq (s) {
            n++
          }
        }
        if y + 1 < ny {
          if w.old[nx * (y + 1) + x].Eq (s) {
            n++
          }
        }
        if species.NNeighbours == 4 {
          return n
        }
        if y > 0 && x+1 < nx {
          if w.old[nx * (y - 1) + x + 1].Eq (s) {
            n++
          }
        }
        if y > 0 && x > 0 {
          if w.old[nx * (y - 1) + x - 1].Eq (s) {
            n++
          }
        }
        if y + 1 < ny && x > 0 {
          if w.old[nx * (y + 1) + x - 1].Eq (s) {
            n++
          }
        }
        if y + 1 < ny && x + 1 < nx {
          if w.old[nx * (y + 1) + x + 1].Eq (s) {
            n++
          }
        }
        return n
      })
      w.spec[nx*y+x].Write (y0 + uint(y), 2 * uint(x))
    }
  }
}

func (w *world) Edit() {
  s := w.spec[nx * uint(w.line) + uint(w.column)]
  sa := s
  y, x := scr.MousePos()
  ya, xa := y, x
  loop:
  for {
    s.Write (uint(y0 + w.line), uint(2 * w.column))
    ya, xa = y, x
    sa = s
    c, d := kbd.Command()
    scr.MousePointer(true)
    y, x = scr.MousePos()
    if y0 <= y && y < uint(ny) + y0 && x < 2 * uint(nx) {
      w.line, w.column = uint16(y - y0), uint16(x/2)
      s = w.spec[nx * uint(w.line) + uint(w.column)]
    }
    if ya >= y0 {
      sa.Write (ya, 2 * (xa / 2))
    }
    switch c {
    case kbd.Esc:
      break loop
    case kbd.Enter:
      w.modify()
      w.Write()
    case kbd.Here:
      if d == 0 {
        s.Set (1)
      } else {
        s.Set (2)
      }
    case kbd.There:
      s.Set (0)
    }
  }
  w.Fin()
}

func (w *world) String() string {
  s := ""
  for y := uint(0); y < ny; y++ {
    for x := uint(0); x < nx; x++ {
      s += w.spec[nx * y + x].String()
    }
    s += string(10)
  }
  return s
}

func (w *world) Defined (s string) bool {
  if uint(len(s)) != ny * (nx + 1) {
    return false
  }
  i := 0
  for y := uint(0); y < ny; y++ {
    for x := uint(0); x < nx; x++ {
      w.spec[nx * y + x].Defined (string(s[i]))
      i++
    }
    i++
  }
  w.line, w.column = uint16(ny / 2), uint16(nx)
  return true
}
