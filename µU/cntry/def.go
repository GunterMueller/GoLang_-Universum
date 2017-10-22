package cntry

// (c) Christian Maurer   v. 170919 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
const ( // Format
  Tld = iota // Top-Level-Domain, z.B. "de",          "it"
  Long       // Bezeichnung,      z.B. "Deutschland", "Italien"
  Tel        // Telefon-Vorwahl,  z.B. 49,            39
  Car        // Kfz-Kennzeichen,  z.B. "D",           "I"
  Ioc        // IOC-Code,         z.B. "GER",         "ITA"
  Fifa       // FIFA-Code,        z.B. "GER",         "ITA"
  NFormats
)
type
  Continent byte; const (
  Europa = iota
  Afrika
  Amerika
  Asien
  Ozeanien
  NContinents
)
type
  Country interface {

  Object
  col.Colourer
  Editor
  Formatter
  Stringer
  Printer

  InContinent (c Continent) bool
}

// Returns a new empty (undefined) Country.
func New() Country { return new_() }
