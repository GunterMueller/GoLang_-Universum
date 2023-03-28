package prog

// (c) Christian Maurer   v. 221213 - license see µU.go

import (
  "os"
  . "µU/obj"
  "µU/env"
  "µU/kbd"
  "µU/str"
  "µU/scr"
  "µU/col"
  "µU/box"
  "µU/errh"
  "µU/N"
  "µU/stk"
  "µU/pseq"
  "µU/files"
  "mini/line"
  "mini/reg"
)
const (
  EOC = byte(10)
  lenProgname = 8
  l0 = 2
  cLine = 4; cName = 10; cAccu = 48; cStack = 48; cReg = 66
//            1         2         3         4         5         6         7

//  01234567890123456789012345678901234567890123456789012345678901234567890123456789
//  Programm:                                       Akkumulatoren:    Registerwerte:
//
//   1  _ ___ _    27  _ ___ _    53  _ ___ _       ax = _________    a = _________
//   2  _ ___ _    28  _ ___ _    54  _ ___ _       bx = _________    b = _________
//   3  _ ___ _    29  _ ___ _    55  _ ___ _                         c = _________
//   4  _ ___ _    30  _ ___ _    56  _ ___ _       Statusregister:   d = _________
//   5  _ ___ _    31  _ ___ _    57  _ ___ _                         e = _________
//   6  _ ___ _    32  _ ___ _    58  _ ___ _       zf = 0            f = _________
//   7  _ ___ _    33  _ ___ _    59  _ ___ _       cf = 0            g = _________
//   8  _ ___ _    34  _ ___ _    60  _ ___ _                         h = _________
//   9  _ ___ _    35  _ ___ _    61  _ ___ _       Stack: _________  i = _________
//  ...
//  26  _ ___ _    52  program    78  _ ___ _                         z = _________
)
type
  program struct {
            file pseq.PersistentSequence
            name string
         strings []string
           lines []line.Line
              pc uint
         actLine line.Line
        stepwise bool
                 }
var (
  pcstack = stk.New (uint(0))
  maskbx, namebx, numberbx = box.New(), box.New(), box.New()
)

func writeMask() {
  maskbx.Wd (9)
  maskbx.Write ("Programm:", 0, 0)
  maskbx.Wd (14)
  maskbx.Write ("Akkumulatoren:", 0, cAccu)
  maskbx.Write ("Registerwerte:", 0, cReg)
  maskbx.Wd (4)
  maskbx.Write ("ax =", 2, cAccu)
  maskbx.Write ("bx =", 3, cAccu)
  maskbx.Wd (15)
  maskbx.Write ("Statusregister:", 5, cAccu)
  maskbx.Wd (4)
  maskbx.Write ("zf =", 7, cAccu)
  maskbx.Write ("cf =", 8, cAccu)
  maskbx.Wd (10)
  maskbx.Write ("Stack top:", 10, cStack)
  maskbx.Wd (1)
  s := ""
  for i := uint(0); i < reg.R; i++ {
    s = string(byte('a') + byte(i))
    maskbx.Write (s, 2 + i, cReg)
    maskbx.Write ("=", 2 + i, cReg + 2)
  }
}

func init() {
  maskbx.Wd (9)
  maskbx.Colours (col.LightWhite(), col.Black())
  namebx.Wd (lenProgname)
  namebx.Colours (col.Yellow(), col.Black())
  numberbx.Wd (2)
}

func new_() Program {
  files.Cds()
  x := new (program)
  x.file = pseq.New (byte(0))
  if env.NArgs() == 0 {
    x.name = "prog"
  } else {
    x.name = env.Arg (1)
  }
  x.strings = make([]string, 0)
  x.lines = make([]line.Line, 0)
  x.actLine = line.New()
  x.stepwise = true
  writeMask()
  return x
}

func (x *program) Empty() bool {
  return len(x.lines) == 0
}

func (x *program) GetLines() {
  namebx.Write (x.name, 0, cName)
  filename := x.name + ".mini"
  if ! files.IsFile (filename) {
    errh.Error0 ("Die Datei " + filename + " existiert nicht.")
    os.Exit(1)
  }
  x.file.Name (filename)
  x.strings = make([]string, 0)
  var s Stream
  for i := uint(0); i < x.file.Num(); i++ {
    x.file.Seek (i)
    s = append (s, x.file.Get().(byte))
  }
  var k uint
  x.strings, k = str.SplitByte (string(s), EOC)
  if k > 3 * reg.R {
    x.strings = x.strings[:78]
  }
}

func (x *program) Parse() (string, uint) {
  n := uint(len(x.strings))
  for i := uint(0); i < n; i++ {
    if ! x.actLine.Defined (x.strings[i]) {
      return x.strings[i], i
    }
    x.lines = append (x.lines, x.actLine.Clone().(line.Line))
    x.strings[i] = x.lines[i].String()
  }
  return "", n
}

