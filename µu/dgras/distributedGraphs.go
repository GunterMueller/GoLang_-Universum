package dgras

// (c) Christian Maurer   v. 171010 - license see µu.go

import (
  "µu/col"
  "µu/scr"
  "µu/nat"
  "µu/bnat"
  "µu/vtx"
  "µu/edg"
  "µu/gra"
  "µu/host"
  "µu/nchan"
  "µu/dgra"
)
const
  max = 16
var
  hs = make([]host.Host, max)

func init() {
  for i := 0; i < max; i++ {
    hs[i] = host.Localhost()
//    switch i { case 2, 6, 8: hs[i] = host.New(); hs[i].Defined("mars") }
  }
}

func graph (dir bool, l, c []int, es [][]uint, m, id uint) dgra.DistributedGraph {
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
/*
    h := char.New()
    h.SetByte ('A' + byte(i) - 1)
    v[i] = vtx.New (h, 1, 1)
*/
    v[i].Set (int(scr.Wd1()) * c[i], int(scr.Ht1()) * l[i])
    v[i].Colours (cf, cb); v[i].ColoursA (ca, cb)
    g.Ins (v[i])
  }
  for i := uint(0); i < n; i++ {
    for _, j := range es[i] {
      g.Ex2 (v[i], v[j])
      if ! g.Edged() {
        q, r := uint32(n * i + j), uint32(n * j + i)
        p0 := uint32(nchan.Port(n, i, j, 0) - nchan.Port0)
        p := p0 + 1<<10 * (q + 1<<10 * r)
        e := edg.New(dir, p)
        e.SetPos0 (v[i].Pos()); e.SetPos1 (v[j].Pos())
        e.Label(false)
        e.Colours (cf, cb); e.ColoursA (ca, cb)
        g.Edge (e)
      }
    }
  }
  g.Ex (v[id])
  g = g.Star()
  g.SetWrite (vtx.W, edg.W)
  d := dgra.New (g)
  n = g.Num() - 1
  h := make([]host.Host, n)
  for i := uint(0); i < n; i++ { h[i] = host.Localhost() }
  d.SetHosts (h) // XXX
  d.SetDiameter(m)
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

func g3 (i uint) dgra.DistributedGraph {
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
                  []uint { 0, 2 },
                  []uint { 0, 1 } }
  return graph (false, l, c, e, 1, i)
}

func g3dir (i uint) dgra.DistributedGraph {
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
  return graph (true, l, c, e, 1, i)
}

func g4 (i uint) dgra.DistributedGraph {
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
                  []uint { 0, 2 },
                  []uint { 0, 1, 3 },
                  []uint { 2 } }
  return graph (false, l, c, e, 2, i)
}

func g4flat (i uint) dgra.DistributedGraph {
/*
4  0 ------ 1 ------ 2 ------ 3

            1         2
  012345678901234567890123456789
*/
  l := []int { 4,  4,  4,  4 }
  c := []int { 1, 10, 19, 28 }
  e := [][]uint { []uint { 1 },
                  []uint { 0, 2 },
                  []uint { 1, 3 },
                  []uint { 2 } }
  return graph (false, l, c, e, 3, i)
}


func g4full (i uint) dgra.DistributedGraph {
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
                  []uint { 0, 2, 3 },
                  []uint { 0, 1, 3 },
                  []uint { 0, 1, 2 } }
  return graph (false, l, c, e, 1, i)
}

func g5 (i uint) dgra.DistributedGraph {
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
                  []uint { 0, 2, 4 },
                  []uint { 1 },
                  []uint { 0, 4 },
                  []uint { 1, 3 }}
  return graph (false, l, c, e, 3, i)
}

func g6 (i uint) dgra.DistributedGraph {
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
                  []uint { 0, 5 },
                  []uint { 0, 3, 4 },
                  []uint { 0, 2, 5 },
                  []uint { 2, 5 },
                  []uint { 1, 3, 4 }}
  return graph (false, l, c, e, 2, i)
}

func g8 (i uint) dgra.DistributedGraph {
/*
1        1 --------- 4 -------- 7
2      /               \
3    /                   \
4  0 --------- 3 --------- 6
5    \       /           /
6      \   /           /
7        2 --------- 5 

            1         2         3
  01234567890123456789012345678901
*/
  l := []int { 4, 1, 7,  4,  1,  7,  4,  1 }
  c := []int { 1, 7, 7, 13, 19, 19, 25, 30 }
  e := [][]uint { []uint { 1, 2, 3 },
                  []uint { 0, 4 },
                  []uint { 0, 3, 5 },
                  []uint { 0, 2, 6 },
                  []uint { 1, 6, 7 },
                  []uint { 2, 6 },
                  []uint { 3, 4, 5 },
                  []uint { 4 } }
  return graph (false, l, c, e, 4, i)
}

func g8dir (i uint) dgra.DistributedGraph {
/*
0       1 ------> 4
1      *  \      *  \
2    /      *  /      *
3  0         3         6 -----> 7
4    \         \      *
5      *         *  /
7       2 ------> 5

            1         2         3
  01234567890123456789012345678901
*/
  l := []int { 4, 1, 7,  4,  1,  7,  4,  4 }
  c := []int { 1, 6, 6, 11, 11, 15, 21, 30 }
  e := [][]uint { []uint { 1, 2 },
                  []uint { 3, 4 },
                  []uint { 5 },
                  []uint { 4, 5 },
                  []uint { 6 },
                  []uint { 6 },
                  []uint { 7 },
                  []uint { } }
  return graph (true, l, c, e, 4, i)
}

