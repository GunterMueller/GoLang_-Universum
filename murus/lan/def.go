package lan

// (c) murus.org  v. 161229 - license see murus.go

import (
  . "murus/obj"
//  "murus/host"
  "murus/gra"
)

// Pre: as is a collection of numbers from 0 to len(as)-1,
//      ns with len(ns)=len(as) is their neighbourhood relation.
//  Returns a new local area network defined by as and ns;
//  its graph Net is directed iff d == true.
func New (as []Any, d bool, ns [][]uint) LocalAreaNetwork { return new_(as,d,ns) }

type
  LocalAreaNetwork interface {

//  Returns the graph defined by x.
  Net() gra.Graph

  Number (i uint) uint

  SetDiameter (d uint) // TODO get rid of this nonsense

// Returns the diameter of the network of all involved nodes.
  Diameter() uint // TODO corresponding graph algorithm

// Pre: i < number of hosts in x.
// Returns the subnet of x consisting of its i-th host and his neighbours.
  Star (i uint) gra.Graph

/*
// TODO Spec
  SetDistances (s []uint)

// TODO Spec
  Distances() []uint
*/
}
