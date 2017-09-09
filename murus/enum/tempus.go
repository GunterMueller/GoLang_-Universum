package enum

// (c) Christian Maurer   v. 140522 - license see murus.go

const (
  UndefTempus = uint8(iota)
  Präsens
  Imperfekt
  FuturI
  Perfekt
  Plusquamperfekt
  FuturII
  NTempora
)
var
  lTempus, sTempus []string =
  []string { "", "Präsens", "Imperfekt", "Futur I", "Perfekt", "Plusquamp.", "Futur II" },
  []string { "", "Präs.", "Impf.", "Fut.I", "Perf.", "Plusq.", "Fut.II" }

func init() {
  l[Tempus], s[Tempus] = lTempus, sTempus
}
