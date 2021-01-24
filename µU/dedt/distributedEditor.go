package dedt

// (c) Christian Maurer   v. 210122 - license see µU.go

import (
  . "µU/obj"
  "µU/fmon"
)
const (
  readerIn = iota
  editorIn
  readerOut
  editorOut
)
type
  editor struct {
                fmon.FarMonitor
                Object
                }

func new_(o Object, h string, p uint16, s bool) DistributedEditor {
  var nR, nE uint
  x := new(editor)
  ps := func (a Any, i uint) bool {
          switch i {
          case readerIn:
            return nE == 0
          case editorIn:
            return nR == 0 && nE == 0
          }
          return true // readerOut, editorOut
        }
  fs := func (a Any, i uint) Any {
          switch i {
          case readerIn:
            nR++
            return Clone (x.Object)
          case readerOut:
            nR--
            return nR
          case editorIn:
            nE = 1
            return Clone (x.Object)
          case editorOut:
            nE = 0
            x.Object.Copy (a)
          }
          return nE
        }
  x.FarMonitor = fmon.New (o, 4, fs, ps, h, p, s)
  x.Object = Clone(o).(Object)
  return x
}

func (x *editor) ReaderIn() Object {
  return x.F (x.Object, readerIn).(Object)
}

func (x *editor) ReaderOut() {
  x.F (x.Object, readerOut)
}

func (x *editor) EditorIn() Object {
  return x.F (x.Object, editorIn).(Object)
}

func (x *editor) EditorOut (o Object) {
  x.F (o, editorOut)
}

func (x *editor) Fin() {
  x.FarMonitor.Fin()
}
