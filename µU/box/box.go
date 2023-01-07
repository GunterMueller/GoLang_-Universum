package box

// (c) Christian Maurer   v. 220804 - license see µU.go

import (
  "µU/char"
  "µU/ker"
  "µU/scr/shape"
  "µU/str"
  "µU/kbd"
  "µU/col"
  "µU/scr"
)
const
  space = byte(' ')
type
  box struct {
             string "content of the box"
       width,
       start uint
      cF, cB col.Colour
overwritable,
   graphical,
 transparent,
   numerical bool
       index uint
             kbd.Comm
       depth uint
             }
var
  edited bool

func new_() Box {
  x := new (box)
  x.width = 0
  x.cF, x.cB = col.StartCols()
  x.Comm = kbd.None
  return x
}

func (x *box) Wd (n uint) {
  x.width = n
}

func (x *box) SetNumerical() {
  x.numerical = true
}

func (x *box) Transparence (t bool) {
  x.transparent = t
}

func (x *box) ScrColours() {
  x.cF, x.cB = scr.ScrColF(), scr.ScrColB()
}

func (x *box) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *box) Cols() (col.Colour, col.Colour) {
  return x.cF, x.cB
}

func (x *box) ColourF (f col.Colour) {
  x.cF = f
}

func (x *box) ColourB (b col.Colour) {
  x.cB = b
}

func (b *box) String() string {
  return b.string
}

func (b *box) Defined (s string) bool {
  return uint(len(b.string)) < b.width
}

func (x *box) Write (s string, l, c uint) {
  nl, nc := scr.NLines(), scr.NColumns()
  if l >= nl { ker.Panic2 ("box.Write: l ==", l, ">= NLines ==", nl) }
  if c >= nc { ker.Panic2 ("box.Write: c ==", c, ">= NColumns ==", nc) }
  char.ToHellWithUTF8 (&s)
  n, w := uint(len (s)), x.width
  if c + w > nc { x.width = nc - c }
  if x.width == 0 { x.width = n }
  if x.width > n { x.width = n }
  if x.width < n { str.Norm (&s, x.width) }
  if x.numerical { str.Move (&s, false) }
  scr.Lock()
  if x.transparent { scr.Transparence (true) }
  scr.Colours (x.cF, x.cB)
  scr.Write (s, l, c)
  if x.transparent { scr.Transparence (false) }
  scr.Unlock()
  x.width = w
}

func (b *box) WriteGr (s string, x, y int) {
  if uint(y) >= scr.Ht() { return }
  if uint(x) >= scr.Wd() - scr.Wd1() { return }
  char.ToHellWithUTF8 (&s)
  n, w := uint(len (s)), b.width
  if b.width == 0 { b.width = n }
  if uint(x) + b.width * scr.Wd1() > scr.Wd() {
    b.width = (scr.Wd() - uint(x)) / scr.Wd1()
  }
  if b.width > n { b.width = n }
  if b.width < n { str.Norm (&s, b.width) }
  if b.numerical { str.Move (&s, false) }
  scr.Lock()
  if b.transparent { scr.Transparence (true) }
  scr.Colours (b.cF, b.cB)
  scr.WriteGr (s, x, y)
  if b.transparent { scr.Transparence (false) }
  scr.Unlock()
  b.width = w
}

func (x *box) Clr (l, c uint) {
  if x.width == 0 { return }
  scr.Lock()
  f, b := scr.Cols()
  scr.Colours (scr.ScrCols())
  scr.WriteGr (str.New (x.width), int(scr.Wd1() * c), int(scr.Ht1() * l))
  scr.Colours (f, b)
  scr.Unlock()
}

func (x *box) Start (c uint) {
  x.start = 0
  if x.start > 0 && x.start < x.width {
    x.start = c
  }
}

func (b *box) write (s string, x, y uint) {
  scr.Lock()
  scr.Colours (b.cF, b.cB)
  if b.transparent { scr.Transparence (true) }
  y1 := b.width
  if y1 > uint(len (s)) { y1 = uint(len (s)) }
  for x1 := b.index; x1 < y1; x1++ {
    if b.graphical {
      scr.Write1Gr (s[x1], int(x + scr.Wd1() * x1), int(y))
    } else {
      scr.Write1 (s[x1], y / scr.Ht1(), x / scr.Wd1() + x1)
    }
  }
  if b.transparent { scr.Transparence (false) }
  scr.Unlock()
}

