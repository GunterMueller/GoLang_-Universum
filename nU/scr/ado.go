package scr

// (c) Christian Maurer   v. 171229 - license see nU.go

import
  "nU/col"

func Wd() uint { return a.Wd() }
func Ht() uint { return a.Ht() }

func NLines() uint { return a.NLines() }
func NColumns() uint { return a.NColumns() }

func Cls() { a.Cls() }

func Colours (f, b col.Colour) { a.Colours(f,b) }
func ColourF (f col.Colour) { a.ColourF(f) }
func ColourB (b col.Colour) { a.ColourB(b) }

func Warp (l, c uint) { a.Warp(l,c) }
func Switch (on bool) { a.Switch(on) }
func Fin() { a.Fin() }

func Write1 (b byte, l, c uint) { a.Write1(b,l,c) }
func Write (s string, l, c uint) { a.Write(s,l,c) }
func WriteNat (n, l, c uint) { a.WriteNat(n,l,c) }

func Line (l, c, l1, c1 uint) { a.Line(l,c,l1,c1) }
func Circle (l, c, r uint) { a.Circle(l,c,r) }

func Lock() { a.Lock() }
func Unlock() { a.Unlock() }
