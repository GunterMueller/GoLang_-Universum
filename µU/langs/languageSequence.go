package langs

// (c) Christian Maurer   v. 170810 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/errh"
  "µU/font"
  "µU/pbox"
  "µU/bnat"
  "µU/enum"
//  "µU/subject"
)
const (
  min = 2
  max = 4
)
const ( // Format
  Short = iota
  Long
  NFormats
)
type
  languageSequence struct {
              lang [max]enum.Enumerator
          from, to [max]bnat.Natural
            cF, cB col.Colour
                   Format
                   }
var (
  bx = box.New()
  pbx = pbox.New()
  lLa, cLa, lFr, cFr, lTo, cTo [NFormats][max]uint
)

func init() {
  for n := uint(0); n < max; n++ {
    lLa[Short][n] = 0; cLa[Short][n] = 11 * n
    lFr[Short][n] = 0; cFr[Short][n] = 11 * n + 3
    lTo[Short][n] = 0; cTo[Short][n] = 11 * n + 6
    lLa[Long][n] = n; cLa[Long][n] = 0
    lFr[Long][n] = n; cFr[Long][n] = 23
    lTo[Long][n] = n; cTo[Long][n] = 37
  }
}

func new_() LanguageSequence {
  x := new (languageSequence)
  for n := uint(0); n < max; n++ {
    x.lang[n] = enum.New (enum.Subject)
    x.from[n], x.to[n] = bnat.New (2), bnat.New (2)
  }
  x.cF, x.cB = scr.StartCols()
  return x
}

func (x *languageSequence) imp (Y Any) *languageSequence {
  y, ok := Y.(*languageSequence)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *languageSequence) Empty() bool {
  for n := uint(0); n < max; n++ {
    if ! x.lang[n].Empty() { return false }
  }
  return true
}

func (x *languageSequence) Clr() {
  for n := uint(0); n < max; n++ {
    x.lang[n].Clr()
    x.from[n].Clr()
    x.to[n].Clr()
  }
}

func (x *languageSequence) Eq (Y Any) bool {
  y := x.imp (Y)
  for n := uint(0); n < max; n++ {
    if ! x.lang[n].Eq (y.lang[n]) {
      return false
    }
    if ! x.from[n].Eq (y.from[n]) || ! x.to[n].Eq (y.to[n]) {
      return false
    }
  }
  return true
}

func (x *languageSequence) Less (Y Any) bool {
  return false
}

func (x *languageSequence) Copy (Y Any) {
  y := x.imp (Y)
  for n := uint(0); n < max; n++ {
    x.lang[n].Copy (y.lang[n])
    x.from[n].Copy (y.from[n])
    x.to[n].Copy (y.to[n])
  }
}

func (x *languageSequence) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *languageSequence) Num (lg []enum.Enumerator, from, to[]uint) uint {
  n := uint(0)
  for {
   if x.lang[n].Empty() {
      break
    } else if n < max - 1 {
      n ++
    } else {
      break
    }
  }
  if n == 0 { return 0 }
  lg = make ([]enum.Enumerator, n)
  from, to = make ([]uint, n), make ([]uint, n)
  for i := uint(0); i < n; i++ {
    lg[i].Copy (x.lang[i])
    from[i], to[i] = x.from[i].Val(), x.to[i].Val()
  }
  return n
}

func (x *languageSequence) Colours (f, b col.Colour) {
  for n := uint(0); n < max; n++ {
    x.lang[n].Colours (f, b)
    x.from[n].Colours (f, b)
    x.to[n].Colours (f, b)
  }
}

func (x *languageSequence) writeMask (l, c uint) {
  bx.ScrColours()
  switch x.Format { case Short:
/*        1         2         3         4
012345678901234567890123456789012345678901
_ (__-__)  e (__-__)  g (__-__)  _ (__-__) */
    bx.Wd (7)
    for n := uint(0); n < max; n++ {
      bx.Write ("(  -  )", l, c + cFr[x.Format][n] - 1)
    }
  case Long:
/*        1         2         3
0123456789012345678901234567890123456789
___________ von Klasse __ bis Klasse __ */
    bx.Wd (10)
    for n := uint(0); n < max; n++ {
      bx.Write ("von Klasse", l + lFr[x.Format][n], c + cFr[x.Format][n] - 11)
      bx.Write ("bis Klasse", l + lTo[x.Format][n], c + cTo[x.Format][n] - 11)
    }
  }
}

func (x *languageSequence) GetFormat() Format {
  return x.Format
}

func (x *languageSequence) SetFormat (f Format) {
  x.Format = f
  for n := uint(0); n < max; n++ {
    if x.Format == Short {
      x.lang[n].SetFormat (enum.Short)
    } else {
      x.lang[n].SetFormat (enum.Long)
    }
  }
}

