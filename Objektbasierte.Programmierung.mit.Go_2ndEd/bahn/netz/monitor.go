package netz

// (c) Christian Maurer    v. 230306 - license see µU.go  

import (
  . "µU/obj"
  "µU/host"
  "µU/fmon"
)
const (
  freigeben = uint(iota)
  besetzen
  besetzt
  nOps
)
type
  Einfahrt interface {

  EinfahrtFreigeben (n uint)
  EinfahrtBesetzen (n uint)
  EinfahrtBesetzt (n uint) bool
}
type
  mon struct {
             fmon.FarMonitor
             }
var (
  server = host.Localhost().String()
  monitor Einfahrt
  aktiv = false
  einfahrtFrei [A*A]bool
)

func init() {
  for n := uint(0); n < A; n++ {
    for i := uint(0); i < AnzahlNachbarn (n); i++ {
      einfahrtFrei[A * n + Nachbar (n, i)] = true
    }
  }
}

func aktivieren() {
  monitor = New (server, 2345, MeinBahnhof >= N - 1)
}

func New (h string, p uint16, s bool) Einfahrt {
  fs := func (a any, i uint) any {
          n := a.(uint)
          switch i {
          case freigeben:
            einfahrtFrei[n] = true
          case besetzen:
            einfahrtFrei[n] = false
          case besetzt:
            if ! einfahrtFrei[n] { return uint(1) }
          }
          return uint(0)
        }
  m := new(mon)
  m.FarMonitor = fmon.New (uint(0), nOps, fs, AllTrueSp, h, p, s)
  return m
}

func (m *mon) EinfahrtFreigeben (n uint) {
  m.F (n, freigeben)
}

func (m *mon) EinfahrtBesetzen (n uint) {
  m.F (n, besetzen)
}

func (m *mon) EinfahrtBesetzt (n uint) bool {
  return m.F (n, besetzt).(uint) == 1
}

func einfahrtFreigeben (n uint) {
  monitor.EinfahrtFreigeben (n)
}

func einfahrtBesetzen (n uint) {
  monitor.EinfahrtBesetzen (n)
}

func einfahrtBesetzt (n uint) bool {
  return monitor.EinfahrtBesetzt (n)
}
