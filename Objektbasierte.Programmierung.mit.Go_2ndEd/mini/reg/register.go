package reg

// (c) Christian Maurer   v. 221213 - license see µU.go

import (
  "µU/env"
  "µU/kbd"
  "µU/str"
  "µU/col"
  "µU/box"
  "µU/N"
  "µU/errh"
  "µU/files"
)
type
  register struct {
                  string "name of register" // len == 1
                  uint "value of register"
                  }
var (
  bx = box.New()
  set [R]*register
  colReg = col.LightGreen()
)

func init() {
  bx.Colours (colReg, col.Black())
  for i := uint(0); i < R; i++ {
    set[i] = new_().(*register)
  }
  files.Cd (env.Gosrc() + "/mini")
}

func new_() Register {
  x := new(register)
  x.string = " "
  return x
}

func (x *register) Empty() bool {
  return x.string == " "
}

func (x *register) Clr() {
  x.string = " "
  x.uint = 0
}

func (x *register) Defined (s string) bool {
  n := str.ProperLen (s)
  if n == 0 || n > 1 {
    x.Clr()
    return false
  }
  if s == " " {
    x.Clr()
    return true
  }
  if "a" <= s && s <= "z" {
    x.string = s
//  x.uint unverändert
    return true
  }
  x.Clr()
  return false
}

func (x *register) index() uint {
  return uint(x.string[0] - 'a')
}

func (x *register) Val() uint {
  if x.string == " " { return 0 }
  return set[x.index()].uint
}

func (x *register) SetVal (n uint) {
  if x.string == " " { return }
  set[x.index()].uint = n
}

func (x *register) String() string {
  return x.string
}

func (x *register) write (l, c uint) {
  N.Colours (colReg, col.Black())
  N.Write (x.uint, l, c + 4)
}

func (x *register) Write (l, c uint) {
  if x.string == " " { return }
  x.write (l, c)
}

func (x *register) Edit (l, c uint) {
  if x.string == " " { return }
  N.SetWd (D)
  N.Edit (&x.uint, l, c + 4)
  x.uint = x.uint % M
}

func writeAll (l, c uint) {
  N.Colours (colReg, col.Black())
  for i := uint(0); i < R; i++ {
    N.Write (set[i].uint, l + i, c + 4)
  }
}

func editAll (l, c uint) {
  errh.Hint ("bewegen: Pfeiltasten     fertig: Umschalt- + Eingabetaste oder Fluchttaste (Esc)")
  N.SetWd (D)
  i := uint(0)
  loop:
  for {
    N.Edit (&set[i].uint, l + i, c + 4)
    switch k, d := kbd.LastCommand(); k {
    case kbd.Esc:
      break loop
    case kbd.Enter:
      if d == 0 {
        if i + 1 < R {
          i++
        } else {
          break loop
        }
      } else {
        break loop
      }
    case kbd.Down:
      if i + 1 < R {
        i++
      }
    case kbd.Up:
      if i > 0 {
        i--
      }
    case kbd.Pos1:
      i = 0
    case kbd.End:
      i = R - 1
    }
  }
  errh.DelHint()
}
