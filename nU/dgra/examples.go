package dgra

// (c) Christian Maurer   v. 171125 - license see nU.go

func g4 (i uint) DistributedGraph {
/*
         0 --- 1
         |   /
         | /
         2 --- 3
*/
  e := [][]uint { []uint { 1, 2 },
                  []uint { 0, 2 },
                  []uint { 0, 1, 3 },
                  []uint { 2 } }
  return newg (false, e, 2, i)
}

func g4flat (i uint) DistributedGraph {
  e := [][]uint { []uint { 1 },
                  []uint { 0, 2 },
                  []uint { 1, 3 },
                  []uint { 2 } }
  return newg (false, e, 3, i)
}


func g4full (i uint) DistributedGraph {
  e := [][]uint { []uint { 1, 2, 3 },
                  []uint { 0, 2, 3 },
                  []uint { 0, 1, 3 },
                  []uint { 0, 1, 2 } }
  return newg (false, e, 1, i)
}

func g8 (i uint) DistributedGraph {
/*
       0 ----- 1 ----- 2
     /           \
   3 ----- 4 ----- 5
     \   /       /
       6 ----- 7 
*/
  e := [][]uint { []uint { 1, 3 },
                  []uint { 0, 2, 5 },
                  []uint { 1 },
                  []uint { 0, 4, 6 },
                  []uint { 3, 5, 6 },
                  []uint { 1, 4, 7 },
                  []uint { 3, 4, 7 },
                  []uint { 5, 6 } }
  return newg (false, e, 4, i)
}

func g8dir (i uint) DistributedGraph {
/*
       0 ------> 1
      * \       * \
     /   \     /   \
    /     *   /     *
   2        3        4 ------> 5
    \         \     *
     \         \   /
      *         * /
       6 ------> 7
*/
  e := [][]uint { []uint { 1, 3 },
                  []uint { 4 },
                  []uint { 5 },
                  []uint { 0, 6 },
                  []uint { 1, 7 },
                  []uint { 5 },
                  []uint { },
                  []uint { 7 } }
  return newg (true, e, 4, i)
}

func g8cyc (i uint) DistributedGraph {
  e := [][]uint { []uint { 1, 2 },
                  []uint { 3, 4 },
                  []uint { 5 },
                  []uint { 4, 5 },
                  []uint { 2, 6 },
                  []uint { 0, 6 },
                  []uint { 1, 7 },
                  []uint { 4 } }
  return newg (true, e, 4, i)
}

func g8ring (i uint) DistributedGraph {
  e := [][]uint { []uint { 2, 7 },
                  []uint { 3, 5 },
                  []uint { 0, 5 },
                  []uint { 1, 6 },
                  []uint { 6, 7 },
                  []uint { 1, 2 },
                  []uint { 3, 4 },
                  []uint { 0, 4 } }
  return newg (false, e, 4, i)
}

func g8ringdir (i uint) DistributedGraph {
  e := [][]uint { []uint { 7 },
                  []uint { 5 },
                  []uint { 0 },
                  []uint { 1 },
                  []uint { 6 },
                  []uint { 2 },
                  []uint { 3 },
                  []uint { 4 } }
  return newg (true, e, 7, i)
}

func g8full (i uint) DistributedGraph {
  e := make ([][]uint, 8)
  for j := uint(0); j < 8; j++ {
    e[j] = make ([]uint, 0)
    for k:= uint(0); k < 8; k++ {
      if k != j { e[j] = append (e[j], k) }
    }
  }
  return newg (false, e, 1, i)
}

func g12 (i uint) DistributedGraph {
/*
        7 ------- 11 ------ 8
      / | \     /    \       \
    /   |   \ /         \      \
   3    6    4    10 --- 0 ---- 2
    \   |      \   |   /
      \ |        \ | /
        9 -------- 1 ---- 5
*/
  e := [][]uint { []uint { 1, 2, 10, 11 },
                  []uint { 0, 4, 5, 9, 10 },
                  []uint { 0, 8 },
                  []uint { 7, 9 },
                  []uint { 1, 7, 11 },
                  []uint { 1 },
                  []uint { 7, 9 },
                  []uint { 3, 4, 6, 11 },
                  []uint { 2, 11 },
                  []uint { 1, 3, 6 },
                  []uint { 0, 1 },
                  []uint { 0, 4, 7, 8 } }
  return newg (false, e, 4, i)
}

// Returns the number of a in the ring
func num (s []uint, a uint) uint {
  n, k := uint(len(s)), uint(0)
  for i := uint(0); i < n; i++ {
    if s[i] == a {
      k = i
      break
    }
  }
  return k
}

func g12ringdir (i uint) DistributedGraph {
  m := uint(12)
  s := []uint { 0,  4, 10,  6,  1, 11,  3,  8,  5,  2,  9,  7 }
  e := make([][]uint, m)
  for i := uint(0); i < m; i++ {
    e[i] = []uint { s[(num (s, i) + 1) % m] }
  }
  return newg (true, e, 6, i)
}

func g12full (i uint) DistributedGraph {
  e := make ([][]uint, 12)
  for j:= uint(0); j < 12; j++ {
    e[j] = make ([]uint, 0)
    for k:= uint(0); k < 12; k++ {
      if k != j { e[j] = append (e[j], k) }
    }
  }
  return newg (false, e, 1, i)
}
