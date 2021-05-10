package audio

// (c) Christian Maurer   v. 210510 - license see µU.go

import (
  . "µU/obj"
  "µU/str"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/text"
  "µU/audio/gebiet"
  "µU/audio/medium"
)
const (
  len0 = 30
  len1 = 60
)
type
  ordnung int; const (
  nachGebiet = iota
  nachMedium
  nachKomponist
  nOrdnungen
)
type
  audio struct {
               gebiet.Gebiet
               medium.Medium
     komponist,
          werk,
    komponist1,
         werk1,
     orchester,
      dirigent,
        solist text.Text
               }
var
  aktuelleOrdnung = nachGebiet

func (x *audio) imp (Y Any) *audio {
  y, ok := Y.(*audio)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func new_() Audio {
  x := new (audio)
  x.Gebiet = gebiet.New()
  x.Medium = medium.New()
  x.komponist = text.New (len0)
  x.werk = text.New (len1)
  x.komponist1 = text.New (len0)
  x.werk1 = text.New (len1)
  x.orchester = text.New (len1)
  x.dirigent = text.New (len0)
  x.solist = text.New (len0)
  x.komponist.Colours (col.Yellow(), col.Red())
  x.komponist1.Colours (col.Yellow(), col.Red())
  x.werk.Colours (col.LightWhite(), col.DarkGreen())
  x.werk1.Colours (col.LightWhite(), col.DarkGreen())
  x.orchester.Colours (col.LightWhite(), col.DarkGray())
  x.dirigent.Colours (col.LightWhite(), col.DarkGray())
  x.solist.Colours (col.Yellow(), col.Red())
  return x
}

func (x *audio) Empty() bool {
  return x.Gebiet.Empty() && x.Medium.Empty() &&
         x.komponist.Empty() && x.werk.Empty() &&
         x.komponist1.Empty() && x.werk1.Empty() &&
         x.orchester.Empty() && x.dirigent.Empty() && x.solist.Empty()
}

func (x *audio) Clr() {
  x.Gebiet.Clr()
  x.Medium.Clr()
  x.komponist.Clr()
  x.werk.Clr()
  x.komponist1.Clr()
  x.werk1.Clr()
  x.orchester.Clr()
  x.dirigent.Clr()
  x.solist.Clr()
}

func (x *audio) Eq (Y Any) bool {
  y := x.imp(Y)
  return x.Gebiet.Eq (y.Gebiet) &&
         x.Medium.Eq (y.Medium) &&
         x.komponist.Eq (y.komponist) && x.werk.Eq (y.werk) &&
         x.komponist1.Eq (y.komponist1) && x.werk1.Eq (y.werk1) &&
         x.dirigent.Eq (y.dirigent) &&
         x.orchester.Eq (y.orchester) &&
         x.solist.Eq (y.solist)
}

func (x *audio) Copy (Y Any) {
  y := x.imp(Y)
  x.Gebiet = y.Gebiet
  x.Medium = y.Medium
  x.komponist.Copy (y.komponist)
  x.werk.Copy (y.werk)
  x.komponist1.Copy (y.komponist1)
  x.werk1.Copy (y.werk1)
  x.dirigent.Copy (y.dirigent)
  x.orchester.Copy (y.orchester)
  x.solist.Copy (y.solist)
}

func (x *audio) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *audio) Less (Y Any) bool {
  y := x.imp(Y)
  switch aktuelleOrdnung {
  case nachGebiet:
    if x.Gebiet.Eq (y.Gebiet) {
      if x.komponist.Eq (y.komponist) {
        if x.werk.Eq (y.werk) {
          return x.Medium.Less (y.Medium)
        }
        return x.werk.Less (y.werk)
      }
      return x.komponist.Less (y.komponist)
    }
    return x.Gebiet.Less (y.Gebiet)
  case nachMedium:
    if x.Medium.Eq (y.Medium) {
      if x.Gebiet.Eq (y.Gebiet) {
        if x.komponist.Eq (y.komponist) {
          return x.werk.Less (y.werk)
        }
        return x.Gebiet.Less (y.Gebiet)
      }
      return x.komponist.Less (y.komponist)
    }
    return x.Medium.Less (y.Medium)
  case nachKomponist:
    if x.komponist.Eq (y.komponist) {
      if x.werk.Eq (y.werk) {
        return x.Medium.Less (y.Medium)
      }
      return x.werk.Less (y.werk)
    }
    return x.komponist.Less (y.komponist)
  }
  return false
}

const (
  lg  =  1; cg  = 10
  lm  =  1; cm  = 49
  lk  =  3; ck  = 10
  lw  =  5; cw  = 10
  lk1 =  7; ck1 = 10
  lw1 =  9; cw1 = 10
  lo  = 11; co  = 10
  ld  = 13; cd  = 10
  ls  = 13; cs  = 49
)
/*/       1         2         3         4         5         6         7
01234567890123456789012345678901234567890123456789012345678901234567890123456789

   Gebiet ________                        Medium ___

Komponist ______________________________

     Werk ____________________________________________________________

Komponist ______________________________

     Werk ____________________________________________________________

Orchester ____________________________________________________________________

 Dirigent ______________________________  Solist ______________________________

/*/
func writeMask() {
  scr.Colours (col.LightGray(), col.Black())
  scr.Write ("Gebiet",    lg,  3)
  scr.Write ("Medium",    lm, 42)
  scr.Write ("Komponist", lk,  0)
  scr.Write ("Werk",      lw,  5)
  scr.Write ("Komponist", lk1, 0)
  scr.Write ("Werk",      lw1, 5)
  scr.Write ("Orchester", lo,  0)
  scr.Write ("Dirigent",  ld,  1)
  scr.Write ("Solist",    ls, 42)
}

