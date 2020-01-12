package show

// (c) Christian Maurer   v. 191019 - license see µU.go

import
  "µU/env"

func arg() Mode {
  switch env.Arg1() {
  case 'f':
    return Fly
  case 'l':
    return Look
  case 'w':
    return Walk
  }
  return Look
}
