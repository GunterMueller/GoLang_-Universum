package dgra

// (c) Christian Maurer   v. 171130 - license see µU.go

type
  ElectAlg byte; const (
  ChangRoberts = ElectAlg(iota)
  Peterson
  DolevKlaweRodeh
  HirschbergSinclair
  Maurer
  Maurerfm
  DFSelect
  DFSelectfm
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
  case Maurer:
    x.maurer()
  case Maurerfm:
    x.maurerfm()
  case DFSelect:
    x.dfselect()
  case DFSelectfm:
    x.dfselectfm()
  }
  return x.leader
}
