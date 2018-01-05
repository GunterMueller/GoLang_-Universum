package scr

// (c) Christian Maurer   v. 171229 - license see nU.go

import "nU/col"

// Zeichenbreite/-höhe des Standardfonts
const (Wd1 = uint(8); Ht1 = uint(16)) // ggf. anpassen

type Screen interface {

// Liefert die Pixelbreite/-höhe von x.
  Wd() uint; Ht() uint

// Liefert die Anzahl der Textzeilen/-spalten von x.
  NLines() uint; NColumns() uint

// x ist gelöscht. Der Kursor von x hat die Position (0, 0).
  Cls()

// f und b ist die aktuelle Vordergrund-/Hintergrundfarbe für Schreiboperationen.
  Colours (f, b col.Colour)
  ColourF (f col.Colour); ColourB (b col.Colour)

// Der Kursor von x ist genau dann sichtbar, wenn on == true.
  Switch (on bool)

// Vor.: l < NLines, c < NColumns.
// Der Kursor von x hat die Position (Zeile, Spalte) == (l, c).
// (0, 0) ist die linke obere Ecke von x.
  Warp (l, c uint)

// Der Kursor von x hat die Position (l, 0), wobei l die
// erste Zeile ist, die nicht von Ausgaben in x benutzt wurde.
  Fin()

// Vor.: 32 <= b < 127, l < NLines, c + 1 < NColumns. 
// b ist auf x ab Position (l, c) ausgegeben.
  Write1 (b byte, l, c uint)

// Vor.: l < NLines, c + len(s) < NColumns. 
// s ist auf x ab Position (l, c) ausgegeben.
  Write (s string, l, c uint)

// Vor.: c + Anzahl der Ziffern von n < NColumns, l < NLines.
// n ist auf x ab Position (l, c) ausgegeben.
  WriteNat (n, l, c uint)

// Vor.: l, l1 < NLines, c, c1 < NColumns.
// Ein Linie aus den Zeichen "-", "|", "/" und "\"
// von (l,c) nach (l1,c1) ist auf x ausgegeben.
  Line (l, c, l1, c1 uint)

// Vor.: l, c >= r, l + r < NLines; l + c < NColumns.
// Ein Kreis aus dem Zeichen "*" mit dem Mittelpunkt (l, c)
// und dem Radius r ist auf x ausgegeben.
  Circle (l, c, r uint)

// Gewährleistet den gegenseitigen Ausschluss
// bei nebenläufigen Schreiboperationen
  Lock(); Unlock()
}

func New() Screen { return new_() }
