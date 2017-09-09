package kbd

// (c) Christian Maurer  v. 170903 - license see murus.go

import (
  "murus/spc"
  "murus/xwin"
  "murus/z"
  "murus/mouse"
  "murus/navi"
)
const (
// PIPE_BUF = 256
// raw key codes
// alphanumeric keyboard: control keys
  shiftL    =  42
  shiftR    =  54
  shiftLock =  58
  ctrlL     =  29
  ctrlR     =  97
  altL      =  56
  altR      = 100
  esc       =   1
  back      =  14
  tab       =  15
  enter     =  28
// alphanumeric keyboard:
  f1        =  59; f2  = 60; f3 = 61; f4 = 62
  f5        =  63; f6  = 64; f7 = 65; f8 = 66
  f9        =  67; f10 = 68
  f11       =  87; f12 = 88
// numeric keypad:
  numEnter  =  96
  num0      =  82; num1 = 79; num2 = 80; num3 = 81; num4 = 75
  num5      =  76; num6 = 77; num7 = 71; num8 = 72; num9 = 73
  numSep    =  83
  numMinus  =  74; numPlus = 78; numTimes = 55; numDiv = 98
  left      = 105
  right     = 106
  up        = 103
  down      = 108
// ? = 101; ? = 112
// ? = 117; ? = 118
// ? = 120; ? = 121; ? = 122; ? = 123; ? = 124
// special keypad:
  pgUp      = 104
  pgDown    = 109
  pos1      = 102
  end       = 107
  ins       = 110
  del       = 111 // XXX
// ? = 89; ? = 90; ? = 91; ? = 92; ? = 93; ? = 94; ? = 95
  prt       =  99
  roll      =  70
  pause     = 119
  numOnOff  =  69
  onOff     = 113; lower = 114; louder = 115
  doofL     = 125; doofM = 126; doofR = 127
  noKeycodes = 128
//  toolbox   = 501 // 501 % 256 = 245, % 128 = 117
  pgRight   = 158 // only under X
  pgLeft    = 159 // only under X
  off       = 128
// combinations:
  shiftLoff = shiftL + off
  shiftRoff = shiftR + off
  shiftLockoff = shiftLock + off
  ctrlLoff  = ctrlL + off
  ctrlRoff  = ctrlR + off
  altLoff   = altL + off
  altRoff   = altR + off
  doofLoff  = doofL + off // d(o,o)f
  doofMoff  = doofM + off
  doofRoff  = doofR + off
  function  = 143
)
var (
  bb,       // key
  bB,       // key + Shift
  aa []byte // key + AltGr
  aA []byte // key + Shift + AltGr
  kK [noKeycodes]Comm
  lastbyte byte
  lastcommand Comm
  lastdepth uint
//  shift, shiftFix, ctrl, alt, altGr /* , numOnOff */, fn, lBut, mBut, rBut bool
)

func init() {
  //           0         1         2         3         4         5         6         7         8         9
  //           012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678
  bb = []byte("  1234567890 '  qwertzuiop +  asdfghjkl  ^ #yxcvbnm,.- +               789-456+1230,  <           /")
  //                       ß             ü            öä
  bb[12] = z.Sz
  bb[26] = z.Ue
  bb[39] = z.Oe
  bb[40] = z.Ae

  //           0         1         2         3         4         5         6         7         8         9
  //           012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678
  bB = []byte("  !  $%&/()=?`  QWERTZUIOP *  ASDFGHJKL    'YXCVBNM;:_ *               789-456+1230,  >           /")
  //               §                     Ü            ÖÄ°
  bB [4] = z.Paragraph
  bB[26] = z.Ü
  bB[39] = z.Ö
  bB[40] = z.Ä
  bB[41] = z.Degree

  //           0         1         2         3         4         5         6         7         8         9
  //           012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678
  aa = []byte("     $%&/()=?`  @WERTZUIOP ~  ASDFGHJKL    'YX VBN ;:_ ~               {[]-456+123},  |           /")
  //             ¹²³             €       ü            öä    ¢©   µ
  aa [2] = z.ToThe1
  aa [3] = z.ToThe2
  aa [4] = z.ToThe3
  aa[18] = z.Euro
  aa[26] = z.Ue
  aa[39] = z.Oe
  aa[40] = z.Ae
  aa[45] = z.Cent
  aa[50] = z.Mue

  //           0         1         2         3         4         5         6         7         8         9
  //           012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678
  aA = []byte("  ! $%&/()=?`  @WERTZUIOP ~  ASDFGHJKL    'YXCVBNM;:_ ~               {[]-456+123},  |           /")
  //                             ®             ª       ¬    ¢   º ÷                    ± ¤     £ ×
  aA[18] = z.Registered
  aA[33] = z.Female
  aA[41] = z.Negate
  aA[46] = z.Copyright
  aA[50] = z.Male
  aA[52] = z.Division
  aA[73] = z.PlusMinus
  aA[75] = z.Euro
  aA[81] = z.Pound
  aA[83] = z.Times

  for b:= 0; b < noKeycodes; b++ { kK[b] = Esc }
  kK[esc]    = Esc
  kK[f1]     = Help
  kK[f2]     = Search
  kK[f3]     = Act
  kK[f4]     = Cfg
  kK[f5]     = Mark
  kK[f6]     = Demark
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
  kK[pgUp]   = Up
  kK[left]   = Left
  kK[right]  = Right
  kK[end]    = End
  kK[down]   = Down
  kK[pgDown] = Down
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
  kK[onOff]  = OnOff
  kK[lower]  = Lower
  kK[louder] = Louder
  lastbyte, lastcommand, lastdepth = 0, None, 0
  underX = xwin.UnderX()
  if underX {
    xpipe = make (chan xwin.Event)
    go catchX()
  } else {
    initConsole()
  }
}

func isAlpha (n uint) bool {
  switch n { case 41, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13, // ^ 1 2 3 4 5 6 7 8 9 0 ß '
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
  switch n { case f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12:
    return true
  }
  return false
}

func isCmd (n uint) bool {
  switch n { case esc, back, tab, enter,
                  left, right, up, down, pgUp, pgDown, pos1, end,
                  ins, del,
                  prt, roll, pause,
                  onOff, lower, louder,
                  numEnter:
    return true
  }
  return isF (n)
}

func isKeypad (n uint) bool {
  switch n { case num0, num1, num2, num3, num4, num5, num6, num7, num8, num9, numSep:
    return true
  }
  return false
}

func read() (byte, Comm, uint) {
  var (b byte; c Comm; d uint)
  if underX {
    inputX (&b, &c, &d)
  } else {
    input (&b, &c, &d)
  }
  return b, c, d
}

func mouseEx() bool {
  if underX {
    return true
  }
  return mouse.Ex()
}

func byte_() byte {
  b:= byte(0)
  for {
    b, _, _ = read()
    if b != 0 {
      break
    }
  }
  return b
}

func command() (Comm, uint) {
  var ( c Comm; d uint )
  for {
    _, c, d = read()
//    if b == 0 { break }
    if c != None { break }
  }
  return c, d
}

func readNavi() (spc.GridCoord, spc.GridCoord) {
  return navi.Read()
}

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
  c0, d0:= lastcommand, lastdepth
  var ( c Comm; d uint )
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

func confirmed (w bool) bool {
  c0, d0:= lastcommand, lastdepth
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
  switch c, d:= LastCommand(); c {
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
