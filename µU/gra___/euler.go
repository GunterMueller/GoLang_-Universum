package gra

// (c) Christian Maurer   v. 231110 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/rand"
)

func existsnb (s []*neighbour, p Pred) (*neighbour, bool) {
  for _, a := range s {
    if p (a) {
      return a, true
    }
  }
  return nil, false
}

func notTraversedNeighbour (v *vertex) *neighbour {
  for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
    if n.outgoing && ! n.edgePtr.bool {
      return n
    }
  }
  return nil
}

func notTraversed (a any) bool {
  return notTraversedNeighbour (a.(*neighbour).from) != nil
}

func (x *graph) Euler() bool {
  if ! x.totallyConnected() {
    return false // TODO Fleury's algorithm
  }
  p := x.colocal
  a := x.local
  x.colocal = x.vAnchor
  x.local = x.vAnchor
  e := uint(0)
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
// check for existence of Euler cycles (iff graph
// not directed:
//   if each vertex has an even number of neighbours,
// directed:
//   if at any vertex the number of outgoing edges is equal to the number of incoming ones)
// or of Euler paths (iff graph
// not directed:
//   if there are exactly two vertices with odd number of neighbours,
// directed:
//   if exactly one vertex has one more outgoing than incoming edges
//   and exactly one vertex has one more incoming than outgoint edges)
    z := uint(0)
    z1 := uint(0)
    nb := v.nbPtr.nextNb
    for nb != v.nbPtr {
      if nb.outgoing {
        z++
      } else {
        z1++
      }
      nb = nb.nextNb
    }
    if x.bool {
      if z == z1 + 1 {
        if x.colocal == x.vAnchor {
          x.colocal = v
          e++
        } else {
          x.colocal = p
          x.local = a
          return false
        }
      } else if z1 == z + 1 {
        if x.local == x.vAnchor {
          x.local = v
          e++
        } else {
          x.colocal = p
          x.local = a
          return false
        }
      }
    } else { // ! x.bool
      if z % 2 == 1 {
        if x.colocal == x.vAnchor {
          x.colocal = v
        } else if x.local == x.vAnchor {
          x.local = v
        } else {
          x.colocal = p
          x.local = a
          return false
        }
        e++
      }
    }
  }
  switch e {
  case 0: // Euler cycle with random starting vertex
    x.colocal = x.vAnchor.nextV
    n := rand.Natural (uint(x.nVertices))
    for n > 0 {
      x.colocal = x.colocal.nextV
      n--
    }
    x.local = x.colocal
  case 1:
    x.colocal = p
    x.local = a
    return false
  case 2: // Euler path from colocal to local vertex
    ;
  default:
    ker.Shit()
  }
  x.ClrMarked()
  x.eulerPath = nil
  x.colocal.bool = true
  v := x.colocal
  v.bool = true
//  for j := 0; j <= 9; j** { for a := false TO true { writeE (E.any, a); ker.Msleep (100) } }
// attempt, to find an Euler path/cycle "by good luck":
  var nb *neighbour
  for {
    nb = notTraversedNeighbour (v)
    if nb == nil { ker.Oops() }
    // writeE (N.edgePtr.any, true)
    //  for j := 0; j <= 9; j++ { for a := false; a <= true; a++ { writVe (N.to.any, a); ker.Msleep (100) } } };
    nb.edgePtr.bool = true
    v = nb.to
    v.bool = true
    x.eulerPath = append (x.eulerPath, nb)
    if v == x.local { break }
  }
// errh.Error0("erster Wegabschnitt gefunden");
// as long there are edges not yet traversed,
// look for vertices in the Euler path, from which such edges go out,
// and find more cycles starting there and insert them into the Euler path:
  for {
    nb, ok := existsnb (x.eulerPath, notTraversed)
    if ! ok { break }
    // for j := 0; j <= 9; j++ { for a := false; a <= true; a++ { // nonsense
    //   x.writeE (nb.edgePtr.any, a); ker.Msleep (100) } }
    v = nb.from
    v1 := v
    for {
      nb = notTraversedNeighbour (v)
      if nb == nil { ker.Oops() }
    // writeE (N.edgePtr.any, true)
    // for j := 0 TO 9 { for a := false TO true { writeV (N.to.any, a); ker.Msleep (100) } }
      nb.edgePtr.bool = true
      v = nb.to
      v.bool = true
      x.eulerPath = append (x.eulerPath, nb)
      if v == v1 { break } // found one mor cycle
    // errh.Error0("weiterer Teil eines Eulerwegs gefunden")
    }
  }
  if x.demo [Euler] {
    x.writeV (x.colocal.v, true)
    wait()
    for i := uint(0); i < uint(len (x.eulerPath)); i++ {
      nb = x.eulerPath[i]
      x.writeE (nb.edgePtr.e, true)
      if nb.edgePtr.nbPtr0 == nb {
        x.writeV (nb.edgePtr.nbPtr1.from.v, true)
      } else {
        x.writeV (nb.edgePtr.nbPtr0.from.v, true)
      }
      if i + 1 < uint(len (x.eulerPath)) {
        wait()
      }
    }
  }
  return true
}
