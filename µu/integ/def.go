package integ

// (c) Christian Maurer   v. 150215 - license see µu.go

import
  "µu/col"

// Specifications analogously to those in µu/nat.

func Wd (z int) uint { return wd(z) }

func Integer (s string) (int, bool) { return integer(s) }

func String (z int) string { return string_(z) }

func StringFmt (z int, w uint) string { return stringFmt(z,w) }

func Colours (f, b col.Colour) { Colours(f,b) }

func Write (z int, l, c uint) { }

func SetWd (w uint) { setWd(w) }

func Edit (z *int, l, c uint) { edit(z,l,c) }
