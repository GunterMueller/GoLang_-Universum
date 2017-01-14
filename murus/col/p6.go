package col

// (c) murus.org  v. 140527 - license see murus.go

func p6Encode (A, P []byte) {
  switch bitDepth {
  case 4:
    // TODO
  case 8:
    // TODO
  case 15:
    // TODO
  case 16: // TODO: might be nonsense, has to be checked !
    P[0] = A[1] << 3
    P[1] = A[0] << 5 + A[1] >> 5 << 5
    P[2] = A[0] >> 3
  case 24, 32:
    P[0] = A[2]
    P[1] = A[1]
    P[2] = A[0]
  default:
    for i:= 0; i < P6; i++ {
      P[i] = byte(0)
    }
  }
}

func p6Colour (A []byte) Colour {
  P:= make ([]byte, P6)
  p6Encode (A, P)
  var c Colour
  c.R, c.G, c.B = A[0], A[1], A[2]
  return c
}
