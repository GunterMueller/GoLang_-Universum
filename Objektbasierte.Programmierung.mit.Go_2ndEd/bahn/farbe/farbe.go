package farbe

// (c) Christian Maurer   v. 230107 - license see µU.go

import (
  "µU/scr"
  "µU/linewd"
  "µU/col"
  . "bahn/konstanten"
)

func init() {
  Init()
  scr.SetLinewidth (linewd.Thicker)
  Fordergrundfarbe, Hintergrundfarbe = col.LightWhite(), col.Black()
  Nichtfarbe= col.Gray()
  Freifarbe = col.Green()
  Besetztfarbe = col.Yellow()
  Zugfarbe = col.Red()
  Fahrtfarbe, Langsamfahrtfarbe, Haltfarbe =
    col.Verkehrsgrün(), col.Verkehrsgelb(), col.Verkehrsrot()
/*/
  Fordergrundfarbe, Hintergrundfarbe = col.Black(), col.LightWhite()
  scr.ScrColours (col.Black(), col.LightWhite()); scr.Cls()
  Nichtfarbe = col.LightGray()
  Freifarbe = col.Black()
  Besetztfarbe = col.DarkOrange()
/*/
}
