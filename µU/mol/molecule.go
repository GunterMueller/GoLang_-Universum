package mol

// (c) Christian Maurer   v. 230326 - license see µU.go

import (
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/errh" 
  "µU/N" 
  "µU/pseq"
  "µU/atom"
  "µU/stru"
)
type
  molecule struct {
                  uint // length of a
                a []atom.Atom
                  }
var (
  help = []string {"    Komponente erzeugen: rechte Maustaste",
//                   "    Schriftfarbe ändern: F10             ",
//                   "Hintergrundfarbe ändern: Umschalt-F10    ",
                   "                 fertig: Esc             "}
  index []uint
  nIndices uint
  actIndex uint
  file pseq.PersistentSequence
)

func new_() Molecule {
  x := new(molecule)
  x.a = make([]atom.Atom, 0)
  return x
}

func (x *molecule) imp (Y any) *molecule {
  y, ok := Y.(*molecule)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *molecule) Eq (Y any) bool {
  y := x.imp (Y)
  for i := uint(0); i < y.uint; i++ {
    if ! x.a[i].Eq (y.a[i]) {
      return false
    }
  }
  return true
}

func (x *molecule) Copy (Y any) {
  y := x.imp (Y)
  x.uint = y.uint
  x.a = make([]atom.Atom, x.uint)
  for i := uint(0); i < x.uint; i++ {
    x.a[i] = atom.New()
    x.a[i].Copy (y.a[i])
  }
}

func (x *molecule) Clone() any {
  y := new_().(*molecule)
  y.Copy (x)
  return y
}

func (x *molecule) Less (Y any) bool {
  y := x.imp (Y)
  if x.a[actIndex] != y.a[actIndex] { return x.a[actIndex].Less (y.a[actIndex]) }
  for i := uint(0); i < nIndices; i++ {
    if i != actIndex {
      if x.a[i] != y.a[i] {
        return x.a[i].Less (y.a[i])
      }
    }
  }
  return false
}

func (x *molecule) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *molecule) Sub (Y any) bool {
  i := actIndex
  return x.a[i].Sub (x.imp(Y).a[i])
}

func (x *molecule) Empty() bool {
  for i := uint(0); i < x.uint; i++ {
    if ! x.a[i].Empty() {
      return false
    }
  }
  return true
}

func (x *molecule) Clr() {
  for i := uint(0); i < x.uint; i++ {
    x.a[i].Clr()
  }
}

func (x *molecule) Codelen() uint {
  n := uint(1)
  for i := uint(0); i < x.uint; i++ {
    n += x.a[i].Codelen()
  }
  return n
}

func (x *molecule) Encode() Stream {
  s := make(Stream, x.Codelen())
  copy (s[0:1], Encode(uint8(x.uint)))
  i := uint(1)
  for j := uint(0); j < x.uint; j++ {
    a := x.a[j].Codelen()
    x.a[j].Write (0, 0)
    copy (s[i:i+a], x.a[j].Encode())
    i += a
  }
  return s
}

func (x *molecule) Decode (s Stream) {
  x.uint = uint(Decode(uint8(0), s[0:1]).(uint8))
  i := uint(1)
  for j := uint(0); j < x.uint; j++ {
    a := x.a[j].Codelen()
    x.a[j].Decode (s[i:i+a])
    i += a
  }
}

func (x *molecule) Write (l, c uint) {
  for i := uint(0); i < x.uint; i++ {
    x.a[i].Write (l, c)
  }
}

func (x *molecule) Edit (l, c uint) {
  x.Write (0, 0)
  i := uint(0)
  loop:
  for {
    x.a[i].Edit (l, c)
    cmd, d := kbd.LastCommand()
    switch cmd {
    case kbd.Esc:
      break loop
    case kbd.Enter:
      if d == 0 {
        if i + 1 < x.uint {
          i++
        } else {
          break loop
        }
      } else {
        break loop
      }
    case kbd.Down:
      if i + 1 < x.uint {
        i++
      } else {
        i = 0
      }
    case kbd.Up:
      if i > 0 {
        i--
      } else {
        i = x.uint - 1
      }
    case kbd.Pos1:
      i = 0
    case kbd.End:
      i = x.uint - 1
    }
  }
  x.Write (0, 0)
}

