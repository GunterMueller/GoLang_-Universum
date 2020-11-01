package box

// (c) Christian Maurer   v. 201016 - license see µU.go

import (
  "strconv"
  "µU/z"
  "µU/ker"
  . "µU/shape"
  . "µU/str"
  "µU/kbd"
  "µU/col"
  "µU/scr"
)
const
  space = byte(' ')
type
  box struct {
             string
       width,
       start uint
      cF, cB col.Colour
overwritable,
   graphical,
 transparent,
   numerical,
 TRnumerical,
   usesMouse bool
       index uint
             kbd.Comm
       depth uint
             }
var
  edited bool = true

func new_() Box {
  x := new (box)
  x.width = 0 // scr.NColumns() // does not work, if no scr.New was called before
  x.cF, x.cB = scr.StartCols()
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

func (x *box) UseMouse() {
  x.usesMouse = true
}

func (x *box) ScrColours() {
  x.cF, x.cB = scr.ScrColF(), scr.ScrColB()
}

func (x *box) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
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
  if l >= nl { ker.Panic ("box.Write: l == " + strconv.Itoa (int(l)) + " >= NLines == " + strconv.Itoa (int(nl))) }
  if c >= nc { ker.Panic ("box.Write: c == " + strconv.Itoa (int(c)) + " >= NColumns == " + strconv.Itoa (int(nc))) }
//  Wd (&s, s)
  z.ToHellWithUTF8 (&s)
  n, b := uint(len (s)), x.width
  if c + b > nc { x.width = nc - c }
  if x.width == 0 { x.width = n }
  if x.width > n { x.width = n }
  if x.width < n { Norm (&s, x.width) }
//  Norm (&s, x.width)
  if x.numerical || x.TRnumerical { Move (&s, false) }
  scr.Lock()
  scr.Colours (x.cF, x.cB)
  if x.transparent { scr.Transparence (true) }
  scr.Write (s, l, c)
  if x.transparent { scr.Transparence (false) }
  scr.Unlock()
  x.width = b
}

func (B *box) WriteGr (s string, x, y int) {
  if uint(y) >= scr.Ht() { return }
  if uint(x) >= scr.Wd() - scr.Wd1() { return }
  z.ToHellWithUTF8 (&s)
  n, b := uint(len (s)), B.width
  if B.width == 0 { B.width = n }
  if uint(x) + B.width * scr.Wd1() > scr.Wd() {
    B.width = (scr.Wd() - uint(x)) / scr.Wd1()
  }
  if B.width > n { B.width = n }
  if B.width < n { Norm (&s, B.width) }
  if B.numerical || B.TRnumerical { Move (&s, false) }
  scr.Lock()
  scr.Colours (B.cF, B.cB)
  if B.transparent { scr.Transparence (true) }
  scr.WriteGr (s, x, y)
  if B.transparent { scr.Transparence (false) }
  scr.Unlock()
  B.width = b
}

func (x *box) Clr (l, c uint) {
  if x.width == 0 { return }
  scr.Lock()
  f, b := scr.Cols()
  scr.Colours (scr.ScrCols())
  scr.WriteGr (New (x.width), int(scr.Wd1() * c), int(scr.Ht1() * l))
  scr.Colours (f, b)
  scr.Unlock()
}

