package fontsize

// (c) Christian Maurer   v. 230112 - license see µU.go

type
  Size byte; const ( // for prt  for screen
  Tiny = Size(iota)  // cmtt6     7 *  5 px
  Small              // cmtt8    10 *  6 px
  Normal             // cmtt10   16 *  8 px
  Big                // cmtt12   24 * 12 px
  Large              // cmtt14   28 * 14 px
  Huge               // cmtt17   32 * 16 px
  NSizes
)
