package dgra

// (c) Christian Maurer   v. 200120 - license see µU.go

import (
  "µU/env"
  "µU/col"
  "µU/scr"
  "µU/nat"
  "µU/bnat"
  "µU/vtx"
  "µU/edg"
  "µU/gra"
//  "µU/host"
  "µU/nchan"
)

func newg (dir bool, l, c []int, es [][]uint, m, id uint) DistributedGraph {
  cf, ca, cb := col.Blue(), col.Red(), col.LightWhite()
  n := uint(len(l))
  if n != uint(len(es)) || n != uint(len(c)) { panic("len's different") }
  wd := nat.Wd (n)
  g := gra.New (dir, vtx.New(bnat.New(wd), wd, 1), edg.New(dir, uint32(nchan.Port0)))
  v := make([]vtx.Vertex, n)
  for i := uint(0); i < n; i++ {
    b := bnat.New(wd)
    b.SetVal(i)
    v[i] = vtx.New (b, wd, 1)
    v[i].Set (int(scr.Wd1()) * c[i], int(scr.Ht1()) * l[i])
    v[i].Colours (cf, cb); v[i].ColoursA (ca, cb)
    g.Ins (v[i])
  }
  for i := uint(0); i < n; i++ {
    for _, j := range es[i] {
      g.Ex2 (v[i], v[j])
      if ! g.Edged() {
        e := edg.New (dir, uint32(nchan.Port(n, i, j, 0)))
        e.SetPos0 (v[i].Pos()); e.SetPos1 (v[j].Pos())
        e.Label(false)
        e.Colours (cf, cb); e.ColoursA (ca, cb)
        g.Edge (e)
      }
    }
  }
  g.Ex (v[id])
  g.SetWrite (vtx.W, edg.W)
  g = g.Star()
  d := new_(g).(*distributedGraph)
  d.setSize (n)
  h := make([]string, n) // make([]host.Host, n)
  for i := uint(0); i < n; i++ {
    h[i] = env.Localhost()
  }
  d.setHosts (h)
  d.diameter = m
  return d
}

// Returns the number of a in the ring
func num (x []uint, a uint) uint {
  n, k := uint(len(x)), uint(0)
  for i := uint(0); i < n; i++ {
    if x[i] == a {
      k = i
      break
    }
  }
  return k
}

func g (x []uint) []uint {
  n := uint(len(x))
  s := make([]uint, 0)
  for i:= uint(0); i < n; i++ {
    s = append(s, x[(num(x, i) + 1) % n])
  }
  return s
}

func g3 (i uint) DistributedGraph {
/*
  screen design for mode.MVGA (10 x 30)

1                1
2               / \
3             /     \
4           /         \
5         /             \
6       /                 \
7      2 ----------------- 0

            1         2
  012345678901234567890123456789
*/
  l := []int {  7,  1, 7 }
  c := []int { 25, 15, 5 }
  e := [][]uint { []uint { 1, 2 },
                  []uint { 2 },
                  []uint { } }
  return newg (false, l, c, e, 1, i)
}

func g3dir (i uint) DistributedGraph {
/*
1                1
2               / \
3             /     \
4           /         \
5         /             \
6       *                 *
7      2 ----------------> 0

            1         2
  012345678901234567890123456789
*/
  l := []int {  7,  1, 7 }
  c := []int { 25, 15, 5 }
  e := [][]uint { []uint { },
                  []uint { 0, 2 },
                  []uint { 0 } }
  return newg (true, l, c, e, 1, i)
}

func g4 (i uint) DistributedGraph {
/*
1        0 ------------- 1
2        |             /
3        |          /
4        |       /
5        |    /
6        | /
7        2 ------------- 3

            1         2
  012345678901234567890123456789
*/
  l := []int { 1,  1, 7,  7 }
  c := []int { 7, 23, 7, 23 }
  e := [][]uint { []uint { 1, 2 },
                  []uint { 2 },
                  []uint { 3 },
                  []uint { } }
  return newg (false, l, c, e, 2, i)
}

func g4flat (i uint) DistributedGraph {
/*
4  0 ------ 1 ------ 2 ------ 3

            1         2
  012345678901234567890123456789
*/
  l := []int { 4,  4,  4,  4 }
  c := []int { 1, 10, 19, 28 }
  e := [][]uint { []uint { 1 },
                  []uint { 2 },
                  []uint { 3 },
                  []uint { } }
  return newg (false, l, c, e, 3, i)
}


