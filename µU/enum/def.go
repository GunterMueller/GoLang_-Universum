package enum

// (c) Christian Maurer   v. 221001 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
const
  M = uint(27) // maximal length of the strings
type
  Enum interface { // sequences of strings of the same length

// Pre: All s[i] are pairwise different and have a length <= n,
//      where n is the stringlength given by New.
  Set (s ...string)

// x has the strings edited by the user.
// They are saved in a file whose name is determined by n.
  SetEdit (n string, l, c uint)

// Pre: All k[i] are pairwise different and have a length <= k,
//      where k is the shortcutlength given by Newk.
//      No k[i] contains blank spaces.
  Setk (k ...string)

// x has the strings and shortcuts edited by the user.
// They are saved in files whose names are determined by n.
  SetEditk (n string, l, c uint)

  Get (n string)

  Editor
  col.Colourer

// x is the object selected by the user.
  Select (l, c uint)

// String returns the selected string.
  Stringer

  TeXer

  Print (l, c uint)

// Returns the number of strings of x given by Set.
  Num() uint

// Returns the width of the strings of x.
  Width() uint
}

// Pre: 0 < k <= M.
// Returns a new enumerator with stringlength n.
func New (n uint) Enum { return new_(n) }

// Pre: 0 < k < n <= M.
// Returns a new enumerator with stringlength n
// and shortcutlength k.
func Newk (n, k uint) Enum { return newk(n,k) }
