package masks

// (c) Christian Maurer   v. 220816 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/kbd"
  "µU/str"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/errh"
  "µU/pseq"
  "µU/masks/internal"
)
type
  masks struct {
             m []mask.Mask
               bool // named
          w, h uint
               }
var (
  help = []string{"             neue Maske: Enter         ",
                  "                 fertig: Esc           "}
  bx = box.New()
  file pseq.PersistentSequence
)

func new_() Masks {
  x := new(masks)
  x.m = make([]mask.Mask, 0)
  x.bool = false
  return x
}

func (x *masks) imp (Y any) *masks {
  y, ok := Y.(*masks)
  if ! ok { TypeNotEqPanic (y, Y) }
  return y
}

func (x *masks) Empty() bool {
  return file.Empty()
}

func (x *masks) Clr() {
  x.m = make([]mask.Mask, 0)
}

func (x *masks) Name (name string) {
  str.OffSpc (&name)
  file = pseq.New (mask.New())
  file.Name (name + ".msk")
  n := file.Num()
  x.m = make([]mask.Mask, n)
  x.bool = true
}

func (x *masks) check() {
  if ! x.bool { ker.Panic ("masks not named") }
}

func (x *masks) Write() {
  x.check()
  for i := uint(0); i < file.Num(); i++ {
    file.Seek (i)
    x.m[i] = file.Get().(mask.Mask)
    x.m[i].Write()
  }
}

func (x *masks) Edit() {
  x.check()
  errh.Hint ("Masken editieren")
  x.Write()
  for {
    k, _ := kbd.Command()
    errh.DelHint()
    l, c := scr.MousePos()
    switch k {
    case kbd.Esc:
      errh.DelHint()
      return
    case kbd.Here:
      i := uint(len(x.m))
      x.m = append (x.m, mask.New())
      x.m[i].Place (l, c)
      x.m[i].Edit()
      file.Seek (i)
      file.Put (x.m[i])
    case kbd.Help:
      errh.Help (help)
    case kbd.Go:
      l0 := scr.NLines() - 1
      scr.Colours (col.White(), col.Black())
      scr.Write ("       ", l0, 0)
      scr.WriteNat (l, l0, 0)
      scr.WriteNat (c, l0, 4)
    }
  }
}

func (x *masks) Print() {
  x.check()
  for i := uint(0); i < file.Num(); i++ {
    file.Seek (i)
    x.m[i] = file.Get().(mask.Mask)
    x.m[i].Print()
  }
}
