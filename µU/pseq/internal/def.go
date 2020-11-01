package internal

// (c) Christian Maurer   v. 201014 - license see µU.go

import
  . "µU/obj"
type
  File interface {

  Fin()
  Name (n string)
  Rename (n string)
  Empty() bool
  Clr()
  Length() uint64
  Seek (p uint64)
  Position() uint64

  Read (s Stream) (int, error)
  Write (s Stream) (int, error)
}

func DirectLength (n string) uint64 { return directLength(n) }
func Erase (n string) { erase(n) }

func New() File { return new_() }
