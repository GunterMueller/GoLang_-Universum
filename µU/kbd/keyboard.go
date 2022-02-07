package kbd

// (c) Christian Maurer   v. 220203 - license see µU.go

import (
  "µU/env"
  "µU/ker"
  . "µU/obj"
  "µU/char"
  "µU/scr"
//  "µU/navi"
)
const (
// control keys
  esc       =   1
  enter     =  28
  shiftL    =  42
  shiftR    =  54
  shiftLock =  58
  ctrlL     =  29
  ctrlR     =  97
  altL      =  56
  altR      = 100
  back      =  14
  tab       =  15
// function keys
  f1        =  59
  f2        =  60
  f3        =  61
  f4        =  62
  f5        =  63
  f6        =  64
  f7        =  65
  f8        =  66
  f9        =  67
  f10       =  68
  f11       =  87
  f12       =  88
// numeric keypad
  numEnter  =  96
  num0      =  82
  num1      =  79
  num2      =  80
  num3      =  81
  num4      =  75
  num5      =  76
  num6      =  77
  num7      =  71
  num8      =  72
  num9      =  73
  numSep    =  83
  numMinus  =  74
  numPlus   =  78
  numTimes  =  55
  numDiv    =  98
// direction keys
  left      = 105
  right     = 106
  up        = 103
  down      = 108
// special keys:
  pgRight   = 113 // only under X
  pgLeft    = 114 // only under X
  pgUp      = 104
  pgDown    = 109
  pos1      = 102
  end       = 107
  ins       = 110
  del       = 111
  prt       =  99
  roll      =  70
  pause     = 119
  numOnOff  =  69
//  onOff     = 113
//  lower     = 114
//  louder    = 115
  doofL     = 125
  doofR     = 126
  windoof   = 127
// AltGr-keys:
  backslash =  92
  braceL    = '{'
  braceR    = '}'
  bracketL  = '['
  bracketR  = ']'
// UTF-8-symbols:
  degree     = char.Degree
  twoSup     = char.ToThe2
  threeSup   = char.ToThe3
  Ä          = char.Ä
  Ö          = char.Ö
  Ü          = char.Ü
  ä          = char.Ae
  ö          = char.Oe
  ü          = char.Ue
  ß          = char.Sz
  euro       = char.Euro
  cent       = char.Cent
  mu         = char.Mu
  paragraph  = char.Paragraph
  registered = char.Registered
  female     = char.Female
  copyright  = char.Copyright
  male       = char.Male
  division   = char.Division

//  toolbox   = 501 // 501 % 256 = 245, % 128 = 117
  noKeycodes= 128
// combinations:
  off       = 128
  shiftLoff = shiftL + off
  shiftRoff = shiftR + off
  shiftLockoff = shiftLock + off
  ctrlLoff  = ctrlL + off
  ctrlRoff  = ctrlR + off
  altLoff   = altL + off
  altRoff   = altR + off
//  function  = 143
)
var (
  under_C, under_X, under_S bool
  shift, ctrl, alt, altGr, /* fn, */ mouseL, mouseM, mouseR bool
  bb, aa Stream
  kK [noKeycodes]Comm
  lastbyte byte
  lastcommand Comm
  lastdepth uint
)

func init() {
  //           0         1         2         3         4         5         6         7         8         9
  //           012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678
  bb = Stream("  1234567890 '  qwertzuiop +  asdfghjkl  ^ #yxcvbnm,.- +               789-456+1230,  <           /")
  //                       ß             ü            öä
  bb[12] = ß
  bb[26] = ü
  bb[39] = ö
  bb[40] = ä
  //           0         1         2         3         4         5         6         7         8         9
  //           012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678
  aa = Stream("  !  $%&/()=?`  QWERTZUIOP ~  ASDFGHJKL    'YXCVBNM;:_ *               789-456+1230,  |           /")
  //             ¹²³             €       Ü            ÖÄ    ¢©   µ
  aa [3] = '"'
  aa [4] = paragraph
  aa[26] = Ü
  aa[39] = Ö
  aa[40] = Ä
  aa[41] = degree
  for b := 0; b < noKeycodes; b++ {
    kK[b] = Esc
  }
  kK[esc]    = Esc
  kK[f1]     = Help
  kK[f2]     = Search
  kK[f3]     = Act
  kK[f4]     = Cfg
  kK[f5]     = Mark
  kK[f6]     = Unmark
  kK[f7]     = Cut
  kK[f8]     = Copy
  kK[f9]     = Paste
  kK[f10]    = Red
  kK[f11]    = Green
  kK[f12]    = Blue
  kK[back]   = Back
  kK[tab]    = Tab
  kK[enter]  = Enter
  kK[prt]    = Print
  kK[pos1]   = Pos1
  kK[up]     = Up
  kK[pgUp]   = PgUp
  kK[left]   = Left
  kK[pgLeft] = PgLeft
  kK[right]  = Right
  kK[pgRight]= PgRight
  kK[end]    = End
  kK[down]   = Down
  kK[pgDown] = PgDown
  kK[ins]    = Ins
  kK[del]    = Del
  kK[roll]   = Roll
  kK[num7]   = kK[pos1]
  kK[num8]   = kK[up]
  kK[num9]   = kK[pgUp]
  kK[num4]   = kK[left]
  kK[num6]   = kK[right]
  kK[num7]   = kK[end]
  kK[num8]   = kK[down]
  kK[num9]   = kK[pgDown]
  kK[num0]   = kK[ins]
  kK[numSep] = kK[del]
  kK[numEnter] = kK[enter]
  kK[pause]  = Pause
//  kK[onOff]  = OnOff
//  kK[lower]  = Lower
//  kK[louder] = Louder
  lastbyte, lastcommand, lastdepth = 0, None, 0
  under_C, under_X = env.UnderC(), env.UnderX()
  if under_X {
    xpipe = make(chan scr.Event)
    go catchX()
  } else {
    if under_C {
      initConsole()
    } else {
      ker.Panic ("no X, no C")
    }
  }
}