func writeLinenumber (i, l, c uint, marked bool) {
  f := col.LightWhite(); if marked { f = col.Red() }
  numberbx.Colours (f, col.Black())
               numberbx.Write (N.StringFmt (i, 2, false), l, c)
}

func (x *program) Write() {
  reg.WriteAll (l0, cReg)
  line.WriteStatus (l0, cAccu)
  line.WriteStack (l0, cStack)
  n := uint(len(x.lines))
  if n == 0 { return }
  for i := uint(0); i < n; i++ {
    l, c := l0 + i % reg.R, i / reg.R
    writeLinenumber (i, l, c * 15, false)
    x.lines[i].Write (l, cLine + c * 15)
  }
}

func (x *program) Edit() {
  x.Write()
  i := uint(0)
  loop:
  for {
    l := l0 + i % reg.R
    c := cLine + 15 * (i / reg.R)
    x.lines[i].Edit (l, c)
    if x.lines[i].Empty() {
      if i + 1 == uint(len(x.lines)) {
        i--
        n := len(x.lines) - 1
        x.lines = x.lines[:n]
        continue
      } else {
        x.Write()
      }
    }
    switch k, _ := kbd.LastCommand(); k {
    case kbd.Esc:
      break loop
    case kbd.Enter:
      if i + 1 < uint(len(x.lines)) {
        i++
      } else {
        if i + 1 < 3 * reg.R {
          x.lines = append (x.lines, line.New())
          i++
          writeLinenumber (i, l0 + i % reg.R, 15 * (i / reg.R), false)
        }
      }
    case kbd.Down:
      if i + 1 < uint(len(x.lines)) {
        i++
        if i % reg.R == 0 { x.Write() }
      }
    case kbd.Up:
      if i > 0 {
        i--
      }
    case kbd.Pos1:
      i = 0
    case kbd.End:
      i = uint(len(x.lines)) - 1
    case kbd.Del:
      x.lines = append (x.lines[:i], x.lines[i+1:]...)
      scr.Clr (l0, 0, 3 * 14, reg.R)
      x.Write()
      if i >= uint(len(x.lines)) {
        i--
      }
    case kbd.Ins:
      newline := make([]line.Line, 1)
      newline[0] = line.New().Clone().(line.Line)
      tail := append (newline, x.lines[i:]...)
      x.lines = append (x.lines[:i], tail...)
      scr.Clr (l0, 0, 3 * 14, reg.R)
      x.Write()
    }
  }
  x.file.Clr()
  for i := uint(0); i < uint(len(x.lines)); i++ {
    s := x.lines[i].String()
    n := str.ProperLen (s)
    s += string(EOC)
    for i := uint(0); i <= n; i++ {
      x.file.Ins (s[i])
    }
  }
}

func (x *program) number (b byte) (uint, bool) {
  for i := uint(0); i < uint(len(x.lines)); i++ {
    if label, ok := x.lines[i].Marked(); ok && label == b {
      return i, true
    }
  }
  return 0, false
}

func (x *program) protocol() {
  x.Write()
  l, c := l0 + x.pc % reg.R, x.pc / reg.R
  writeLinenumber (x.pc, l, c, true)
  loop:
  for {
    c, _ := kbd.Command ()
    switch c {
    case kbd.Enter:
      break loop
    case kbd.Esc:
      x.stepwise = false
      errh.DelHint()
      break loop
    }
  }
  writeLinenumber (x.pc, l, c, false)
}

func (x *program) Run() {
  if x.Empty () {
    return
  }
  reg.EditAll (l0, cReg)
  errh.Hint ("Enter für nächsten Schritt, Esc für schnellen Durchlauf")
  x.pc = 0
  x.actLine = x.lines[x.pc]
  loop:
  for {
    if x.actLine.IsRet() {
      if pcstack.Empty() {
        break loop
      } else {
        x.pc = pcstack.Pop().(uint)
      }
    }
    x.actLine = x.lines[x.pc]
    if x.stepwise {
      x.protocol()
    }
    target := x.actLine.Run()
    reg.WriteAll (l0, cReg)
    line.WriteStatus (l0, cAccu)
    if target == line.EmptyLabel {
      x.pc++
    } else {
      if x.actLine.IsCall() {
        x.pc++
        pcstack.Push (x.pc)
      }
      if n, ok := x.number (target); ok {
        x.pc = n
      } else {
        errh.Error0 ("Programm abgebrochen: Sprungmarke kommt nicht vor")
        return
      }
    }
  }
  reg.WriteAll (l0, cReg)
  errh.DelHint()
  errh.Error0 ("Programm ausgeführt")
}
