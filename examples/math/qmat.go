package main

// (c) Christian Maurer   v. 241011 - license see µU.go

import (
  "µU/scr"
  "µU/errh"
  "µU/qmat"
//  "µU/Q"
)

func main() {
  scr.NewWH (0, 0, 888, 880); defer scr.Fin()
  const n = 8
//  const n = 2
  A := qmat.New (n, n)
  A.Set1 (2, 1, 3, 1, 6, 5, 7, 4,
          7, 6, 5, 8, 3, 4, 1, 2,
          1, 3, 2, 4, 1, 3, 1, 6,
          1, 1, 4, 5, 7, 6, 3, 2,
          5, 7, 4, 3, 6, 2, 1, 1,
          3, 4, 6, 2, 5, 1, 1, 7,
          6, 4, 7, 1, 7, 3, 5, 2,
          1, 4, 7, 2, 1, 3, 6, 5)
/*/
  A.Set1 (2, 1,
          3, 4)
/*/
  A.Write (0, 0)
  errh.Error0 ("A")

  A0 := A.Clone().(qmat.QMatrix)
  A1 := A.Clone().(qmat.QMatrix)
  A2 := A.Clone().(qmat.QMatrix)
  A3 := A.Clone().(qmat.QMatrix)

/*/
  d := Q.New()
  d = A.Det()
  _, a, b := d.Vals()
  errh.Error2 ("det(A) =", a, "/", b)
/*/

  A.Invert()
  A.Write (n + 1, 0)
  errh.Error0 ("A^-1")
  A.Mul (A0)
  A.Write (2 * (n + 1), 0)
  errh.Error0 ("A * A^-1")

  A1.Quot (qmat.Unit (n, n), A2)
  A1.Write (3 * (n + 1), 0)
  errh.Error0 ("E / A")
  A1.Write (3 * (n + 1), 0)

  B := qmat.New (n, n)
  B.Set1 ( 5,-1, 7,-9, 6, 5,-4, 8,
           0, 1,-2, 7,-1, 3, 7, 2,
           4,-1, 2, 5, 3,-9, 5,-9,
          -7, 1,-5, 1,-7, 8, 3, 4,
           3,-1, 4,-3, 9,-2, 1,-6,
          -1, 1, 6, 4,-4, 4,-6, 7,
           1,-1, 1,-2, 4, 1,-1,-1,
          -7, 1,-9, 6,-2, 6, 2, 3)
/*/
  B.Set1 ( 1, 0,
           2, 3)
/*/

  B.Write (4 * (n + 1), 0)
  errh.Error0 ("B")

  A3.Mul (B)
  A3.Write (5 * (n + 1), 0)
  errh.Error0 ("A * B")
}
