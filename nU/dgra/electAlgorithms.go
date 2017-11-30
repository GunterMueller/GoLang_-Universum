package dgra

// (c) Christian Maurer   v. 171125 - license see nU.go

type ElectAlg byte; const (
  ChangRoberts = ElectAlg(iota)
  Peterson
  DolevKlaweRodeh
  HirschbergSinclair
//  Maurer
//  FmMaurer
  DFSE
  FmDFSE
)

func (x *distributedGraph) SetElectAlgorithm (a ElectAlg) {
  x.ElectAlg = a
}

func (x *distributedGraph) ElectAlgorithm() ElectAlg {
  return x.ElectAlg
}

func (x *distributedGraph) Leader() uint {
  switch x.ElectAlg {
  case ChangRoberts:
    x.changRoberts()
  case Peterson:
    x.peterson()
  case DolevKlaweRodeh:
    x.dolevKlaweRodeh()
  case HirschbergSinclair:
    x.hirschbergSinclair()
//  case Maurer:
//    x.maurer()
//  case FmMaurer:
//    x.fmMaurer()
  case DFSE:
    x.dfse()
  case FmDFSE:
    x.fmdfse()
  }
  return x.leader
}
