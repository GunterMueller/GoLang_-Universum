package lan

// (c) murus.org  v. 170101 - license see murus.go

import (
  "murus/ker"
  . "murus/obj"
  "murus/bnat"
  "murus/node"
  "murus/gra"
)
type
  localAreaNetwork struct {
                 directed bool
                      net gra.Graph
                      nat []bnat.Natural
                     star []gra.Graph
                 diameter uint
//                distances []uint
                          }

func new_(as []Any, d bool, ns [][]uint) LocalAreaNetwork {
  x := new (localAreaNetwork)
  n := uint(len(as))
  if n != uint(len(ns)) { ker.Oops() }
  x.directed = d
  x.net = gra.New (x.directed, as[0], nil)
  x.net.Define (as, ns)
  x.nat = make([]bnat.Natural, n)
  x.star = make([]gra.Graph, n)
  for i, a := range as {
    x.star[i] = x.net.Star (a)
    x.nat[i] = a.(node.Node).Content().(bnat.Natural)
  }
//  x.distances = make([]uint, 0)
  return x
}

func (x *localAreaNetwork) Number (i uint) uint {
  if i >= x.net.Num() { ker.Oops() }
  return x.nat[i].Val()
}

func (x *localAreaNetwork) SetDiameter (d uint) {
  x.diameter = d
}

func (x *localAreaNetwork) Diameter() uint {
  return x.diameter
}

func (x *localAreaNetwork) Net() gra.Graph {
  return x.net.Clone().(gra.Graph)
}

func (x *localAreaNetwork) Star (i uint) gra.Graph {
  if i >= x.net.Num() { ker.Oops() }
  return x.star[i]
}

/*
func (x *localAreaNetwork) SetDistances (s []uint) {
  x.distances = make([]uint, len(s))
  copy(x.distances, s)
}

func (x *localAreaNetwork) Distances() []uint {
  ds := make([]uint, len(x.distances))
  copy(ds, x.distances)
  return ds
}
*/
