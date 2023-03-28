package main

// (c) Christian Maurer   v. 220830 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/env"
  "µU/kbd"
  "µU/scr"
  "µU/str"
  "µU/errh"
  "µU/files"
  "µU/pseq"
  "µU/set"
  "µU/masks"
  "µU/mol"
)

func sub (x, y Rotator) bool {
  return x.(mol.Molecule).Sub (y.(mol.Molecule))
}

func main() {
  help := []string {" vor-/rückwärts: Pfeiltaste ab-/aufwärts",
                    "zum Anfang/Ende: Pos1/Ende              ",
                    " Eintrag ändern: Enter                  ",
                    "       einfügen: Einfg                  ",
                    "      entfernen: Entf                   ",
                    "         suchen: F2                     ",
                    "       umordnen: F3                     ",
                    "        beenden: Esc                    "}
  for i, he := range (help) { help[i] = str.Lat1 (he) }
  files.Cds()
  ms := masks.New()
  name := env.Arg(1)
  str.OffSpc (&name)
  if name == "" {
    ker.Panic ("Dem Aufruf von inferno muss der Name als Parameter mitgegeben werden!")
  }
  ms.Name (name)
  if pseq.Length (name + mol.Suffix) > 0 && env.NArgs() > 1 {
    ker.Panic ("Nach " + name + " keine weiteren Parameter eingeben!")
  }
  h_file := pseq.New (uint(0))
  h_file.Name (name + ".h.dat")
  var w, h uint
  if ms.Empty() {
    if env.NArgs() < 3 {
      ker.Panic ("Nach " + name + " auch die Zeilenzahl und Spaltenzahl als Parameter eingeben!")
    }
    h, w = env.N(2), env.N(3)
    if h <= 1 {
      ker.Panic ("Der zweite Parameter (die Zeilenzahl) muss größer als 1 sein!")
    }
    if w <= 24 {
      ker.Panic ("Der dritte Parameter (die Spaltenzahl) muss größer als 24 sein!")
    }
    h_file.Seek (0); h_file.Put (h)
    h_file.Seek (1); h_file.Put (w)
  } else {
    h_file.Seek (0); h = h_file.Get().(uint)
    h_file.Seek (1); w = h_file.Get().(uint)
  }
  h_file.Fin()
  scr.NewWH (2, 24, 8 * w, 16 * h); defer scr.Fin()
  scr.Name ("inferno " + name)
  if ms.Empty() {
    ms.Edit()
  } else {
    ms.Write()
  }
  m := mol.New()
  if pseq.Length (name + mol.Suffix) == 0 {
    m.Construct (name)
  } else {
    m = mol.Constructed (name)
  }
  m.Write (0, 0)
  file := pseq.New (m)
  file.Name (name + ".seq")
  all := set.New (m)
  for i := uint(0); i < file.Num(); i++ {
    file.Seek (i)
    m = file.Get().(mol.Molecule)
    all.Ins (m)
  }
  errh.Hint ("Hilfe: F1   Ende: Esc")
//  m.DefineName (name)
  all.Jump (false)
  if all.Empty() {
    for {
      m.Clr()
      m.Edit (0, 0)
      if m.Empty() {
        // return ?
      } else {
        all.Ins (m)
        break
      }
    }
  }
  loop:
  for {
    m = all.Get().(mol.Molecule)
    m.Write (0, 0)
    switch c, _ := kbd.Command(); c {
    case kbd.Esc:
      break loop
    case kbd.Help:
      errh.Help (help)
    case kbd.Enter:
      m1 := m.Clone().(Rotator)
      m.Edit (0, 0)
      if ! m.Eq (m1) {
        all.Del()
        all.Put (m)
      }
    case kbd.Up:
      all.Step (false)
    case kbd.Down:
      all.Step (true)
    case kbd.Pos1:
      all.Jump (false)
    case kbd.End:
      all.Jump (true)
    case kbd.Ins:
      m.Clr()
      m.Edit (0, 0)
      if m.Empty() {
        ker.Panic ("Molekül ist leer")
      }
      all.Ins (m) 
    case kbd.Del:
      if errh.Confirmed() {
        all.Del()
      }
    case kbd.Search:
      m.Clr()
      m.Edit (0, 0)
      if ! m.Empty() {
        all.Jump (false)
        loop1:
        for {
          m1 := all.Get().(Rotator)
          if sub (m, m1) {
            m1.(Editor).Write (0, 0)
            for {
              switch c1, _ := kbd.Command(); c1 {
              case kbd.Esc:
                break loop1
              case kbd.Down:
                for ! all.Eoc (true) {
                  all.Step (true)
                  m2 := all.Get().(Rotator)
                  if sub (m, m2) {
                    m2.(Editor).Write (0, 0)
                    break
                  }
                }
              case kbd.Up:
                for ! all.Eoc (false) {
                  all.Step (false)
                  m2 := all.Get().(Rotator)
                  if sub (m, m2) {
                    m2.(Editor).Write (0, 0)
                    break
                  }
                }
              case kbd.Del:
                if errh.Confirmed() {
                  all.Del()
                }
              }
            }
          }
          if all.Eoc (true) {
            all.Jump (false)
            break
          }
          all.Step (true)
        }
      }
    case kbd.Act:
      m.Rotate()
      all.Sort()
    case kbd.Print:
      ms.Print()
      m.Print()
    }
  }
  errh.DelHint()
  file.Clr()
  all.Trav (func (a any) { file.Ins (a.(mol.Molecule)) } )
  file.Fin()
}
