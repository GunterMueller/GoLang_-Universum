package lans

// (c) murus.org  v. 170111 - license see murus.go
//
// >>> Modify and/or extend according to your needs !

import (
  . "murus/obj"
  "murus/col"
  "murus/nat"
  "murus/bnat"
  "murus/sort"
  "murus/node"
  "murus/lan"
)

func init() {
  nw8()
  nw8dir()
  nw8cyc()
  nw8ring()
  nw8ringdir()
  nw8full()
  nw10()
  nw12()
  nw12ringdir()
  nw12full()
  nw18()
  nw24()
}

func nodes (l, c []int) []Any {
  n := uint(len(l))
  wd := nat.Wd(n)
  a := make ([]Any, n)
  for i:= uint(0); i < n; i++ {
    k := bnat.New(wd); k.SetVal(i)
    a[i] = node.New(k, wd, 1)
    a[i].(node.Node).Set (8 * c[i], 16 * l[i])
    a[i].(node.Node).Colours (col.Blue, col.LightWhite) // XXX
    a[i].(node.Node).ColoursA (col.Red, col.LightWhite) // XXX
  }
  return a
}

func sortNeighbours (n [][]uint) {
  for i:= uint(0); i < uint(len(n)); i++ {
    ns := make ([]Any, len(n[i]))
    for j, k := range n[i] {
      ns[j] = k
    }
    sort.Sort (ns)
    for j, k := range ns {
      n[i][j] = k.(uint)
    }
  }
}

// Returns the number of a in the ring
func h(x []uint, a uint) uint {
  n, k := uint(len(x)), uint(0)
  for i := uint(0); i < n; i++ {
    if x[i] == a {
      k = i
      break
    }
  }
  return k
}

func g(x []uint) []uint {
  n := uint(len(x))
  s := make([]uint, 0)
  for i:= uint(0); i < n; i++ {
    s = append(s, x[(h(x, i) + 1) % n])
  }
  return s
}

func nw8() {
/*
  screen design for mode.MVGA (10 x 30)

1       1 ------- 4
2     /   \     /   \
3   /       \ /       \
4  0         3         6 ---- 7
5   \         \       /
6     \         \   /
7       2 ------- 5

            1         2
  012345678901234567890123456789
*/
  l := []int { 4, 1, 7,  4,  1,  7,  4,  4 }
  c := []int { 1, 6, 6, 11, 16, 16, 21, 28 }
  n := [][]uint { []uint { 1, 2 },
                  []uint { 0, 3, 4 },
                  []uint { 0, 5 },
                  []uint { 1, 4, 5 },
                  []uint { 1, 3, 6 },
                  []uint { 2, 3, 6 },
                  []uint { 4, 5, 7 },
                  []uint { 6 } }
//  sortNeighbours (n)
  Network8 = lan.New(nodes (l, c), false, n)
  Network8.SetDiameter(4)
}

func nw8dir() {
/*
  screen design for mode.MVGA (10 x 30)

1       1 ------- 4
2     /   \     /   \
3   /       \ /       \
4  0         3         6 ---- 7
5   \         \       /
6     \         \   /
7       2 ------- 5

            1         2
  012345678901234567890123456789
*/
  l := []int { 4, 1, 7,  4,  1,  7,  4,  4 }
  c := []int { 1, 6, 6, 11, 16, 16, 21, 28 }
  n := [][]uint { []uint { 1, 2 },
                  []uint { 3, 4 },
                  []uint { 5 },
                  []uint { 4, 5 },
                  []uint { 6 },
                  []uint { 6 },
                  []uint { 7 },
                  []uint { } }
  Network8dir = lan.New(nodes (l, c), true, n)
  Network8dir.SetDiameter(4)
}

func nw8cyc() {
  l := []int { 4, 1, 7,  4,  1,  7,  4,  4 }
  c := []int { 1, 6, 6, 11, 16, 16, 21, 28 }
  n := [][]uint { []uint { 1, 2 },
                  []uint { 3, 4 },
                  []uint { 5 },
                  []uint { 4, 5 },
                  []uint { 2, 6 },
                  []uint { 0, 6 },
                  []uint { 1, 7 },
                  []uint { 4 } }
  Network8cyc = lan.New(nodes (l, c), true, n)
  Network8cyc.SetDiameter(4)
}

