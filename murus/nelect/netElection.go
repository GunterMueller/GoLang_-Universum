package nelect

// (c) murus.org  v. 170111 - license see murus.go
//
// >>> experimental version, still some stuff TODO

import (
  "murus/ker"
  . "murus/obj" // maurer
  "murus/errh"
//  "murus/nat"
  "murus/node"
  "murus/gra"
  "murus/host"
  "murus/bnat"
  "murus/nchan"
  . "murus/nelect/internal"
)
type
  netElection struct {
                     uint "number of hosts involved"
                   n uint
               graph gra.Graph
                  me,
                root uint
              meHost,
              hR, hL host.Host
//                  nb []node.Node
              meNode,
            nbR, nbL node.Node
            nnR, nnL bnat.Natural
            nrR, nrL uint
                     bnat.Natural "me"
                 chR,
                 chL nchan.NetChannel
                     bool "demo mode"
                     }

func new_(g gra.Graph, h []host.Host, n, root uint /* , ds []uint */) NetElection {
  x := new(netElection)
  x.uint = n
  x.graph = g.Clone().(gra.Graph)
  x.n = x.graph.NumLoc()
  x.root = root
  x.meHost = host.Localhost()
  x.meNode = x.graph.Get().(node.Node)
  if x.graph.CoNeighbour(0) == nil { ker.Oops() } // x.graph is not directed
/*
  x.nb = make([]node.Node, x.uint)
  for i:= uint(0); i < x.uint; i++ { x.nb[i] = make([]node.Node, x.uint) }
*/
  x.nbR, x.nbL = x.graph.Neighbour(0).(node.Node), x.graph.CoNeighbour(0).(node.Node)
  x.nnR, x.nnL = x.nbR.Content().(bnat.Natural), x.nbL.Content().(bnat.Natural)
  x.nrR, x.nrL = x.nnR.Val(), x.nnL.Val()
//  println("nrR =", x.nrR, ";  nrL =", x.nrL)
  x.Natural = x.meNode.Content().(bnat.Natural)
  x.hR, x.hL = h[x.nrR], h[x.nrL]
  x.me = x.Natural.Val()
  if ! x.meHost.Eq(h[x.me]) { ker.Oops() }
//  if x.graph.CoNeighbour(0) == nil { ker.Oops() }
  return x
}

func (x *netElection) Demo() {
  x.bool = true
}

func (x *netElection) Do (alg Algorithm) {
/*
  for process id:

          :id           :nrR
  ... ----------> id ----------> nrR
              chL    chR
  nhL            Host            nhR
*/
  const p = 50000
  pR, pL := p + uint16(x.nrR), p + uint16(x.me)
// keep the pairing of the constructors to avoid deadlocks:2
  var nix Any = 0
  if alg == Maurer { nix = nil }
  if alg == HirschbergSinclair { nix = NewMsg() }
  if x.nrL < x.nrR {
    x.chR = nchan.New (nix, x.me, x.nrR, x.hR, pR)
    x.chL = nchan.New (nix, x.me, x.nrL, x.hL, pL)
  } else {
    x.chL = nchan.New (nix, x.me, x.nrL, x.hL, pL)
    x.chR = nchan.New (nix, x.me, x.nrR, x.hR, pR)
  }
  x.graph.Locate (true)
  x.Write()
  var winner uint
  switch alg {
  case Peterson:
    winner = uint(x.peterson())
  case PetersonImproved:
    winner = uint(x.petersonImproved())
  case ChangRoberts:
    winner = x.changRoberts()
  case HirschbergSinclair:
    winner = x.hirschbergSinclair()
  case Maurer:
    winner = x.maurer (x.graph, x.root)
    if x.graph.ExPred (func (a Any) bool { // just to mark the winner
                         return a.(node.Node).Content().(bnat.Natural).Val() == winner
                       }) {
      // paletti
    } else {
      ker.Oops()
    }
    x.Write()
  }
  errh.Error ("The winner is", winner)
}

func (x *netElection) Graph() gra.Graph {
  return x.graph
}

func (x *netElection) Write() {
  x.graph.Trav3CondDir (node.O, node.O3dir)
}
