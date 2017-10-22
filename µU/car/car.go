package car

// (c) Christian Maurer   v. 140210 - license see µU.go

import (
  "µU/col"
  "µU/scr"
)
var
  car = [...]string {
   "                         *      ",
   "                         *      ",
   "      ************       *      ",
   "     ***************     *      ",
   "    ***      *      *    *      ",
   "   ***       *       *   *      ",
   "  ***        *        *  *      ",
   " ****        *         * *      ",
   "******************************* ",
   "**************  ************** *",
   "* **************************** *",
   "* ***************************** ",
   "******************************* ",
   " *****************************  ",
   "     *****          *****       ",
   "      ***            ***        " }


func draw (right bool, c col.Colour, X, Y int) {
  xs, ys:= make ([]int, W * H), make ([]int, W * H)
  for y:= 0; y < H; y++ {
    for x:= 0; x < W; x++ {
      i:= W * y + x
      if car[y][x] == '*' {
        if right {
          xs[i] = X + x
        } else {
          xs[i] = X + W - 1 - x
        }
        ys[i] = Y + y
      }
    }
  }
  scr.Lock()
  scr.ColourF (c)
  scr.Points (xs, ys)
//  scr.Flush()
  scr.Unlock()
}
