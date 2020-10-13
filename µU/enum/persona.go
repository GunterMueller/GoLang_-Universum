package enum

// (c) Christian Maurer   v. 201011 - license see µU.go

const (
  UndefPersona = uint8(iota)
  Erste
  Zweite
  Dritte
  NPersonae
)
var
  lPersona = []string {"", "1.", "2.", "3."}

func init() {
  l[Persona] = lPersona
  s[Persona] = lPersona
}
