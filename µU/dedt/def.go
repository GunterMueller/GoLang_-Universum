package dedt

// (c) Christian Maurer   v. 210122 - license see µU.go

import
  . "µU/obj"
type
  DistributedEditor interface {

  ReaderIn() Object
  EditorIn() Object
  ReaderOut()
  EditorOut (o Object)
  Fin()
}

func New (o Object, h string, p uint16, s bool) DistributedEditor { return new_(o,h,p,s) }