func (b *box) done (s *string, x, y uint) bool {
  switch b.Comm {
  case kbd.Enter, kbd.Esc:
    return true
  case kbd.Back:
    switch b.depth {
    case 0:
      if b.index > 0 {
        b.index--
        str.Rem (s, b.index, 1)
        *s += " "
      }
    case 1:
      b.index = 0
      *s = str.New (b.width)
      if b.overwritable {
        b.overwritable = ! b.overwritable
      }
    default:
      return true
    }
    b.write (*s, x, y)
  case kbd.Left:
    if b.depth == 0 {
      if b.index > 0 {
        b.index--
      }
    } else {
      return true
    }
  case kbd.Right:
    if b.depth == 0 {
      if b.index < b.width - 1 {
        b.index++
      }
    } else {
      return true
    }
  case kbd.Up, kbd.Down:
    return true
  case kbd.PgLeft, kbd.PgRight, kbd.PgUp, kbd.PgDown:
    return true
  case kbd.Pos1:
    if b.depth == 0 {
      b.index = 0
    } else {
      return true
    }
  case kbd.End:
    if b.depth == 0 {
      b.index = b.width
      for {
        if b.index == 0 { break }
        if (*s)[b.index-1] == space {
          b.index--
        } else {
          break
        }
      }
    } else {
      return true
    }
  case kbd.Tab:
    return true
  case kbd.Del:
    switch b.depth { case 0:
      if b.index < str.ProperLen (*s) {
        str.Rem (s, b.index, 1)
        *s += " "
      }
    case 1:
      if b.overwritable {
        b.index = 0
        *s = str.New (b.width)
      } else {
        return true
      }
    default:
      return true
    }
    b.write (*s, x, y)
  case kbd.Ins:
    if b.depth == 0 {
      b.overwritable = ! b.overwritable
    } else {
      return true
    }
  case kbd.Help, kbd.Search, kbd.Act, kbd.Cfg, kbd.Mark, kbd.Unmark,
       kbd.Cut, kbd.Copy, kbd.Paste, kbd.Red, kbd.Green, kbd.Blue:
    return true
  case kbd.Print:
    return true
  }
  return false
}

func (b *box) possible (s *string, x, y uint) bool {
  if b.index < b.width {
    if b.overwritable { return true }
    if (*s)[b.width - 1] == space {
      if ! b.overwritable { // move s one to the right and write again
        str.InsSpace (s, b.index) // -> this operation
        *s = (*s)[:b.width]       // -> to str
        b.write (*s, x, y)
      }
      return true
    }
  } else { // b.index >= b.width
    // editNumber
  }
  return false
}

func (b *box) editText (imGraphikmodus bool, s *string, x, y uint) {
  var c byte
  var cursorshape shape.Shape
  b.graphical = imGraphikmodus
  str.Norm (s, b.width)
  b.overwritable = ! str.Empty (*s)
  b.index = 0
  b.write (*s, x, y)
  b.overwritable = ! str.Empty (*s)
  b.write (*s, x, y)
  if b.start > 0 && b.start < b.width {
    b.index = b.start
    b.start = 0
  }
  cf, cb := scr.Cols() // Warp may destroy the colours
  for {
    if b.overwritable {
      cursorshape = shape.Block
    } else {
      cursorshape = shape.Understroke
    }
    if b.graphical {
      scr.WarpGr (x + scr.Wd1() * b.index, y, cursorshape)
    } else {
      scr.Warp (y / scr.Ht1(), x / scr.Wd1() + b.index, cursorshape)
    }
    c, b.Comm, b.depth = kbd.Read()
    edited = c != 0
    if b.graphical {
      scr.WarpGr (x + scr.Wd1() * b.index, y, shape.Off)
    } else {
      scr.Warp (y / scr.Ht1(), x / scr.Wd1() + b.index, shape.Off)
    }
    if b.Comm == kbd.None {
      if b.index == b.width {
        // see editNumber
      } else {
        if b.possible (s, x, y) {
          str.Replace1 (s, b.index, c)
          scr.Lock()
          scr.Colours (b.cF, b.cB)
          if b.graphical {
            scr.Write1Gr (c, int(x + scr.Wd1() * b.index), int(y))
          } else {
            scr.Write1 (c, y / scr.Ht1(), x / scr.Wd1() + b.index)
          }
          scr.Unlock()
          b.index++
        }
      }
    } else {
      if b.done (s, x, y) {
        break
      }
    }
  }
  scr.Colours (cf, cb)
}

// Pre: n > 0, len(s) >= n - 1.
// Returns true, if s contains a character != ' ' in a position < n.
func leftNotEmpty (s string, n uint) bool {
  if n == 0 || len (s) + 1 < int(n) { return false }
  for i := 0; i < int(n) - 1; i++ {
    if s[i] != ' ' { return true }
  }
  return false
}

func (b *box) Edit (s *string, l, c uint) {
  if l >= scr.NLines() { return }
  if c >= scr.NColumns() { return }
  n, w := uint(len (*s)), b.width
  if c + w > scr.NColumns() { b.width = scr.NColumns() - c }
  if b.width == 0 { b.width = n }
  if b.width < n { str.Norm (s, b.width) }
  b.graphical = false
  if b.numerical {
    b.editNumber (false, s, scr.Wd1() * c, scr.Ht1() * l)
  } else {
    b.editText (false, s, scr.Wd1() * c, scr.Ht1() * l)
  }
  b.width = w
//  C, D := kbd.LastCommand(); println ("box Edit: C, D ==", C.String(), D) // XXX
}

func (b *box) EditGr (s *string, x, y int) {
  if uint(y) >= scr.Ht() { return }
  if uint(x) >= scr.Wd() - scr.Wd1() { return }
  n, w := uint(len (*s)), b.width
  if uint(x) + b.width * scr.Wd1() > scr.Wd() {
    b.width = (scr.Wd() - uint(x)) / scr.Wd1()
  }
  if b.width == 0 { b.width = n }
  if b.width < n { str.Norm (s, b.width) }
  b.graphical = true
  if b.numerical {
    b.editNumber (true, s, uint(x), uint(y))
  } else {
    b.editText (true, s, uint(x), uint(y))
  }
  b.width = w
}

func Edited() bool {
  return edited
}
