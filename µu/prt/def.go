package prt

// (c) Christian Maurer   v. 121119 - license see µu.go

import
  "µu/font"
var
//  PrintCommand string = "lp -o fit-to-page -d pomona"
  PrintCommand string = "lp"

// The actual fontsize is f.
func Switch (f font.Size) {_switch(f) }

// Returns the actual fontsize.
func Fontsize () font.Size { return size }

// Returns the number of lines per page.
func NLines () uint { return nL[size] }

// Returns the number of columns per line.
func NColumns () uint { return nC[size] }

// Secret.
func Voffset (mm uint) { voffset(mm) }

// Secret.
func Footline (s string) { footline(s) }

// Pre: l < maxL; c + 1 < maxC. f ist ein Font.
// b is n line l, column c in font f in the printer buffer.
func Print1 (b byte, l, c uint, f font.Font) { print1(b, l, c, f) }

// Pre: l < maxL, c + len(s) < maxC. f is a font.
// s is in line l from column c in font f in the printer buffer.
func Print (s string, l, c uint, f font.Font) { print(s, l, c, f) }

// Pre: TeX is installed.
// All lines of the printer buffer are printed; the printer buffer is not empty.
func GoPrint() { goPrint() }
