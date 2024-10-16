package main

// (c) Christian Maurer   v. 241005 - license see µU.go

import (
  "µU/mode"
  "µU/col"
  "µU/kbd"
  "µU/scr"
  "µU/errh"
  "µU/prt"
  . "µU/day"
)

func main() {
  scr.New (0, 0, mode.TXT); defer scr.Fin()
  scr.Name ("Wie alt bist du ?")
  var today, birthday = New(), New()
  today.Update()
  birthday.Colours (col.FlashWhite(), col.Blue())
  scr.ColourF (col.Yellow())
  scr.Write ("Ihr Geburtsdatum:", 12, 0)
  birthday.Edit (12, 18)
  if birthday.Empty() {
    birthday.Update()
  } else {
    scr.Write (" war ein ", 12, 26)
    birthday.SetFormat (WD)
    birthday.Write (12, 35)
    errh.Error ("So viele Tage sind Sie heute alt:", birthday.Distance (today))
  }
  scr.Colours (scr.ScrCols())
  scr.Cls()
  errh.Hint (" vor-/rückwärts: Pfeiltasten               fertig: Esc ")
  var ( c kbd.Comm; t uint )
  neu:= true
  loop: for {
    if neu {
      birthday.WriteYear (0, 0)
      neu = false
    }
    switch c, t = kbd.Command(); c { case kbd.Esc:
      break loop
    case kbd.Down:
      if t == 0 {
        birthday.Inc (Yearly)
      } else {
        birthday.Inc (Decadic)
      }
      neu = true
    case kbd.Up:
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
