package showppm

// (c) Christian Maurer   v. 220320 - license see µU.go

import (
  "µU/obj"
  "µU/scr"
  "µU/ppm"
  "µU/pseq"
)

func ppm2dat (name string) {
  i := ppm.New(); i.Load (name)
  scr.WriteImage (i.Colours(), 0, 0)
  s := scr.Screenshot (0, 0, scr.Wd(), scr.Ht())
  file := pseq.New(byte(0))
  file.Name (name + ".dat")
  file.Clr()
  for i := 0; i < len(s); i++ {
    file.Seek (uint(i))
    file.Put (s[i])
  }
  file.Fin()
}

func showImage (name string, x, y int) {
  if scr.UnderX() {
    filename := name + ".dat"
    file := pseq.New (make([]byte, pseq.Length(filename)))
    file.Name (filename)
    file.Seek (0)
    s := file.Get().(obj.Stream)
    scr.Decode (s, x, y)
//    scr.Save1()
  } else {
    i := ppm.New()
    i.Load (name)
    scr.WriteImage (i.Colours(), x, y)
  }
}
