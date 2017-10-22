package clk

// (c) Christian Maurer   v. 171005 - license see µU.go

import (
  . "µU/ker"
  . "µU/obj"
  "µU/str"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/font"
  "µU/pbox"
  "µU/errh"
  "µU/nat"
)
const (
  maxlength = 8 // maximal Formatlength for "Hh_mm_ss"
  ms = 60 // minutes / hour = seconds / minute
  hd = 24 // hours per day
  sd = hd * ms * ms // seconds per day
)
type
  clocktime struct {
              hour, // <= hd // 24 for the empty clocktime
            minute, // < ms
            second uint // < ms
                   Format
            cF, cB col.Colour
                   font.Font
                   }
var (
  textlength [NFormats]uint
  currentTime, clock = new_().(*clocktime), new_().(*clocktime)
  bx, clockbx = box.New(), box.New()
  line, column uint
  pbx = pbox.New()
)

func init() {
  textlength[Hh_mm] =    5
  textlength[Hh_mm_ss] = maxlength
  textlength[Mm_ss] =    5
  clock.Format = Hh_mm_ss
  line = 1000
//  SetAttributes(Zero, scr.NColumns() - textlength[clock.Format], col.HintF, col.HintB)
}

func new_() Clocktime {
  x := new(clocktime)
  x.hour = hd
  x.cF, x.cB = scr.StartCols()
  x.Format = Hh_mm
  return x
}

func (x *clocktime) imp(Y Any) *clocktime {
  y, ok := Y.(*clocktime)
  if ! ok { TypeNotEqPanic(x, Y) }
  return y
}

func (x *clocktime) Update() {
  x.hour, x.minute, x.second = UpdateTime()
}

func (x *clocktime) Empty() bool {
  return x.hour == hd
}

func (x *clocktime) Clr() {
  x.hour = hd
  x.minute, x.second = 0, 0
}

func (x *clocktime) Copy (Y Any) {
  y := x.imp(Y)
  x.hour, x.minute, x.second = y.hour, y.minute, y.second
}

func (x *clocktime) Clone() Any {
  y := new_()
  y.Copy(x)
  return y
}

func (x *clocktime) internalCode() uint {
  c := (ms * x.hour + x.minute) * ms + x.second
  if c > sd {
    println("jaul", c) // ; Panic("jaul")
    c = sd - 1
  }
  return c
}

func (x *clocktime) Eq (Y Any) bool {
  return x.internalCode() == x.imp(Y).internalCode()
}

func less (C, C1 uint) bool {
  if C == sd {
    return C1 != sd
  } else if C1 == sd {
    return false
  }
  return C < C1
}

func (x *clocktime) Less (Y Any) bool {
  return x.internalCode() < x.imp(Y).internalCode()
}

func (x *clocktime) Elapsed() bool {
  currentTime.Update()
  return x.internalCode() < currentTime.internalCode()
}

func (x *clocktime) Distance (Y Clocktime) uint {
  y := Y.(*clocktime)
  if x.Empty() || y.Empty() { return MaxNat() }
  c, d := x.internalCode(), y.internalCode()
  if d > c {
    return d - c
  }
  return c - d
}

func (x *clocktime) NSeconds() uint {
  if x.Empty() { return MaxNat() }
  return x.internalCode()
}

func (x *clocktime) Hours() uint {
  if x.Empty() { return hd }
  return x.hour
}

func (x *clocktime) Minutes() uint {
  if x.Empty() { return ms }
  return x.minute
}

func (x *clocktime) Seconds() uint {
  if x.Empty() { return ms }
  return x.second
}

func (x *clocktime) Inc (Y Clocktime) {
  y := Y.(*clocktime)
  if x.Empty() || y.Empty() { return }
  c := x.internalCode()
  c+= y.internalCode()
  c = c % sd
  x.second = c % ms
  c = c / ms
  x.minute, x.hour = c % ms, c / ms
}

func (x *clocktime) Dec (Y Clocktime) {
  y := Y.(*clocktime)
  if x.Empty() || y.Empty() { return }
  c := x.internalCode() + sd
  c-= y.internalCode()
  c = c % sd
  x.second = c % ms
  c = c / ms
  x.minute, x.hour = c % ms, c / ms
}

func (x *clocktime) SetFormat(f Format) {
  if f < NFormats {
    x.Format = f
  }
}

func (x *clocktime) GetFormat() Format {
  return x.Format
}

func (x *clocktime) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *clocktime) String() string {
  if x.Empty() {
    return str.Clr(textlength[x.Format])
  }
  s := nat.StringFmt(x.minute, 2, true)
  if x.Format <= Hh_mm_ss {
    s = nat.StringFmt(x.hour, 2, true) + "." + s
  }
  if x.Format >= Hh_mm_ss {
    s += ":" + nat.StringFmt(x.second, 2, true)
  }
  return s
}

