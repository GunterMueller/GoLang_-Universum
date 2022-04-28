package fontsize

// (c) Christian Maurer   v. 220420 - license see µU.go

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
const
  M = 6 // len names
var
  Name []string

// Returns a string of len 2, that uniquely defines f and s.
func Code (f Font, s Size) string { return code(f,s) }

// Returns the width resp. the height of a font in size s;
// for prt in pt and for scr in px.
func Wd (s Size) uint { return wd(s) }
func Ht (s Size) uint { return ht(s) }
