package internal

// (c) Christian Maurer   v. 220715 - license see µU.go

import
  . "µU/obj"
type
  File interface {

  Fin()
  Name (n string)
  Rename (n string)
  Empty() bool
  Clr()
  Length() uint
  Seek (p uint)
  Position() uint

  Read (s Stream) (int, error)
  Write (s Stream) (int, error)
}

func DirectLength (n string) uint { return directLength(n) }
func Erase (n string) { erase(n) }

func New() File { return new_() }