func nw8ring() {
/*
  screen design for mode.MVGA (10 x 30)

1           1        3
2  5                          6
3 
4
5
6  2                          4
7           0        7

            1         2
  012345678901234567890123456789
*/
  l := []int {  7,  1,  6,  1,  6,  2,  2,  7 }
  c := []int { 10, 10,  1, 19, 28,  1, 28, 19 }
  n := [][]uint { []uint { 7, 2 },
                  []uint { 5, 3 },
                  []uint { 0, 5 },
                  []uint { 1, 6 },
                  []uint { 6, 7 },
                  []uint { 2, 1 },
                  []uint { 3, 4 },
                  []uint { 4, 0 } }
  Network8ring = lan.New(nodes (l, c), false, n)
  Network8ring.SetDiameter(4)
}

func nw8ringdir() {
/*
  screen design for mode.MVGA (10 x 30)
1           6        2
2  4                          0
3
4   tid-00  ntid-00  nntid-00
5           sent-00
6  1                          3
7           5        7
8
9  active/relay       round 00
4 tid -00   ntid -00   nntid -00

            1         2
  012345678901234567890123456789
*/
  l := []int {  2,  6,  1,  6,  2,  7,  1,  7 }
  c := []int { 28,  1, 19, 28,  1, 10, 10, 19 }
  n := [][]uint { []uint { 2 },
                  []uint { 5 },
                  []uint { 6 },
                  []uint { 0 },
                  []uint { 1 },
                  []uint { 7 },
                  []uint { 4 },
                  []uint { 3 } }
  Network8ringdir = lan.New(nodes (l, c), true, n)
  Network8ringdir.SetDiameter(4)
}

func nw8full() {
  l := []int { 4, 1, 7,  4,  1,  7,  4,  4 }
  c := []int { 1, 6, 6, 11, 16, 16, 21, 28 }
  n := make ([][]uint, 8)
  for i:= uint(0); i < 8; i++ {
    n[i] = make ([]uint, 0)
    for j:= uint(0); j < 8; j++ {
      if j != i {
        n[i] = append (n[i], j)
      }
    }
  }
  Network8full = lan.New(nodes (l, c), false, n)
  Network8full.SetDiameter(1)
}

func nw10() {
/*
1        1 ---- 4 ----- 7
2      /          \       \
3    /              \       \
4  0 ------- 3 ------ 6 ----- 9
5    \     /        /           
6      \ /        /              
7       2 ------ 5 ------ 8

            1         2
  012345678901234567890123456789
*/
  l := []int { 4, 1, 7,  4,  1,  7,  4,  1,  7,  4 }
  c := []int { 1, 7, 6, 11, 14, 15, 20, 22, 24, 28 }
  n := [][]uint { []uint { 1, 2, 3 },
                  []uint { 0, 4 },
                  []uint { 0, 3, 5 },
                  []uint { 0, 2, 6 },
                  []uint { 1, 6, 7 },
                  []uint { 2, 6, 8 },
                  []uint { 3, 4, 5, 9 },
                  []uint { 4, 9 },
                  []uint { 5 },
                  []uint { 6, 7 } }
//  sortNeighbours (n)
  Network10 = lan.New(nodes (l, c), false, n)
  Network10.SetDiameter(4)
}

func nw12() {
/*
  screen design for mode.MVGA (10 x 30)

1        7 ------ 11 --- 8
2      / | \     /   \    \
3    /   |   \ /       \    \
4   3    6    4   10 -- 0 --- 2
5    \   |     \   |  /
6      \ |       \ | /
7        9 ------- 1 -- 5
            1         2
  012345678901234567890123456789
*/
  l := []int {  4,  7,  4,  4,  4,  7,  4,  1,  1,  7,  4,  1 }
  c := []int { 22, 17, 28,  2, 12, 22,  7,  7, 22,  7, 17, 17 }
  n := [][]uint { []uint { 1, 2, 10, 11 },
                  []uint { 0, 4, 5, 9, 10 },
                  []uint { 0, 8 },
                  []uint { 6, 7, 9 },
                  []uint { 1, 7, 11 },
                  []uint { 1 },
                  []uint { 7, 9 },
                  []uint { 3, 4, 6, 11 },
                  []uint { 2, 11 },
                  []uint { 1, 3, 6 },
                  []uint { 0, 1 },
                  []uint { 0, 4, 7, 8 } }
  Network12 = lan.New(nodes(l, c), false, n)
  Network12.SetDiameter(4)
}

