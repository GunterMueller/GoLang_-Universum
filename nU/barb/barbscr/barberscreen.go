package barbscr

// (c) Christian Maurer   v. 180101 - license see nU.go

import ( "time"; "math/rand"; "nU/col"; "nU/scr")

const (
  line = 1
  column0 = 3; columnB = 20; columnC = 30
)
var (
  N, K uint
  barberSleeping = true
  tx = []string {
         "ch        ",
         " ch       ",
         "  ch      ",
         "   ch     ",
         "r   ch    ",
         "hr   ch   ",
         "chr   ch  ",
         " chr   ch ",
         "  chr   ch",
         "   chr   h",
         "    chr   ",
         "     chr  ",
         "      chr ",
         "       chr",
         "        ch",
         "         c",
         "h         " }
)

func drawCustomer (column uint, veryHairy bool, F col.Colour) {
  scr.Lock ()
  if veryHairy {
    scr.ColourF (F)
    scr.Write (" w ", line, column)
    scr.Write (" § ", line + 1, column)
  } else {
    scr.ColourF (col.LightGreen())
    scr.Write ("   ", line, column)
    scr.Write (" ö ", line + 1, column)
  }
  scr.Write ("/O\\", line + 2, column)
  scr.Write ("|=|", line + 3, column)
  scr.Unlock ()
}

func takeSeatInWaitingRoom() {
  drawCustomer (columnC + 4 * N, true, col.LightRed())
  N++
  if N == 10 { N = 0 }
}

func w() {
  scr.Unlock ()
  const t = 500
  time.Sleep (time.Duration(rand.Intn(t) * 1e6))
  scr.Lock()
}

func cut () {
  scr.Lock(); defer scr.Unlock()
  scr.ColourF (col.Orange())
  scr.Write ("ö   ", line    , columnB - 4)
  scr.Write ("O=8<", line + 1, columnB - 4)
  scr.Write ("|   ", line + 2, columnB - 4)
  scr.Write ("|_  ", line + 3, columnB - 4)
  scr.ColourF (col.Yellow())
  scr.Write (  "8<", line + 1, columnB - 2)
  w()
  scr.ColourF (col.Orange())
  scr.Write ("ö 8<", line    , columnB - 4)
  scr.Write ("O/  ", line + 1, columnB - 4)
  scr.Write ("|\\  ",line + 2, columnB - 4)
  scr.Write ("||_ ", line + 3, columnB - 4)
  scr.ColourF (col.Yellow())
  scr.Write (  "8<", line    , columnB - 2)
  w()
  scr.Write ("    ", line    , columnB - 4)
  scr.Write ("    ", line + 1, columnB - 4)
  scr.Write ("    ", line + 2, columnB - 4)
  scr.Write ("    ", line + 3, columnB - 4)
  scr.ColourF (col.Orange())
  scr.Write (">8 ö", line    , columnB + 3)
  scr.Write ("  \\O",line + 1, columnB + 3)
  scr.Write ("   |", line + 2, columnB + 3)
  scr.Write ("  _|", line + 3, columnB + 3)
  scr.ColourF (col.Yellow())
  scr.Write (">8"    , line    , columnB + 3)
  w()
  scr.ColourF (col.Orange())
  scr.Write ("   ö", line    , columnB + 3)
  scr.Write (">8=O", line + 1, columnB + 3)
  scr.Write ("  /|", line + 2, columnB + 3)
  scr.Write (" _||", line + 3, columnB + 3)
  scr.ColourF (col.Yellow())
  scr.Write (">8"    , line + 1, columnB + 3)
  w()
  scr.Write ("    ", line    , columnB + 3)
  scr.Write ("    ", line + 1, columnB + 3)
  scr.Write ("    ", line + 2, columnB + 3)
  scr.Write ("    ", line + 3, columnB + 3)
}

func getNextCustomer() {
  barberSleeping = false
  drawCustomer (columnC + 4 * K, true, col.Black())
  K ++
  if K == 10 { K = 0 }
  time.Sleep (1e9)
  drawCustomer (columnB, true, col.LightGreen())
  for i:= 0; i <= 15; i++ { cut () }
  drawCustomer (columnB, false, col.LightGreen())
  time.Sleep (2e9)
  drawCustomer (columnB, true, col.Black())
  barberSleeping = true
}

func snore() {
  scr.Lock(); defer scr.Unlock()
  for {
    if barberSleeping {
      scr.ColourF (col.Orange())
    } else {
      scr.ColourF (col.Black())
    }
    scr.Write ("o    ", line + 1, column0)
    if barberSleeping {
    } else {
      scr.Write ("          ", line + 1, column0 + 1)
    }
    scr.Write ("|\\_  ", line + 2, column0)
    scr.Write ("   |_", line + 3, column0)
    if barberSleeping {
      scr.ColourF (col.LightBlue())
    } else {
      scr.ColourF (col.Black())
    }
    scr.Write ("\\",    line + 2, column0 - 1)
    scr.Write ("|==\\", line + 3, column0 - 1)
    if barberSleeping {
      scr.ColourF (col.White())
    }
    for i:= 0; i <= 16; i++ {
      if barberSleeping {
        scr.ColourF (col.White())
      } else {
        scr.ColourF (col.Black())
      }
      scr.Write (tx[i], line + 1, column0 + 2)
      scr.Warp (5, 0)
      scr.Switch (false)
      scr.Unlock ()
      time.Sleep (7e7)
      scr.Lock()
    }
  }
}