func g4full (i uint) DistributedGraph {
/*
1             /--------  2
2           /         /  |
3         /         /    |
4       0 ------ 1       |
5         \         \    |
6           \         \  |
7             \--------  3

            1         2
  012345678901234567890123456789
*/
  l := []int {  4,  4,  1,  7 }
  c := []int {  6, 15, 23, 23 }
  e := [][]uint { []uint { 1, 2, 3 },
                  []uint { 2, 3 },
                  []uint { 3 },
                  []uint { } }
  return newg (false, l, c, e, 1, i)
}

func g5 (i uint) DistributedGraph {
/*
1    0 ---------- 1 ---------- 2
2    |            |
3    |            |
4    |            |
5    |            |
6    |            |
7    3 ---------- 4

            1         2         3
  01234567890123456789012345678901
*/
  l := []int { 1,  1, 1,  7,  7 }
  c := []int { 3, 16, 3, 16, 29 }
  e := [][]uint { []uint { 1, 3 },
                  []uint { 2, 4 },
                  []uint { },
                  []uint { 4 },
                  []uint { }}
  return newg (false, l, c, e, 3, i)
}

func g5ring (i uint) DistributedGraph {
/*
1            3
2          /   \
3       /         \
4    /               \
5  0                   1
6  |                   |
7   \                 /
8    \               /
9     \             /
10     4 --------- 2

            1         2
  0123456789012345678901
*/
  l := []int {  1, 5, 10, 18,  5 }
  c := []int { 11, 1,  5, 17, 21 }
  e := [][]uint { []uint { 3, 4 },
                  []uint { 2, 3 },
                  []uint { 1, 4 },
                  []uint { 0, 1 },
                  []uint { 0, 2 } }
  return newg (false, l, c, e, 4, i)
}

func g5ringdir (i uint) DistributedGraph {
/*
1            3
2          /   \
3       /         \
4    /               \
5  0                   1
6  |                   |
7   \                 /
8    \               /
9     \             /
10     4 --------- 2

            1         2
  0123456789012345678901
*/
  l := []int {  1, 5, 10, 18,  5 }
  c := []int { 11, 1,  5, 17, 21 }
  e := [][]uint { []uint { 4 },
                  []uint { 3 },
                  []uint { 1 },
                  []uint { 0 },
                  []uint { 2 } }
  return newg (false, l, c, e, 4, i)
}

func g5full (i uint) DistributedGraph {
/*
1            0
2          // \\
3       /  /   \  \
4    /    /     \    \
5  1 ----/-------\---- 4
6  |  \ /         \ /  |
7   \  / \       / \  /
8    \ |    \ /    | /
9     \|  /     \  |/
10     2 --------- 3

            1         2
  0123456789012345678901
*/
  l := []int {  1, 5, 10, 18,  5 }
  c := []int { 11, 1,  5, 17, 21 }
  e := [][]uint { []uint { 1, 2, 3, 4 },
                  []uint { 2, 3, 4 },
                  []uint { 3, 4 },
                  []uint { 4 },
                  []uint { } }
  return newg (false, l, c, e, 1, i)
}

func g6 (i uint) DistributedGraph {
/*
1        /---- 1 ----\
2      /               \
3    /                   \
4  0 --------- 3 --------- 5
5    \       /           /
6      \   /           /
7        2 --------- 4 

            1         2         3
  01234567890123456789012345678901
*/
  l := []int { 4,  1, 7,  4,  7,  4 }
  c := []int { 1, 13, 7, 13, 19, 25 }
  e := [][]uint { []uint { 1, 2, 3 },
                  []uint { 5 },
                  []uint { 3, 4 },
                  []uint { 5 },
                  []uint { 5 },
                  []uint { }}
  return newg (false, l, c, e, 2, i)
}

func g6full (i uint) DistributedGraph {
/*
2          /-- 1 --\
3        /     |     \
4      /       |       \
5    /         |         \
6  0 --------- 3 --------- 5
7    \       /   \       /
8     \     /     \     /
9      \   /       \   /
10       2 --------- 4 

            1         2
  01234567890123456789012345
*/
  l := []int { 6,  2, 10,  6, 10,  6 }
  c := []int { 1, 13,  7, 13, 19, 25 }
  e := [][]uint { []uint { 1, 2, 3, 4, 5 },
                  []uint { 2, 3, 4, 5 },
                  []uint { 3, 4, 5 },
                  []uint { 4, 5 },
                  []uint { 5 },
                  []uint { }}
  return newg (false, l, c, e, 1, i)
}

