package gra

// (c) Christian Maurer   v. 220702 - license see nU.go

import . "nU/obj"

/*    vertex           neighbour                        neighbour            vertex
   ___________                                                            ___________ 
  /           \         /----------------------------------------------->/           \
  |    any    |        /                                                 |    any    |
  |___________|<--------------------------------------------------\      |___________|
  |           |      /                                             \     |           |
  |   nbPtr---|-----------\                                  /-----------|---nbPtr   |
  |___________|    /       |                                |        \   |___________|
  |           |   |        |                                |         |  |           |
  |   bool    |   |        v              edge              V         |  |   bool    |
  |___________|   |   ___________      __________      ___________    |  |___________|
  |           |   |  /           \    /          \    /           \   |  |           |
  |   dist    |   |  | edgePtr---|--->|   any    |<---|--edgePtr  |   |  |   dist    |
  |_____ _____|   |  |___________|    |__________|    |___________|   |  |_____ _____|
  |           |   |  |           |    |          |    |           |   |  |           |
  |predecessor|<-----|---from    |<---|--nbPtr0  |    |   from----|----->|predecessor|
  |___________|   |  |___________|    |__________|    |___________|   |  |___________|
  |           |   \  |           |    |          |    |           |   |  |           |
  |   repr    |    --|----to     |    |  nbPtr1 -|--->|    to-----|--/   |   repr    |
  |___________|      |___________|    |__________|    |___________|      |___________|
  |           |      |           |    |          |    |           |      |           |
  |   nextV---|->    | outgoing  |    |   bool   |    | outgoing  |      |   nextV---|->
  |___________|      |___________|    |__________|    |___________|      |___________|
  |           |      |           |    |          |    |           |      |           |
<-|---prevV   |      |  nextNb---|->  |  nextE---|->  |  nextNb---|->  <-|---prevV   |
  \___________/      |___________|    |__________|    |___________|      \___________/
                     |           |    |          |    |           |
                   <-|---prevNb  |  <-|---prevE  |  <-|---prevNb  |
                     \___________/    \__________/    \___________/

The vertices of a graph are represented by structs,
whose field "any" represents the "real" vertex.
All vertices are connected in a doubly linked list with anchor cell,
that can be traversed to execute some operation on all vertices of the graph.

The edges are also represented by structs,
whose field "any" is a variable of a type that implements Valuator.
Also all edges are connected in a doubly linked list with anchor cell.

For a vertex v one finds all outgoing and incoming edges
with the help of a further doubly linked ringlist of neighbour(hoodrelation)s
  nb1 = v.nbPtr, nb2 = v.nbPtr.nextNb, nb3 = v.nbPtr.nextNb.nextNb etc.
by the links outgoing from the nbi (i = 1, 2, 3, ...)
  nb1.edgePtr, nb2.edgePtr, nb3.edgePtr etc.
In directed graphs the edges outgoing from a vertex are exactly those ones
in the neighbourlist, for which outgoing == true.

For an edge e one finds its two vertices by the links
  e.nbPtr0.from = e.nbPtr1.to und e.nbPtr0.to = e.nbPtr1.from. */

type vertex struct {
  any "content of the vertex"
  nbPtr *neighbour
  bool "marked"
  dist uint32
  nextV, prevV *vertex
}
type vCell struct {
  vPtr *vertex
  next *vCell
}
type edge struct {
  any "attribute of the edge"
  nbPtr0, nbPtr1 *neighbour
  bool "marked"
  nextE, prevE *edge
}
type neighbour struct {
  edgePtr *edge
  from, to *vertex
  outgoing bool
  nextNb, prevNb *neighbour
}
type graph struct {
  bool "directed"
  nVertices, nEdges uint32
  vAnchor, colocal, local *vertex
  eAnchor *edge
  write CondOp
  write2 CondOp2
}
func newVertex (a any) *vertex {
  v := new(vertex)
  v.any = Clone(a)
  v.nextV, v.prevV = v, v
  return v
}

func new_(d bool, v, e any) Graph {
  CheckAtomicOrObject(v)
  x := new (graph)
  x.bool = d
  x.vAnchor = newVertex(v)
  if e == nil {
    e = uint(1)
  }
  CheckUintOrValuator (e)
  x.eAnchor = newEdge (e)
  x.colocal, x.local = x.vAnchor, x.vAnchor
  x.write = CondIgnore
  x.write2 = CondIgnore2
  return x
}

func (x *graph) imp (Y any) *graph {
  y, ok := Y.(*graph)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}
