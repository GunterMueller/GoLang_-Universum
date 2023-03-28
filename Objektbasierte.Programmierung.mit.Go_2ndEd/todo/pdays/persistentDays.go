package pdays

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "µU/env"
  "µU/str"
  "µU/set"
  "µU/day"
  "µU/pseq"
  "µU/files"
)
const
  suffix = "day"
type
  persistentDays struct {
                        set.Set
                        pseq.PersistentSequence
                        string "name"
                        bool "changed"
                        }

func init() {
  files.Cd (env.Gosrc() + "/todo/")
}

func New() PersistentDays {
  x := new (persistentDays)
  d := day.New()
  x.PersistentSequence = pseq.New (d)
  x.Set = set.New (d)
  return x
}

func (x *persistentDays) Name (s string) {
  x.string = s
  str.OffSpc (&x.string)
  x.PersistentSequence.Name (x.string + "." + suffix)
  x.Set.Clr()
  x.PersistentSequence.Trav (func (a any) { x.Set.Ins (a.(day.Calendarday)) })
}

func (x *persistentDays) Empty() bool {
  return x.Set.Empty()
}

func (x *persistentDays) Clr() {
  x.Set.Clr()
}

func (x *persistentDays) Num() uint {
  return x.Set.Num()
}

func (x *persistentDays) Ex (d day.Calendarday) bool {
  return x.Set.Ex (d)
}

func (x *persistentDays) Ins (d day.Calendarday) {
  if ! x.Set.Ex (d) {
    x.Set.Ins (d)
    x.bool = true
  }
}

func (x *persistentDays) Del (d day.Calendarday) {
  x.bool = x.Set.Ex (d)
  if x.bool {
    x.Set.Del()
  }
}

func (x *persistentDays) Fin() {
  if x.bool {
    x.PersistentSequence.Clr()
    x.Set.Trav (func (a any) { x.PersistentSequence.Ins (a.(day.Calendarday)) })
    x.PersistentSequence.Fin()
  }
}