func g8cyc (i uint) dgra.DistributedGraph {
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
  return graph (true, l, c, e, 4, i)
}

func g8ring (i uint) dgra.DistributedGraph {
/*
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
  e := [][]uint { []uint { 2, 7 },
                  []uint { 3, 5 },
                  []uint { 0, 5 },
                  []uint { 1, 6 },
                  []uint { 6, 7 },
                  []uint { 1, 2 },
                  []uint { 3, 4 },
                  []uint { 0, 4 } }
  return graph (false, l, c, e, 4, i)
}

func g8ringdir (i uint) dgra.DistributedGraph {
/*
1           1        3
2  5                          6
3
4   tid-00  ntid-00  nntid-00
5           sent-00
6  2                          4
7           0        7
8
9  active/relay       round 00
4 tid -00   ntid -00   nntid -00

            1         2
  012345678901234567890123456789
*/
  l := []int {  7,  1,  6,  1,  6,  2,  2,  7 }
  c := []int { 10, 10,  1, 19, 28,  1, 28, 19 }
  e := [][]uint { []uint { 7 },
                  []uint { 5 },
                  []uint { 0 },
                  []uint { 1 },
                  []uint { 6 },
                  []uint { 2 },
                  []uint { 3 },
                  []uint { 4 } }
  return graph (true, l, c, e, 7, i)
}

func g8full (i uint) dgra.DistributedGraph {
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
  return graph (false, l, c, e, 1, i)
}

func g10 (i uint) dgra.DistributedGraph {
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
                  []uint { 0, 4 },
                  []uint { 0, 3, 5 },
                  []uint { 0, 2, 6 },
                  []uint { 1, 6, 7 },
                  []uint { 2, 6, 8 },
                  []uint { 3, 4, 5, 9 },
                  []uint { 4 },
                  []uint { 5, 9 },
                  []uint { 6, 8 } }
  return graph (false, l, c, e, 4, i)
}

func g12 (i uint) dgra.DistributedGraph {
/*
1       7 ------- 11 ---- 8
2     / | \     /    \      \
3   /   |   \ /        \      \
4  3    6    4    10 --- 0 --- 2
5   \   |      \   |   /
6     \ |        \ | /
7       9 -------- 1 ---- 5
            1         2         3
  01234567890123456789012345678901
*/
  l := []int {  4,  7,  4,  4,  4,  7,  4,  1,  1,  7,  4,  1 }
  c := []int { 22, 17, 28,  2, 12, 22,  7,  7, 22,  7, 17, 17 }
  e := [][]uint { []uint { 1, 2, 10, 11 },
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
  return graph (false, l, c, e, 4, i)
}

func g12ringdir (i uint) dgra.DistributedGraph {
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
  return graph (true, l, c, e, 6, i)
}

func g12full (i uint) dgra.DistributedGraph {
  l := []int {  4,  7,  4,  4,  4,  7,  4,  1,  1,  7,  4,  1 }
  c := []int { 22, 17, 28,  2, 12, 22,  7,  7, 22,  7, 17, 17 }
  e := make ([][]uint, 12)
  for j:= uint(0); j < 12; j++ {
    e[j] = make ([]uint, 0)
    for k:= uint(0); k < 12; k++ {
      if k != j { e[j] = append (e[j], k) }
    }
  }
  return graph (false, l, c, e, 1, i)
}

func g16 (i uint) dgra.DistributedGraph {
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
                  []uint { 0, 2, 5, 9 },
                  []uint { 14, 15 },
                  []uint { 2, 3, 13 },
                  []uint { 0, 8, 12, 14 },
                  []uint { 2, 10 },
                  []uint { 0, 2, 6, 10 },
                  []uint { 3, 12, 15 },
                  []uint { 2, 7, 8 },
                  []uint { 1, 15 },
                  []uint { 0, 6, 9, 15 },
                  []uint { 2, 5 },
                  []uint { 1, 4, 6 },
                  []uint { 4, 9, 11, 12 } }
  return graph (false, l, c, e, 5, i)
}

func g16dir (i uint) dgra.DistributedGraph {
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
  return graph (true, l, c, e, 5, i)
}

func g16ring (i uint) dgra.DistributedGraph {
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
  e := [][]uint { []uint { 7, 10 },
                  []uint { 6, 9 },
                  []uint { 8, 14 },
                  []uint { 7, 14 },
                  []uint { 10, 13 },
                  []uint { 12, 15 },
                  []uint { 1, 11 },
                  []uint { 0, 3 },
                  []uint { 2, 12 },
                  []uint { 1, 13 },
                  []uint { 0, 4 },
                  []uint { 6, 15 },
                  []uint { 5, 8 },
                  []uint { 4, 9 },
                  []uint { 2, 4 },
                  []uint { 5, 11 } }
  return graph (false, l, c, e, 8, i)
}

func g16ringdir (i uint) dgra.DistributedGraph {
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
  return graph (true, l, c, e, 6, i)
}

func g16full (i uint) dgra.DistributedGraph {
  l := []int {  4,  1,  7,  7,  1, 6, 1,  6,  7,  1,  2, 2, 7,  1,  7, 4 }
  c := []int { 28, 10, 14, 22, 22, 3, 6, 26, 10, 14, 26, 3, 6, 18, 18, 1 }
  m := uint(len(l))
  e := make ([][]uint, m)
  for j:= uint(0); j < m; j++ {
    e[j] = make ([]uint, 0)
    for k:= uint(0); k < m; k++ { if k != i { e[j] = append (e[j], k) } }
  }
  return graph (false, l, c, e, 1, i)
}
