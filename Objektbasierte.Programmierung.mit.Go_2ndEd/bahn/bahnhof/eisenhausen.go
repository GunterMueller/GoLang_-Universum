package bahnhof

// (c) Christian Maurer   v. 230109 - license see µU.go

import (
  . "bahn/kilo"
  . "bahn/richtung"
  b "bahn/block"
  s "bahn/signal"
  "bahn/netz"
)

func (x *bahnhof) eisenhausen() {
/*
00  02  04  06  08  10  12  14  16  18  20  22  24  26  28  30  32  34  36  38  40  Y0 +
  01  03  05  07  09  11  13  15  17  19  21  23  25  27  29  31  33  35  37  39  41

   [------14-------------5-----------------4-------------------9--------24--------]    4
                        /                                       \
----------13-----1--43-4-------------------3-------------------810------23--------]    5
Bahnstadt         \   /                                       /
----------12-------2-3---------------------2-----------------7----------22--------]    6
                    \                                       /                
                     \---------------------1---------------6------------21--------]    7
*/

const (
  m = Mit; g = Gegen; k = NK
  L = Links; G = Gerade; R = Rechts
)

//  Gleis  Nr   Art   Lage Zeile Spalte
//         |     |    | Länge|   | Nummernspalte
//         |     |    |   |  |   |   |
//         |     |    |   |  |   |   |  Signal Gegen
//         |     |    |   |  |   |   |   Typ     Stellung    
//         |     |    |   |  |   |   |    |   |    |   Spalte
//         |     |    |   |  |   |   |    |   |    |     |  Signal Mit
//         |     |    |   |  |   |   |    |   |    |     |   Typ  |  Stellung
//         |     |    |   |  |   |   |    |   |    |     |    |   |    |   Spalte
//         |     |    |   |  |   |   |    |   |    |     |    |   |    |     |
  x.gleis ( 1, b.Dfg, G, 18, 7, 11, 21, s.NT, k, s.NS,   1, s.T2, m, s.Hp0, 29)
  x.gleis ( 2, b.Dfg, G, 19, 6, 11, 21, s.NT, k, s.NS,   2, s.T2, m, s.Hp0, 29)
  x.gleis ( 3, b.Dfg, G, 19, 5, 12, 21, s.T1, g, s.Hp0, 12, s.NT, g, s.Hp0,  0)
  x.gleis ( 4, b.Dfg, G, 18, 4, 13, 21, s.T2, g, s.Hp0, 13, s.NT, k, s.NS,   0)
  x.gleis (12, b.EfM, G,  9, 6,  0,  5, s.NT, k, s.NS,   0, s.T1, m, s.Hp0,  9)
  x.gleis (13, b.AfG, G,  8, 5,  0,  5, s.T1, g, s.Hp0,  0, s.NT, k, s.NS,   0)
  x.gleis (43, b.Dfg, G,  2, 5,  9,  0, s.NT, k, s.NS,   0, s.NT, k, s.NS,   0)
  x.gleis (14, b.AsG, G, 10, 4,  2,  5, s.NT, k, s.NS,   0, s.T1, m, s.Hp0, 11)
  x.gleis (21, b.AsM, G, 11, 7, 30, 36, s.T1, g, s.Hp0, 30, s.NT, k, s.NS,   0)
  x.gleis (22, b.AsM, G, 10, 6, 31, 36, s.T2, g, s.Hp0, 31, s.NT, k, s.NS,   0)
  x.gleis (23, b.AsM, G,  8, 5, 33, 36, s.T2, g, s.Hp0, 33, s.NT, k, s.NS,   0)
  x.gleis (24, b.AsM, G,  9, 4, 32, 36, s.T2, g, s.Hp0, 32, s.NT, k, s.NS,   0)

  x.knick ( 1, g,     R, 7, 10)

  x.gleisBesetzen (1)
  x.gleisBesetzen (3)

  nachbar[12] = netz.Bahnstadt
  nachbar[13] = netz.Bahnstadt

//  Weiche  Nr  Kilometrierung
//           |  | Lage
//           |  |  | Richtung
//           |  |  |  | Stellung
//           |  |  |  |  | Zeile
//           |  |  |  |  |  | Spalte
//           |  |  |  |  |  |   | 
  x.weiche ( 1, m, G, R, G, 5,  8)
  x.dkw    ( 2,    R,    G, 6,  9)
  x.weiche ( 3, m, G, L, G, 6, 10)
  x.dkw    ( 4,    L,    G, 5, 11)
  x.weiche ( 5, g, G, L, G, 4, 12)
  x.weiche ( 6, m, G, L, G, 7, 29)
  x.dkw    ( 7,    L,    G, 6, 30)
  x.weiche ( 8, g, G, L, G, 5, 31)
  x.weiche ( 9, m, G, R, G, 4, 31)
  x.weiche (10, g, G, R, G, 5, 32)

  x.verbinden()
}
