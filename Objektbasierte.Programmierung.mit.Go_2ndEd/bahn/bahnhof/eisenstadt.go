package bahnhof

// (c) Christian Maurer   v. 230109 - lices.NSe see µU.go

import (
  . "bahn/kilo"
  . "bahn/richtung"
  b "bahn/block"
  s "bahn/signal"
  "bahn/netz"
)

func (x *bahnhof) eisenstadt() {
/*
00  02  04  06  08  10  12  14  16  18  20  22  24  26  28  30  32  34  36  38  40  Y0 +
  01  03  05  07  09  11  13  15  17  19  21  23  25  27  29  31  33  35  37  39  41

   [------14-------------3-----------------4---------------\                           4
                        /                                   \  
----------13-----------2-------------------3-----------------6----------23----------   5
Bahnstadt                                                                Eisenhausen
----------12-----------1-------------------2-----------------5----------22----------   6
                        \                                   /                
                         \-----------------1---------------4------------21------]      7
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
  x.gleis ( 1, b.Dfg, G, 16, 7, 13, 21, s.NT, k, s.NS,   0, s.T2, m, s.Hp0, 29)
  x.gleis ( 2, b.Dfg, G, 18, 6, 12, 21, s.NT, k, s.Hp0,  0, s.T2, m, s.Hp0, 29)
  x.gleis ( 3, b.Dfg, G, 18, 5, 12, 21, s.T1, g, s.Hp0, 12, s.NT, k, s.Hp0,  0)
  x.gleis ( 4, b.Dfg, G, 16, 4, 13, 21, s.T2, g, s.Hp0, 13, s.NT, k, s.NS,   0)
  x.gleis (12, b.EfM, G, 11, 6,  0,  5, s.NT, k, s.NS,   0, s.T1, m, s.Hp0, 11)
  x.gleis (13, b.AfG, G, 11, 5,  0,  5, s.T1, g, s.Hp0,  0, s.NT, k, s.Hp0,  0)
  x.gleis (14, b.AsG, G, 10, 4,  2,  5, s.NT, k, s.NS,   0, s.T1, m, s.Hp0, 11)
  x.gleis (21, b.AsM, G, 10, 7, 30, 36, s.T1, g, s.Hp0, 30, s.NT, k, s.NS,   0)
  x.gleis (22, b.AfM, G, 11, 6, 31, 36, s.NT, k, s.Hp0,  0, s.T1, m, s.Hp0, 41)
  x.gleis (23, b.EfG, G, 11, 5, 31, 36, s.T2, g, s.Hp0, 31, s.NT, k, s.Hp0,  0)

  x.knick ( 1, g, R,     7, 12)
  x.knick ( 4, m, R,     4, 29)
  
  x.gleisBesetzen (1)
  x.gleisBesetzen (4)

  nachbar[12] = netz.Bahnstadt
  nachbar[13] = netz.Bahnstadt
  nachbar[22] = netz.Eisenhausen
  nachbar[23] = netz.Eisenhausen

//  Weiche  Nr  Kilometrierung
//           |  | Lage
//           |  |  | Richtung
//           |  |  |  | Stellung
//           |  |  |  |  | Zeile
//           |  |  |  |  |  | Spalte
//           |  |  |  |  |  |   | 
  x.weiche ( 1, m, G, R, G, 6, 11)
  x.weiche ( 2, m, G, L, G, 5, 11)
  x.weiche ( 3, g, G, L, G, 4, 12)
  x.weiche ( 4, m, G, L, G, 7, 29)
  x.weiche ( 5, g, G, L, G, 6, 30)
  x.weiche ( 6, g, G, R, G, 5, 30)

  x.verbinden()
}
