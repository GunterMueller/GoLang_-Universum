package main

// (c) Christian Maurer   v. 220818 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/env"
  "µU/scr"
  "µU/pseq"
  "µU/files"
  "µU/set"
  "µU/collop"
  "µU/book"
)

func sub (x, y Rotator) bool {
  return x.(book.Book).Sub (y.(book.Book))
}

func main() {
  scr.NewWH (0, 0, 80 * 8, 10 * 16); defer scr.Fin()
  files.Cds()
  file := pseq.New (byte(0))
  file.Name (env.Call() + ".dat")
  n := file.Num()
  s, b := "", book.New()
  all := set.New (b)
  for i := uint(0); i < n; i++ {
    file.Seek (i)
    a := file.Get().(byte)
    if a == byte(10) {
      if ! b.Defined (s) { ker.Panic ("kein Buch: " + s) }
      all.Ins (b)
      s = ""
    } else {
      s += string(a)
    }
  }
  collop.Operate (all, b, sub)
  tex := "\\input prv\n\\font\\rm cmr12 scaled\\magstephalf \\rm\n\\pagenumbers\n\n"
  tex += "\\centerline{(w = Wohnzimmer, p = Papas Zimmer, g = Gästezimmer,}\n"
  tex += "\\centerline{l/r = linker/rechter Bücherschrank,}\n"
  tex += "\\centerline{n = n-tes Regal von unten, v/h = vorne/hinten)}\n\n"
  all.Trav (func (a any) { tex += a.(TeXer).TeX() })
  tex += "\n\\bye"
  texfile := pseq.New(byte(0))
  texfile.Name (env.Call() + ".tex")
  texfile.Clr()
  for i := 0; i < len(tex); i++ { texfile.Ins (tex[i]) }
  texfile.Fin()
  file.Clr()
  all.Trav (func (a any) { s = a.(book.Book).String() + string(byte(10))
              for i := uint(0); i < uint(len(s)); i++ { file.Ins (s[i]) }
            })
  file.Fin()
}