func g8a (i uint) DistributedGraph {
/*
2        1 --------- 4 -------- 7
3       /             \
4     /                 \
5   /                     \
6  0 --------- 3 --------- 6
7   \         /           /
8     \     /           /
9       \ /           /
10       2 --------- 5 

            1         2         3
  01234567890123456789012345678901
*/
  l := []int { 6, 2, 10,  6,  2, 10,  6,  2 }
  c := []int { 1, 7,  7, 13, 19, 19, 25, 31 }
  e := [][]uint { []uint { 1, 2, 3 },
                  []uint { 0, 4 },
                  []uint { 0, 3, 5 },
                  []uint { 0, 2, 6 },
                  []uint { 1, 6, 7 },
                  []uint { 2, 6 },
                  []uint { 3, 4, 5 },
                  []uint { 4 } }
  return newg (false, l, c, e, 4, i)
}

func g8 (i uint) DistributedGraph {
/*
2        0 --------- 1 -------- 2
3       /             \
4     /                 \
5   /                     \
6  3 --------- 4 --------- 5
7   \         /           /
8     \     /           /
9       \ /           /
10       6 --------- 7 

            1         2         3
  01234567890123456789012345678901
*/
  l := []int { 2,  2,  2, 6,  6,  6, 10, 10 }
  c := []int { 7, 19, 31, 1, 13, 25,  7, 19 }
  e := [][]uint { []uint { 1, 3 },
                  []uint { 2, 5 },
                  []uint { },
                  []uint { 4, 6 },
                  []uint { 5, 6 },
                  []uint { 7 },
                  []uint { 7 },
                  []uint { } }
  return newg (false, l, c, e, 4, i)
}

func g8dir (i uint) DistributedGraph {
/*
2        0 --------> 1 --------> 2
3       *             *         /
4     /                 \     /
5   /                     \ *
6  3 <-------- 4 --------> 5
7   \         /           /
8     \     /           /
9       * *           *
10       6 --------> 7 

            1         2         3
  012345678901234567890123456789012
*/
  l := []int { 2,  2,  2,  6,  6,  6, 10, 10 }
  c := []int { 7, 19, 22,  1, 13, 25,  7, 19 }
  e := [][]uint { []uint { 1 },
                  []uint { 2 },
                  []uint { 5 },
                  []uint { 0, 6 },
                  []uint { 3, 5, 6 },
                  []uint { 7 },
                  []uint { 7 },
                  []uint { } }
  return newg (true, l, c, e, 3, i)
}

func g8cyc (i uint) DistributedGraph {
  l := []int { 4, 1, 7,  4,  1,  7,  4,  4 }
  c := []int { 1, 6, 6, 11, 16, 16, 21, 28 }
  e := [][]uint { []uint { 1, 2 },
                  []uint { 3, 4 },
                  []uint { 5 },
                  []uint { 4, 5 },
                  []uint { 2, 6 },
                  []uint { 0, 6 },
                  []uint { 1, 7 },
                  []uint { 4 } }
  return newg (true, l, c, e, 4, i)
}

func g8ring (i uint) DistributedGraph {
/*
2            1        3
3
4   5                          6
5
6
7
8   2                          4
9
10           0        7

            1         2         3
  01234567890123456789012345678901
*/
  l := []int { 10,  2,  8,  2,  8,  4,  4, 10 }
  c := []int { 11, 11,  2, 20, 29,  2, 29, 20 }
  e := [][]uint { []uint { 7 },
                  []uint { 5 },
                  []uint { 0 },
                  []uint { 1 },
                  []uint { 6 },
                  []uint { 2 },
                  []uint { 3 },
                  []uint { 4 } }
  return newg (false, l, c, e, 4, i)
}

func g8ringdir (i uint) DistributedGraph {
/*
1            1        3
2
3   5                          6
4   tid-00  ntid-00  nntid-00
5           sent-00
6  active/relay       round 00
7 tid -00   ntid -00   nntid -00
8   2                          4
9 
10           0        7

            1         2         3
  01234567890123456789012345678901
*/
  l := []int { 10,  1,  8,  1,  8,  3,  3, 10 }
  c := []int { 11, 11,  2, 20, 29,  2, 29, 20 }
  e := [][]uint { []uint { 7 },
                  []uint { 5 },
                  []uint { 0 },
                  []uint { 1 },
                  []uint { 6 },
                  []uint { 2 },
                  []uint { 3 },
                  []uint { 4 } }
  return newg (true, l, c, e, 7, i)
}