func (x *clocktime) defined(h, m, s uint) bool {
  if h < hd { x.hour = h } else { return false }
  if m < ms { x.minute = m } else { return false }
  if s < ms { x.second = s } else { return false }
  x.hour, x.minute, x.second = h, m, s
  return true
}

func (x *clocktime) Defined(t string) bool {
  x.Clr()
  if str.Empty(t) { return true }
  n, ss, P, L := nat.DigitSequences(t)
  if n == 0 || n > 3 { return false }
  if n == 3 {
    if x.Format == Hh_mm { return false }
  }
  if L[0] >= textlength[x.Format] { return false }
  h, m, s, ok := uint(0), uint(0), uint(0), false
  if n == 1 {
    if h, ok = nat.Natural(str.Part(t, P[0], 2)); ! ok { return false }
    if L[0] > 2 {
      if m, ok = nat.Natural(str.Part(t, P[0] + 2, 2)); ! ok { return false }
      if L[0] > 4 {
        if s, ok = nat.Natural(str.Part(t, P[0] + 4, 2)); ! ok { return false }
      }
    }
  } else {
    if h, ok = nat.Natural(ss[0]); ! ok { return false }
    if m, ok = nat.Natural(ss[1]); ! ok { return false }
    if n == 2 && x.Format == Mm_ss {
      s, m, h = m, h, 0
    }
    if n == 3 {
      if _, ok := nat.Natural(ss[2]); ! ok { return false }
    }
  }
  return x.defined(h, m, s)
}

func (x *clocktime) Set(h, m, s uint) bool {
  if h < hd && m < ms && s < ms {
    x.hour, x.minute, x.second = h, m, s
    return true
  }
  x.Clr()
  return false
}

func (x *clocktime) SetSeconds(s uint) bool {
  if s < sd {
    m := s % ms
    h := m % ms
    s %= h * m
    x.hour, x.minute, x.second = h, m, s
    return true
  }
  x.Clr()
  return false
}

func (x *clocktime) Write(l, c uint) {
  bx.Wd(textlength[x.Format])
  bx.Colours(x.cF, x.cB)
  bx.Write(x.String(), l, c)
}

func (x *clocktime) Edit(l, c uint) {
  x.Write(l, c)
  s := x.String()
  err := uint(0)
  for {
    bx.Edit(&s, l, c)
    if x.Defined(s) {
      s = x.String()
      bx.Write(s, l, c)
      return
    } else {
      err ++
      switch err { case 1: // --> errh.WriteError
        errh.Error0("Die Uhrzeitangabe ist unverständlich!") // , l + 1, c)
      case 2:
        errh.Error0("Die Angabe ist immer noch unklar!") // , l + 1, c)
      case 3:
        errh.Error0("Jetzt pass doch mal uff, Du Trottel!") // l + 1, c)
      case 4:
        errh.Error0("Was soll der Quatsch? Ist das eine Uhrzeit?") // l + 1, c)
      case 5:
        errh.Error0("Bist Du zu doof, eine Uhrzeit einzugeben?") // l + 1, c)
      default:
        errh.Error0("Vergiss es ...") // , l + 1, c)
        x.Update()
        x.Write(l, c)
        return
      }
    }
  }
}

func (x *clocktime) SetFont(f font.Font) {
  x.Font = f
}

func (x *clocktime) Print(l, c uint) {
  pbx.SetFont(x.Font)
  pbx.Print(x.String(), l, c)
}

func (x *clocktime) Codelen() uint {
  return 4 // Codelen(uint32(0))
}

func (x *clocktime) Encode() []byte {
  bs := make([]byte, x.Codelen())
  copy(bs, Encode(uint32(x.internalCode())))
  return bs
}

func (x *clocktime) Decode(bs []byte) {
  n := uint(Decode(uint32(0), bs).(uint32))
  x.second = n % ms
  n /= ms
  x.minute = n % ms
  n /= ms
  if n > hd {
    x.hour = hd
  } else {
    x.hour = n
  }
}

func SetAttributes(l, c uint, f, b col.Colour) {
  line, column = l, c
  clockbx.Colours(f, b)
}

func Show() {
  if line >= scr.NLines() {
    SetAttributes(Zero, scr.NColumns() - textlength[clock.Format], col.HintF(), col.HintB())
  }
  for {
    clock.Update()
    clockbx.Write(clock.String(), line, column)
    Sleep(1) // not precise, but good enough for practical purposes
             // more precise would be: sleep until AlarmClock rings
  }
}
