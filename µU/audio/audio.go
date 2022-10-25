package audio

// (c) Christian Maurer   v. 221021 - license see µU.go

import (
  . "µU/obj"
  "µU/str"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/errh"
  "µU/text"
  "µU/enum"
)
const (
  len0 = 30
  len1 = 60
  lenf = 13
  lenm =  2
  sep = ';'
  seps = ";"
)
const (
  fieldIndex = iota
  composerIndex
  mediumIndex
  workIndex
  nIndices
)
type
  audio struct {
         field,
        medium enum.Enum
      composer,
          work,
     composer1,
         work1,
     orchestra,
     conductor,
       soloist text.Text
               }
var
  actIndex = fieldIndex

func (x *audio) imp (Y any) *audio {
  y, ok := Y.(*audio)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func new_() Audio {
  x := new (audio)
  x.field = enum.Newk (lenf, 1)
  x.field.Set ("alte Musik", "Barock", "Klassik", "Romantik", "neue Musik", "Pop/Beat/Rock",
               "Folklore", "Jazz", "Italien", "Weihnachten", "Kinder")
  x.field.Setk ("a", "b", "k", "r", "n", "p", "f", "j", "i", "w", "c")
  x.medium = enum.Newk (lenm, 1)
  x.medium.Set ("LP", "SP", "CD")
  x.medium.Setk ("l", "s", "c")
  x.composer = text.New (len0)
  x.work = text.New (len1)
  x.composer1 = text.New (len0)
  x.work1 = text.New (len1)
  x.orchestra = text.New (len1)
  x.conductor = text.New (len0)
  x.soloist = text.New (len0)
  x.composer.Colours (col.Yellow(), col.Red())
  x.composer1.Colours (col.Yellow(), col.Red())
  x.work.Colours (col.LightWhite(), col.DarkGreen())
  x.work1.Colours (col.LightWhite(), col.DarkGreen())
  x.orchestra.Colours (col.LightWhite(), col.DarkGray())
  x.conductor.Colours (col.LightWhite(), col.DarkGray())
  x.soloist.Colours (col.Yellow(), col.Red())
  return x
}

func (x *audio) Empty() bool {
  return x.field.Empty() && x.medium.Empty() &&
         x.composer.Empty() && x.work.Empty() &&
         x.composer1.Empty() && x.work1.Empty() &&
         x.orchestra.Empty() && x.conductor.Empty() && x.soloist.Empty()
}

func (x *audio) Clr() {
  x.field.Clr()
  x.medium.Clr()
  x.composer.Clr(); x.work.Clr()
  x.composer1.Clr(); x.work1.Clr()
  x.orchestra.Clr()
  x.conductor.Clr()
  x.soloist.Clr()
}

func (x *audio) Eq (Y any) bool {
  y := x.imp(Y)
  return x.field.Eq (y.field) &&
         x.medium.Eq (y.medium) &&
         x.composer.Eq (y.composer) && x.work.Eq (y.work) &&
         x.composer1.Eq (y.composer1) && x.work1.Eq (y.work1) &&
         x.conductor.Eq (y.conductor) &&
         x.orchestra.Eq (y.orchestra) &&
         x.soloist.Eq (y.soloist)
}

func (x *audio) Copy (Y any) {
  y := x.imp(Y)
  x.field.Copy (y.field)
  x.medium.Copy (y.medium)
  x.composer.Copy (y.composer)
  x.work.Copy (y.work)
  x.composer1.Copy (y.composer1)
  x.work1.Copy (y.work1)
  x.conductor.Copy (y.conductor)
  x.orchestra.Copy (y.orchestra)
  x.soloist.Copy (y.soloist)
}

func (x *audio) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *audio) Less (Y any) bool {
  y := x.imp(Y)
  switch actIndex {
  case fieldIndex:
    if ! x.field.Eq (y.field) {
      return x.field.Less (y.field)
    }
    if ! x.medium.Eq (y.medium) {
      return x.medium.Less (y.medium)
    }
    if ! x.composer.Eq (y.composer) {
      return x.composer.Less (y.composer)
    }
    if ! x.work.Eq (y.work) {
      return x.work.Less (y.work)
    }
  case mediumIndex:
    if ! x.medium.Eq (y.medium) {
      return x.medium.Less (y.medium)
    }
    if ! x.field.Eq (y.field) {
      return x.field.Less (y.field)
    }
    if ! x.composer.Eq (y.composer) {
       return x.composer.Less (y.composer)
    }
    if ! x.work.Eq (y.work) {
      return x.work.Less (y.work)
    }
  case composerIndex:
    if ! x.composer.Eq (y.composer) {
       return x.composer.Less (y.composer)
    }
    if ! x.field.Eq (y.field) {
      return x.field.Less (y.field)
    }
    if ! x.medium.Eq (y.medium) {
      return x.medium.Less (y.medium)
    }
    if ! x.work.Eq (y.work) {
      return x.work.Less (y.work)
    }
  case workIndex:
    if ! x.work.Eq (y.work) {
      return x.work.Less (y.work)
    }
    if ! x.field.Eq (y.field) {
      return x.field.Less (y.field)
    }
    if ! x.medium.Eq (y.medium) {
      return x.medium.Less (y.medium)
    }
    if ! x.composer.Eq (y.composer) {
       return x.composer.Less (y.composer)
    }
  }
  return false
}