func (x *molecule) Print() {
  for i := uint(0); i < x.uint; i++ {
    l, c := x.a[i].Pos()
    x.a[i].Print (l, c)
  }
}

func (x *molecule) defineIndices() {
/*/ example for len(x.a) = 6:
  If x.a[i].IsIndex() for the numbers marked by "*",
 
    0   1   2   3   4   5
  |---|---|---|---|---|---|
  |   | * | * |   | * |   |
  |---|---|---|---|---|---|
 
  then nIndices = 3, index[0] = 1, index[1] = 2 and index [2] = 4.
/*/
  nIndices = 0
  for i := uint(0); i < x.uint; i++ {
    if x.a[i].IsIndex() {
      index = append (index, uint(0))
      index[nIndices] = i
      nIndices++
    }
  }
  actIndex = index[0]
}

func (x *molecule) Construct (name string) {
  errh.Hint ("Molekülkonstruktion")
  i := uint(0)
  loop:
  for {
    x.Write (0, 0)
    cmd, _ := kbd.Command()
    scr.MousePointer (true)
    l, c := scr.MousePos()
    switch cmd {
    case kbd.Esc:
      if nIndices == 0 {
        errh.Error0 ("kein Index !")
      } else {
        break loop
      }
    case kbd.Here:
      x.uint++
      a := atom.New()
      x.a = append (x.a, a)
      x.a[i] = a
      x.a[i].Place (l, c)
      x.a[i].Select()
      if x.a[i].Typ() == atom.Enum {
        x.a[i].EnumName (name + "." + N.String(i))
        x.a[i].EnumSet (l, c)
      }
      x.a[i].EditIndex()
      x.a[i].Index (x.a[i].IsIndex())
      if x.a[i].IsIndex() { nIndices++ }
/*/
      errh.Hint ("Schriftfarbe auswählen")
      x.a[i].SelectColF()
/*/
      errh.Hint ("Hintergrundfarbe auswählen")
      x.a[i].SelectColB()
      errh.Hint ("Molekülkonstruktion")
      i++
    case kbd.Go:
      l0 := scr.NLines() - 1
      scr.Colours (col.FlashWhite(), col.Black())
      scr.Write ("       ", l0, 0)
      scr.WriteNat (l, l0, 0)
      scr.WriteNat (c, l0, 4)
    }
  }
  errh.DelHint()
  x.defineIndices()
// store the structure of x
  file = pseq.New (stru.New())
  file.Name (name + Suffix)
  for i := uint(0); i < x.uint; i++ {
    s := stru.New()
    w := x.a[i].Width()
    s.Define (x.a[i].Typ(), w)
    l, c := x.a[i].Pos()
    s.Place (l, c)
    f, b := x.a[i].Cols()
    s.Colours (f, b)
    s.Index (x.a[i].IsIndex())
    file.Seek (i)
    file.Put (s)
  }
}

// Returns the molecule constructed from the stored structure
func constructed (name string) Molecule {
  file = pseq.New (stru.New())
  filename := name + Suffix
  file.Name (filename)
  m := new_().(*molecule)
  num := file.Num()
  m.uint = num
  m.a = make([]atom.Atom, num)
  for i := uint(0); i < num; i++ {
    file.Seek (i)
    s := file.Get().(stru.Structure)
    m.a[i] = atom.New()
    m.a[i].Define (s.Typ(), s.Width())
    if m.a[i].Typ() == atom.Enum {
      filename = name + "." + N.String(i)
      m.a[i].EnumName (filename)
      m.a[i].EnumGet()
    }
    l, c := s.Pos()
    m.a[i].Place (l, c)
    f, b := s.Cols()
    m.a[i].Colours (f, b)
    m.a[i].Index (s.IsIndex())
  }
  m.defineIndices()
  return m
}

func (x *molecule) NumAtoms() uint {
  return uint(len(x.a))
}

func (x *molecule) Index() Func {
  return func (a any) any {
    x, ok := a.(*molecule)
    if ! ok { TypeNotEqPanic (x, a) }
    return actIndex
  }
}

func (x *molecule) Rotate() {
  actIndex = (actIndex + 1) % nIndices
}
