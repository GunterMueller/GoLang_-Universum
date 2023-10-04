package main

// (c) Christian Maurer   v. 230924 - license see µU.go

// Calculations with roman numbers

import (
  "µU/str"
  "µU/scr"
  "µU/errh"
  "µU/rn"
)
const (
  cs = 22
  cz = 34
  kl = "1. Zahl < 2. Zahl"
)
  
func edit (l uint) rn.RomanNatural {
  e := "röm. Zahl eingeben"
  x := rn.New0()
  errh.Hint (e)
  for {
    x.Edit (l, 0)
    if x.Undef() {
      scr.Write (str.Lat1("Wert zu groß"), l, cs)
    } else {
      scr.WriteNat (x.Val(), l, cz)
      s := "2. Zahl   ="; if l == 1 { s = "1. Zahl   =" }
      scr.Write (s, l, cs)
      break
    }
  }
  return x
}

func main() {
  scr.NewWH (0, 0, 312, 168); defer scr.Fin()
  errh.Head ("römische Zahlen")
  l, c := uint(1), uint(0)
  x := edit (l)
  l++
  y := edit (l)
  l += 2
  s := rn.New0()
  s.Sum (x, y)
  if s.Undef() {
    scr.Write (str.Lat1("Summe zu groß"), l, c + cs)
  } else {
    s.Write (l, c)
    scr.Write ("Summe     =", l, c + cs)
    scr.WriteNat (s.Val(), l, c + cz)
  }
  l++
  d := rn.New0()
  d.Diff (x, y)
  d.Write (l, c)
  if d.Undef() {
    scr.Write (kl, l, c + cs)
  } else {
    d.Write (l, c)
    scr.Write ("Differenz =", l, c + cs)
    scr.WriteNat (d.Val(), l, c + cz)
  }
  l++
  p := rn.New0()
  p.Prod (x, y)
  p.Write (l, c)
  if p.Undef() {
    scr.Write (str.Lat1("Produkt zu groß"), l, c + cs)
  } else {
    p.Write (l, c)
    scr.Write ("Produkt   =", l, c + cs)
    scr.WriteNat (p.Val(), l, c + cz)
  }
  l++
  v := rn.New0()
  v.Div (x, y)
  v.Write (l, c)
  if v.Undef() {
    scr.Write (kl, l, c + cs)
  } else {
    v.Write (l, c)
    scr.Write ("Div       =", l, c + cs)
    scr.WriteNat (v.Val(), l, c + cz)
  }
  l++
  m := rn.New0()
  m.Mod (x, y)
  m.Write (l, c)
  if m.Undef() {
    scr.Write (kl, l, c + cs)
  } else {
    m.Write (l, c)
    scr.Write ("Mod       =", l, c + cs)
    scr.WriteNat (m.Val(), l, c + cz)
  }
  errh.Error0 ("Summe, Differenz, Produkt, Div und Mod")
}
