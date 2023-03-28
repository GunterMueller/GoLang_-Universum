package net

// (c) Christian Maurer   v. 230112 - license see µU.go

import (
  "µU/ker"
  "µU/kbd"
  "µU/str"
  "µU/col"
  "µU/fontsize"
  "µU/scr"
  "µU/N"
  "µU/errh"
//  "µU/ppm"
  "µU/gra"
  . "bus/line"
  "bus/track"
  "bus/stat"
)
var (
  station, trk = stat.New(), track.New()
  netgraph = gra.New (false, station, trk)
  lastX, lastY float64
  help = []string {"Start und                         ",
                   "Ziel auswählen: linke Maustaste   ",
                   "größer/kleiner: Eingabe-/Rücktaste",
                   "   verschieben: Pfeiltasten       ",
                   "       beenden: Esc               "}
)

func init() {
  for i, h := range (help) { help[i] = str.Lat1 (h) }
  constructNet()
}

func write (blinkend bool) {
  scr.Buf (true)
  netgraph.Trav1Cond (func (a any, b bool) { a.(track.Track).Write (b) })
  netgraph.TravCond (func (a any, b bool) { a.(stat.Station).Write (b) })
  scr.Buf (false)
}

func selected1() bool {
  dummy := stat.New()
//  ok := false
  for {
    c, _ := kbd.Command()
    scr.MousePointer (true)
    switch c {
    case kbd.Esc:
      return false
    case kbd.Enter, kbd.Back, kbd.Left, kbd.Right, kbd.Up, kbd.Down:
      dummy.EditScale()
      write (false)
    case kbd.Help:
      errh.Help (help)
    case kbd.Here:
//    case kbd.To:
      if netgraph.ExPred (func (a any) bool { return a.(stat.Station).UnderMouse() }) {
        return true
      }
/*/
    case kbd.Drag, kbd.Drop, kbd.Move:
      dummy.EditScale()
      write (false)
/*/
/*/
    case kbd.Here:
      ok == netgraph.ExPred (func (a any) bool { return a.(stat.Station).UnderMouse() })
    case kbd.There:
      x, y = scr.MousePosGr()
      if ok {
        st = netgraph.Get().(stat.Station)
        st.Rescale (uint(x), uint(y))
        netgraph.Put (st)
        write (false)
      }
/*/
    case kbd.Print:
      errh.DelHint()
//      ppm.Print ("U- und S-Bahn")
    }
  }
  return false
}

func selected() bool {
  loop:
  for {
    f := scr.ActFontsize()
    scr.SetFontsize (fontsize.Normal)
    errh.Hint ("Start auswählen     (Klick mit linker Maustaste)")
    scr.SetFontsize (f)
    if selected1() { // Start aktuell
      netgraph.Get().(stat.Station).Write (true)
      netgraph.Locate (true) // Start postaktuell
      write (false)
    } else {
      break
    }
    f = scr.ActFontsize()
    scr.SetFontsize (fontsize.Normal)
    errh.Hint ("Ziel auswählen     (Klick mit linker Maustaste)")
    scr.SetFontsize (f)
    for {
      if selected1() { // Ziel aktuell
        if ! netgraph.Located() {
          netgraph.Get().(stat.Station).Write (true)
          errh.DelHint()
          return true
        }
      } else {
        break loop
      }
    }
  }
  write (false)
  return false
}

func shortestPath() {
  const
    maxU = 8
  var (
    startline, destline [maxU]Line
    startnumber, destnumber [maxU]uint
    imin, kmin uint
  )
  t1, t2 := netgraph.Get2()
  st1, st2 := t1.(stat.Station), t2.(stat.Station)
  startline [0], startnumber [0] = st1.Line(), st1.Number()
  destline [0], destnumber [0] = st2.Line(), st2.Number()
  ss, zz := 1, 1
  netgraph.Trav (func (a any) {
               st := a.(stat.Station)
               if st.Equiv (st1) {
                 startline [ss], startnumber [ss] = st.Line(), st.Number()
                 ss ++
               } else if st.Equiv (st2) {
                 destline [zz], destnumber [zz] = st.Line(), st.Number()
                 zz ++
               }
            })
  nmin := uint(ker.MaxNat)
  for i := 0; i < ss; i++ {
    for k := 0; k < zz; k++ {
      l, j := startline [i], startnumber [i]
      l1, j1 := destline [k], destnumber [k]
      if ! netgraph.ExPred2 (func (a any) bool {
                               s := a.(stat.Station)
                               return s.Line() == l && s.Number() == j
                             }, func (a any) bool {
                                  s := a.(stat.Station)
                                  return s.Line() == l1 && s.Number() == j1
                                }) {
        ker.ToDo()
      }
      netgraph.FindShortestPath()
      lm := netgraph.LenMarked()
      if lm < nmin {
        nmin = lm
        imin, kmin = uint(i), uint(k)
      }
    }
  }
  l, j := startline [imin], startnumber [imin]
  l1, j1 := destline [kmin], destnumber [kmin]
  if ! netgraph.ExPred2 (func (a any) bool {
                           s := a.(stat.Station)
                           return s.Line() == l && s.Number() == j
                         }, func (a any) bool {
                              s := a.(stat.Station)
                              return s.Line() == l1 && s.Number() == j1
                            }) {
    ker.ToDo()
  }
  netgraph.FindShortestPath()
  write (true)
  scr.Colours (col.HintF(), col.HintB())
  scr.Transparence (false)
  na := netgraph.LenMarked()
  f := scr.ActFontsize()
  scr.SetFontsize (fontsize.Normal)
  scr.Write ("Fahrzeit etwa " + N.String (na) + " Minuten", 0, 0)
  scr.SetFontsize (f)
  scr.Transparence (true)
}
