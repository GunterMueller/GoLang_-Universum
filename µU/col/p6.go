package col

// (c) Christian Maurer   v. 201227 - license see µU.go

import
  . "µU/obj"

func p6Encode (a, p Stream) {
  switch bitDepth {
  case 4:
    // TODO
  case 8:
    // TODO
  case 15:
    // TODO
  case 16: // TODO: might be nonsense, has to be checked !
    p[0] = a[1] << 3
    p[1] = a[0] << 5 + a[1] >> 5 << 5
    p[2] = a[0] >> 3
  case 24, 32:
    p[0] = a[0]
    p[1] = a[2]
    p[2] = a[1]
  default:
    p[0] = 0
    p[1] = 0
    p[2] = 0
  }
}