func (x *audio) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *audio) String() string {
  s := x.field.String()
  str.OffSpc1 (&s)
  s += seps
  t := x.medium.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.composer.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.work.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.composer1.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.work1.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.orchestra.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.conductor.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.soloist .String()
  str.OffSpc1 (&t)
  s += t + seps
  return s
}

func (x *audio) Defined (s string) bool {
  ss, n := str.SplitByte (s, sep)
  if n != 9 { errh.Error (s, n) }
  if ! x.field.Defined (ss[0]) {
    errh.Error0 (ss[0])
    return false
  }
  if ! x.medium.Defined (ss[1]) {
    errh.Error0 (ss[1])
    return false
  }
  if ! x.composer.Defined (ss[2]) {
    errh.Error0 (ss[2])
    return false
  }
  if ! x.work.Defined (ss[3]) {
    errh.Error0 (ss[3])
    return false
  }
  if ! x.composer1.Defined (ss[4]) {
    errh.Error0 (ss[4])
    return false
  }
  if ! x.work1.Defined (ss[5]) {
    errh.Error0 (ss[5])
    return false
  }
  if ! x.orchestra.Defined (ss[6]) {
    errh.Error0 (ss[6])
    return false
  }
  if ! x.conductor.Defined (ss[7]) {
    errh.Error0 (ss[7])
    return false
  }
  if ! x.soloist.Defined (ss[8]) {
    errh.Error0 (ss[8])
    return false
  }
  return true
}

func (x *audio) Sub (Y any) bool {
  y := x.imp(Y)
  if ! x.field.Empty() {
    return x.field.Eq (y.field)
  }
  if ! x.medium.Empty() {
    return x.medium.Eq (y.medium)
  }
  if ! x.composer1.Empty() {
    return x.composer1.Sub (y.composer1)
  }
  if ! x.work.Empty() {
    return x.work.Sub (y.soloist)
  }
  return false
}

const (
  lg  =  1; cg  = 10
  lm  =  1; cm  = 49
  lc  =  3; cc  = 10
  lw  =  5; cw  = 10
  lc1 =  7; cc1 = 10
  lw1 =  9; cw1 = 10
  lo  = 11; co  = 10
  ld  = 13; cd  = 10
  ls  = 13; cs  = 49
)

/*        1         2         3         4         5         6         7
01234567890123456789012345678901234567890123456789012345678901234567890123456789

   Gebiet ___________                     Medium __

Komponist ______________________________

     Werk ____________________________________________________________

Komponist ______________________________

     Werk ____________________________________________________________

Orchester ____________________________________________________________________

 Dirigent ______________________________  Solist ______________________________ */

