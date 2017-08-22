package main

// (c) murus.org  v. 140615 - license see murus.go

import (
  "murus/mode"
  "murus/col"
  "murus/kbd"
  "murus/scr"; "murus/errh"
  "murus/prt"
  . "murus/day"
)

func main() {
/*
  scr.(0, 0, mode.TXT)
  scr.New (32, 32, TXT)
  scr.New (192, 144, TXT)
*/
  scr.New (32, 32, mode.TXT); defer scr.Fin()
  var today, birthday = New(), New()
  today.Actualize()
  birthday.Colours (col.LightWhite, col.Blue)
  scr.ColourF (col.Yellow)
  scr.Write ("Ihr Geburtsdatum:", 12, 0)
  birthday.Edit (12, 18)
  if birthday.Empty() {
    birthday.Actualize()
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
