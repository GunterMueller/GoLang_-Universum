package ptr

// (c) Christian Maurer   v. 210312 - license see µU.go

type
  Pointer byte; const ( // appearence of the mouse "arrow"
  Standard = Pointer(iota)
  Gumby
  Hand
  Gobbler
  Watch
  NPointers
)

// Returns the number of p in /usr/include/X11/cursorfont.h.
func Code (p Pointer) uint { return code(p) }