func (x *box) Start (c uint) {
  x.start = c
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
    switch b.depth { case 0:
      if b.index > 0 {
        b.index--
        Rem (s, b.index, 1)
        *s += " "
      }
    case 1:
      b.index = 0
      *s = New (b.width)
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
      if b.index < ProperLen (*s) {
        Rem (s, b.index, 1)
        *s += " "
      }
    case 1:
      if b.overwritable {
        b.index = 0
        *s = New (b.width)
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
  case kbd.Help, kbd.Search, kbd.Act, kbd.Cfg, kbd.Mark, kbd.Demark,
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
        InsSpace (s, b.index) // -> this operation
        *s = (*s)[:b.width]   // -> to str
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
  var char byte
  var cursorshape Shape
  b.graphical = imGraphikmodus
// if b.usesMouse { scr.SwitchMouseCursor (true) }
  Norm (s, b.width)
  b.overwritable = ! Empty (*s)
  b.index = 0
  b.write (*s, x, y)
  b.overwritable = ! Empty (*s)
  b.write (*s, x, y)
  if b.start > 0 && b.start < b.width {
    b.index = b.start
    b.start = 0
  } else {
    b.index = 0
  }
  cf, cb := scr.Cols() // Warp may destroy the colours
  for {
    if b.overwritable {
      cursorshape = Block
    } else {
      cursorshape = Understroke
    }
    if b.graphical {
      scr.WarpGr (x + scr.Wd1() * b.index, y, cursorshape)
    } else {
      scr.Warp (y / scr.Ht1(), x / scr.Wd1() + b.index, cursorshape)
    }
    for {
      char, b.Comm, b.depth = kbd.Read()
      if b.Comm < kbd.Nav { // kbd.Go {
        break
      }
    }
    edited = char != 0
    if b.graphical {
      scr.WarpGr (x + scr.Wd1() * b.index, y, Off)
    } else {
      scr.Warp (y / scr.Ht1(), x / scr.Wd1() + b.index, Off)
    }
    if b.Comm == kbd.None {
      if b.index == b.width {
        // see editNumber
      } else {
        if b.possible (s, x, y) {
          Replace1 (s, b.index, char)
          scr.Lock()
          scr.Colours (b.cF, b.cB)
          if b.graphical {
            scr.Write1Gr (char, int(x + scr.Wd1() * b.index), int(y))
          } else {
            scr.Write1 (char, y / scr.Ht1(), x / scr.Wd1() + b.index)
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
// if B.usesMouse { scr.SwitchMouseCursor (false) }
}

// Precondition: n > 0, len (S) >= n - 1.
// Returns true, if S contains a character != ' ' in a position < n.
func leftNotEmpty (s string, n uint) bool {
  if n == 0 || len (s) + 1 < int(n) { return false }
  for i := 0; i < int(n) - 1; i++ {
    if s[i] != ' ' { return true }
  }
  return false
}

type
  stati byte; const (
  start = iota
  bp // before '.'
  ap // after '.'
  ee // after 'E', i.e. in exponent
)
var
  status stati

func getStatus (s *string) {
  if _, ok := Pos (*s, 'E'); ok {
    status = ee
  } else if _, ok := Pos (*s, '.'); ok {
    status = ap
  } else if Empty (*s) {
    status = start
  } else {
    status = bp
  }
}

func (b *box) doneNumerical (s *string, x, y uint) bool {
  switch b.Comm {
  case kbd.Enter, kbd.Esc:
    return true
/*
  case kbd.Left:
    if b.depth == 0 {
      if b.index > 0 && leftNotEmpty (s, b.index) {
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
    }
    return true
  case kbd.Down, kbd.Up: 
    return |
  case kbd.Pos1:
    if b.depth == 0 {
      b.index = 0
      for {
        if b.index == b.width { break }
        if s [b.index] != space { break }
        b.index++
      }
    } else {
      return true
    }
  case kbd.End:
    if b.depth == 0 {
      b.index = b.width
    } else {
      return true
    }
*/
  case kbd.Back:
    switch b.depth {
    case 0:
      if b.overwritable {
        if b.index == 0 {
        } else {
          Rem (s, b.index - 1, 1)
          *s += " "
        }
      } else if b.index < b.width {
        Rem (s, b.index, 1)
        *s += " "
        b.index++
      } else if b.index == b.width {
        Rem (s, b.width - 1, 1)
        *s = " " + *s
      }
    case 1:
      *s = New (b.width)
      status = start
      b.index = b.width
    default:
      return true
    }
    getStatus (s)
    if b.index < b.width {
      b.write (*s, x, y)
    } else {
      i := b.index
      b.index = 0
      b.write (*s, x, y)
      b.index = i
    }
  case kbd.Del:
    switch b.depth { case 0:
      if b.overwritable {
        if b.index == 0 {
        } else {
          Rem (s, b.index - 1, 1)
          *s += " "
        }
      } else if b.index < b.width {
        Rem (s, b.index, 1)
        *s += " "
        b.index++
      } else if b.index == b.width {
        Rem (s, b.width - 1, 1)
        *s = " " + *s
      }
    case 1:
      *s = New (b.width)
      b.index = b.width
    default:
      return true
    }
    if b.index < b.width {
      b.write (*s, x, y)
    } else {
      i := b.index
      b.index = 0
      b.write (*s, x, y)
      b.index = i
    }
/*
  case kbd.Ins:
    if b.depth == 0 {
      if b.overwritable {
        b.overwritable = false
      } else if i < b.width {
        b.overwritable = true
      }
    } else {
      return true
    }
*/
  case kbd.Help, kbd.Search, kbd.Act, kbd.Cfg, kbd.Mark, kbd.Demark,
       kbd.Cut, kbd.Copy, kbd.Paste, kbd.Red, kbd.Green, kbd.Blue:
    return true
  }
  return false
}

func (b *box) possibleNumerical (s *string, x, y uint) bool {
  if b.index < b.width {
    panic ("uff") // return false
    if b.overwritable { return true }
    if (*s)[b.width - 1] == ' ' {
      // if ! overwritable, shift s one to the right and Write
      InsSpace (s, b.index)
      b.write (*s, x, y)
      return true
    }
  } else { // overwritable == false
    i := uint(0)
    for {
      if i + 2 == b.width {
        break
      }
      if (*s)[i] == '0' && (*s)[i + 1] == '0' {
        Replace1 (s, i, ' ')
      } else {
        break
      }
      i++
    }
    if (*s)[0] == ' ' {
      if b.width > 1 {
        Rem (s, 0, 1)
        *s += " "
      }
      return true
    }
  }
  return false
}

func (B *box) editNumber (imGraphikmodus bool, s *string, x, y uint) {
  var (
    char byte
    cursorshape Shape
    temp uint
    firstTime bool
  )
  B.graphical = imGraphikmodus
//  if B.usesMouse { scr.SwitchMouseCursor (true) }
  Norm (s, B.width)
  B.overwritable = ! Empty (*s)
  Move (s, false)
  B.index = 0
  B.write (*s, x, y)
  B.index = B.width
  if B.TRnumerical {
    firstTime = true
    edited = false
    // Zahl beim ersten Lesen eines Zeichens zurücksetzen, s.u.
  } else {
    edited = true
  }
  for {
    getStatus (s)
    if B.overwritable {
      cursorshape = Block
    } else {
      cursorshape = Understroke
    }
    if B.graphical {
      scr.WarpGr (x + scr.Wd1() * B.index, y, cursorshape)
    } else {
      scr.Warp (y / scr.Ht1(), x / scr.Wd1() + B.index, cursorshape) // Off
    }
    for {
      char, B.Comm, B.depth = kbd.Read()
      switch char { case 0: // Command
        break
      case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
        if B.TRnumerical {
          if firstTime {
            *s = New (B.width)
            status = start
            firstTime = false
            edited = true
          }
        }
        if status == start {
          status = bp
          break
        } else if status == ee {
          if temp, ok := Pos (*s, 'E'); ok {
            if temp >= B.width - 3 { // not more than 2 digits after 'E'
              break
            }
          }
        } else {
          break
        }
      case '-':
        if B.TRnumerical {
          kbd.DepositCommand (kbd.None)
          kbd.DepositByte (char)
          return
        } else {
          if Empty (*s) || (*s)[B.width - 1] == 'E' {
            break
          }
        }
      case  '.', ',':
        if status == bp {
          status = ap
          break
        }
      case 'E':
        if B.numerical || B.TRnumerical {
          if status == ap && // noch Platz für zwei Zeichen
             (*s)[0] == space && (*s)[1] == space {
            status = ee
            if B.numerical {
              break
            } else {
              Rem (s, B.width - 2, 2)
              *s = *s + "E+"
              char = 0
              break
            }
          }
        }
      case 'v':
        char = 0
        if B.TRnumerical { // || B.numerical {
          if status == bp || status == ap {
            temp = 0
            for (*s)[temp] == space { temp++ }
            if (*s)[temp] == '-' {
              Replace1 (s, temp, '+')
              break
            } else if (*s)[temp] == '+' {
              Replace1 (s, temp, '-')
              break
            } else if temp > 0 {
              Replace1 (s, temp - 1, '-')
              break
            }
          } else if status == ee {
            if temp, ok := Pos (*s, 'E'); ok {
              if (*s)[temp + 1] == '-' {
                Replace1 (s, temp + 1, '+')
                break
              } else if (*s)[temp + 1] == '+' {
                Replace1 (s, temp + 1, '-')
                break
              }
            }
          }
        }
      default:
        if B.TRnumerical {
   // >>> Besser wäre dies nur für den Fall, dass 'Zeichen' ein Funktionszeichen aus dem Zahlen-Modul ist:
          kbd.DepositCommand (kbd.None)
          kbd.DepositByte (char)
          return
        }
      }
    }
    if B.graphical {
      scr.WarpGr (x + scr.Wd1() * B.index, y, Off)
    } else {
      scr.Warp (y / scr.Ht1(), x / scr.Wd1() + B.index, Off)
    }
    if B.Comm == kbd.None {
      if B.index == B.width {
        if B.overwritable {
          B.overwritable = false
        }
        if char == 0 { // change of sign or exponent
          temp = B.index
          B.index = 0
          B.write (*s, x, y)
          B.index = temp
        } else if B.possibleNumerical (s, x, y) {
          temp = B.index
          B.index = 0
          B.write (*s, x, y)
          B.index = temp
          Replace1 (s, B.index - 1, char)
          scr.Lock()
          scr.Colours (B.cF, B.cB)
          if B.graphical {
            scr.Write1Gr (char, int(x + scr.Wd1() * (B.index - 1)), int(y))
          } else {
            scr.Write1 (char, y / scr.Ht1(), x / scr.Wd1() + B.index - 1)
          }
          scr.Unlock()
        } else {
        }
      } else {
        // see editText
      }
    } else {
      if B.doneNumerical (s, x, y) {
        break
      }
    }
  }
// if B.usesMouse { scr.SwitchMouseCursor (false) }
}

/*
func isDigit (b byte) bool {
  return '0' <= b && b <= '9'
}

func (b *box) editNumber1 (imGraphikmodus bool, s *string, x, y uint) {
  for uint(len (*s)) < b.width { *s = " " + *s }; b.Write (*s, x, y); for *s != "" && (*s)[0] == ' ' { *s = (*s)[1:] }; if *s == " " { *s = "" }
  var char byte
  if b.graphical {
    scr.WarpGr (x + scr.Wd1() * b.width, y, scr.Understroke)
  } else {
    scr.Warp (y / scr.Ht1(), x / scr.Wd1() + b.width, scr.Understroke)
  }
  loop: for {
    l := uint(len (*s))
    char, b.Comm, b.depth = kbd.Read()
    switch b.Comm {
    case kbd.None:
      if isDigit (char) && l < b.width {
        *s += string(char)
      }
    case kbd.Esc:
      break loop
    case kbd.Enter:
      break loop
    case kbd.Back, kbd.Del:
      if l > 0 {
        *s = (*s)[:l-1]
      }
    }
    for uint(len (*s)) < b.width { *s = " " + *s }; b.Write (*s, x, y); for *s != "" && (*s)[0] == ' ' { *s = (*s)[1:] }; if *s == " " { *s = "" }
  }
  if b.graphical {
    scr.WarpGr (x + scr.Wd1() * b.width, y, Off)
  } else {
    scr.Warp (y / scr.Ht1(), x / scr.Wd1() + b.width, Off)
  }
}
*/

func (b *box) Edit (s *string, l, c uint) {
  if l >= scr.NLines() { return }
  if c >= scr.NColumns() { return }
  n, w := uint(len (*s)), b.width
  if c + w > scr.NColumns() { b.width = scr.NColumns() - c }
  if b.width == 0 { b.width = n }
//  if b.width > n { b.width = n }
  if b.width < n { Norm (s, b.width) }
  b.graphical = false
//  scr.WarpMouse (l, c)
  if b.numerical || b.TRnumerical {
    b.editNumber (false, s, scr.Wd1() * c, scr.Ht1() * l)
  } else {
    b.editText (false, s, scr.Wd1() * c, scr.Ht1() * l)
  }
  b.width = w
//  scr.MousePointer (false)
}

func (b *box) EditGr (s *string, x, y int) {
  if uint(y) >= scr.Ht() { return }
  if uint(x) >= scr.Wd() - scr.Wd1() { return }
  n, w := uint(len (*s)), b.width
  if uint(x) + b.width * scr.Wd1() > scr.Wd() {
    b.width = (scr.Wd() - uint(x)) / scr.Wd1()
  }
  if b.width == 0 { b.width = n }
  if b.width < n { Norm (s, b.width) }
//  if b.width > n { b.width = n }
  b.graphical = true
  if b.numerical || b.TRnumerical {
    b.editNumber (true, s, uint(x), uint(y))
  } else {
    b.editText (true, s, uint(x), uint(y))
  }
  b.width = w
}

func Edited() bool {
  return edited
}
