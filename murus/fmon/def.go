package fmon

// (c) murus.org  v. 170104 - license see murus.go

//     Nichtsequentielle Programmierung mit Go 1 kompakt,
//     Kapitel 8, insbesondere Abschnitt 8.3

import (
  . "murus/obj"
  "murus/host"
)

// Pre: XXX
// Returns a new far monitor.
func New (a Any, n uint, fs FuncSpectrum, ps PredSpectrum,
     h host.Host, p uint16, s bool) FarMonitor {
          return new_(a,n,fs,ps,h,p,s) }

func NewS (a Any, n uint, fs FuncSpectrum, ps PredSpectrum,
     h host.Host, p uint16, s bool, stmt Stmt) FarMonitor {
          return news(a,n,fs,ps,h,p,s,stmt) }

type
  FarMonitor interface {
// Specifications: Buy my book and read chapter 8.

  F (a Any, i uint) Any
//  F0 (a Any, i uint)
//  S (a Any, i uint, c chan Any) // experimental

  Fin()
}
