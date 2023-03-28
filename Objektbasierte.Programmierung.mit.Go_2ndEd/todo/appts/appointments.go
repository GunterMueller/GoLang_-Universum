package appts

// (c) Christian Maurer   v. 221021 - license see µU.go

import (
  "sort"
  . "µU/obj"
  "µU/kbd"
  "µU/font"
  "µU/errh"
  "µU/day"
  "µU/seq"
  "µU/stk"
  . "todo/help"
  "todo/appt"
  "todo/attr"
)
const
  nAppointments = uint(21)
type
  appointments struct {
                 appt []appt.Appointment
                      }
var (
  actualFormat day.Period
  emptyAppointment = appt.New()
  clipboard = seq.New (emptyAppointment)
  trash = stk.New (emptyAppointment)
  dl uint
)

func init() {
  clipboard.Sort()
//  SetFormat (day.Daily)
}

func (x *appointments) imp (Y any) *appointments {
  y, ok := Y.(*appointments)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func New() Appointments {
  x := new (appointments)
  x.appt = make([]appt.Appointment, nAppointments)
  for i := uint(0); i < nAppointments; i++ {
    x.appt[i] = appt.New()
  }
  return x
}

func (x *appointments) Empty() bool {
  for i := uint(0); i < nAppointments; i++ {
    if ! x.appt[i].Empty() {
      return false
    }
  }
  return true
}

func (x *appointments) Clr() {
  for i := uint(0); i < nAppointments; i++ {
    x.appt[i].Clr()
  }
}

func (x *appointments) Eq (Y any) bool {
  y := x.imp (Y)
  for i := uint(0); i < nAppointments; i++ {
    if ! x.appt[i].Eq (y.appt[i]) {
      return false
    }
  }
  return true
}

func (x *appointments) Less (Y any) bool {
  return false
}

func (x *appointments) Leq (Y any) bool {
  return false
}

func (x *appointments) Copy (Y any) {
  y := x.imp (Y)
  for i := uint(0); i < nAppointments; i++ {
    x.appt[i].Copy (y.appt[i])
  }
}

func (x *appointments) Clone() any {
  y := New()
  y.Copy (x)
  return y
}

func (x *appointments) SetFormat (p day.Period) {
  actualFormat = p
  dl = 1
//  if x.num() == 0 { return }
  switch actualFormat {
  case day.Daily:
    x.appt[0].SetFormat (appt.Long)
  case day.Weekly:
    x.appt[0].SetFormat (appt.Short)
  case day.Monthly:
    x.appt[0].SetFormat (appt.Short)
    dl = 0
    // aus den Appointments werden nur die Appointmentattribute gebraucht
  default:
    errh.Error ("seq.SetFormat default: Quatsch", 0)
//    x.appt[0].SetFormat (appt.GanzKurz)
//    dl = 0
  }
}

func (x *appointments) HasWord() bool {
  for i := uint(0); i < nAppointments; i++ {
    if x.appt[i].HasWord() {
      return true
    }
  }
  return false
}

func (x *appointments) num() uint {
  n := uint(0)
  for i := uint(0); i < nAppointments; i++ {
    if ! x.appt[i].Empty() {
      n++
    }
  }
  return n
}

var
  set attr.AttrSet = attr.NewSet()

func (x *appointments) Write (l, c uint) {
  switch actualFormat {
  case day.Daily, day.Weekly:
    for i := uint(0); i < nAppointments; i++ {
      x.appt[i].Write (l + dl * i, c)
    }
  case day.Monthly:
    set.Clr()
    for i := uint(0); i < nAppointments; i++ {
      set.Ins (x.appt[i].Attrib())
    }
    set.Write (l + 1, c, x.HasWord())
  }
}

func (x *appointments) SetFont (f font.Font) {
// dummy
}

func (x *appointments) Print (l, c uint) {
  switch actualFormat {
  case day.Daily:
    for i := uint(0); i < nAppointments; i++ {
      x.appt[i].Print (l + i, c)
    }
  }
}

func (x *appointments) postEdit (l, c uint) {
  sort.Slice (x.appt, func (i, j int) bool { return x.appt[i].Less (x.appt[j]) })
  x.Write (l, c)
}

func (x *appointments) Edit (l, c uint) {
  first := true
  if actualFormat != day.Daily { return }
  index := uint(0)
  defer x.postEdit (l, c)
  for {
    x.Write (l, c)
    x.appt[index].Edit (l + dl * index, c)
    C, D := appt.PostEdit()
    switch C {
    case kbd.Esc:
      return
    case kbd.Enter:
      switch D {
      case 0, 1:
        if index + 1 < nAppointments - 1 {
          index++
        } else {
          return
        }
      default:
        return
      }
    case kbd.Up, kbd.PgUp:
      if index > 0 {
        index--
        break
      }
    case kbd.Down, kbd.PgDown:
      if index + 1 < nAppointments - 1 {
        index++
        break
      }
    case kbd.Pos1:
      index = 0
    case kbd.End:
      index = nAppointments - 1
      for x.appt[index].Empty() && index > 0 {
        index--
      }
    case kbd.Del:
      trash.Push (x.appt[index])
      x.appt[index].Clr()
    case kbd.Ins:
      n := x.num()
      if ! trash.Empty() && n < nAppointments - 1 {
        x.appt[n - 1] = trash.Pop().(appt.Appointment)
      }
    case kbd.Help:
      Help()
    case kbd.Search, kbd.Act, kbd.Cfg:
      // ??
    case kbd.Cut:
      if first {
        first = false
        clipboard.Clr()
      }
      clipboard.Ins (x.appt[index])
      x.appt[index].Clr()
    case kbd.Copy:
      clipboard.Ins (x.appt[index])
    case kbd.Paste:
      if ! clipboard.Empty() {
        n := x.num()
        // alle aus clipboard an x.appt anhängen, solange das geht
        for i := uint(0); i < clipboard.Num() && n + i < nAppointments; i++ {
          clipboard.Seek (i)
          x.appt[n + i] = clipboard.Get().(appt.Appointment)
        }
      }
    case kbd.Print:
      return
    }
  }
}

func (x *appointments) Codelen() uint {
  return nAppointments * x.appt[0].Codelen()
}

func (x *appointments) Encode() Stream {
  b := make (Stream, x.Codelen())
  c := x.appt[0].Codelen()
  a := uint(0)
  for i := uint(0); i < nAppointments; i++ {
    copy (b[a:a+c], x.appt[i].Encode())
    a += c
  }
  return b
}

func (x *appointments) Decode (b Stream) {
  c := x.appt[0].Codelen()
  a := uint(0)
  for i := uint(0); i < nAppointments; i++ {
    x.appt[i].Decode (b[a:a+c])
    a += c
  }
}