func g8full (i uint) DistributedGraph {
//  l := []int { 4, 1, 7,  4,  1,  7,  4,  4 }
//  c := []int { 1, 6, 6, 11, 16, 16, 21, 28 }
  l := []int {  7,  1,  6,  1,  6,  2,  2,  7 }
  c := []int { 10, 10,  1, 19, 28,  1, 28, 19 }
  e := make ([][]uint, 8)
  for j := uint(0); j < 8; j++ {
    e[j] = make ([]uint, 0)
    for k:= uint(0); k < 8; k++ {
      if k != j { e[j] = append (e[j], k) }
    }
  }
  return newg (false, l, c, e, 1, i)
}

func g10 (i uint) DistributedGraph {
/*
1       1 ------ 4 ------ 7
2     /            \
3    /              \
4  0 ------ 3 ------ 6 ------ 9
5    \     /        /        /
6     \   /        /        /
7       2 ------ 5 ------ 8

            1         2         3
  01234567890123456789012345678901
*/
  l := []int { 4, 1, 7,  4,  1,  7,  4,  1,  7,  4 }
  c := []int { 1, 6, 6, 10, 15, 15, 19, 24, 24, 28 }
  e := [][]uint { []uint { 1, 2, 3 },
                  []uint { 4 },
                  []uint { 3, 5 },
                  []uint { 6 },
                  []uint { 6, 7 },
                  []uint { 6, 8 },
                  []uint { 9 },
                  []uint { },
                  []uint { 9 },
                  []uint { } }
  return newg (false, l, c, e, 4, i)
}

func g12 (i uint) DistributedGraph {
/*
2        0 -------- 1------ 2
3      / | \       / \       \
4     /  |  \    /     \      \
5    /   |   \ /         \     \
6   3 -- 4 -- 5 --- 6 --- 7 --- 8
7    \   |     \    |    /
8     \  |       \  |  /
9      \ |         \|/
10       9-------- 10 ----- 11
            1         2         3
  01234567890123456789012345678901
*/
//              0   1   2   3   4   5   6   7   8   9  10  11
  l := []int {  3,  2,  3,  5,  6,  7,  5,  6,  6,  9, 10,  9 }
  c := []int {  7, 18, 26,  2,  7, 12, 18, 24, 30,  7, 18, 27 }
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
  return newg (false, l, c, e, 4, i)
}

func g12ringdir (i uint) DistributedGraph {
/*
1         x     x     x
2    x                     x
3
4  x                          0
5
6    x                     x
7         x     x     x

            1         2         3
  01234567890123456789012345678901
*/
  m := uint(12)
//  s := []uint { 0,  7,  8,  2,  9,  6,  4, 10,  1,  3, 11,  5 }
  s := []uint { 0,  4, 10,  6,  1, 11,  3,  8,  5,  2,  9,  7 }
  nn := g(s)
  e := make([][]uint, m)
  for i := uint(0); i < m; i++ {
    e[i] = []uint { nn[i] }
  }
  y := []int {  4,  2,  1,  1,  1,  2,  4,  6,  7,  7,  7,  6 }
  x := []int { 28, 25, 20, 14,  8,  3,  2,  3,  8, 14, 20, 25 }
  l, c := make([]int, m), make([]int, m)
  for j := uint(0); j < m; j++ { l[s[j]], c[s[j]] = y[j], x[j] }
  return newg (true, l, c, e, 6, i)
}

func g12full (i uint) DistributedGraph {
  l := []int {  4,  7,  4,  4,  4,  7,  4,  1,  1,  7,  4,  1 }
  c := []int { 22, 17, 28,  2, 12, 22,  7,  7, 22,  7, 17, 17 }
  e := make ([][]uint, 12)
  for j:= uint(0); j < 12; j++ {
    e[j] = make ([]uint, 0)
    for k:= uint(0); k < 12; k++ {
      if k != j { e[j] = append (e[j], k) }
    }
  }
  return newg (false, l, c, e, 1, i)
}

func g16 (i uint) DistributedGraph {
/*
1     5-----3--------9-----15
2    / \   / \       |   / | \
3   /   \ /   \      | /   |   \
4  13----2     0----12     4   11
5      / | \   | \   |     |    |
6    /   |   \ |   \ |     |    |
7  7----10-----8-----6----14----1

            1         2         3
  01234567890123456789012345678901
*/
  l := []int {  4,  7,  4,  1,  4,  1,  7,  7,  7,  1,  7,  4,  4,  4,  7,  1 }
  c := []int { 13, 30,  7, 11, 23,  4, 19,  1, 13, 18,  7, 30, 19,  1, 25, 25 }
  e := [][]uint { []uint { 3, 6, 8, 12 },
                  []uint { 11, 14 },
                  []uint { 3, 5, 7, 8, 10, 13 },
                  []uint { 5, 9 },
                  []uint { 14, 15 },
                  []uint { 13 },
                  []uint { 8, 12, 14 },
                  []uint { 10 },
                  []uint { 10 },
                  []uint { 12, 15 },
                  []uint { },
                  []uint { 15 },
                  []uint { 15 },
                  []uint { },
                  []uint { },
                  []uint { } }
  return newg (false, l, c, e, 5, i)
}

