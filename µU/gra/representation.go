package gra

// (c) Christian Maurer   v. 171118 - license see µU.go
//
// >>>  References:
// >>>  CLR  = Cormen, Leiserson, Rivest        (1990)
// >>>  CLRS = Cormen, Leiserson, Rivest, Stein (2001)

import (
//  "reflect"
//  "sort"
//  "µU/ker"
  . "µU/obj"
//  "µU/str"
//  "µU/rand"
//  "µU/kbd"
//  "µU/errh"
//  "µU/pseq"
//  "µU/adj"
)

/*    vertex           neighbour                        neighbour            vertex
   ___________                                                            ___________ 
  /           \         /----------------------------------------------->/           \
  |    Any    |        /                                                 |    Any    |
  |___________|<--------------------------------------------------\      |___________|
  |           |      /                                             \     |           |
  |   nbPtr---|-----------\                                  /-----------|---nbPtr   |
  |___________|    /       |                                |        \   |___________|
  |           |   |        |                                |         |  |           |
  |   bool    |   |        v              edge              V         |  |   bool    |
  |___________|   |   ___________      __________      ___________    |  |___________|
  |     |     |   |  /           \    /          \    /           \   |  |     |     |
  |dist | time|   |  | edgePtr---|--->|   Any    |<---|--edgePtr  |   |  |dist | time|
  |_____|_____|   |  |___________|    |__________|    |___________|   |  |_____|_____|
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
whose field "Any" represents the "real" vertex.
All vertices are connected in a doubly linked list with anchor cell,
that can be traversed to execute some operation on all vertices of the graph.

The edges are also represented by structs,
whose field "Any" is a variable of a type that implements Valuator.
Also all edges are connected in a doubly linked list with anchor cell.

For a vertex v one finds all outgoing and incoming edges
with the help of a further doubly linked ringlist of neighbour(hoodrelation)s
  nb1 = v.nbPtr, nb2 = v.nbPtr.nextNb, nb3 = v.nbPtr.nextNb.nextNb etc.
by the links outgoing from the nbi (i = 1, 2, 3, ...)
  nb1.edgePtr, nb2.edgePtr, nb3.edgePtr etc.
In directed graphs the edges outgoing from a vertex are exactly those ones
in the neighbourlist, for which outgoing == true.

For an edge e one finds its two vertices by the links
  e.nbPtr0.from = e.nbPtr1.to und e.nbPtr0.to = e.nbPtr1.from.

Semantics of some variables, that are "hidden" in fields of vAnchor:
  vAnchor.time0: in that the "time" is incremented for each search step
  vAnchor.acyclic: (after call of search) == true <=> graph has no cycles. */

const (
  suffix = "gra"
  inf = uint32(1<<32 - 1)
)
type (
  vertex struct {
                Any "content of the vertex"
          nbPtr *neighbour
                bool "marked"
        acyclic bool   // for the development of design patterns by clients
           dist,       // for breadth first search/Dijkstra and use in En/Decode
   time0, time1 uint32 // for applications of depth first search
    predecessor,       // for back pointers in depth first search and in ways
           repr,       // for the computation of connected components
   nextV, prevV *vertex
                }

  vCell struct {
          vPtr *vertex
          next *vCell
               }

  edge struct {
              Any "attribute of the edge"
       nbPtr0,
       nbPtr1 *neighbour
              bool "marked"
 nextE, prevE *edge
              }

  neighbour struct {
           edgePtr *edge
          from, to *vertex
          outgoing bool
    nextNb, prevNb *neighbour
                   }

  graph struct {
          name,
      filename string
          bool "directed"
     nVertices,
        nEdges uint32
       vAnchor,
       colocal,
         local *vertex
       eAnchor *edge
          path []*vertex // XXX
     eulerPath []*neighbour
          demo Demoset
        writeV,
        writeE CondOp
               }
)
type
  nSeq []*neighbour

func newVertex (a Any) *vertex {
  v := new(vertex)
  v.Any = Clone(a)
  v.time1 = inf // for applications of depth first search
  v.dist = inf
  v.repr = v
  v.nextV, v.prevV = v, v
  return v
}

func new_(d bool, v, e Any) Graph {
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
  x.writeV, x.writeE = CondIgnore, CondIgnore
  return x
}

func (x *graph) imp (Y Any) *graph {
  y, ok := Y.(*graph)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}
