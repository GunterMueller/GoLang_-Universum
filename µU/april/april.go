package april

// (c) Christian Maurer   v. 240401 - license see µU.go

import (
  "strconv"
  "µU/time"
  "µU/str"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/pseq"
  "µU/day"
)
const (
  f = "/tmp/april"
  t = ".tmp"
  doof = "d(o,o)f"
)
var (
  x0, y0, y1 uint
  ok bool
  doofcol = col.Pink()
  b = box.New()
  v day.Calendarday = day.New()
)

func rot (n int) { var b byte
  for i := 0; ! ok; i++ {
    switch i % 4 {
    case 0:
      b = '|'
    case 1:
      b = '/'
    case 2:
      b = '-'
    case 3:
      b = '\\'
    }
    scr.Write1 (b, y1, x0 + 64); time.Msleep (50)
  }
}

func doodle (c col.Colour, n int) {
  scr.ScrColourB (c); scr.Cls(); time.Msleep(50)
  scr.ScrColourB (col.Blue()); scr.Cls(); scr.ScrColourF (col.LightGreen())
  b.Wd (70)
  b.Colours (col.LightGreen(), col.Blue())
  b.Write ("The µU-tool to remove " + doof + " is going to be executed, i.e.", y0, x0)
  b.Write ("your disk will be completely reformatted, one moment please ...", y0 + 2, x0)
  const m = 1<<16
  x := str.New (m)
  ok = false
  for i := 0; i < 10 * n; i++ { if i == 0 { go rot (n) }
    file := pseq.New (x)
    file.Name (f + strconv.Itoa(i) + t)
    file.Clr()
    file.Ins (x)
    file.Fin()
  }
  ok = true
  scr.ColourF (col.FlashWhite())
}

func first() {
  scr.ScrColourB (col.Blue()); scr.Cls()
  v.Set (1, 4, 2024)
  v.SetFormat (day.Yymmdd)
  x0, y0 = scr.NColumns() / 2 - 35, scr.NLines() / 2 - 1
  y1 = y0 + 2
  scr.MousePointer (false)
  b.Write ("Found evil software: " + doof, y0, x0)
  time.Sleep (5)
  doodle (col.FlashWhite(), 2); doodle (col.LightYellow(), 5)
  doodle (col.Yellow(), 3); doodle (col.LightOrange(), 5)
  doodle (col.Orange(), 8); doodle (col.LightRed(), 3)
  s := str.New (70); b.Write (s, y0, x0); b.Write (s, y1, x0)
  b.Colours (col.LightGreen(), col.Blue())
  b.Write ("The µU-tool has removed " + doof + " - your disk is reformatted :-)", y0, x0)
  b.ColourF (col.FlashWhite())
  b.Write ("Please install Linux, TeX, Go and µU completely new !", y1, x0)
  time.Sleep(30)
}
