package line

// (c) Christian Maurer   v. 221213 - license see µU.go

import (
  . "µU/obj"
  "µU/ker"
  "µU/str"
  "µU/col"
  "µU/box"
  "µU/errh"
  "µU/N"
  "µU/stk"
  "mini/reg"
)
const (
  comment = byte('#')
  maxLen = 8
)
type
  argtype byte; const (
  noArg = argtype (iota)
  regArg
  labArg
)
type
  line struct {
        label byte
              Instruction
              reg.Register
       target byte
              }
type
  errorType byte; const (
  noError = errorType (iota)
  tooManyIds
  noInst
  argTooMuch
  regMissing
  jumpMissing
  noJumpLabel
  nErrors
)
var (
  ax, bx, zf, cf uint
  stack = stk.New (uint(0))
  txt [nInstructions + 1]string
  typ [nInstructions + 1]argtype
  bx1 = box.New()
  errorText [nErrors]string
  fehler errorType
)

func def (i Instruction, s string, t argtype) {
  txt[i], typ[i] = s, t
}

func init() {
  bx1.Wd (maxLen)
  bx1.Colours (col.Yellow(), col.Black())
  def (NOP,  "   ", noArg)
  def (LDA,  "lda", regArg)
  def (STA,  "sta", regArg)
  def (LDB,  "ldb", regArg)
  def (STB,  "stb", regArg)
  def (EXA,  "exa", regArg)
  def (EXB,  "exb", regArg)
  def (INA,  "ina", noArg)
  def (DEA,  "dea", noArg)
  def (INC,  "inc", regArg)
  def (DEC,  "dec", regArg)
  def (SHL,  "shl", regArg)
  def (SHR,  "shr", regArg)
  def (ADD,  "add", regArg)
  def (ADC,  "adc", regArg)
  def (SUB,  "sub", regArg)
  def (MUL,  "mul", regArg)
  def (DIV,  "div", regArg)
  def (CMP,  "cmp", regArg)
  def (JMP,  "jmp", labArg)
  def (JE,   "je",  labArg)
  def (JNE,  "jne", labArg)
  def (JC,   "jc",  labArg)
  def (JNC,  "jnc", labArg)
  def (PUSH, "push",regArg)
  def (POP,  "pop", regArg)
  def (CLC,  "clc", noArg)
  def (STC,  "stc", noArg)
  def (CMC,  "cmc", noArg)
  def (CALL, "call",labArg)
  def (RET,  "ret", noArg)
  errorText[noError] =     ""
  errorText[tooManyIds] =  "zu viele Bezeichner"
  errorText[noInst] =      "kein Befehl"
  errorText[argTooMuch] =  "Argument zuviel"
  errorText[regMissing] =  "Register fehlt"
  errorText[jumpMissing] = "Sprungmarke fehlt"
  errorText[noJumpLabel] = "keine Sprungmarke"
}

func new_() Line {
  x := new (line)
  x.label = EmptyLabel
  x.Register = reg.New()
  x.target = EmptyLabel
  return x
}

func (x *line) imp (Y any) *line {
  y, ok := Y.(*line)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *line) Clr() {
  x.label = EmptyLabel
  x.Instruction = NOP
  x.Register.Clr()
  x.target = EmptyLabel
}

func (x *line) Empty() bool {
  return x.Instruction == NOP
}

func (x *line) Eq (Y any) bool {
  return x.String() == x.imp(Y).String()
}

func (x *line) Copy (Y any) {
  x.Defined (x.imp (Y).String())
}

func (x *line) Clone() any {
  y := New ()
  y.Copy (x)
  return y
}

func (x *line) IsCall() bool {
  return x.Instruction == CALL
}

func (x *line) IsRet() bool {
  return x.Instruction == RET
}

func (x *line) Marked() (byte, bool) {
  return x.label, x.label != EmptyLabel
}

func (x *line) String() string {
  s := string(x.label) + " " + txt[x.Instruction]
  if ! x.Register.Empty () {
    s += " " + x.Register.String()
  }
  s += " " + string(x.target)
  str.OffSpc1 (&s)
  return s
}

func (x *line) Write (l, c uint) {
  bx1.Write (x.String(), l, c)
}

func writeStatus (l, c uint) {
  N.Colours (col.LightGreen(), col.Black())
  N.Write (ax, l,     c + 5)
  N.Write (bx, l + 1, c + 5)
  N.Colours (col.LightRed(), col.Black())
  N.Write (zf, l + 5, c + 5)
  N.Write (cf, l + 6, c + 5)
}

func writeStack (l, c uint) {
  N.Colours (col.LightBlue(), col.Black())
  if ! stack.Empty() {
    k := stack.Pop().(uint)
    stack.Push (k)
    N.Colours (col.LightGreen(), col.Black())
    N.Write (k, l + 10, c + 0)
  }
}

func isLabel (b byte) bool {
  return 'A' <= b && b <= 'Z'
}

