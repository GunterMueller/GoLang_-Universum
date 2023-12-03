package dgra

// (c) Christian Maurer   v. 231110 - license see µU.go

import
  "µU/gra"

func (x *distributedGraph) Esel() [][]uint {
  return x.esel
}

func g3 (i uint) DistributedGraph {
  return new_(gra.G3 (i))
}

func g3dir (i uint) DistributedGraph {
  return new_(gra.G3dir (i))
}

func g4 (i uint) DistributedGraph {
  return new_(gra.G4 (i))
}

func g4flat (i uint) DistributedGraph {
  return new_(gra.G4flat (i))
}

func g4full (i uint) DistributedGraph {
  return new_(gra.G4full (i))
}

func g4star (i uint) DistributedGraph {
  return new_(gra.G4star (i))
}

func g5 (i uint) DistributedGraph {
  return new_(gra.G5 (i))
}

func g5ring (i uint) DistributedGraph {
  return new_(gra.G5ring (i))
}

func g5ringdir (i uint) DistributedGraph {
  return new_(gra.G5ringdir (i))
}

func g5full (i uint) DistributedGraph {
  return new_(gra.G5full (i))
}

func g6 (i uint) DistributedGraph {
  return new_(gra.G6 (i))
}

func g6full (i uint) DistributedGraph {
  return new_(gra.G6full (i))
}

func g8 (i uint) DistributedGraph {
  return new_(gra.G8 (i))
}

func g8a (i uint) DistributedGraph {
  return new_(gra.G8 (i))
}

func g8dir (i uint) DistributedGraph {
  return new_(gra.G8dir (i))
}

func g8cyc (i uint) DistributedGraph {
  return new_(gra.G8cyc (i))
}

func g8ring (i uint) DistributedGraph {
  return new_(gra.G8ring (i))
}

func g8ringdir (i uint) DistributedGraph {
  return new_(gra.G8ringdir (i))
}

func g8full (i uint) DistributedGraph {
  return new_(gra.G8full (i))
}

func g9dir (i uint) DistributedGraph {
  return new_(gra.G9dir (i))
}

func g9dsdir (i uint) DistributedGraph {
  return new_(gra.G9dsdir (i))
}

func g9a (i uint) DistributedGraph {
  return new_(gra.G9a (i))
}

func g9b (i uint) DistributedGraph {
  return new_(gra.G9b (i))
}

func g10 (i uint) DistributedGraph {
  return new_(gra.G10 (i))
}

func g12 (i uint) DistributedGraph {
  return new_(gra.G12 (i))
}

func g12ringdir (i uint) DistributedGraph {
  return new_(gra.G12ringdir (i))
}

func g12full (i uint) DistributedGraph {
  return new_(gra.G12full (i))
}

func g16 (i uint) DistributedGraph {
  return new_(gra.G16 (i))
}

func g16dir (i uint) DistributedGraph {
  return new_(gra.G16dir (i))
}

func g16ring (i uint) DistributedGraph {
  return new_(gra.G16ring (i))
}

func g16ringdir (i uint) DistributedGraph {
  return new_(gra.G16ringdir (i))
}

func g16full (i uint) DistributedGraph {
  return new_(gra.G16full (i))
}
