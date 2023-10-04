package main

// (c) Christian Maurer   v. 230217 - license see µU.go

import (
  "µU/mode"
  "µU/scr"
  "µU/errh"
  . "µU/li"
)

func main() {
  scr.New (0, 0, mode.XGA); defer scr.Fin()
  var m, k uint
  a := New (1<<63 - 1)
  a.Write (0, 0)
  errh.Error0 ("maxInt")

  m, k = uint(5555), uint(2345)
  a.Binom (m, k)
  scr.Cls()
  a.Write (0, 0)
  errh.Error2 ("", m, "über", k)

  a = New(0)
  m = uint(888)
  a.Fak (m)
  scr.Cls()
  a.Write (0, 0)
  errh.Error ("Fakultät von", m)

  a = New(0)
  m, k = uint(2222), uint(666)
  a.LowFak (m, k)
  scr.Cls()
  a.Write (0, 0)
  errh.Error2 ("fallende Faktorielle", m, "über", k)

  a = New(2)
  e := 6789
  b := New(e)
  a.Pow (b)
  scr.Cls()
  a.Write (0, 0)
  errh.Error ("2 hoch", uint(b.Val ()))

  a = New(0)
  m, k = uint(1111), uint(7)
  a.Stirl2 (m, k)
  scr.Cls()
  a.Write (0, 0)
  errh.Error2 ("Stirling-Koeffizient 2. Art", m, "über", k)
}
// anzahl
