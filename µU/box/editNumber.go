package box

import (
  . "µU/scr/shape"
  "µU/str"
  "µU/kbd"
  "µU/scr"
)
const (
  start = iota
  bp // before '.'
  ap // after '.'
  ee // immediately after 'E', i.e. in exponent
  ae // after 'E' plus next character
)

func (b *box) doneNumerical (s *string, x, y uint) bool {
  return true
/*/
  switch b.Comm {
  case kbd.Enter, kbd.Esc:
    return true
  case kbd.Back:
    switch b.depth {
    case 0:
      if b.overwritable {
        if b.index == 0 {
        } else {
          str.Rem (s, b.index - 1, 1)
          *s += " "
        }
      } else if b.index < b.width {
        str.Rem (s, b.index, 1)
        *s += " "
        b.index++
      } else if b.index == b.width {
        str.Rem (s, b.width - 1, 1)
        *s = " " + *s
      }
    case 1:
      *s = str.New (b.width)
      status = start
      b.index = b.width
    default:
      return true
    }
//    getStatus (s)
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
          str.Rem (s, b.index - 1, 1)
          *s += " "
        }
      } else if b.index < b.width {
        str.Rem (s, b.index, 1)
        *s += " "
        b.index++
      } else if b.index == b.width {
        str.Rem (s, b.width - 1, 1)
        *s = " " + *s
      }
    case 1:
      *s = str.New (b.width)
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
//  case kbd.Ins:
//    if b.depth == 0 {
//      if b.overwritable {
//        b.overwritable = false
//      } else if i < b.width {
//        b.overwritable = true
//      }
//    } else {
//      return true
//    }
  case kbd.Help, kbd.Search, kbd.Act, kbd.Cfg, kbd.Mark, kbd.Demark,
       kbd.Cut, kbd.Copy, kbd.Paste, kbd.Red, kbd.Green, kbd.Blue:
    return true
  }
  return false
/*/
}

func (b *box) possibleNumerical (s *string, x, y uint) bool {
  return true
/*/
  if b.index < b.width {
    panic ("uff") // return false
    if b.overwritable { return true }
    if (*s)[b.width - 1] == ' ' {
      // if ! overwritable, shift s one to the right and Write
      str.InsSpace (s, b.index)
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
        str.Replace1 (s, i, ' ')
      } else {
        break
      }
      i++
    }
    if (*s)[0] == ' ' {
      if b.width > 1 {
        str.Rem (s, 0, 1)
        *s += " "
      }
      return true
    }
  }
  return false
/*/
}

func (b *box) editNumber (imGraphikmodus bool, s *string, x, y uint) {
  b.graphical = imGraphikmodus
  n := b.width
  str.Norm (s, n)
  str.Move (s, false)
  b.write (*s, x, y)
  edited = false
  status := start
  cursorshape := Understroke
  b.index = n - 1
  var c byte
  if b.graphical {
    scr.WarpGr (x + scr.Wd1() * n, y, cursorshape)
  } else {
    scr.Warp (y / scr.Ht1(), x / scr.Wd1() + n, cursorshape)
  }
  for {
    if b.graphical {
//      scr.WriteGr (*s, y, x) // TODO
    } else {
      scr.Write (*s, x, y)
    }
    c, b.Comm, b.depth = kbd.Read()
    ok := true
    if b.Comm == kbd.None {
      switch c { // character
      case '+', '-':
        if status == start {
        } else if status == ee {
          status = ae
        } else {
          ok = false
        }
      case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
        if status == start {
          status = bp
        } else if status == ee {
          status = ae
        }
      case  '.', ',':
        if status == ap {
          ok = false
        } else {
          if status == bp {
            status = ap
            break
          }
        }
      case 'e', 'E':
        if status == ae {
          ok = false
        } else {
          status = ee
        }
      default:
        ok = false
      }
      if ok {
        *s = *s + string(c)
        *s = (*s)[1:]
/*/
        if b.graphical {
          scr.Write1Gr (c, int(x + scr.Wd1() * b.index), int(y))
        } else { // Command
          scr.Write1 (c, y / scr.Ht1(), x / scr.Wd1() + b.index)
        }
        str.Replace1 (s, b.index, c)
        b.index++
/*/
      }
    } else { // Command
      switch b.Comm {
      case kbd.Enter, kbd.Esc:
        kbd.DepositCommand (b.Comm)
//        kbd.DepositByte (c) // XXX necessary ?
        return
      case kbd.Back, kbd.Del:
        *s = " " + (*s)[:n-1]
      case kbd.Down, kbd.Up:
        // XXX as in case Enter ?
      case kbd.Help, kbd.Search, kbd.Act, kbd.Cfg,
           kbd.Mark, kbd.Demark, kbd.Cut,
           kbd.Copy, kbd.Paste:
        // XXX as in case Enter ?
      case kbd.Red, kbd.Green, kbd.Blue:
        // XXX as in case Enter ?
      }
    }
  }
}

/*
func (b *box) editNumber1 (imGraphikmodus bool, s *string, x, y uint) {
  for uint(len (*s)) < b.width { *s = " " + *s }
  b.Write (*s, x, y)
  for *s != "" && (*s)[0] == ' ' { *s = (*s)[1:] }; if *s == " " { *s = "" }
  var c byte
  if b.graphical {
    scr.WarpGr (x + scr.Wd1() * b.width, y, scr.Understroke)
  } else {
    scr.Warp (y / scr.Ht1(), x / scr.Wd1() + b.width, scr.Understroke)
  }
  loop: for {
    l := uint(len (*s))
    c, b.Comm, b.depth = kbd.Read()
    switch b.Comm {
    case kbd.None:
      if char.IsDigit (c) && l < b.width {
        *s += string(c)
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
