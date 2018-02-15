package dgra

// (c) Christian Maurer   v. 171228 - license see nU.go

func g8 (i uint) DistributedGraph {
/*     0 --- 1 --- 2
      /       \
     /         \
    3 --- 4 --- 5
     \   /     /
      \ /     /
       6 --- 7      */
  l := []uint { 0, 0, 0, 3, 3, 3, 6, 6 }
  c := []uint { 4,10,16, 1, 7,13, 4,10 }
  e := [][]uint { []uint { 1, 3 },
                  []uint { 0, 2, 5 },
                  []uint { 1 },
                  []uint { 0, 4, 6 },
                  []uint { 3, 5, 6 },
                  []uint { 1, 4, 7 },
                  []uint { 3, 4, 7 },
                  []uint { 5, 6 } }
/*
  j := "jupiter"; h := []string { j, "uranus", j, j, "saturn", j, j, j }
  return newg (false, l, c, e, h, 4, i)
*/
  return newg (false, l, c, e, nil, 4, i)
}

func g8dr (i uint) DistributedGraph {
/*      6 ---- 4 ---- 7
       /               \
      /                 \
     3                   0
      \                 /
       \               /
        1 ---- 5 ---- 2      */
  l := []uint {  3, 6,  6, 3,  0,  6, 0,  0 }
  c := []uint { 21, 4, 18, 1, 11, 11, 4, 18 }
  e := [][]uint { []uint { 7 },
                  []uint { 5 },
                  []uint { 0 },
                  []uint { 1 },
                  []uint { 6 },
                  []uint { 2 },
                  []uint { 3 },
                  []uint { 4 } }
  return newg (true, l, c, e, nil, 4, i)
}

func g12 (i uint) DistributedGraph {
/*      0 --- 1 --- 2
       /|\   / \     \
      / | \ /   \     \
     3  4  5  6--7 --- 8
      \ |   \ | /
       \|    \|/
        9 --- 10 -- 11      */
  l := []uint { 0, 0, 0, 3, 3, 3, 3, 3, 3, 6, 6, 6 }
  c := []uint { 5,11,17, 2, 5, 8,11,14,20, 5,11,17 }
  e := [][]uint { []uint { 1, 3, 4, 5 },
                  []uint { 2, 5, 7 },
                  []uint { 8 },
                  []uint { 9 },
                  []uint { 9 },
                  []uint { 10 },
                  []uint { 7, 10 },
                  []uint { 8, 10 },
                  []uint { },
                  []uint { 10 },
                  []uint { 11 },
                  []uint { } }
  return newg (false, l, c, e, nil, 4, i)
}

func g12dr (i uint) DistributedGraph {
/*      6----10---4----9----7
       /                     \
      /                       \
     3                         0
      \                       /
       \                     /
        11---1----5----8----2     */
  l := []uint {  3, 6,  6, 3,  0,  6, 0,  0,  6,  0, 0, 6 }
  c := []uint { 27, 9, 24, 1, 14, 14, 4, 24, 19, 19, 9, 4 }
  e := [][]uint { []uint { 7 },
                  []uint { 5 },
                  []uint { 0 },
                  []uint { 11 },
                  []uint { 10 },
                  []uint { 8 },
                  []uint { 3 },
                  []uint { 9 },
                  []uint { 2 },
                  []uint { 4 },
                  []uint { 6 },
                  []uint { 1 } }
  return newg (true, l, c, e, nil, 6, i)
}
