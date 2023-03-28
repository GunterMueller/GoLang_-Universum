package bahnhof

// (c) Christian Maurer   v. 230304 - license see µU.go

import (
  . "bahn/kilo"
  . "bahn/richtung"
  b "bahn/block"
  s "bahn/signal"
  "bahn/netz"
)

func (x *bahnhof) bahnheim() {
/*
00  02  04  06  08  10  12  14  16  18  20  22  24  26  28  30  32  34  36  38  40   Y0 +
  01  03  05  07  09  11  13  15  17  19  21  23  25  27  29  31  33  35  37  39  41

   [------14-------------3-----------------4---------------\                            4
                        /                                   \  
   [------13-----------2-------------------3-----------------8----------23----------    5
                      /                                       \           Bahnhausen 
   [------12---------1-4-------------------2-----------------7-9--------22----------    6
                        \                                   /                
   [------11-------------5-----------------1---------------6------------21------]       7
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
//         |     |    |   |  |   |   |    |   |    |     |   Typ     Stellung
//         |     |    |   |  |   |   |    |   |    |     |    |   |    |   Spalte
//         |     |    |   |  |   |   |    |   |    |     |    |   |    |     |
  x.gleis ( 1, b.Dfg, G, 16, 7, 13, 21, s.NT, k, s.NS,   0, s.T2, m, s.Hp0, 29)
  x.gleis ( 2, b.Dfg, G, 18, 6, 12, 21, s.NT, k, s.Hp0,  0, s.T2, m, s.Hp0, 29)
  x.gleis ( 3, b.Dfg, G, 18, 5, 12, 21, s.T1, g, s.Hp0, 12, s.NT, k, s.Hp0,  0)
  x.gleis ( 4, b.Dfg, G, 16, 4, 13, 21, s.T2, g, s.Hp0, 13, s.T2, g, s.Hp0, 28)
  x.gleis (11, b.AsG, G, 10, 7,  2,  5, s.NT, k, s.NS,   0, s.T1, m, s.Hp0, 12)
  x.gleis (12, b.AsG, G,  8, 6,  2,  5, s.NT, k, s.NS,   0, s.T1, m, s.Hp0, 10)
  x.gleis (13, b.AsG, G,  9, 5,  2,  5, s.NT, k, s.NS,   0, s.T1, m, s.Hp0, 10)
  x.gleis (14, b.AsG, G, 10, 4,  2,  5, s.NT, k, s.NS,   0, s.T1, m, s.Hp0, 11)
  x.gleis (21, b.AsM, G, 10, 7, 30, 36, s.T1, g, s.Hp0, 30, s.NT, k, s.NS,   0)
  x.gleis (22, b.AfM, G, 10, 6, 32, 36, s.NT, k, s.Hp0,  0, s.T1, m, s.Hp0, 41)
  x.gleis (23, b.EfG, G, 11, 5, 31, 36, s.T2, g, s.Hp0, 31, s.NT, k, s.Hp0,  0)

  x.knick ( 4, m, R,     4, 29)

  x.gleisBesetzen (2)
  x.gleisBesetzen (4)

  nachbar[22] = netz.Bahnhausen
  nachbar[23] = netz.Bahnhausen

//  Weiche  Nr  Kilometrierung
//           |  | Lage
//           |  |  | Richtung
//           |  |  |  | Stellung
//           |  |  |  |  | Zeile
//           |  |  |  |  |  | Spalte
//           |  |  |  |  |  |   | 
  x.weiche ( 1, m, G, L, G, 6, 10)
  x.dkw    ( 2,    L,    G, 5, 11)
  x.weiche ( 3, g, G, L, G, 4, 12)
  x.weiche ( 4, m, G, R, G, 6, 11)
  x.weiche ( 5, g, G, R, G, 7, 12)
  x.weiche ( 6, m, G, L, G, 7, 29)
  x.weiche ( 7, g, G, L, G, 6, 30)
  x.dkw    ( 8,    R,    G, 5, 30)
  x.weiche ( 9, g, G, R, G, 6, 31)

  x.verbinden()
}