func g16dir (i uint) DistributedGraph {
/*
1     5*----3----*9----*15
2    / \   / *     \   / | \
3   *   * *   \     * *  *  *
4  13*---2     0---*12   4   11
5      / |*   / \   *    *    |
6     *  * \ *   * /     |    *
7   7*--10*-8----*6----*14---*1
*/
  l := []int {  4,  7,  4,  1,  4,  1,  7,  7,  7,  1,  7,  4,  4,  4,  7,  1 }
  c := []int { 13, 28,  8, 12, 23,  4, 16,  1, 11, 20,  6, 28, 18,  2, 22, 28 }
  e := [][]uint { []uint { 3, 6, 8, 12 },
                  []uint { },
                  []uint { 7, 10, 13 },
                  []uint { 2, 5, 9 },
                  []uint { },
                  []uint { 2, 13 },
                  []uint { 12, 14 },
                  []uint { },
                  []uint { 2, 6, 10 },
                  []uint { 12, 15 },
                  []uint { 7 },
                  []uint { 1 },
                  []uint { },
                  []uint { },
                  []uint { 1, 4 },
                  []uint { 4, 11, 12 } }
  return newg (true, l, c, e, 5, i)
}

func g16ring (i uint) DistributedGraph {
/*
1       6   1   9  13   4
2   11                     10
3
4 15                          0
5
6    5                      7
7      12   8   2  14   3

            1         2
  012345678901234567890123456789
*/
  l := []int {  4,  1,  7,  7,  1, 6, 1,  6,  7,  1,  2, 2, 7,  1,  7, 4 }
  c := []int { 28, 10, 14, 22, 22, 3, 6, 26, 10, 14, 26, 3, 6, 18, 18, 1 }
  e := [][]uint { []uint { 10 },
                  []uint { 6 },
                  []uint { 14 },
                  []uint { 7 },
                  []uint { 13 },
                  []uint { 12 },
                  []uint { 11 },
                  []uint { 0 },
                  []uint { 2 },
                  []uint { 1 },
                  []uint { 4 },
                  []uint { 15 },
                  []uint { 8 },
                  []uint { 9 },
                  []uint { 3 },
                  []uint { 5 } }
  return newg (false, l, c, e, 8, i)
}

func g16ringdir (i uint) DistributedGraph {
/*
1       x   x   x   x   x
2    x                      x
3
4  x                          0
5
6    x                      x
7       x   x   x   x   x

            1         2
  012345678901234567890123456789
*/
  m := uint(16)
  s := make([]uint, m)
  s = []uint { 0, 13, 4, 10, 15,  6,  1, 11,  3,  8,  5, 12,  2,  9, 14,  7 }
  s = []uint { 0,  1, 2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13, 14, 15 }
  s = []uint { 0, 14, 5,  8,  3, 12, 13,  6, 11,  2,  1, 15,  7,  4, 10,  9 }
  s = []uint { 0,  8, 1,  9,  2, 10,  3, 11,  4, 12,  5, 13,  6, 14,  7, 15 }
  nn := g(s)
  e := make([][]uint, m)
  for j := uint(0); j < m; j++ { e[j] = []uint { nn[j] } }
  y := []int {  4,  2,  1,  1,  1,  1, 1, 2, 4, 6, 7,  7,  7,  7,  7,  6 }
  x := []int { 28, 26, 22, 18, 14, 10, 6, 3, 1, 3, 6, 10, 14, 18, 22, 26 }
  l, c := make([]int, m), make([]int, m)
  for j := uint(0); i < m; j++ { l[s[j]], c[s[j]] = y[j], x[j] }
  return newg (true, l, c, e, 6, i)
}

func g16full (i uint) DistributedGraph {
  l := []int {  4,  1,  7,  7,  1, 6, 1,  6,  7,  1,  2, 2, 7,  1,  7, 4 }
  c := []int { 28, 10, 14, 22, 22, 3, 6, 26, 10, 14, 26, 3, 6, 18, 18, 1 }
  m := uint(len(l))
  e := make ([][]uint, m)
  for j:= uint(0); j < m; j++ {
    e[j] = make ([]uint, 0)
    for k:= uint(0); k < m; k++ { if k != i { e[j] = append (e[j], k) } }
  }
  return newg (false, l, c, e, 1, i)
}
