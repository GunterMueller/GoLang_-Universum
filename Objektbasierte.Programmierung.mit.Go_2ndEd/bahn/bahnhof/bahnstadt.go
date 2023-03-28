package bahnhof

// (c) Christian Maurer   v. 230305 - license see µU.go

import (
  . "bahn/kilo"
  . "bahn/richtung"
  b "bahn/block"
  s "bahn/signal"
  "bahn/netz"
)

func (x *bahnhof) bahnstadt() {
/*
00  02  04  06  08  10  12  14  16  18  20  22  24  26  28  30  32  34  36  38  40  Y0 +
  01  03  05  07  09  11  13  15  17  19  21  23  25  27  29  31  33  35  37  39  41

   [--------18--------------10---------8--------------18----28------26----38----]      0
                            /                           \           /
   [--------17-------------9-----------7----------------19--27----25------37----}      1
                          /                               \
   [--------16-----------8-------------6------------------202122----------36--------   2
                        /                                   /  \           Eisenheim
                      15 [-------------6------------------16    25                     3
                      /                                   /      \
------14-------1-----6-----------------4----------------15--24----23------34--------   4
Bahnhausen      \   /                                   /           \     Eisenstadt
------13---------2-5-------------------3--------------14----23------24----33--------   5
                  \                                   /                      
                   3-------------------2------------13------22----]                    6
                    \                              /                               
         [--11-------4-----------------1----------12--------21----]                    7
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
  x.gleis ( 1, b.Dfg, G, 14, 7, 11, 19, s.T2, g, s.Hp0, 11, s.T2, m, s.Hp0, 25)
  x.gleis ( 2, b.Dfg, G, 16, 6, 10, 19, s.T2, g, s.Hp0, 10, s.T2, m, s.Hp0, 25)
  x.gleis ( 3, b.Dfg, G, 17, 5, 10, 19, s.NT, k, s.NS,  10, s.T1, m, s.Hp0, 26)
  x.gleis ( 4, b.Dfg, G, 17, 4, 11, 19, s.T1, g, s.Hp0, 11, s.NT, k, s.NS,   0)
  x.gleis ( 5, b.AsG, G, 16, 3, 13, 19, s.NT, k, s.NS,   0, s.T2, m, s.Hp0, 28)
  x.gleis ( 6, b.Dfg, G, 16, 2, 13, 19, s.T2, g, s.Hp0, 13, s.T2, m, s.Hp0, 29)
  x.gleis ( 7, b.Dfg, G, 14, 1, 14, 19, s.T2, g, s.Hp0, 14, s.T2, m, s.Hp0, 28)
  x.gleis ( 8, b.Dfg, G, 12, 0, 15, 19, s.T2, g, s.Hp0, 15, s.T2, m, s.Hp0, 27)
  x.gleis (11, b.AsG, G,  5, 7,  5,  6, s.NT, k, s.NS,   0, s.T1, m, s.Hp0, 10)
  x.gleis (13, b.EfM, G,  8, 5,  0,  4, s.NT, k, s.NS,   0, s.T2, m, s.Hp0,  7)
  x.gleis (14, b.AfG, G,  7, 4,  0,  4, s.NT, k, s.NS,   0, s.T1, g, s.Hp0,  0)
  x.gleis (15, b.Dfg, L,  1, 3, 11,  0, s.NT, k, s.NS,   0, s.NT, k, s.NS,   0)
  x.gleis (16, b.AsG, G, 10, 2,  2,  6, s.NT, k, s.NS,   0, s.T1, m, s.Hp0, 12)
  x.gleis (17, b.AsG, G, 11, 1,  2,  6, s.NT, k, s.NS,   0, s.T1, m, s.Hp0, 12)
  x.gleis (18, b.AsG, G, 12, 0,  2,  6, s.NT, k, s.NS,   0, s.T2, m, s.Hp0, 13)
  x.gleis (21, b.AsM, G,  8, 7, 26, 30, s.T1, g, s.Hp0, 26, s.NT, k, s.NS,   0)
  x.gleis (22, b.AsM, G,  7, 6, 27, 30, s.T1, g, s.Hp0, 27, s.NT, k, s.NS,   0)
  x.gleis (23, b.Dfg, G,  6, 5, 28, 30, s.T2, g, s.Hp0, 28, s.T1, m, s.Hp0, 34)
  x.gleis (24, b.Dfg, G,  4, 4, 29, 30, s.T2, g, s.Hp0, 29, s.T1, m, s.Hp0, 33)
  x.gleis (25, b.Dfg, R,  1, 3, 32,  0, s.NT, k, s.NS,   0, s.NT, k, s.NS,   0)
  x.gleis (27, b.Dfg, G,  4, 1, 29, 30, s.T2, g, s.Hp0, 29, s.T1, m, s.Hp0, 32)
  x.gleis (28, b.Dfg, G,  6, 0, 28, 30, s.T2, g, s.Hp0, 28, s.T1, m, s.Hp0, 33)
  x.gleis (33, b.AfM, G,  7, 5, 35, 37, s.NT, k, s.NS,   0, s.T1, m, s.Hp0, 41)
  x.gleis (34, b.EfG, G,  8, 4, 34, 37, s.T2, g, s.Hp0, 34, s.NT, k, s.NS,   0)
  x.gleis (36, b.EAM, G, 10, 2, 32, 37, s.T2, g, s.Hp0, 32, s.T1, m, s.Hp0, 41)
  x.gleis (37, b.AsM, G,  6, 1, 34, 37, s.T2, g, s.Hp0, 34, s.NT, k, s.NS,   0)
  x.gleis (38, b.AsM, G,  5, 0, 35, 37, s.T2, g, s.Hp0, 34, s.NT, k, s.NS,   0)
  x.gleis (44, b.Dfg, G,  2, 4,  8,  0, s.NT, k, s.NS,   0, s.NT, k, s.NS,   0)

  x.gleisBesetzen (2)
  x.gleisBesetzen (5)
  x.gleisBesetzen (7)
  x.gleisBesetzen (17)
  x.gleisBesetzen (29)

  nachbar[13] = netz.Bahnhausen
  nachbar[14] = netz.Bahnhausen
  nachbar[33] = netz.Eisenstadt
  nachbar[34] = netz.Eisenstadt
  nachbar[36] = netz.Eisenheim

//  Weiche  Nr  Kilometrierung
//           |  |  Lage
//           |  |  |  Richtung
//           |  |  |  |  Stellung
//           |  |  |  |  | Zeile
//           |  |  |  |  |  |  Spalte
//           |  |  |  |  |  |   | 
  x.weiche ( 1, m, G, R, G, 4,  7)
  x.dkw    ( 2,    R,    G, 5,  8)
  x.weiche ( 3, m, R, L, G, 6,  9)
  x.weiche ( 4, g, G, R, G, 7, 10)
  x.weiche ( 5, m, G, L, G, 5,  9)
  x.dkw    ( 6,    L,    G, 4, 10)
  x.dkw    ( 8,    L,    G, 2, 12)
  x.dkw    ( 9,    L,    G, 1, 13)
  x.weiche (10, g, G, L, G, 0, 14)
  x.weiche (12, m, G, L, G, 7, 25)
  x.dkw    (13,    L,    G, 6, 26)
  x.dkw    (14,    L,    G, 5, 27)
  x.dkw    (15,    L,    G, 4, 28)
  x.weiche (16, g, L, R, G, 3, 29)
  x.weiche (18, m, G, R, G, 0, 27)
  x.dkw    (19,    R,    G, 1, 28)
  x.weiche (20, g, G, R, G, 2, 29)
  x.weiche (21, g, G, L, L, 2, 30)
  x.weiche (22, m, G, R, G, 2, 31)
  x.dkw    (23,    R,    G, 4, 33)
  x.weiche (24, g, G, R, R, 5, 34)
  x.weiche (25, m, G, L, L, 1, 33)
  x.weiche (26, g, G, L, L, 0, 34)

  x.verbinden()
}
