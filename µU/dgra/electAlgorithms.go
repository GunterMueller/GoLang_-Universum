package dgra

// (c) Christian Maurer   v. 171112 - license see µU.go

type // algorithms to elect a leader in a ring
  ElectAlg byte; const (
  ChangRoberts = ElectAlg(iota)
  Peterson
  DolevKlaweRodeh
  HirschbergSinclair
  Maurer
  FmMaurer
  DFSE
  FmDFSE
)
