package track

// (c) Christian Maurer   v. 230309 - license see µU.go

import (
  . "µU/obj"
  "bus/line"
)
type
  Track interface { // Verbindung mit Linie und natürlicher Zahl
                    // als Wert (mittlere Fahrzeit in Minuten)
  Object
  Valuator

// x gehört zur Linie l und hat den Wert f.
  Def (l line.Line, v uint)

// Die Positionen der Endpunkte der Strecke von x
// sind durch (x, y) und (x1, y1) gegeben.
  SetPos (x, y, x1, y1 float64)

// x ist auf dem Bildschirm ausgegeben, für b == true
// in der Farbe seiner Linie, sonst in schwarz.
  Write (b bool)
}
