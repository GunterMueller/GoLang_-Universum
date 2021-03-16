package ptr

// (c) Christian Maurer   v. 210311 - license see µU.go

func code (p Pointer) uint { // /usr/include/X11/cursorfont.h
  switch p {
  case Standard:
    break
  case Gumby:
    return 56
  case Hand:
    return 58
  case Gobbler:
    return 54
  case Watch:
    return 150
  }
  return 68
}