func writeMask (l, c uint) {
  scr.Colours (col.LightGray(), col.Black())
  scr.Write ("Gebiet",    l + lg,  c +  3)
  scr.Write ("medium",    l + lm,  c + 42)
  scr.Write ("Komponist", l + lc,  c +  0)
  scr.Write ("Werk",      l + lw,  c +  5)
  scr.Write ("Komponist", l + lc1, c +  0)
  scr.Write ("Werk",      l + lw1, c +  5)
  scr.Write ("Orchester", l + lo,  c +  0)
  scr.Write ("Dirigent",  l + ld,  c +  1)
  scr.Write ("Solist",    l + ls,  c + 42)
}

var
  maskWritten = false

func (x *audio) Write (l, c uint) {
  if ! maskWritten {
    writeMask (l, c)
    maskWritten = true
  }
  x.field.Write (l + lg, c + cg)
  x.medium.Write (l + lm, c + cm)
  x.composer.Write (l + lc, c + cc)
  x.work.Write (l + lw, c + cw)
  x.composer1.Write (l + lc1, c + cc1)
  x.work1.Write (l + lw1, c + cw1)
  x.orchestra.Write (l + lo, c + co)
  x.conductor.Write (l + ld, c + cd)
  x.soloist.Write (l + ls, c + cs)
}

func containsSep (t text.Text) bool {
  _, c := str.Pos (t.String(), sep)
  return c
}

func edit (t text.Text, s string, l, c uint) {
  for {
    t.Edit (l, c)
    if containsSep (t) {
      errh.Error0 (s + " darf kein " + seps + " enthalten")
    } else {
      break
    }
  }
}

func (x *audio) Edit (l, c uint) {
  x.Write (l, c)
  i := 0
  loop:
  for {
    x.Write (l, c)
    switch i {
    case 0:
      x.field.Edit (l + lg, c + cg)
    case 1:
      x.medium.Edit (l + lm, c + cm)
    case 2:
      edit (x.composer, "Komponist", l + lc, c + cc)
      if ! x.composer.Empty() {
        if co, _ := kbd.LastCommand(); co == kbd.Tab {
          for i := 0; i < len(k); i++ {
            if x.composer.Sub0 (composer[i]) {
              x.composer.Copy (composer[i])
              x.composer.Write (l + lc, c + cc)
              break
            }
          }
        }
      }
    case 3:
      edit (x.work, "Werk", l + lw, c + cw)
      s := x.work.String()
      if str.ProperLen (s) == 1 {
        switch s[0] {
        case 'K':
          x.work.Defined ("Klavierkonzert")
        case 'V':
          x.work.Defined ("Violinkonzert")
        }
        x.work.Write (l + lw, c + cw)
      }
    case 4:
      edit (x.composer, "Komponist", l + lc1, c + cc1)
      if ! x.composer1.Empty() {
        if co, _ := kbd.LastCommand(); co == kbd.Tab {
          for i := 0; i < len(k); i++ {
            if x.composer1.Sub0 (composer[i]) {
              x.composer1.Copy (composer[i])
              x.composer1.Write (l + lc1, c + cc1)
              break
            }
          }
        }
      }
    case 5:
      edit (x.work, "Werk", l + lw1, c + cw1)
      if ! x.work1.Empty() {
        s := x.work1.String()
        if str.ProperLen (s) == 1 {
          switch s[0] {
          case 'K':
            x.work1.Defined ("Klavierkonzert")
          case 'V':
            x.work1.Defined ("Violinkonzert")
          }
        }
      }
      x.work1.Write (l + lw1, c + cw1)
    case 6:
      edit (x.orchestra, "Orchester", l + lo, c + co)
    case 7:
      x.conductor.Edit (l + ld, c + cd)
      if ! x.conductor.Empty() {
        for i := 0; i < len(con); i++ {
          if x.conductor.Sub0 (conductor[i]) {
            x.conductor.Copy (conductor[i])
            x.conductor.Write (l + ld, c + cd)
            break
          }
        }
      }
      edit (x.conductor, "Dirigent", l + ld, c + cd)
    case 8:
      edit (x.soloist, "Soloist", l + ls, c + cs)
    }
    switch k, _ := kbd.LastCommand(); k {
    case kbd.Esc:
      break loop
    case kbd.Enter, kbd.Down:
      if i < 11 {
        i++
      } else {
        break loop
      }
    case kbd.Back, kbd.Up:
			if i > 0 {
        i--
      }
    }
  }
}

