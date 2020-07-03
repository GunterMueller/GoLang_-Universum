package dgra

type ElectAlg byte; const (
  ChangRoberts = ElectAlg(iota)
  Peterson
  DolevKlaweRodeh
  HirschbergSinclair
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
  case DFSelect:
    x.dfselect()
  case DFSelectfm:
    x.dfselectfm()
  }
  return x.leader
}
