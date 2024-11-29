package main

// (c) Christian Maurer   v. 241108 - license see µU.go

// TODO remove bugs

import (
  "µU/kbd"
  "µU/str"
  "µU/col"
  "µU/mode"
  "µU/scr"
  "µU/box"
  "µU/N"
  "µU/files"
)

func main() {
  scr.New (0, 0, mode.XGA); defer scr.Fin()
  const x = 4 // Anzahl der Ausgabespalten
  S0 := scr.NColumns() / x
  bx := box.New()
  path := "."
  var (
    name string
    typ files.Type
  )
  extloop:
  for {
    files.Cd (path)
    path = files.ActualPath ()
    bx.Wd (scr.NColumns())
    bx.Colours (col.FlashWhite(), col.Magenta())
    str.Norm (&path, scr.NColumns()) // work around bug in µU/box
    bx.Write (path, 0, 0)
    typ = files.Dir
    a := files.Num1 (typ)
    z, s := uint(1), uint(0)
    bx.Wd (S0)
    bx.Colours (col.FlashWhite(), col.Blue())
    for i := uint(0); i < scr.NLines () - 2; i++ {
      if i < a {
        name = files.Names1 (typ)[i]
      } else {
        name = str.New (S0)
      }
      bx.Write (name, z, s)
      if z + 2 < scr.NLines () {
        z++
      } else {
        z = 1; s+= S0
      }
    }
    typ = files.File
    a = files.Num1 (typ)
    z = 1; s = S0 - 1
    bx.Colours (col.Yellow(), col.Black())
    for i := uint(0); i < x * (scr.NLines () - 2); i++ {
      if i < a {
        name = files.Names1(typ)[i]
      } else {
        name = str.New (S0)
      }
      bx.Write (name, z, s)
      if i < a {
        k := 1027 // XXX
        name = N.StringFmt (uint(k), 6, true)
        bx.Colours (col.Black(), col.FlashMagenta())
        bx.Write (name, z, s + 20)
        bx.Colours (col.FlashWhite(), col.Blue())
      }
      if z + 2 < scr.NLines () {
        z++
      } else {
        z = 1; s+= S0
      }
    }
    typ = files.Dir
    a = files.Num1 (typ)
println ("a", a)
    l, c := uint(0), uint(0)
    intloop:
    for {
      name = files.Names1(typ)[l]
      bx.Colours (col.Yellow(), col.Red())
//      bx.Write (str.New (S0), 1 + l, 0) // XXX
      bx.Write (name, 1 + l, c)
      cmd, _ := kbd.Command()
      bx.Colours (col.FlashWhite(), col.Blue())
      bx.Write (name, 1 + l, c)
      switch cmd {
      case kbd.Esc:
        break extloop
      case kbd.Back:
        name = ".."
        break intloop
      case kbd.Enter:
        break intloop
      case kbd.Down:
        if l + 1 < scr.NLines() - 1 {
          l++
        } else {
          l = 0; c += S0 - 1
        }
      case kbd.Up:
println ("l", l)
        if l > 1 {
          l--
        } else {
          l = scr.NLines() - 2
          c -= S0 - 1
        }
      }
    }
    path = name
  }
}
