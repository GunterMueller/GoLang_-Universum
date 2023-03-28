package konstanten

// (c) Christian Maurer   v. 230306 - license see µU.go

var
  Y0, H1, H2, W1, W2 int
const
  NZeilen, NSpalten = 16, 42

// Definiert die Werte Variablen unter Berücksichtigung
// der Größe des verfügbaren Bildschirms.
func Init() { init_() }