func (x *line) Defined (s string) bool {
  x.Clr()
  fehler = noError
  str.OffSpc1 (&s)
  nn := uint(len (s))
  if nn > maxLen { s = str.Part (s, 0, nn) }
  n, ts, _ := str.Split (s)
  if n == 0 {
    return true
  }
  if ts[0] == string(comment) {
    x.Instruction = NOP
    return true
  }
  posInst, posArg := uint(0), uint(0) // <= 2
  var lab byte
  switch n {
  case 1: // no label, only inst (without argument)
    lab = EmptyLabel
  case 2: // label and inst (without argument) or inst with argument
    if len (ts[0]) > 1 { /* error */ }
    lab = ts[0][0]
    if isLabel (lab) { // label and inst (without argument)
      posInst = 1
    } else { // no label, but inst with argument
      lab = EmptyLabel
      posArg = 1
    }
  case 3: // label and inst with argument
    if len (ts[0]) > 1 { /* error */ }
    lab = ts[0][0]
    posInst, posArg = 1, 2
  default:
    fehler = tooManyIds
    return false
  }
  if isLabel (lab) {
    x.label = lab
  }
  x.Instruction = Instruction(1)
  for {
    if str.Eq (ts[posInst], txt[x.Instruction]) {
      break
    } else if x.Instruction < nInstructions {
      x.Instruction++
    } else {
      x.Instruction = NOP
      fehler = noInst
      return false
    }
  }
  switch typ[x.Instruction] {
  case noArg:
    if posArg > 0 {
      fehler = argTooMuch
      return false
    }
  case labArg:
    if posArg == 0 {
      fehler = jumpMissing
      return false
    }
    if ! isLabel (ts[posArg][0]) {
      fehler = noJumpLabel
      return false
    }
    x.target = ts[posArg][0]
  case regArg:
    if ! x.Register.Defined (ts[posArg]) {
      fehler = regMissing
      return false
    }
  }
  return true
}

func (x *line) Edit (l, c uint) {
  x.Write (l, c)
  s := x.String()
  for {
    bx1.Edit (&s, l, c)
    if x.Defined (s) {
      break
    } else {
      errh.Error0 (errorText[fehler])
    }
  }
}

func (x *line) Run() byte {
  v := x.Register.Val()
  m := EmptyLabel
  switch x.Instruction {
  case NOP:
    ;
  case LDA:
    ax = v
  case STA:
    x.Register.SetVal (ax)
  case LDB:
    bx = v
  case STB:
    x.Register.SetVal (bx)
  case EXA:
    v, ax = ax, v
    x.Register.SetVal (v)
  case EXB:
    v, bx = bx, v
    x.Register.SetVal (v)
  case INA:
    if ax + 1 < reg.M {
      ax++
    } else {
      ax = 0
    }
    if ax == 0 {
      zf = 1
    } else {
      zf = 0
    }
  case DEA:
    if ax > 0 {
      ax--
    } else {
      ax = reg.M - 1
    }
    if ax == 0 {
      zf = 1
    } else {
      zf = 0
    }
  case SHR:
    cf = v % 2
    v = v / 2
    x.Register.SetVal (v)
  case SHL:
    v = (2 * v) % reg.M
    x.Register.SetVal (v)
    if 2 * v < reg.M {
      cf = 0
    } else {
      cf = 1
    }
  case INC:
    if v + 1 < reg.M {
      v++
    } else {
      v = 0
    }
    x.Register.SetVal (v)
    if v == 0 {
      zf = 1
    } else {
      zf = 0
    }
  case DEC:
    if v == 0 {
      v = reg.M - 1
    } else {
      v--
    }
    x.Register.SetVal (v)
    if v == 0 {
      zf = 1
    } else {
      zf = 0
    }
  case ADD:
    ax += v
    if ax < reg.M {
      cf = 0
    } else {
      ax = ax % reg.M
      cf = 1
    }
  case ADC:
    ax += (v + cf)
    if ax < reg.M {
      cf = 0
    } else {
      ax = ax % reg.M
      cf = 1
    }
  case SUB:
    if v <= ax {
      ax -= v
      cf = 0
    } else {
      ax = reg.M - (v - ax)
      cf = 1
    }
  case MUL:
    p := ax * v
    ax, bx = p % 1e9, p / 1e9
    if bx == 0 {
      cf = 0
    } else {
      cf = 1
    }
  case DIV:
    if v == 0 {
      ker.Panic ("Division durch 0")
    }
    ax, bx = ax / v, ax % v
  case CMP:
    if v == ax {
      zf = 1
    } else {
      zf = 0
      if v < ax {
        cf = 1
      } else { // v > ax
        cf = 0
      }
    }
  case JMP:
    m = x.target
  case JE:
    if zf == 1 {
      m = x.target
    }
  case JNE:
    if zf == 0 {
      m = x.target
    }
  case JC:
    if cf == 1 {
      m = x.target
    }
  case JNC:
    if cf == 0 {
      m = x.target
    }
  case PUSH:
    stack.Push (v)
  case POP:
    if stack.Empty() {
      ker.Panic ("Programmabbruch: der Stapel ist leer")
    }
    v = stack.Pop().(uint)
    x.Register.SetVal (v)
  case CLC:
    cf = 0
  case STC:
    cf = 1
  case CMC:
    cf = 1 - cf
  case CALL:
    m = x.target
  case RET:
    // fertig
  }
  return m
}
