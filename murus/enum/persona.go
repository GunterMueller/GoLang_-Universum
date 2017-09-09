package enum

// (c) Christian Maurer   v. 170419 - license see murus.go

const (
  UndefPersona = uint8(iota)
  Erste
  Zweite
  Dritte
  NPersonae
)
var (
  lPersona, sPersona []string =
  []string { "", "1.", "2.", "3." },
  lPersona
)

func init() {
  l[Persona], s[Persona] = lPersona, sPersona
}
