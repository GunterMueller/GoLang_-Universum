package dgra

import "nU/gra"

type DistributedGraph interface {

  gra.Graph

// r ist die Wurzel von x.
  SetRoot (r uint)

// Voraussetzung für alle folgenden Methoden: Die Rechner von x sind definiert.
  Me() uint
  Root() uint
  Parent() uint
  Children() string
  Time() uint
  Time1() uint

  SetHeartbeatAlgorithm (a HeartbeatAlg) // s. heartbeatAlgorithms.go
  HeartbeatAlgorithm() HeartbeatAlg
  Heartbeat()

  SetTravAlgorithm (a TravAlg) // s. travAlgorithms.go
  TravAlgorithm() TravAlg

  SetElectAlgorithm (a ElectAlg) // s. electAlgorithms.go
  ElectAlgorithm() ElectAlg
  Leader() uint
}

// Vor.: Die Werte der Kanten von g, um nchan.Port0 erhöht sind die Ports
//       für 1-1-Verbindungen zwischen den Ecke, die sie verbinden.
// Liefert einen neuen verteilten Graph mit zugrundeliegendem Graph g.
func New (g gra.Graph) DistributedGraph { return new_(g) }

// Beispiele verteilter Graphen:

// G_ liefert den Stern des verteilten Graphen, der durch g_ definiert ist,
// wobei die Ecke mit der Identität i das Zentrum ist.
func G8 (i uint) DistributedGraph { return g8(i) }
func G8dirring (i uint) DistributedGraph { return g8dr(i) }
func G12 (i uint) DistributedGraph { return g12(i) }
func G12dirring (i uint) DistributedGraph { return g12dr(i) }
