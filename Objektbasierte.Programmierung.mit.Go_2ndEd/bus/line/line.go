package line

// (c) Christian Maurer   v. 230112 - license see µU.go

import (
  "µU/mode"
  "µU/col"
  "µU/fontsize"
  "µU/scr"
)

func init() {
  scr.NewWH (0, 0, 1594, 1158)
  if scr.ActMode() <= mode.XGA {
    scr.SetFontsize (fontsize.Tiny)
  } else {
    scr.SetFontsize (fontsize.Small)
  }
  scr.ScrColourB (col.LightWhite())
  scr.Cls()
}