func nw12ringdir() {
  max := uint(12)
/*
  screen design for mode.MVGA (10 x 30)

1         x     x     x
2    x                     x
3 
4  x                          0
5
6    x                     x
7         x     x     x

            1         2
  012345678901234567890123456789
*/
//  s := []uint { 0,  7,  8,  2,  9,  6,  4, 10,  1,  3, 11,  5 }
  s := []uint { 0,  4, 10,  6,  1, 11,  3,  8,  5,  2,  9,  7 }
  nn := g(s)
  n := make([][]uint, max)
  for i := uint(0); i < max; i++ {
    n[i] = []uint { nn[i] }
  }
  y := []int {  4,  2,  1,  1,  1,  2,  4,  6,  7,  7,  7,  6 }
  x := []int { 28, 25, 20, 14,  8,  3,  2,  3,  8, 14, 20, 25 }
  l, c := make([]int, max), make([]int, max)
  for i:= uint(0); i < max; i++ { l[s[i]], c[s[i]] = y[i], x[i] }
  Network12ringdir = lan.New(nodes (l, c), true, n)
//  Network12ringdir.SetDistances(s)
  Network12ringdir.SetDiameter(6)
}

func nw12full() {
  l := []int {  4,  7,  4,  4,  4,  7,  4,  1,  1,  7,  4,  1 }
  c := []int { 22, 17, 28,  2, 12, 22,  7,  7, 22,  7, 17, 17 }
  n := make ([][]uint, 12)
  for i:= uint(0); i < 12; i++ {
    n[i] = make ([]uint, 0)
    for j:= uint(0); j < 12; j++ {
      if j != i {
        n[i] = append (n[i], j)
      }
    }
  }
  Network12full = lan.New(nodes (l, c), false, n)
  Network12full.SetDiameter(4)
}

func nw18() {
/*
1   5   17----3----9   15   16
2   |   /     |       / |    |
3   | /       |     /   |    |
4  13----2    0---12    4   11
5       /|    |    |    |    |
6     /  |    |    |    |    |
7   7---10    8--- 6---14----1
*/
  l := []int {  4,  7,  4,  1,  4,  1,  7,  7,  7,  1,  7,  4,  4,  4,  7,  1,  1,  1 }
  c := []int { 12, 27,  7, 12, 22,  2, 17,  2, 12, 17,  7, 27, 17,  2, 22, 22, 27,  7 }
  n := [][]uint { []uint { 3, 8, 12 },
                  []uint { 11, 14 },
                  []uint { 7, 10, 13 },
                  []uint { 0, 9, 17 },
                  []uint { 14, 15 },
                  []uint { 13 },
                  []uint { 8, 12, 14 },
                  []uint { 2, 10 },
                  []uint { 0, 6 },
                  []uint { 3 },
                  []uint { 2, 7 },
                  []uint { 1, 16 },
                  []uint { 0, 6, 15 },
                  []uint { 2, 5, 17 },
                  []uint { 1, 4, 6 },
                  []uint { 4, 12 },
                  []uint { 11 },
                  []uint { 3, 13 } }
  Network18 = lan.New(nodes (l, c), false, n)
  Network18.SetDiameter(11)
}

func nw24() {
/*
1   5   17----3----9   22   16
2   | /       |      /  |    |
3  13----2    0---12    4   11
4     /  |    |         |    |
5   7---23    8---20---14----6
6   |      \  | \    /       |
7  18   10---19    1   21---15
*/
  l := []int {  3,  7, 3,  1,  3, 1,  5, 5,  5,  1, 7,  3,  3, 3,  5,  7,  1, 1, 7,  7,  5,  7,  1,  5 }
  c := []int { 12, 17, 7, 12, 22, 2, 27, 2, 12, 17, 7, 27, 17, 2, 22, 27, 27, 7, 2, 12, 17, 22, 22,  7 }
  n := [][]uint { []uint { 3, 8, 12 },
                  []uint { 8, 14 },
                  []uint { 7, 13, 23 },
                  []uint { 0, 9, 17 },
                  []uint { 14, 22 },
                  []uint { 13 },
                  []uint { 11, 14, 15 },
                  []uint { 2, 18, 23 },
                  []uint { 0, 19, 20 },
                  []uint { 3 },
                  []uint { 19 },
                  []uint { 6, 16 },
                  []uint { 0, 22 },
                  []uint { 2, 5, 17 },
                  []uint { 1, 4, 6, 20 },
                  []uint { 6, 21 },
                  []uint { 11 },
                  []uint { 3, 13 },
                  []uint { 7 },
                  []uint { 8, 10, 23 },
                  []uint { 8, 14 },
                  []uint { 15 },
                  []uint { 2, 14 },
                  []uint { 2, 7 } }
  Network24 = lan.New(nodes (l, c), false, n)
  Network24.SetDiameter(10)
}