func isAlpha (n uint) bool {
  switch n {
  case 41, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13, // ^ 1 2 3 4 5 6 7 8 9 0 ß '
        16,17,18,19,20,21,22,23,24,25,26,27,   //  Q W E R T Z U I O P Ü +
         30,31,32,33,34,35,36,37,38,39,40,43,  //   A S D F G H J K L Ö Ä #
        86,44,45,46,47,48,49,50,51,52,53,      //  < Y X C V B N M , . -
                       57,                     // space
       numMinus, numPlus, numTimes, numDiv:    // keypad
    return true
  }
  return false
}

func isF (n uint) bool {
  switch n {
  case f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12:
    return true
  }
  return false
}

func isCmd (n uint) bool {
  switch n {
  case esc, back, tab, enter,
       left, right, up, down, pgUp, pgDown, pos1, end,
       ins, del,
       prt, roll, pause,
//       onOff, lower, louder,
       numEnter:
    return true
  }
  return isF (n)
}

func isKeypad (n uint) bool {
  switch n {
  case num0, num1, num2, num3, num4, num5, num6, num7, num8, num9, numSep:
    return true
  }
  return false
}

func read() (byte, Comm, uint) {
  var (b byte; c Comm; d uint)
  if under_X {
    inputX (&b, &c, &d)
  } else { // under_C
    inputC (&b, &c, &d)
  }
  return b, c, d
}

func byte_() byte {
  b := byte(0)
  for {
    b, _, _ = read()
    if b != 0 {
      break
    }
  }
  return b
}

func command() (Comm, uint) {
  c, d := None, uint(0)
  for {
    _, c, d = read()
    if c != None { break }
  }
  return c, d
}

/*
func readNavi() (spc.GridCoord, spc.GridCoord) {
  return navi.Read()
}
*/

func lastByte() byte {
  return lastbyte
}

func lastCommand() (Comm, uint) {
  return lastcommand, lastdepth
}

func depositCommand (c Comm) {
  lastcommand = c
}

func depositByte (b byte) {
  lastbyte = b
}

func wait (b bool) bool {
  c0, d0 := lastcommand, lastdepth
  var (c Comm; d uint)
  for {
    c, d = Command()
    if b {
      if c == Enter /* || c == Here */ { break }
    } else {
      if c == Esc || c == Back /* || c == There */ { break }
    }
  }
  lastcommand, lastdepth = c0, d0
  return d == 0
}

func waitFor (command Comm, depth uint) {
  c0, d0 := lastcommand, lastdepth
  for {
    c, d := Command()
    if c == command && d == depth { break }
  }
  lastcommand, lastdepth = c0, d0
}

func quit() {
  c0 := lastcommand
  for {
    switch c, _ := Command(); c {
    case Enter, Esc, Back:
      break
    }
  }
  lastcommand = c0
}

func confirmed (w bool) bool {
  c0, d0 := lastcommand, lastdepth
  var ( c Comm; d, dmin uint )
  if w {
    dmin = 1
  } else {
    dmin = 0
  }
  var b bool
  for {
    c, d = Command()
    if c == Enter {
      if d >= dmin {
        b = true
        break
      }
    } else if c == Esc {
      if d >= dmin {
        b = false
        break
      }
    }
  }
  lastcommand, lastdepth = c0, d0
  return b
}

/*
func Control (n uint, i *uint) {
  switch c, d := LastCommand(); c {
  case Esc:
    break loop
  case Enter:
    if d == 0 {
      if *i + 1 < n {
        *i ++
      } else {
        break loop
      }
    } else {
      break loop
    }
  case Down:
    if *i + 1 < n {
      *i ++
    } else {
      *i = 0
    }
  case Up:
    if *i > 0 {
      *i --
    } else {
      *i = n - 1
    }
  case Pos1:
    *i = 0
  case End:
    *i = n - 1
  }
}
 */
