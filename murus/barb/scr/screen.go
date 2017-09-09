package scr

// (c) Christian Maurer   v. 170818 - license see murus.go

import (
  "murus/ker"
  "murus/rand"
  "murus/col"
  "murus/scr"
)
const (
  line = 1
  column0 = 3; columnB = 20; columnC = 30
)
var (
  N, K uint
  barberSleeping bool
  tx []string
)

func drawCustomer (column uint, veryHairy bool, F col.Colour) {
  scr.Lock ()
  if veryHairy {
    scr.ColourF (F)
    scr.Write (" w ", line, column)
    scr.Write (" § ", line + 1, column)
  } else {
    scr.ColourF (col.LightGreen)
    scr.Write ("   ", line, column)
    scr.Write (" ö ", line + 1, column)
  }
  scr.Write ("/O\\", line + 2, column)
  scr.Write ("|=|", line + 3, column)
  scr.Unlock ()
}

func takeSeatInWaitingRoom() {
  drawCustomer (columnC + 4 * N, true, col.CinnabarRed)
  N++
  if N == 10 { N = 0 }
}

func w() {
  scr.Unlock ()
  const t = 500
  ker.Msleep (rand.Natural (t))
  scr.Lock ()
}

func cut () {
  scr.Lock(); defer scr.Unlock()
  scr.ColourF (col.LightOrange)
  scr.Write ("ö   ", line    , columnB - 4)
  scr.Write ("O=8<", line + 1, columnB - 4)
  scr.Write ("|   ", line + 2, columnB - 4)
  scr.Write ("|_  ", line + 3, columnB - 4)
  scr.ColourF (col.Yellow)
  scr.Write (  "8<", line + 1, columnB - 2)
  w()
  scr.ColourF (col.LightOrange)
  scr.Write ("ö 8<", line    , columnB - 4)
  scr.Write ("O/  ", line + 1, columnB - 4)
  scr.Write ("|\\  ",line + 2, columnB - 4)
  scr.Write ("||_ ", line + 3, columnB - 4)
  scr.ColourF (col.Yellow)
  scr.Write (  "8<", line    , columnB - 2)
  w()
  scr.Clr (line, columnB - 4, 4, 4)
  scr.ColourF (col.LightOrange)
  scr.Write (">8 ö", line    , columnB + 3)
  scr.Write ("  \\O",line + 1, columnB + 3)
  scr.Write ("   |", line + 2, columnB + 3)
  scr.Write ("  _|", line + 3, columnB + 3)
  scr.ColourF (col.Yellow)
  scr.Write (">8"    , line    , columnB + 3)
  w()
  scr.ColourF (col.LightOrange)
  scr.Write ("   ö", line    , columnB + 3)
  scr.Write (">8=O", line + 1, columnB + 3)
  scr.Write ("  /|", line + 2, columnB + 3)
  scr.Write (" _||", line + 3, columnB + 3)
  scr.ColourF (col.Yellow)
  scr.Write (">8"    , line + 1, columnB + 3)
  w()
  scr.Clr (line, columnB + 3, 4, 4)
}

func getNextCustomer() {
  barberSleeping = false
  drawCustomer (columnC + 4 * K, true, col.Black)
  K ++
  if K == 10 { K = 0 }
  ker.Sleep (1)
  drawCustomer (columnB, true, col.LightGreen)
  for i:= 0; i <= 15; i++ { cut () }
  drawCustomer (columnB, false, col.LightGreen)
  ker.Sleep (2)
  drawCustomer (columnB, true, col.Black)
  barberSleeping = true
}

func snore() {
  scr.Lock(); defer scr.Unlock()
  for {
    if barberSleeping {
      scr.ColourF (col.LightOrange)
    } else {
      scr.ColourF (col.Black)
    }
    scr.Write ("o    ", line + 1, column0)
    if barberSleeping {
    } else {
      scr.Write ("          ", line + 1, column0 + 1)
    }
    scr.Write ("|\\_  ", line + 2, column0)
    scr.Write ("   |_", line + 3, column0)
    if barberSleeping {
      scr.ColourF (col.XLightSkyBlue)
    } else {
      scr.ColourF (col.Black)
    }
    scr.Write ("\\",    line + 2, column0 - 1)
    scr.Write ("|==\\", line + 3, column0 - 1)
    if barberSleeping {
      scr.ColourF (col.Silver)
    }
    for i:= 0; i <= 16; i++ {
      if barberSleeping {
        scr.ColourF (col.Silver)
      } else {
        scr.ColourF (col.Black)
      }
      scr.Write (tx[i], line + 1, column0 + 2)
      scr.Unlock ()
      const t = 200 * 1000
      ker.Usleep (rand.Natural (t))
      scr.Lock()
    }
  }
}

func init() {
  scr.NewWH (0, 0, 480, 6 * 16)
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
  go snore()
}
