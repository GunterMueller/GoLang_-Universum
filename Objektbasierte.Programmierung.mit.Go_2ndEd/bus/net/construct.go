package net

// (c) Christian Maurer   v. 230114 - license see µU.go

import (
  "µU/ker"
  "µU/str"
  "µU/errh"
  . "bus/line"
  "bus/stat"
)
const
  tU = 5 // stark vereinfacht: universelle Zeit zum Umsteigen

// Aktuelle Ecke ist (l1, k1), postaktuelle Ecke ist (l, k).
func edge1 (l Line, k uint, l1 Line, k1, t uint, direkt bool) {
  if l == Zoo && l1 == BG { return }
  if l == BG&& l1 == Zoo { return }
  if netgraph.ExPred2 (func (a any) bool {
                         s := a.(stat.Station)
                         return s.Line() == l && s.Number() == k
                       },
                       func (a any) bool {
                         s := a.(stat.Station)
                         return s.Line() == l1 && s.Number() == k1
                       }) {
    b, b1 := netgraph.Get2()
    st, st1:= b.(stat.Station), b1.(stat.Station)
    st.Umstieg()
    st1.Umstieg()
    netgraph.Put2 (st, st1)
    if direkt {
      l = st.Line()
    } else {
      l = Footpath
    }
    x, y := st.Pos()
    x1, y1 := st1.Pos()
    trk.Def (l, t)
    trk.SetPos (x, y, x1, y1)
    netgraph.Edge (trk)
  } else {
    if ! netgraph.ExPred (func (a any) bool {
                            s := a.(stat.Station)
                            return s.Line() == l && s.Number() == k
                          }) {
      errh.Error2 ("missing Line", uint(l),  "Station", k)
    }
    if ! netgraph.ExPred (func (a any) bool {
                            s := a.(stat.Station)
                            return s.Line() == l1 && s.Number() == k1
                          }) {
      errh.Error2 ("missing Line", uint(l1), "Station", k1)
    }
    ker.Oops()
  }
}

func edge (l Line, k uint, l1 Line, k1, t uint) {
  edge1 (l, k, l1, k1, t, imBahnhof)
}

// Aktuelle Ecke ist (l, nr), postaktuelle Ecke ist die, die vorher aktuell war.
func ins (l Line, nr uint, k string, b byte, y, x float64) {
  k = str.Lat1 (k)
  station.Set (l, nr, k, b, y, x)
  lastX, lastY = x, y
  netgraph.Ins (station)
}

// Aktuelle Ecke ist (l, nr), postaktuelle Ecke ist die, die vorher aktuell war.
// t ist die mittlere Fahrzeit vom Bahnhof aus der Zeile davor.
func ins1 (l Line, nr uint, n string, b byte, y, x float64, t uint) {
  x0, y0 := lastX, lastY
  ins (l, nr, n, b, y, x)
  trk.Def (l, t)
  trk.SetPos (x0, y0, x, y)
  netgraph.Edge (trk)
}

// s. ins
func insert (l Line, k uint, l0 Line, k0, u uint) {
  var st stat.Station
  if netgraph.ExPred (func (a any) bool {
                        b := a.(stat.Station)
                        return b.Line() == l0 && b.Number() == k0
                      }) {
    st = netgraph.Get().(stat.Station)
  } else {
    errh.Error2 ("missing: Line", uint(l0), "Station", k0); ker.Oops()
  }
// aktuell: (l0, k0)
  st1 := st.Clone().(stat.Station)
  st1.Renumber (l, k)
  st1.Umstieg()
  netgraph.Ins (st1)
  edge (l0, k0, l, k, u)
// postaktuell: (l0, k0), aktuell: (l, k)
  lastX, lastY = st.Pos()
}

// s. ins1
func insert1 (l Line, k uint, l0 Line, k0, u, t uint) {
// aktuell: vorherige aktuelle
  x0, y0 := lastX, lastY
  netgraph.Relocate()
  var st stat.Station
// aktuell: vorherige aktuelle, postaktuell: alt.
  if netgraph.ExPred (func (a any) bool {
                        b := a.(stat.Station)
                        return b.Line() == l0 && b.Number() == k0
                      }) {
   // aktuell: (L0, n0), postaktuell: vorherige aktuelle
    st = netgraph.Get().(stat.Station)
  } else {
    errh.Error2 ("missing: Line", uint(l0), "Station", k0); ker.Oops()
  }
  netgraph.Relocate()
// aktuell: vorherige aktuelle, postaktuell: (l0, k0)
  st1 := st.Clone().(stat.Station)
  st1.Renumber (l, k)
  st1.Umstieg()
  netgraph.Ins (st1)
// aktuell: (l, k), postaktuell: vorherige aktuelle
  x, y := st1.Pos()
  lastX, lastY = x, y
  trk.Def (l, t)
  trk.SetPos (x0, y0, x, y)
  netgraph.Edge (trk)
// verbunden: vorherige aktuelle mit (l, k), aktuell: (l, k).
  edge (l0, k0, l, k, u)
// postaktuell: (l0, k0), aktuell: (l, k)
  station = st1
}

const (
  imBahnhof = true
  mitFußweg = false
  L = stat.L
  R = stat.R
  O = stat.O
  U = stat.U
)