var maskWritten = false

func (x *audio) Write (l, c uint) {
  if ! maskWritten {
    writeMask()
    maskWritten = true
  }
  x.Gebiet.Write (lg, cg)
  x.Medium.Write (lm, cm)
  x.komponist.Write (lk, ck)
  x.werk.Write (lw, cw)
  x.komponist1.Write (lk1, ck1)
  x.werk1.Write (lw1, cw1)
  x.orchester.Write (lo, co)
  x.dirigent.Write (ld, cd)
  x.solist.Write (ls, cs)
}

func (x *audio) Edit (l, c uint) {
  x.Write (l, c)
  i := 0
  loop:
  for {
    switch i {
    case 0:
      x.Gebiet.Edit (lg, cg)
    case 1:
      x.Medium.Edit (lm, cm)
    case 2:
      x.komponist.Edit (lk, ck)
      if ! x.komponist.Empty() {
        for i := 0; i < len(k); i++ {
          if x.komponist.Sub0 (komponist[i]) {
            x.komponist.Copy (komponist[i])
            break
          }
        }
      }
      x.komponist.Write (lk, ck)
    case 3:
      x.werk.Edit (lw, cw)
      s := x.werk.String()
      if str.ProperLen (s) == 1 {
        switch s[0] {
        case 'K':
          x.werk.Defined ("Klavierkonzert")
        case 'V':
          x.werk.Defined ("Violinkonzert")
        }
        x.werk.Write (lw, cw)
      }
    case 4:
      x.komponist1.Edit (lk1, ck1)
      if ! x.komponist1.Empty() {
        for i := 0; i < len(k); i++ {
          if x.komponist1.Sub0 (komponist[i]) {
            x.komponist1.Copy (komponist[i])
            x.komponist1.Write (lk1, ck1)
            break
          }
        }
      }
    case 5:
      x.werk1.Edit (lw1, cw1)
      if ! x.werk1.Empty() {
        s := x.werk1.String()
        if str.ProperLen (s) == 1 {
          switch s[0] {
          case 'K':
            x.werk1.Defined ("Klavierkonzert")
          case 'V':
            x.werk1.Defined ("Violinkonzert")
          }
        }
      }
      x.werk1.Write (lw1, cw1)
    case 6:
      x.orchester.Edit (lo, co)
    case 7:
      x.dirigent.Edit (ld, cd)
    case 8:
      x.solist.Edit (ls, cs)
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

var letztesGebiet = gebiet.New()

func (x *audio) TeX() string {
  s := ""
  if ! x.Gebiet.Eq (letztesGebiet) {
    letztesGebiet.Copy (x.Gebiet)
    s += "\\medskip{\\bf " + x.Gebiet.TeX() + "}\\medskip\n"
  }
  s += "\\x " + x.Medium.TeX() + " "
  s += "{\\bf " + x.komponist.TeX() + "}\\newline\n" + x.werk.TeX()
  if ! x.komponist1.Empty() {
    s += "\\newline\n{\\bf " + x.komponist1.TeX() + "}\\newline\n" + x.werk1.TeX()
  }
  if ! x.orchester.Empty() {
    s += "\\newline\n" + x.orchester.TeX()
    if ! x.dirigent.Empty() {s += " (" + x.dirigent.TeX() + ") "}
  }
  if ! x.solist.Empty() {s += "\\newline\n{\\bf " + x.solist.TeX() + "}"}
  s += "\n\\smallskip\n"
  return s
}

func (x *audio) Codelen() uint {
  return x.Gebiet.Codelen() + x.Medium.Codelen() +
         len0 + len1 + len0 + len1 +
         len1 + 2 * len0
}

func (x *audio) Encode() Stream {
  s := make(Stream, x.Codelen())
  i, a := uint(0), x.Gebiet.Codelen()
  copy (s[i:i+a], x.Gebiet.Encode())
  i += a
  a = x.Medium.Codelen()
  copy (s[i:i+a], x.Medium.Encode())
  i += a
  a = len0
  copy (s[i:i+a], x.komponist.Encode())
  i += a
  a = len1
  copy (s[i:i+a], x.werk.Encode())
  i += a
  a = len0
  copy (s[i:i+a], x.komponist1.Encode())
  i += a
  a = len1
  copy (s[i:i+a], x.werk1.Encode())
  i += a
  copy (s[i:i+a], x.orchester.Encode())
  i += a
  a = len0
  copy (s[i:i+a], x.dirigent.Encode())
  i += a
  copy (s[i:i+a], x.solist.Encode())
  return s
}

func (x *audio) Decode (s Stream) {
  i, a := uint(0), x.Gebiet.Codelen()
  x.Gebiet.Decode (s[i:i+a])
  i += a
  a = x.Medium.Codelen()
  x.Medium.Decode (s[i:i+a])
  i += a
  a = len0
  x.komponist.Decode (s[i:i+a])
  i += a
  a = len1
  x.werk.Decode (s[i:i+a])
  i += a
  a = len0
  x.komponist1.Decode (s[i:i+a])
  i += a
  a = len1
  x.werk1.Decode (s[i:i+a])
  i += a
  x.orchester.Decode (s[i:i+a])
  i += a
  a = len0
  x.dirigent.Decode (s[i:i+a])
  i += a
  x.solist.Decode (s[i:i+a])
}

func (x *audio) Rotate() {
  aktuelleOrdnung = (aktuelleOrdnung + 1) % nOrdnungen
}

func (x *audio) Index() Func {
  return func (a Any) Any {
    return a
  }
}
