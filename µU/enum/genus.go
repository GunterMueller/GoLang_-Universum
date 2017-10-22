package enum

// (c) Christian Maurer   v. 140527 - license see µU.go

const (
  UndefGenus = uint8(iota)
  M
  F
  Ne
  NGenera
)
var
  lGenus, sGenus []string =
  []string { "", "masculinum", "femininum", "neutrum" },
  []string { "", "m.", "f.", "n." }

func init() {
  l[Genus], s[Genus] = lGenus, sGenus
}
