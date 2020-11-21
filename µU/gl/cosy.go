package gl

// (c) Christian Maurer   v. 201102 - license see µU.go

// import
//   "µU/col"

func coSy (X, Y, Z float64, mit bool) {
/*
  const N = 0.
  cX, cY, cZ := col.LightRed, col.LightGreen, col.LightBlue
  if mit {
    parallelogram (cX, N,-Y,-Z), N, Y,-Z), N,-Y, Z))
  }
  R := X / 128
  R1 := X / 16
  G := X
  G1 := G + 2
  var x float64
  fein := X <= 10
  y := -Y
//  var c0 col.Colour
  for y < Y {
    z := -Z
    for z < Z {
//      if y = 0 {
//        c0 = cY
//      } else if z = 0 {
//        c0 = cZ
//      } else {
//        c0 = cX
//      }
      if fein {
        figure (Points, cX, N, y, z))
      } else {
        Octahedron (cX, N, y, z, R)
      }
      z += 1 // muß gekörnt werden
    }
    y += 1
  }
  figure (Lines, cX, -G1, N, N), G1, N, N))
  sphere (G1, N, N, R1, cX)
  if mit {
    parallelogram (cY, -X, N,-Z), X, N,-Z), -X, N, Z))
  }
  x = -X
  for x < X {
    z := - Z
    for z < Z {
//      if x = 0 {
//        c0 = cX
//      } else if z = 0 {
//        c0 = cZ
//      } else {
//        c0 = cY
//      }
      if fein {
        figure (Points, cY, x, N, z))
      } else {
        octahedron (cY, x, N, z, R)
      }
      z += 1
    }
    x += 1
  }
  figure (Lines, cY, N,-G1, N), N, G1, N))
  sphere (N, G1, N, R1, cY)
  if mit {
    parallelogram (cZ, -X,-Y, N), X,-Y, N), -X, Y, N))
  }
  x = -X
  for x < X {
    y := -Y
    for y < Y {
//      if x = 0 {
//        c0 = cX
//      } else if y = 0 {
//        c0 = cY
//      } else {
//        c0 = cZ
//      }
      if fein {
        figure (Points, cZ, x, y, N))
      } else {
        octahedron (cZ, x, y, N, R)
      }
      y += 1
    }
    x += 1
  }
  figure (Lines, cZ, N, N,-G1), N, N, G1))
  sphere (N, N, G1, R1, cZ)
*/
}