var
  lastfield = enum.New (lenm)
//  lastfield = enum.Newk (lenm, 1)

func texdef() string {
  return "\\def\\n{\\newline} \\def\\p{\\par\\smallpagebreak}\n"
}

func (x *audio) TeX() string {
  s := ""
  if ! x.field.Eq (lastfield) {
    lastfield.Copy (x.field)
    s += "\\bigskip\\line{\\bf\\hfil " + x.field.(TeXer).TeX() + "\\hfil}\\medskip\\nopagebreak\n"
  }
  s += "\\x " + x.medium.TeX() + " "
  if x.composer.Empty() {
    s += "\\leavevmode"
  } else {
    s += "{\\bi " + x.composer.TeX() + "}"
  }
  if ! x.work.Empty() {
    s += "\\n\n" + x.work.TeX() + ""
  }
  if ! x.composer1.Empty() {
    s += "\\n\n{\\bi " + x.composer1.TeX() + "}"
  }
  if ! x.work1.Empty() {
    s += "\\n\n" + x.work1.TeX() + ""
  }
  if ! x.orchestra.Empty() {
    s += "\\n\n" + x.orchestra.TeX()
    if ! x.conductor.Empty() {s += " (" + x.conductor.TeX() + ") "}
  }
  if ! x.soloist.Empty() {
    s += "\\n\n{\\bi " + x.soloist.TeX() + "}"
  }
  s += "\n\\p\n"
  return s
}

func (x *audio) Codelen() uint {
  return x.field.Codelen() + x.medium.Codelen() +
         len0 + len1 + len0 + len1 +
         len1 + 2 * len0
}

func (x *audio) Encode() Stream {
  s := make(Stream, x.Codelen())
  i, a := uint(0), x.field.Codelen()
  copy (s[i:i+a], x.field.Encode())
  i += a
  a = x.medium.Codelen()
  copy (s[i:i+a], x.medium.Encode())
  i += a
  a = len0
  copy (s[i:i+a], x.composer.Encode())
  i += a
  a = len1
  copy (s[i:i+a], x.work.Encode())
  i += a
  a = len0
  copy (s[i:i+a], x.composer1.Encode())
  i += a
  a = len1
  copy (s[i:i+a], x.work1.Encode())
  i += a
  copy (s[i:i+a], x.orchestra.Encode())
  i += a
  a = len0
  copy (s[i:i+a], x.conductor.Encode())
  i += a
  copy (s[i:i+a], x.soloist.Encode())
  return s
}

func (x *audio) Decode (s Stream) {
  i, a := uint(0), x.field.Codelen()
  x.field.Decode (s[i:i+a])
  i += a
  a = x.medium.Codelen()
  x.medium.Decode (s[i:i+a])
  i += a
  a = len0
  x.composer.Decode (s[i:i+a])
  i += a
  a = len1
  x.work.Decode (s[i:i+a])
  i += a
  a = len0
  x.composer1.Decode (s[i:i+a])
  i += a
  a = len1
  x.work1.Decode (s[i:i+a])
  i += a
  x.orchestra.Decode (s[i:i+a])
  i += a
  a = len0
  x.conductor.Decode (s[i:i+a])
  i += a
  x.soloist.Decode (s[i:i+a])
}

func (x *audio) Rotate() {
  actIndex = (actIndex + 1) % nIndices
}

func (x *audio) Index() Func {
  return Id
}
