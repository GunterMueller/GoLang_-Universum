package enum

// (c) Christian Maurer   v. 201007 - license see µU.go

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
var (
  lTempus = []string {"",
                      "Präsens", "Imperfekt", "Futur I",
                      "Perfekt", "Plusquamp.", "Futur II"}
  sTempus = []string {"",
                      "Präs.", "Impf.", "Fut.I",
                      "Perf.", "Plusq.", "Fut.II"}
)

func init() {
  l[Tempus] = lTempus
  s[Tempus] = sTempus
}
