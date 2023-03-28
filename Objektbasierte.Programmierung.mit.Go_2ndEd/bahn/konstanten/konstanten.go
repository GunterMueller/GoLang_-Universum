package konstanten

// (c) Christian Maurer   v. 230306 - license see µU.go

import (
  "µU/env"
  . "µU/scr"
  "µU/errh"
)

func init_() {
  Y0 = 3 * 16
  x, _ := MaxRes()
  W1, H1 = 36, 24
// falls Bildschirm zu klein:
  if x < 1300 { W1, H1 = 30, 20 }
  if x < 1100 { W1, H1 = 24, 16 }
  W2, H2 = W1 / 2, H1 / 2
  y := uint(0)
  if env.NArgs() == 0 { // Server
    NewWH (0, 0, 272, 16)
    Name ("Server ist gestartet")
    errh.Hint ("beenden mit Esc")
  } else { // Bahnhof
    y = (NZeilen * 16 + 28) * (env.N(1) % 3)
    NewWH (0, 24 + y, NSpalten * uint(W1), 8 * uint(H1) + 4 * 16)
//                          42 * 36 = 1512       8 * 24 + 4 * 16 = 256
//                          42 * 30 = 1260       8 * 20 + 4 * 16 = 224
//                          42 * 24 = 1082       8 * 16 + 4 * 16 = 192
  }
}
