package main

// (c) Christian Maurer   v. 241205 - license see µU.go

import (
  "µU/mode"
  "µU/col"
  "µU/kbd"
  "µU/scr"
  "µU/errh"
  "µU/str"
  "µU/prt"
  . "µU/day"
)

func main() {
  scr.New (0, 0, mode.TXT); defer scr.Fin()
  scr.Name ("Wie alt bist Du ?")
  var today, birthday = New(), New()
  today.Update()
  birthday.Colours (col.FlashWhite(), col.Blue())
  scr.ColourF (col.Yellow())
  const l = 12
  scr.Write ("Dein Geburtsdatum:", l, 0)
  birthday.Edit (l, 19)
  if birthday.Empty() {
    birthday.Update()
  } else {
    if birthday.Day() < 10 {
      birthday.SetFormat (D_M_yyyy)
    } else {
      birthday.SetFormat (Dd_M_yyyy)
    }
    s := birthday.String()
    n := uint(len(s))
    scr.Write ("Der " + s + " war ein ", l, 0)
    birthday.SetFormat (WD)
    s = birthday.String()
    str.OffSpc1 (&s)
    scr.Write (s, l, 13 + n)
    scr.Write (".", l, 13 + n + uint(len(s)))
    errh.Errorm ("Du bist heute", birthday.Distance (today), "Tage alt.")
  }
  scr.Colours (scr.ScrCols())
  scr.Cls()
  errh.Hint (" vor-/rückwärts: Pfeiltasten               fertig: Esc ")
  var ( c kbd.Comm; t uint )
  neu := true
  loop:
  for {
    if neu {
      birthday.WriteYear (0, 0)
      neu = false
    }
    switch c, t = kbd.Command(); c {
    case kbd.Esc:
      break loop
    case kbd.Down, kbd.Right:
      if t == 0 {
        birthday.Inc (Yearly)
      } else {
        birthday.Inc (Decadic)
      }
      neu = true
    case kbd.Up, kbd.Left:
      if t == 0 {
        birthday.Dec (Yearly)
      } else {
        birthday.Dec (Decadic)
      }
      neu = true
    case kbd.Print:
      birthday.PrintYear (0, 0)
      prt.GoPrint()
    }
  }
}
