package dgra

// (c) Christian Maurer   v. 241030 - license see µU.go

import (
  "µU/env"
  "µU/vtx"
  "µU/edg"
  "µU/gra"
)

func newg1 (g gra.Graph, id uint) DistributedGraph {
  g.SetWrite (vtx.W, edg.W)
  g.ExVal (id)
  k := g.Num()
  g = g.Star()
  d := new_(g).(*distributedGraph)
  d.size = k
  h := make([]string, k)
  for i := uint(0); i < k; i++ {
    h[i] = env.Localhost()
  }
  d.setHosts (h)
  d.diameter = g.Diameter()
  d.Write()
  return d
}

func g3 (i uint) DistributedGraph {
  return newg1 (gra.G3(), i)
}

func g3dir (i uint) DistributedGraph {
  return newg1 (gra.G3dir(), i)
}

func g4 (i uint) DistributedGraph {
  return newg1 (gra.G4(), i)
}

func g4flat (i uint) DistributedGraph {
  return newg1 (gra.G4flat(), i)
}

func g4ring (i uint) DistributedGraph {
  return newg1 (gra.G4ring(), i)
}

func g4ringdir (i uint) DistributedGraph {
  return newg1 (gra.G4ringdir(), i)
}

func g4full (i uint) DistributedGraph {
  return newg1 (gra.G4full(), i)
}

func g4star (i uint) DistributedGraph {
  return newg1 (gra.G4star(), i)
}

func g4ds (i uint) DistributedGraph {
  return newg1 (gra.G4ds(), i)
}

func g5 (i uint) DistributedGraph {
  return newg1 (gra.G5(), i)
}

func g5ring (i uint) DistributedGraph {
  return newg1 (gra.G5ring(), i)
}

func g5ringdir (i uint) DistributedGraph {
  return newg1 (gra.G5ringdir(), i)
}

func g5full (i uint) DistributedGraph {
  return newg1 (gra.G5full(), i)
}

func g6 (i uint) DistributedGraph {
  return newg1 (gra.G6(), i)
}

func g6full (i uint) DistributedGraph {
  return newg1 (gra.G6full(), i)
}

func g8 (i uint) DistributedGraph {
  return newg1 (gra.G8(), i)
}

func g8a (i uint) DistributedGraph {
  return newg1 (gra.G8a(), i)
}

func g8dir (i uint) DistributedGraph {
  return newg1 (gra.G8dir(), i)
}

func g8ring (i uint) DistributedGraph {
  return newg1 (gra.G8ring(), i)
}

func g8ringdir (i uint) DistributedGraph {
  return newg1 (gra.G8ringdir(), i)
}

func g8ringdirord (i uint) DistributedGraph {
  return newg1 (gra.G8ringdirord(), i)
}

func g8full (i uint) DistributedGraph {
  return newg1 (gra.G8full(), i)
}

func g8ds (i uint) DistributedGraph {
  return newg1 (gra.G8ds(), i)
}

func g9a (i uint) DistributedGraph {
  return newg1 (gra.G9a(), i)
}

func g9b (i uint) DistributedGraph {
  return newg1 (gra.G9b(), i)
}

func g9dir (i uint) DistributedGraph {
  return newg1 (gra.G9dir(), i)
}

func g10 (i uint) DistributedGraph {
  return newg1 (gra.G10(), i)
}

func g12 (i uint) DistributedGraph {
  return newg1 (gra.G12(), i)
}

func g12ringdir (i uint) DistributedGraph {
  return newg1 (gra.G12ringdir(), i)
}

func g12full (i uint) DistributedGraph {
  return newg1 (gra.G12full(), i)
}

func g16 (i uint) DistributedGraph {
  return newg1 (gra.G16(), i)
}

func g16dir (i uint) DistributedGraph {
  return newg1 (gra.G16dir(), i)
}

func g16ring (i uint) DistributedGraph {
  return newg1 (gra.G16ring(), i)
}

func g16ringdir (i uint) DistributedGraph {
  return newg1 (gra.G16ringdir(), i)
}

func g16full (i uint) DistributedGraph {
  return newg1 (gra.G16full(), i)
}