func (x *languageSequence) Write (l, c uint) {
  x.writeMask (l, c)
  for n := uint(0); n < max; n++ {
    x.lang[n].Write (l + lLa[x.Format][n], c + cLa[x.Format][n])
    x.from[n].Write (l + lFr[x.Format][n], c + cFr[x.Format][n])
    x.to[n].Write (l + lTo[x.Format][n], c + cTo[x.Format][n])
  }
}

func (x *languageSequence) multiple (n *uint) bool {
  for i := uint(1); i < max; i++ {
    for k := i + 1; k < i; k++ {
      if ! x.lang[i].Empty() && x.lang[k].Eq (x.lang[i]) {
        *n = k
        return true
      }
    }
  }
  *n = 0
  return false
}

func (x *languageSequence) Edit (l, c uint) {
  x.Write (l, c)
  n := uint(0)
  loop_n:
  for {
    i := uint(0)
    loop_i: for {
      weg := false
      switch i { case 0: // lang
        for {
          for {
            x.lang[n].Edit (l + lLa[x.Format][n], c + cLa[x.Format][n])
            if x.lang[n].Ord() >= 2 && // Englisch
               x.lang[n].Ord() <= 12 { // Griechisch
              break
            } else {
              errh.Error0("keine Fremdsprache")
            }
          }
          if x.lang[n].Empty() {
            if n > 1 {
              weg = true
              break
            } else {
              errh.Error2 ("", n + 1, ". Fremdsprache fehlt", 0)
            }
          } else {
            break
          }
        }
      case 1: // from
        if weg {
          x.from[n].Clr()
          x.from[n].Write (l + lFr[x.Format][n], c + cFr[x.Format][n])
        } else {
          x.from[n].Edit (l + lFr[x.Format][n], c + cFr[x.Format][n])
        }
      case 2: // to
        if weg {
          x.to[n].Clr()
          x.to[n].Write (l + lTo[x.Format][n], c + cTo[x.Format][n])
          } else {
            for {
              x.to[n].Edit (l + lTo[x.Format][n], c + cTo[x.Format][n])
              if x.to[n].Empty() || x.to[n].Val() == 0 && x.to[n].Val() >= 12 {
                errh.Error ("geht nich", x.to[n].Val())
            } else {
              break
            }
          }
        }
      }
      if i < 2 {
        i ++
      } else {
        if ! x.to[n].Less (x.from[n]) {
          break loop_i
        } else {
          i = 1
        }
      }
    } // loop_i
    if n + 1 < max {
      n ++
    } else {
      k := uint(0)
      if x.multiple (&k) {
        errh.Error2 ("Die", k + 1, ". Fremdsprache kommt mehrfach vor", 0)
      } else {
        break loop_n
      }
    }
  }
}

func (x *languageSequence) printMask (l, c uint) {
  switch x.Format { case Short:
    for n := uint(0); n < max; n++ {
      pbx.Print ("(  -  )", l, c + cFr[x.Format][n] - 1)
    }
  case Long:
    for n := uint(0); n < max; n++ {
      pbx.Print ("von Klasse", l + lFr[x.Format][n], c + cFr[x.Format][n] - 11)
      pbx.Print ("bis Klasse", l + lTo[x.Format][n], c + cTo[x.Format][n] - 11)
    }
  }
}

func (x *languageSequence) SetFont (f font.Font) {
  for n := uint(0); n < max; n++ {
    x.lang[n].SetFont (f)
    x.from[n].SetFont (f)
    x.to[n].SetFont (f)
  }
}

func (x *languageSequence) Print (l, c uint) {
  x.printMask (l, c)
  for n := uint(0); n < max; n++ {
    x.lang[n].Print (l + lLa[x.Format][n], c + cLa[x.Format][n])
    x.from[n].Print (l + lFr[x.Format][n], c + cFr[x.Format][n])
    x.to[n].Print (l + lTo[x.Format][n], c + cTo[x.Format][n])
  }
}

func (x *languageSequence) Codelen() uint {
  return max * (x.lang[0].Codelen() + 1)
}

func (x *languageSequence) Encode()[]byte {
  b := make ([]byte, x.Codelen())
  i := uint(0)
  for n := uint(0); n < max; n++ {
    a := x.lang[n].Codelen()
    copy (b[i:i+a], x.lang[n].Encode())
    i += a
    c := byte (x.from[n].Val() + 16 * x.to[n].Val())
    copy (b[i:i+1], Encode (c))
    i ++
  }
  return b
}

func (x *languageSequence) Decode (b[]byte) {
  i := uint(0)
  for n := uint(0); n < max; n++ {
    a := x.lang[n].Codelen()
    x.lang[n].Decode (b[i:i+a])
    i += a
    c := uint(Decode (byte(0), b[i:i+1]).(byte))
    x.from[n].SetVal (c % 16)
    x.to[n].SetVal (c / 16)
    i ++
  }
}