func constructNet() {
  ins     (U1, 10, "Uhlandstr___",                  U, 52.5030, 13.3276)
  ins1    (U1, 11, "  Kurfürstendamm",              O, 52.5038, 13.3314, 1)
  ins1    (U1, 12, "Wittenbergplatz",               R, 52.5018, 13.3430, 2)
  ins1    (U1, 13, "Nollendorfplatz",               L, 52.4994, 13.3535, 2)
  ins1    (U1, 14, "Kurfürstenstr",                 O, 52.5001, 13.3615, 1)
  ins1    (U1, 15, " Gleisdreieck",                 U, 52.4992, 13.3753, 2)
  ins1    (U1, 16, "   Möckernbrücke",              O, 52.4991, 13.3828, 2)
  ins1    (U1, 17, "Hallesches Tor",                U, 52.4977, 13.3908, 1)
  ins1    (U1, 18, "Prinzenstr",                    U, 52.4984, 13.4056, 2)
  ins1    (U1, 19, "Kottbusser Tor",                O, 52.4994, 13.4187, 2)
  ins1    (U1, 20, "Görlitzer Bhf",                 U, 52.4993, 13.4282, 2)
  ins1    (U1, 21, "Schlesisches Tor",              U, 52.5009, 13.4415, 2)
  ins1    (U1, 22, "Warschauer Str.",               L, 52.5049, 13.4490, 2)

  ins     (U2, 10, "Ruhleben",                      O, 52.5256, 13.2413)
  ins1    (U2, 11, "Olympiastadion",                L, 52.5173, 13.2501, 3)
  ins1    (U2, 12, "Neu-Westend",                   R, 52.5162, 13.2592, 2)
  ins1    (U2, 13, "Theodor-Heuss-Platz",           L, 52.5100, 13.2727, 2)
  ins1    (U2, 14, "Kaiserdamm_",                   O, 52.5100, 13.2818, 1)
  ins1    (U2, 15, "Sophie-Charlotte-Platz",        U, 52.5108, 13.2954, 2)
  ins1    (U2, 16, "Bismarckstr____",               O, 52.5117, 13.3058, 1)
  ins1    (U2, 17, "  Deutsche Oper",               U, 52.5119, 13.3108, 1)
  ins1    (U2, 18, "Ernst-Reuter-Platz",            O, 52.5120, 13.3218, 2)
  ins1    (U2, 19, "Zoologischer Garten",           L, 52.5070, 13.3324, 2)
  insert1 (U2, 20, U1, 12,  2, 2) // Wittenbergplatz
  insert1 (U2, 21, U1, 13, tU, 2) // Hollendorfplatz
  ins1    (U2, 22, "Bülowstr",                      U, 52.4975, 13.3635, 1)
  insert1 (U2, 23, U1, 15, tU, 2) // Gleisdreieck
  ins1    (U2, 24, "Mendelssohn-Bartholdy-Park",    L, 52.5039, 13.3749, 1)
  ins1    (U2, 25, "Potsdamer Platz",               L, 52.5093, 13.3779, 2)
  ins1    (U2, 26, "Mohrenstr",                     O, 52.5117, 13.3845, 1)
  ins1    (U2, 27, "Stadtmitte",                    U, 52.5104, 13.3899, 1)
  ins1    (U2, 28, "  Hausvogteiplatz",             O, 52.5133, 13.3962, 2)
  ins1    (U2, 29, "Spittelmarkt____",              U, 52.5113, 13.4043, 1)
  ins1    (U2, 30, "Märkisches Museum",             R, 52.5124, 13.4103, 1)
  ins1    (U2, 31, "Klosterstr",                    U, 52.5172, 13.4125, 2)
  ins1    (U2, 32, "Alexanderplatz",                R, 52.5216, 13.4113, 2)
  ins1    (U2, 33, "Rosa-Luxemburg-Platz",          R, 52.5276, 13.4105, 2)
  ins1    (U2, 34, "Senefelderplatz",               R, 52.5326, 13.4126, 1)
  ins1    (U2, 35, "Eberswalder Str",               R, 52.5417, 13.4122, 2)
  ins1    (U2, 36, "Schönhauser Allee",             R, 52.5494, 13.4141, 2)
  ins1    (U2, 37, "Vinetastr",                     R, 52.5595, 13.4132, 2)
  ins1    (U2, 38, "Pankow",                        R, 52.5678, 13.4120, 2)

  insert  (U3,  1, U1, 22, tU)    // Warschauer Straße
  insert1 (U3,  2, U1, 21, tU, 2) // Schlesisches Tor
  insert1 (U3,  3, U1, 20, tU, 2) // Görlitzer Bahnhof
  insert1 (U3,  4, U1, 19, tU, 2) // Kottbusser Tor
  insert1 (U3,  5, U1, 18, tU, 2) // Prinzenstraße
  insert1 (U3,  6, U1, 17, tU, 2) // Hallesches Tor
  insert1 (U3,  7, U1, 16, tU, 1) // Möckernbrücke
  insert1 (U3,  8, U1, 15, tU, 2) // Gleisdreieck
  insert1 (U3,  9, U1, 14, tU, 2) // Kurfürstenstraße
  insert1 (U3, 10, U1, 13, tU, 1) // Nollendorfplatz
  insert1 (U3, 11, U1, 12, tU, 2) // Wittenbergplatz
  ins1    (U3, 12, "AugsburgerStr",                 L, 52.5005, 13.3366, 2)
  ins1    (U3, 13, "Spichernstr",                   L, 52.4963, 13.3308, 1)
  ins1    (U3, 14, "Hohenzollernplatz",             U, 52.4941, 13.3248, 2)
  ins1    (U3, 15, "Fehrbelliner Platz",            U, 52.4905, 13.3146, 1)
  ins1    (U3, 16, "Heidelberger Platz",            L, 52.4798, 13.3122, 2)
  ins1    (U3, 17, "Rüdesheimer Platz",             L, 52.4730, 13.3145, 2)
  ins1    (U3, 18, "Breitenbachplatz",              L, 52.4667, 13.3084, 1)
  ins1    (U3, 19, "Podbielskiallee",               L, 52.4641, 13.2958, 2)
  ins1    (U3, 20, "Dahlem Dorf",                   R, 52.4573, 13.2897, 1)
  ins1    (U3, 21, "Thielplatz",                    R, 52.4510, 13.2818, 2)
  ins1    (U3, 22, "Oskar-Helene-Heim",             U, 52.4504, 13.2690, 2)
  ins1    (U3, 23, "Onkel Toms Hütte",              L, 52.4499, 13.2531, 2)
  ins1    (U3, 24, "Krumme Lanke ",                 U, 52.4435, 13.2415, 2)

  insert  (U4, 10, U1, 13, tU) // Nollendorfplatz
  ins1    (U4, 11, "Viktoria-Luise-Platz",          U, 52.4960, 13.3428, 2)
  ins1    (U4, 12, "Bayerischer Platz",             U, 52.4885, 13.3401, 2)
  ins1    (U4, 13, "Rathaus Schöneberg",            U, 52.4831, 13.3420, 1)
  ins1    (U4, 14, "Innsbrucker Platz",             U, 52.4784, 13.3429, 2)

//  insert1 (U5, 9, U9, 16, tU, 3) 
//  ins1 (U5,  0, "Fritz-Schloß-Park",                L, 52.5259, 13.3430, 1)
  ins     (U5, 10, "Hauptbahnhof____",              O, 52.5252, 13.3691)
  ins1    (U5, 11, "Bundestag",                     L, 52.5202, 13.3730, 1)
  ins1    (U5, 12, "Brandenburger Tor",             L, 52.5165, 13.3812, 1)
  ins1    (U5, 13, "Unter den Linden",              U, 52.5169, 13.3889, 1)
  ins1    (U5, 14, "Museumsinsel",                  U, 52.5181, 13.4008, 1)
  ins1    (U5, 15, "Rotes Rathaus",                 O, 52.5186, 13.4078, 1)
  insert1 (U5, 16, U2, 32, tU, 3) // Alexanderplatz
  ins1    (U5, 17, "Schillingstr",                  R, 52.5204, 13.4219, 2)
  ins1    (U5, 18, "Strausberger Platz",            R, 52.5180, 13.4330, 2)
  ins1    (U5, 19, "Weberwiese",                    R, 52.5167, 13.4449, 1)
  ins1    (U5, 20, "Frankfurter Tor",               R, 52.5159, 13.4532, 1)
  ins1    (U5, 21, "Samariterstr",                  U, 52.5148, 13.4644, 2)
  ins1    (U5, 22, "Frankfurter Allee",             R, 52.5144, 13.4749, 1)
  ins1    (U5, 23, "Magdalenenstr",                 R, 52.5125, 13.4871, 2)
  ins1    (U5, 24, "Lichtenberg",                   L, 52.5106, 13.4970, 1)
  ins1    (U5, 25, "Friedrichsfelde",               U, 52.5056, 13.5128, 2)
  ins1    (U5, 26, "   Tierpark",                   O, 52.5056, 13.5235, 2)
  ins1    (U5, 27, "Biesdorf-Süd",                  U, 52.4997, 13.5464, 3)
  ins1    (U5, 28, "Elsterwerdaer Platz",           R, 52.5050, 13.5605, 2)
  ins1    (U5, 29, "Wuhletal",                      U, 52.5127, 13.5752, 2)
  ins1    (U5, 30, "Kaulsdorf Nord",                L, 52.5210, 13.5890, 2)
  ins1    (U5, 31, "Neue Grottkauer Str",           L, 52.5285, 13.5906, 2)
  ins1    (U5, 32, "Cottbusser Platz     ",         L, 52.5339, 13.5965, 1)
  ins1    (U5, 33, "Hellersdorf           ",        O, 52.5367, 13.6067, 2)
  ins1    (U5, 34, "Louis-Lewin-Str",               O, 52.5390, 13.6184, 1)
  ins1    (U5, 35, "Hönow",                         U, 52.5384, 13.6333, 2)

  ins     (U6, 10, "Alt Tegel",                     L, 52.5896, 13.2837)
  ins1    (U6, 11, "Borsigwerke",                   L, 52.5818, 13.2906, 2)
  ins1    (U6, 12, "Holzhauser Str",                L, 52.5757, 13.2961, 1)
  ins1    (U6, 13, "Otisstr",                       L, 52.5710, 13.3030, 1)
  ins1    (U6, 14, "Scharnweberstr",                L, 52.5668, 13.3127, 2)
  ins1    (U6, 15, "Kurt-Schumacher-Platz",         R, 52.5641, 13.3278, 1)
  ins1    (U6, 16, "Afrikanische Str",              L, 52.5602, 13.3348, 2)
  ins1    (U6, 17, "Rehberge       ",               L, 52.5568, 13.3410, 1)
  ins1    (U6, 18, "Seestr",                        L, 52.5500, 13.3528, 2)
  ins1    (U6, 19, "Leopoldplatz",                  L, 52.5464, 13.3591, 1)
  ins1    (U6, 20, "Wedding",                       L, 52.5430, 13.3651, 1)
  ins1    (U6, 21, "Reinickendorfer Str",           U, 52.5399, 13.3704, 1)
  ins1    (U6, 22, "Schwartzkopffstr",              L, 52.5355, 13.3769, 1)
  ins1    (U6, 23, "Zinnowitzer Str",               L, 52.5313, 13.3822, 1)
  ins1    (U6, 24, "Oranienburger Tor",             O, 52.5256, 13.3874, 2)
  ins1    (U6, 25, "Friedrichstr",                  L, 52.5201, 13.3880, 1)
  insert1 (U6, 26, U5, 13, tU, 2) // Unter den Linden
  ins1    (U6, 27, "Französische Str.",             L, 52.5148, 13.3892, 1)
  insert1 (U6, 28, U2, 27, tU, 2) // Stadtmitte
  ins1    (U6, 29, "Kochstr ",                      R, 52.5059, 13.3906, 1)
  insert1 (U6, 30, U1, 17, tU, 1) // Hallesches Tor
  ins1    (U6, 31, "Mehringdamm",                   L, 52.4943, 13.3888, 2)
  ins1    (U6, 32, "Platz der Luftbrücke",          R, 52.4852, 13.3860, 2)
  ins1    (U6, 33, "Paradestr",                     O, 52.4782, 13.3859, 1)
  ins1    (U6, 34, "Tempelhof",                     L, 52.4699, 13.3856, 1)
  ins1    (U6, 35, "Alt-Tempelhof",                 R, 52.4659, 13.3858, 2)
  ins1    (U6, 36, "Kaiserin-Augusta-Str",          R, 52.4594, 13.3845, 1)
  ins1    (U6, 37, "Ullsteinstr",                   R, 52.4527, 13.3845, 1)
  ins1    (U6, 38, "Westphalweg",                   R, 52.4465, 13.3857, 2)
  ins1    (U6, 39, "Alt-Mariendorf",                U, 52.4392, 13.3881, 1)

  ins     (U7, 10, "Rathaus Spandau",               L, 52.5353, 13.1999)
  ins1    (U7, 11, "Altstadt Spandau",              O, 52.5390, 13.2056, 1)
  ins1    (U7, 12, "Zitadelle",                     U, 52.5377, 13.2187, 1)
  ins1    (U7, 13, "Haselhorst",                    O, 52.5382, 13.2319, 2)
  ins1    (U7, 14, "Paulsternstr",                  O, 52.5379, 13.2476, 2)
  ins1    (U7, 15, "Rohrdamm",                      O, 52.5372, 13.2622, 1)
  ins1    (U7, 16, "Siemensdamm",                   U, 52.5367, 13.2729, 2)
  ins1    (U7, 17, "Halemweg",                      O, 52.5367, 13.2864, 1)
  ins1    (U7, 18, "Jakob-Kaiser-Platz",            R, 52.5363, 13.2944, 1)
  ins1    (U7, 19, "Jungfernheide",                 L, 52.5308, 13.3005, 2)
  ins1    (U7, 20, "Mierendorffplatz",              R, 52.5266, 13.3050, 1)
  ins1    (U7, 21, "Richard-Wagner-Platz",          R, 52.5161, 13.3067, 2)
  insert1 (U7, 22, U2, 16, tU, 1) // Bismarckstr.
  ins1    (U7, 23, "Wilmersdorfer Str___________",  O, 52.5067, 13.3067, 1)
  ins1    (U7, 24, "Adenauerplatz",                 R, 52.5000, 13.3071, 2)
  ins1    (U7, 25, "Konstanzer Str",                L, 52.4945, 13.3091, 1)
  insert1 (U7, 26, U3, 15, tU, 1) // Fehrbellinger Platz
  ins1    (U7, 27, "Blissestr",                     L, 52.4866, 13.3221, 2)
  ins1    (U7, 28, "Berliner Str",                  U, 52.4874, 13.3310, 1)
  insert1 (U7, 29, U4, 12, tU, 1) // Bayerischer Platz
  ins1    (U7, 30, "Eisenacher Str",                U, 52.4895, 13.3503, 1)
  ins1    (U7, 31, "Kleistpark",                    U, 52.4905, 13.3604, 2)
  ins1    (U7, 32, "Yorckstr",                      L, 52.4922, 13.3678, 1)
  insert1 (U7, 33, U1, 16, tU, 2) // Möckernbrücke
  insert1 (U7, 34, U6, 30, tU, 2) // Mehringdamm
  ins1    (U7, 35, "Gneisenaustr",                  R, 52.4913, 13.3958, 1)
  ins1    (U7, 36, " Südstern",                     U, 52.4891, 13.4077, 2)
  ins1    (U7, 37, "Hermannplatz",                  R, 52.4864, 13.4243, 2)
  ins1    (U7, 38, "Rathaus Neukölln",              R, 52.4819, 13.4339, 1)
  ins1    (U7, 39, "Karl Marx-Str",                 U, 52.4765, 13.4392, 2)
  ins1    (U7, 40, "  Neukölln",                    U, 52.4686, 13.4419, 1)
  ins1    (U7, 41, "Grenzallee",                    U, 52.4631, 13.4445, 2)
  ins1    (U7, 42, "Blaschkoallee",                 L, 52.4516, 13.4495, 1)
  ins1    (U7, 43, "Parchimer Allee",               L, 52.4457, 13.4450, 2)
  ins1    (U7, 44, "Britz Süd",                     L, 52.4377, 13.4475, 1)
  ins1    (U7, 45, "Johannistaler Chaussee",        L, 52.4294, 13.4533, 2)
  ins1    (U7, 46, "Lipschitzallee",                L, 52.4235, 13.4626, 1)
  ins1    (U7, 47, "Wutzkyallee",                   U, 52.4226, 13.4750, 2)
  ins1    (U7, 48, "Zwickauer Damm",                R, 52.4232, 13.4838, 1)
  ins1    (U7, 49, "Rudow",                         L, 52.4158, 13.4966, 1)

  ins     (U8, 10, "Wittenau",                      R, 52.5926, 13.3344)
  ins1    (U8, 11, "Rathaus Reinickendorf",         L, 52.5906, 13.3256, 1)
  ins1    (U8, 12, "Karl-Bonhoeffer-Nervenklinik",  R, 52.5785, 13.3330, 2)
  ins1    (U8, 13, "Lindauer Allee",                L, 52.5753, 13.3396, 1)
  ins1    (U8, 14, "Paracelsus-Bad",                U, 52.5741, 13.3494, 2)
  ins1    (U8, 15, "Residenzstr",                   R, 52.5706, 13.3608, 2)
  ins1    (U8, 16, "Franz Naumann-Platz",           R, 52.5646, 13.3639, 1)
  ins1    (U8, 17, "Osloer Str",                    R, 52.5571, 13.3733, 2)
  ins1    (U8, 18, "Pankstr",                       L, 52.5523, 13.3815, 2)
  ins1    (U8, 19, "Gesundbrunnen",                 L, 52.5488, 13.3891, 1)
  ins1    (U8, 20, "Voltastr",                      R, 52.5425, 13.3927, 2)
  ins1    (U8, 21, "Bernauer Str____",              R, 52.5377, 13.3965, 1)
  ins1    (U8, 22, "Rosenthaler Platz______",       L, 52.5297, 13.4014, 2)
  ins1    (U8, 23, "Weinmeisterstr",                R, 52.5255, 13.4057, 1)
  insert1 (U8, 24, U2, 32, tU, 2) // Alexanderplatz
  ins1    (U8, 25, "Jannowitzbrücke",               R, 52.5151, 13.4186, 2)
  ins1    (U8, 26, "Heinrich-Heine-Platz",          U, 52.5104, 13.4160, 1)
  ins1    (U8, 27, "Moritzplatz",                   R, 52.5036, 13.4107, 2)
  insert1 (U8, 28, U1, 19, tU, 2) // Kottbusser Tor
  ins1    (U8, 29, "Schönleinstr",                  R, 52.4933, 13.4224, 1)
  insert1 (U8, 30, U7, 37, tU, 2) // Hermannplat
  ins1    (U8, 31, "Boddinstr",                     L, 52.4803, 13.4253, 2)
  ins1    (U8, 32, "Leinestr",                      L, 52.4734, 13.4281, 2)
  ins1    (U8, 33, "Hermannstr  ",                  U, 52.4685, 13.4307, 1)

  insert  (U9, 10, U8, 17, tU)    // Osloer Straße
  ins1    (U9, 11, "NauenerPlatz",                  L, 52.5515, 13.3674, 1)
  insert1 (U9, 12, U6, 19, tU, 1) // Leopoldplatz
  ins1    (U9, 13, "Amrumer Str",                   L, 52.5422, 13.3495, 2)
  ins1    (U9, 14, "Westhafen",                     R, 52.5362, 13.3443, 1)
  ins1    (U9, 15, "Birkenstr",                     L, 52.5323, 13.3413, 1)
  ins1    (U9, 16, "Turmstr",                       L, 52.5259, 13.3430, 2)
  ins1    (U9, 17, "Hansaplatz",                    L, 52.5179, 13.3423, 1)
  insert1 (U9, 18, U2, 19, tU, 2) // Zoologischer Garten
  insert1 (U9, 19, U1, 11, tU, 2) // Kurfürstendamm
  insert1 (U9, 20, U3, 13, tU, 1) // Spichernstraße
  ins1    (U9, 21, "Güntzelstr",                    O, 52.4907, 13.3310, 1)
  insert1 (U9, 22, U7, 28, tU, 1) // Berliner Straße
  ins1    (U9, 23, "Bundesplatz",                   U, 52.4781, 13.3286, 2)
  ins1    (U9, 24, "Friedrich-Wilhelm-Platz",       R, 52.4724, 13.3282, 1)
  ins1    (U9, 25, "Walther-Schreiber-Platz",       L, 52.4649, 13.3284, 1)
  ins1    (U9, 26, "Schloßstr",                     L, 52.4609, 13.3249, 2)
  ins1    (U9, 27, "Rathaus Steglitz",              R, 52.4551, 13.3195, 1)

  ins     (S1, 10, "Wannsee",                       R, 52.4213, 13.1797)
  ins1    (S1, 11, "Nikolassee",                    R, 52.4319, 13.1932, 2)
  ins1    (S1, 12, "Schlachtensee",                 O, 52.4400, 13.2150, 3)
  ins1    (S1, 13, "Mexikoplatz",                   R, 52.4369, 13.2331, 2)
  ins1    (S1, 14, "Zehlendorf",                    U, 52.4310, 13.2582, 2)
  ins1    (S1, 15, "Sundgauer Str      ",           R, 52.4363, 13.2738, 2)
  ins1    (S1, 16, "Lichterfelde West        ",     R, 52.4432, 13.2934, 3)
  ins1    (S1, 17, "Botanischer Garten     ",       R, 52.4480, 13.3073, 1)
  insert1 (S1, 18, U9, 27, tU, 3)
  ins1    (S1, 19, "Feuerbachstr",                  R, 52.4633, 13.3327, 1)
  ins1    (S1, 20, "Friedenau",                     R, 52.4704, 13.3412, 2)
  ins1    (S1, 21, "Schöneberg",                    R, 52.4793, 13.3519, 2)
  ins1    (S1, 22, "Julius-Leber-Brücke",           L, 52.4861, 13.3607, 2)
  ins1    (S1, 23, "Yorckstr",                      R, 52.4929, 13.3698, 1)
  ins1    (S1, 24, "Anhalter Bhf",                  R, 52.5045, 13.3823, 2)
  insert1 (S1, 25, U2, 25, tU, 2) // Potsdamer Platz
  insert1 (S1, 26, U5, 12, tU, 2) // Brandenburger Tor
  insert1 (S1, 27, U6, 25, tU, 3) // Friedrichtstraße
  ins1    (S1, 28, "  Oranienburger Str",       U, 52.5250, 13.3929, 2)
  ins1    (S1, 29, "Nordbhf",                       R, 52.5317, 13.3890, 2)
  ins1    (S1, 30, "Humboldthain",                  R, 52.5448, 13.3794, 3)
  insert1 (S1, 31, U8, 19, tU, 2) // Gesundbrunnen
  ins1    (S1, 32, "Bornholmer Str",                R, 52.5545, 13.3980, 2)
  ins1    (S1, 33, "Wollankstr",                    R, 52.5652, 13.3924, 2)
  ins1    (S1, 34, "Schönholz",                     R, 52.5712, 13.3815, 2)
  ins1    (S1, 35, "Wilhelmsruh",                   R, 52.5818, 13.3621, 3)
  ins1    (S1, 36, "Wittenau ",                     R, 52.5970, 13.3343, 3)
  ins1    (S1, 37, "Waidmannslust",                 R, 52.6064, 13.3211, 2)
  ins1    (S1, 38, "Hermsdorf",                     R, 52.6174, 13.3075, 3)
  ins1    (S1, 39, "Frohnau",                       O, 52.6323, 13.2904, 2)
  ins1    (S1, 40, "Hohen Neuendorf",               R, 52.6686, 13.2870, 5)
  ins1    (S1, 41, "Birkenwerder",                  R, 52.6912, 13.2889, 2)
  ins1    (S1, 42, "Borgsdorf",                     R, 52.7157, 13.2764, 4)
  ins1    (S1, 43, "Lehnitz",                       R, 52.7411, 13.2635, 3)
  ins1    (S1, 44, "Oranienburg",                   R, 52.7536, 13.2496, 2)
  edge1   (S1, 23, U7, 32, 10, mitFußweg) // Yorckstraße
  edge1   (S1, 36, U8, 10, 10, mitFußweg) // Wittenau

  ins     (S2,  6, "Bernau",                        R, 52.6756, 13.5917)
  ins1    (S2,  7, "Bernau-Friedenstal",            R, 52.6683, 13.5646, 2)
  ins1    (S2,  8, "Zepernick",                     R, 52.6599, 13.5342, 3)
  ins1    (S2,  9, "Röntgental",                    R, 52.6487, 13.5136, 3)
  ins1    (S2, 10, "Buch",                          O, 52.6358, 13.4916, 3)
  ins1    (S2, 11, "Karow",                         R, 52.6153, 13.4695, 3)
  ins1    (S2, 12, "Blankenburg",                   R, 52.5916, 13.4436, 3)
  ins1    (S2, 13, "Pankow-Heinersdorf",            R, 52.5782, 13.4297, 2)
  insert1 (S2, 14, U2, 38, tU, 1) // Pankow
  insert1 (S2, 15, S1, 32, tU, 2) // Bornholmer Straße
  insert1 (S2, 16, U8, 19, tU, 3) // Gesundbrunnen
  insert1 (S2, 17, S1, 30, tU, 2) // Humboldthain
  insert1 (S2, 18, S1, 29, tU, 3) // Nordbahnhof
  insert1 (S2, 19, S1, 28, tU, 2) // Oranienburger Straße
  insert1 (S2, 21, U6, 25, tU, 2) // Friedrichstraße
  insert1 (S2, 20, S1, 26, tU, 3) // Brandenburger Tor
  insert1 (S2, 22, U2, 25, tU, 2) // Friedrichstraße
  insert1 (S2, 23, S1, 24, tU, 2) // Anhalter Bahnhof
  ins1    (S2, 24, "Yorckstr_",                     R, 52.4917, 13.3718, 2)
  ins1    (S2, 25, "Südkreuz______",                U, 52.4758, 13.3654, 1)
  ins1    (S2, 26, "Priesterweg",                   L, 52.4597, 13.3562, 3)
  ins1    (S2, 27, "Attilastr",                     R, 52.4479, 13.3608, 2)
  ins1    (S2, 28, "Marienfelde",                   R, 52.4241, 13.3747, 3)
  ins1    (S2, 29, "Buckower Chaussee",             R, 52.4103, 13.3829, 3)
  ins1    (S2, 30, "Schichauweg",                   R, 52.3988, 13.3892, 2)
  ins1    (S2, 31, "Lichtenrade",                   R, 52.3877, 13.3963, 3)
  ins1    (S2, 32, "Mahlow",                        R, 52.3609, 13.4076, 5)
  ins1    (S2, 33, "Blankenfelde",                  U, 52.3369, 13.4161, 3)

  ins     (S25, 10, "Hennigsdorf",                  L, 52.6384, 13.2056)
  ins1    (S25, 11, "Heiligensee",                  R, 52.6245, 13.2294, 2)
  ins1    (S25, 12, "Schulzendorf",                 R, 52.6130, 13.2459, 3)
  ins1    (S25, 13, "Tegel",                        R, 52.5881, 13.2898, 4)
  ins1    (S25, 14, "Eichborndamm",                 L, 52.5777, 13.3169, 3)
  ins1    (S25, 15, "Karl-Bonhoeffer-Nervenklinik", U, 52.5781, 13.3294, 2)
  ins1    (S25, 16, "Alt-Reinickendorf",            R, 52.5778, 13.3511, 2)
  insert1 (S25, 17, S1, 34, tU, 2) // Schönholz
  insert1 (S25, 18, S1, 33, tU, 2) // Wollankstraße
  insert1 (S25, 19, S1, 32, tU, 2) // Bornholmer Straße
  insert1 (S25, 20, S1, 31, tU, 2) // Gesundbrunnen
  insert1 (S25, 21, S1, 30, tU, 2) // Humboldthain
  insert1 (S25, 22, S1, 29, tU, 3) // Nordbahnhof
  insert1 (S25, 23, S1, 28, tU, 2) // Oranienburger Straße
  insert1 (S25, 24, S1, 27, tU, 2) // Friedrichstraße
  insert1 (S25, 25, S1, 26, tU, 3) // Brandenburger Tor
  insert1 (S25, 26, S1, 25, tU, 2) // Potsdamer Platz
  insert1 (S25, 27, S1, 24, tU, 2) // Anhalter Bahnhof
  insert1 (S25, 28, U7, 32, tU, 2) // Yorckstraße
  insert1 (S25, 29, S2, 25, tU, 1) // Südkreuz
  insert1 (S25, 30, S2, 26, tU, 3) // Priesterweg
  ins1    (S25, 31, "Südende",                      L, 52.4484, 13.3539, 2)
  ins1    (S25, 32, "Lankwitz",                     L, 52.4387, 13.3418, 2)
  ins1    (S25, 33, "Lichterfelde Ost",             R, 52.4300, 13.3284, 2)
  ins1    (S25, 34, "Osdorfer Str",                 R, 52.4193, 13.3146, 2)
  ins1    (S25, 35, "Lichterfelde Süd",             U, 52.4102, 13.3087, 2)
  ins1    (S25, 36, "Teltow Stadt",                 U, 52.3969, 13.2766, 2)
  edge1   (S25, 13, U6, 10, 10, mitFußweg) // Tegel-Alt Tegel
  edge1   (S25, 13, U8, 12, 10, mitFußweg) // Karl-Bonhoeffer-Nervenklinik

  insert  (S26,  1, S1,  37, tU)    // Waidmannslust
  insert1 (S26,  2, S1,  36, tU, 3) // Wittenau
  insert1 (S26,  3, S1,  35, tU, 3) // Wilhelmsruh
  insert1 (S26,  4, S1,  34, tU, 2) // Schönholz
  insert1 (S26,  5, S1,  33, tU, 2) // Wollankstraße
  insert1 (S26,  6, S1,  32, tU, 2) // Bornholmer Straße
  insert1 (S26,  7, S1,  31, tU, 2) // Gesundbrunnen
  insert1 (S26,  8, S1,  30, tU, 2) // Humboldthain
  insert1 (S26,  9, S1,  29, tU, 3) // Nordbahnhof
  insert1 (S26, 10, S1,  28, tU, 2) // Oranienburger Straße
  insert1 (S26, 11, S1,  27, tU, 2) // Friedrichstraße
  insert1 (S26, 12, S1,  26, tU, 3) // Brandenburger Tor
  insert1 (S26, 13, S1,  25, tU, 2) // Potsdamer Platz
  insert1 (S26, 14, S1,  24, tU, 2) // Anhalter Bahnhof
  insert1 (S26, 15, U7,  32, tU, 2) // Yorckstraße
  insert1 (S26, 16, S2,  25, tU, 1) // Südkreuz
  insert1 (S26, 17, S2,  26, tU, 3) // Priesterweg
  insert1 (S26, 18, S25, 31, tU, 2) // Südende
  insert1 (S26, 19, S25, 32, tU, 2) // Lankwitz
  insert1 (S26, 20, S25, 33, tU, 2) // Lichterfelde Ost
  insert1 (S26, 21, S25, 34, tU, 2) // Osdorfer Str
  insert1 (S26, 22, S25, 35, tU, 2) // Lichterfelde Süd
  insert1 (S26, 23, S25, 36, tU, 2) // Teltow Stadt

  ins     (S3,  1, "Spandau ",                      U, 52.5348, 13.1963)
  ins1    (S3,  2, "Stresow",                       R, 52.5319, 13.2093, 1)
  ins1    (S3,  3, "Pichelsberg",                   U, 52.5102, 13.2276, 4)
  ins1    (S3,  4, "Olympiastadion",                O, 52.5112, 13.2424, 2)
  ins1    (S3,  5, "Heerstr",                       L, 52.5083, 13.2587, 2)
  ins1    (S3,  6, "Messe Süd",                     U, 52.4987, 13.2700, 2)
  ins1    (S3,  7, "Westkreuz",                     L, 52.5008, 13.2840, 2)
  insert1 (S3,  8, U7, 23, tU, 2) // Wilmersdorfer Straße
  ins1    (S3,  9, "Savignyplatz",                  U, 52.5052, 13.3192, 2)
  insert1 (S3, 10, U2, 19, tU, 1) // Zoologischer Garten
  ins1    (S3, 11, "Tiergarten",                    L, 52.5144, 13.3365, 1)
  ins1    (S3, 12, "Bellevue",                      R, 52.5200, 13.3481, 2)
  insert1 (S3, 13, U5, 10, tU, 2) // Hauptbahnhof
  insert1 (S3, 14, U6, 25, tU, 2) // Friedrichstraße
  ins1    (S3, 15, "Hackescher Markt",              O, 52.5226, 13.4023, 1)
  insert1 (S3, 16, U2, 32, tU, 1) // Alexanderplatz
  insert1 (S3, 17, U8, 25, tU, 1) // Jannowitzbrücke
  ins1    (S3, 18, "Ostbahnhof",                    R, 52.5103, 13.4348, 1)
  ins1    (S3, 19, "Warschauer Str       ",         R, 52.5061, 13.4515, 2)
  ins1    (S3, 20, "Ostkreuz",                      L, 52.5031, 13.4693, 1)
  ins1    (S3, 21, "Rummelsburg",                   U, 52.5012, 13.4786, 2)
  ins1    (S3, 22, "       Bbhf Rummelsburg",       U, 52.4933, 13.4979, 3)
  ins1    (S3, 23, "Karlshorst",                    L, 52.4805, 13.5272, 3)
  ins1    (S3, 24, "Wuhlheide",                     L, 52.4686, 13.5543, 2)
  ins1    (S3, 25, "Köpenick",                      L, 52.4586, 13.5815, 3)
  ins1    (S3, 26, "Hirschgarten",                  U, 52.4579, 13.6031, 2)
  ins1    (S3, 27, "    Friedrichshagen",           O, 52.4574, 13.6236, 2)
  ins1    (S3, 28, "Rahnsdorf ",                    R, 52.4516, 13.6901, 5)
  ins1    (S3, 29, "Wilhelmshagen ",                L, 52.4386, 13.7223, 4)
  ins1    (S3, 30, "Erkner",                        U, 52.4293, 13.7507, 3)
  edge1   (S3, 19, U1, 22, 10, mitFußweg) // Warschauer Straße
  edge1   (S3,  1, U7, 10, 10, mitFußweg) // -Rathaus Spandau

  insert  (S41,  0, U8, 19, tU)    // Gesundbrunnen
  insert1 (S41,  1, U6, 20, tU, 2) // Wedding
  insert1 (S41,  2, U9, 14, tU, 3) // Westhafen
  ins1    (S41,  3, "Beusselstr",                   O, 52.5344, 13.3292, 2)
  insert1 (S41,  4, U7, 19, tU, 3) // Jungfernheide
  ins1    (S41,  5, "Westend",                      L, 52.5179, 13.2847, 3)
  ins1    (S41,  6, "Messe Nord",                   U, 52.5077, 13.2835, 1)
  insert1 (S41,  7, S3,  7, tU, 1) // Westkreuz
  ins1    (S41,  8, "Halensee",                     L, 52.4961, 13.2905, 2)
  ins1    (S41,  9, "Hohenzollerndamm",             L, 52.4886, 13.3003, 2)
  insert1 (S41, 10, U3, 16, tU, 2) // Heidelberger Platz
  insert1 (S41, 11, U9, 23, tU, 2) // Bundesplatz
  insert1 (S41, 12, U4, 14, tU, 2) // Innsbrucker Platz
  insert1 (S41, 13, S1, 21, tU, 2) // Schöneberg
  insert1 (S41, 14, S2, 25, tU, 1) // Südkreuz
  insert1 (S41, 15, U6, 33, tU, 2) // Tempelhof
  insert1 (S41, 16, U8, 33, tU, 4) // Hermannstraße
  insert1 (S41, 17, U7, 40, tU, 1) // Neukölln
  ins1    (S41, 18, "Sonnenallee",                  O, 52.4729, 13.4556, 2)
  ins1    (S41, 19, "Treptower Park",               U, 52.4938, 13.4618, 3)
  insert1 (S41, 20, S3, 20, tU, 1) // Ostkreuz
  insert1 (S41, 21, U5, 22, tU, 2) // Frankfurter Allee
  ins1    (S41, 22, "Storkower Str",                R, 52.5238, 13.4646, 2)
  ins1    (S41, 23, "Landsberger Allee",            R, 52.5295, 13.4548, 2)
  ins1    (S41, 24, "Greifswalder Str",             R, 52.5402, 13.4392, 2)
  ins1    (S41, 25, "Prenzlauer Allee",             R, 52.5448, 13.4259, 2)
  insert1 (S41, 26, U2, 36, tU, 2) // Schönhauser Allee
  insert1 (S41, 27, U8, 19, tU, 3) // Gesundbrunnen
  edge1   (S41,  6, U2, 14,  5, mitFußweg) // Westend-Kaiserdamm

//         S42 = Umkehrung von S41

  ins     (S45,  1, "Flughafen BER T1-2",           R, 52.3636, 13.5102)
  ins1    (S45,  2, "Waßmannsdorf",                 L, 52.3693, 13.4633, 4)
  ins1    (S45,  3, "Flughafen BER T5",             L, 52.3911, 13.5130, 4)
  ins1    (S45,  4, "Grünbergallee",                R, 52.3992, 13.5416, 3)
  ins1    (S45,  5, "Altglienicke",                 R, 52.4072, 13.5586, 2)
  ins1    (S45,  6, "Adlershof",                    L, 52.4346, 13.5418, 4)
  ins1    (S45,  7, "Johannisthal",                 L, 52.4467, 13.5238, 2)
  ins1    (S45,  8, " Schöneweide",                 L, 52.4549, 13.5096, 2)
  ins1    (S45,  9, "Baumschulenweg ",              R, 52.4669, 13.4908, 2)
  ins1    (S45, 10, "Köllnische Heide",             O, 52.4697, 13.4685, 4)
  insert1 (S45, 11, U7, 40, tU, 2) // Neukölln
  insert1 (S45, 12, U8, 33, tU, 1) // Hermannstraße
  insert1 (S45, 13, U6, 34, tU, 4) // Tempelhof
  insert1 (S45, 14, S2, 25, tU, 2) // Südkreuz

  ins     (S46,  1, "      Königs Wusterhausen",    U, 52.2964, 13.6315)
  ins1    (S46,  2, "Wildau",                       R, 52.3201, 13.6341, 2)
  ins1    (S46,  3, "Zeuthen",                      R, 52.3488, 13.6274, 3)
  ins1    (S46,  4, "Eichwalde",                    R, 52.3713, 13.6154, 3)
  ins1    (S46,  5, "Grünau",                       R, 52.4124, 13.5743, 5)
  insert1 (S46,  6, S45,  6, tU, 4) // Adlershof
  insert1 (S46,  7, S45,  7, tU, 2) // Johannisthal
  insert1 (S46,  8, S45,  8, tU, 2) // Schöneweide
  insert1 (S46,  9, S45,  9, tU, 2) // Baumschulenweg
  insert1 (S46, 10, S45, 10, tU, 4) // Köllnische Heide
  insert1 (S46, 11, U7,  38, tU, 2) // Neukölln
  insert1 (S46, 12, U8,  33, tU, 1) // Hermannstraße
  insert1 (S46, 13, U6,  34, tU, 4) // Tempelhof
  insert1 (S46, 14, S2,  25, tU, 2) // Südkreuz
  insert1 (S46, 15, S1,  21, tU, 1) // Schöneberg
  insert1 (S46, 16, U4,  14, tU, 1) // Innsbrucker Platz
  insert1 (S46, 17, U9,  23, tU, 2) // Bundesplatz
  insert1 (S46, 18, S41, 10, tU, 2) // Heidelberger Platz
  insert1 (S46, 19, S41,  9, tU, 2) // Hohenzollerndamm
  insert1 (S46, 20, S41,  8, tU, 2) // Halensee
  insert1 (S46, 21, S3,   7, tU, 1) // Westkreuz
  insert1 (S46, 22, S41,  6, tU, 1) // Messe Nord
  insert1 (S46, 23, U2,  12, tU, 1) // Westend

  ins     (S47,  1, "Spindlersfeld",                U, 52.4473, 13.5613)
  ins1    (S47,  2, "Oberspree",                    R, 52.4524, 13.5381, 2)
  insert1 (S47,  3, S45,  8, tU, 4) // Schöneweide
  insert1 (S47,  4, S45,  9, tU, 2) // Baumschulenweg
  insert1 (S47,  5, S45, 10, tU, 4) // Köllnische Heide
  insert1 (S47,  6, U7,  40, tU, 2) // Neukölln
  insert1 (S47,  7, U8,  33, tU, 1) // Hermannstraße

  insert  (S5,  1, S3,  7, tU) // Westkreuz
  ins1    (S5, 11, "Charlottenburg",                U, 52.5020, 13.3000, 2)
  insert1 (S5, 12, S3,  9, tU, 2) // Savignyplatz
  insert1 (S5, 13, U2, 19, tU, 3) // Zoologischer Garten
  insert1 (S5, 14, S3, 11, tU, 1) // Tiergarten
  insert1 (S5, 15, S3, 12, tU, 2)
  insert1 (S5, 16, U5, 10, tU, 3) // Hauptbahnhof
  insert1 (S5, 17, U6, 25, tU, 3) // Friedrichstraße
  insert1 (S5, 18, S3, 15, tU, 2) // Hackescher Markt
  insert1 (S5, 19, U2, 32, tU, 2) // Alexanderplatz
  insert1 (S5, 20, U8, 25, tU, 2) // Jannowitzbrücke
  insert1 (S5, 21, S3, 18, tU, 2) // Ostbahnhof
  insert1 (S5, 22, S3, 19, tU, 3) // Warschauer Straße
  insert1 (S5, 23, S3, 20, tU, 2) // Ostkreuz
  ins1    (S5, 24, "Nöldnerplatz",                  R, 52.5038, 13.4853, 2)
  insert1 (S5, 25, U5, 24, tU, 2) // Lichtenberg
  ins1    (S5, 26, " Friedrichsfelde Ost",          O, 52.5141, 13.5201, 3)
  ins1    (S5, 27, "Biesdorf ",                     O, 52.5130, 13.5560, 3)
  insert1 (S5, 28, U5, 29, tU, 2) // Wuhletal
  ins1    (S5, 29, " Kaulsdorf",                    O, 52.5122, 13.5901, 2)
  ins1    (S5, 30, "Mahlsdorf",                     U, 52.5122, 13.6113, 1)
  ins1    (S5, 31, "Birkenstein",                   R, 52.5157, 13.6477, 3)
  ins1    (S5, 32, "Hoppegarten",                   R, 52.5181, 13.6733, 3)
  ins1    (S5, 33, "Neuenhagen",                    R, 52.5205, 13.6996, 3)
  ins1    (S5, 34, "Fredersdorf",                   R, 52.5264, 13.7613, 4)
  ins1    (S5, 35, "Petershagen Nord",              R, 52.5289, 13.7898, 3)
  ins1    (S5, 36, "Strausberg",                    R, 52.5322, 13.8331, 4)
  ins1    (S5, 37, "Hegermühle",                    R, 52.5487, 13.8666, 4)
  ins1    (S5, 38, "Strausberg Stadt",              L, 52.5766, 13.8878, 3)
  ins1    (S5, 39, "Strausberg Nord",               L, 52.5905, 13.9084, 2)
  edge1   (S5, 11, U7, 23, 10, mitFußweg) // Charlottenburg-Wilmersdorfer Straße

  ins     (S7,  7, "Potsdam Hauptbahnhof",          L, 52.3918, 13.0671)
  ins1    (S7,  8, "Babelsberg",                    U, 52.3914, 13.0928, 4)
  ins1    (S7,  9, "Griebnitzsee",                  R, 52.3945, 13.1274, 3)
  insert1 (S7, 10, S1, 10, tU, 5) // Wannsee
  insert1 (S7, 11, S1, 11, tU, 2) // Nikolassee
  ins1    (S7, 12, "Grunewald",                     L, 52.4882, 13.2610, 7)
  insert1 (S7, 13, S3,  7, tU, 4) // Westkreuz
  insert1 (S7, 14, S5, 11, tU, 2) // Charlottenburg
  insert1 (S7, 15, S3,  9, tU, 2) // Savignyplatz
  insert1 (S7, 16, U2, 19, tU, 3) // Zoologischer Garten
  insert1 (S7, 17, S3, 11, tU, 1) // Tiergarten
  insert1 (S7, 18, S5, 15, tU, 2) // Bellevue
  insert1 (S7, 19, U5, 10, tU, 2) // Hauptbahnhof
  insert1 (S7, 20, S5, 17, tU, 3) // Friedrichstraße
  insert1 (S7, 21, S3, 15, tU, 2) // Hackescher Markt
  insert1 (S7, 22, U2, 32, tU, 2) // Alexanderplatz
  insert1 (S7, 23, U8, 25, tU, 2) // Jannowitzstraße
  insert1 (S7, 24, S3, 18, tU, 2) // Ostbahnhof
  insert1 (S7, 25, S3, 19, tU, 3) // Warschauer Straße
  insert1 (S7, 26, S3, 20, tU, 2) // Ostkreuz
  insert1 (S7, 27, S5, 24, tU, 2) // Nöldnerplatz
  insert1 (S7, 28, S5, 25, tU, 2) // Lichtenberg
  insert1 (S7, 29, S5, 26, tU, 3) // Friedrichsfelde Ost
  ins1    (S7, 30, "Springpfuhl",                   R, 52.5270, 13.5366, 4)
  ins1    (S7, 31, "Poelchaustr",                   R, 52.5358, 13.5355, 2)
  ins1    (S7, 32, "Marzahn",                       R, 52.5436, 13.5413, 1)
  ins1    (S7, 33, "Raoul-Wallenberg-Str",          R, 52.5507, 13.5476, 2)
  ins1    (S7, 34, "Mehrower Allee",                R, 52.5576, 13.5536, 2)
  ins1    (S7, 35, "Ahrensfelde",                   O, 52.5713, 13.5656, 2)

  insert  (S75, 1, S3, 19, tU)    // Warschauer Straße
  insert1 (S75, 2, S3, 20, tU, 1) // Ostkreuz
  insert1 (S75, 3, S5, 24, tU, 2) // Nöldnerplatz
  insert1 (S75, 4, U5, 24, tU, 3) // Lichtenberg
  insert1 (S75, 5, S5, 26, tU, 2) // Friedrichsfelde Ost
  insert1 (S75, 6, S7, 30, tU, 3) // Springpfuhl
  ins1    (S75, 7, "Gehrenseestr",                  L, 52.5565, 13.5248, 4)
  ins1    (S75, 8, "Hohenschönhausen",              L, 52.5663, 13.5125, 2)
  ins1    (S75, 9, "Wartenberg",                    O, 52.5730, 13.5038, 2)

  insert  (S8,  6, S1,  40, tU) // Hohen Neuendorf
  ins1    (S8,  7, "Bergfelde",                     R, 52.6702, 13.3201, 4)
  ins1    (S8,  8, "Schönfließ",                    R, 52.6646, 13.3406, 2)
  ins1    (S8,  9, "Mühlenbeck-Mönchmühle",         R, 52.6548, 13.3859, 3)
  insert1 (S8, 10, S2,  12, tU, 10) // Blankenburg
  insert1 (S8, 11, S2,  13, tU, 2) // Pankow-Heinersdorf
  insert1 (S8, 12, U2,  38, tU, 1) // Pankow
  insert1 (S8, 13, S1,  32, tU, 2) // Bornholmer Straße
  insert1 (S8, 14, U2,  36, tU, 1) // Schönhauser Allee
  insert1 (S8, 15, S41, 25, tU, 2) // Prenzlauer Allee
  insert1 (S8, 16, S41, 24, tU, 2) // Greifswalder Straße
  insert1 (S8, 17, S41, 23, tU, 3) // Landberger Allee
  insert1 (S8, 18, S41, 22, tU, 2) // Storkower Straße
  insert1 (S8, 19, U5,  22, tU, 2) // Frankfurter Allee
  insert1 (S8, 20, S3,  20, tU, 3) // Ostkreuz
  insert1 (S8, 21, S41, 19, tU, 2) // Treptower Park
  ins1    (S8, 22, "Plänterwald",                   R, 52.4785, 13.4733, 1)
  insert1 (S8, 23, S45,  9, tU, 2) // Baumschulenweg
  insert1 (S8, 24, S45,  8, tU, 3) // Schöneweide
  insert1 (S8, 25, S45,  7, tU, 3) // Johannisthal
  insert1 (S8, 26, S45,  6, tU, 3) // Adlershof
  insert1 (S8, 27, S46,  3, tU, 5) // Grünau
  insert1 (S8, 28, S46,  4, tU, 5) // Eichwalde
  insert1 (S8, 29, S46,  3, tU, 3) // Zeuthen

  insert  (S85, 50, S1, 37, tU)    // Waidmannslust
  insert1 (S85, 51, S1, 36, tU, 3) // Wittenau
  insert1 (S85, 52, S1, 35, tU, 3) // Wilhelmsruh
  insert1 (S85, 53, S1, 34, tU, 2) // Schönholz
  insert1 (S85, 54, S1, 33, tU, 2) // Wollankstraße
  insert1 (S85, 55, S1, 32, tU, 2) // Bornholmer Straße

  insert  (S9,  1, S3,   1, tU)    // Spandau
  insert1 (S9,  2, S3,   2, tU, 4) // Stresow
  insert1 (S9,  3, S3,   3, tU, 2) // Pichelsberg
  insert1 (S9,  4, S3,   4, tU, 2) // Olympiastadion
  insert1 (S9,  5, S3,   5, tU, 2) // Heerstraße
  insert1 (S9,  6, S3,   6, tU, 2) // Messe Süd
  insert1 (S9,  7, S3,   7, tU, 2) // Westkreuz
  insert1 (S9,  8, S5,  11, tU, 2) // Charlottenburg
  insert1 (S9,  9, S3,   9, tU, 2) // Savignyplatz
  insert1 (S9, 10, U2,  19, tU, 3) // Zoologischer Garten
  insert1 (S9, 11, S3,  11, tU, 1) // Tiergarten
  insert1 (S9, 12, S5,  15, tU, 2) // Bellevue
  insert1 (S9, 13, U5,  10, tU, 2) // Hauptbahnhof
  insert1 (S9, 14, U6,  25, tU, 3) // Friedrichstraße
  insert1 (S9, 15, S3,  15, tU, 2) // Hackescher Markt
  insert1 (S9, 16, U2,  32, tU, 2) // Alexanderplatz
  insert1 (S9, 17, U8,  25, tU, 2) // Jannowitzbrücke
  insert1 (S9, 18, S3,  18, tU, 2) // Ostbahnhof
  insert1 (S9, 19, S3,  19, tU, 3) // Warschauer Straße
  insert1 (S9, 20, S41, 19, tU, 2) // Treptower Park
  insert1 (S9, 21, S8,  22, tU, 2) // Plänterwald
  insert1 (S9, 22, S45,  9, tU, 3) // Baumschulenweg
  insert1 (S9, 23, S45,  8, tU, 4) // Schöneweide 
  insert1 (S9, 24, S45,  7, tU, 2) // Johannisthal
  insert1 (S9, 25, S45,  6, tU, 2) // Adlershof
  insert1 (S9, 26, S45,  5, tU, 4) // Alt-Glienicke
  insert1 (S9, 27, S45,  4, tU, 2) // Grünbergallee
  insert1 (S9, 28, S45,  3, tU, 3) // Fughafen BER T5
  insert1 (S9, 29, S45,  2, tU, 4) // Waßmannsdorf
  insert1 (S9, 30, S45,  1, tU, 4) // Flughafen BER T1-2

  ins     (Zoo,  1, "",                             R, 52.5057, 13.3391)
  ins1    (Zoo,  2, "",                             R, 52.5059, 13.3387, 0)
  ins1    (Zoo,  3, "",                             R, 52.5057, 13.3368, 0)
  ins1    (Zoo,  4, "",                             R, 52.5058, 13.3357, 0)
  ins1    (Zoo,  5, " Garten",                      R, 52.5067, 13.3342, 0)
  ins1    (Zoo,  6, "logischer_",                   R, 52.5079, 13.3350, 0)
  ins1    (Zoo,  7, "Zoo-",                         R, 52.5092, 13.3349, 0)
  ins1    (Zoo,  8, "",                             R, 52.5109, 13.3356, 0)
  ins1    (Zoo,  9, "",                             R, 52.5111, 13.3358, 0)
  ins1    (Zoo, 10, "",                             R, 52.5110, 13.3365, 0)
  ins1    (Zoo, 11, "",                             R, 52.5100, 13.3394, 0)
  ins1    (Zoo, 12, "",                             R, 52.5094, 13.3412, 0)
  ins1    (Zoo, 13, "",                             R, 52.5089, 13.3427, 0)
  ins1    (Zoo, 14, "",                             R, 52.5088, 13.3424, 0)
  ins1    (Zoo, 15, "",                             R, 52.5086, 13.3430, 0)
  ins1    (Zoo, 16, "",                             R, 52.5081, 13.3433, 0)
  ins1    (Zoo, 17, "",                             R, 52.5064, 13.3439, 0)
  ins1    (Zoo, 18, "",                             R, 52.5059, 13.3416, 0)
  ins1    (Zoo, 19, "",                             R, 52.5061, 13.3415, 0)
  ins1    (Zoo, 20, "",                             R, 52.5060, 13.3407, 0)
  ins1    (Zoo, 22, "",                             R, 52.5057, 13.3397, 0)
  ins1    (Zoo, 23, "",                             R, 52.5056, 13.3390, 0)
  edge1   (Zoo, 23, Zoo, 1, 0, true)

  ins     (BG,  1, "",                              R, 52.4583, 13.3039)
  ins1    (BG,  2, "",                              R, 52.4582, 13.3037, 0)
  ins1    (BG,  3, " nischer",                      R, 52.4545, 13.3000, 0)
  ins1    (BG,  4, "      Garten",                  O, 52.4518, 13.3038, 0)
  ins1    (BG,  5, "",                              R, 52.4500, 13.3056, 0)
  ins1    (BG,  6, "",                              R, 52.4518, 13.3099, 0)
  ins1    (BG,  7, "",                              R, 52.4520, 13.3102, 0)
  ins1    (BG,  8, "",                              R, 52.4536, 13.3082, 0)
  ins1    (BG,  9, "",                              R, 52.4540, 13.3082, 0)
  ins1    (BG, 10, "",                              R, 52.4543, 13.3094, 0)
  ins1    (BG, 11, "Bota-",                         L, 52.4558, 13.3081, 0)
  ins1    (BG, 12, "",                              R, 52.4558, 13.3086, 0)
  ins1    (BG, 13, "",                              R, 52.4576, 13.3065, 0)
  ins1    (BG, 14, "",                              R, 52.4572, 13.3056, 0)
  edge1   (BG, 14, BG, 1, 0, true)

  write (false)
}
