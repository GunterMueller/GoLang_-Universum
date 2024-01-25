package dgra

// (c) Christian Maurer   v. 231221 - license see nU.go

import
  "nU/gra"
type
  DistributedGraph interface {

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
  ParentChildren() string

  HeartbeatMatrix()
  HeartbeatMatrix1()
  HeartbeatGraph()
  HeartbeatGraph1()

  Dfs() // depth first seach showing discover and finish times
  Dfs1() // depth first seach showing the DFS-tree
  Dfsfm() // depth first search without visit phase showing the DFS-tree, with far monitors
  Awerbuch() // simplified DFS-algorithm of Awerbuch, with far monitors
  Awerbuch1() // simplified DFS-algorithm of Awerbuch showing the DFS-tree, with far monitors
  HelaryRaynal() // experimental
  Ring() // construction of a ring using Dfs showing the vertices of the ring, with far monitors
  Ring1() // construction of a ring using Dfs showing the ring, with far monitors
  Bfs() // BFS-algorithm of Zhu/Cheung
  Bfsfm() // breadth first search, with far monitors
  Bfsfm1() // breadth first search showing the BFS-tree, with far monitors

  ChangRoberts()
  Peterson()
  DolevKlaweRodeh()
  HirschbergSinclair()
  Maurerfm()
  Dfselect()
  Dfselectfm()
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
func G8ring (i uint) DistributedGraph { return g8r(i) }
func G8dirring (i uint) DistributedGraph { return g8dr(i) }
func G12 (i uint) DistributedGraph { return g12(i) }
func G12dirring (i uint) DistributedGraph { return g12dr(i) }
