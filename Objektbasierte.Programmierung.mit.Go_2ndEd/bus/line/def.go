package line

// (c) Christian Maurer   v. 230114 - license see µU.go

import
  "µU/col"
type
  Line byte; const (
  Footpath = Line(iota)
  U1; U2; U3; U4; U5; U6; U7; U8; U9
  S1
  S2; S25; S26
  S3
  S41; S45; S46; S47
  S5
  S7; S75
  S8; S85
  S9
  Zoo
  BG
  NLines
)
var (
  Text = []string {"F",
                   "U1", "U2", "U3", "U4", "U5", "U6", "U7", "U8", "U9",
                   "S1",
                   "S2", "S25", "S26",
                   "S3",
                   "S41", "S45", "S46", "S47",
                   "S5",
                   "S7", "S75",
                   "S8", "S85",
                   "S9",
                   "Zoo",
                   "BG"}
  Colour = []col.Colour {col.White (),
                         col.New3n ("U1",   85, 184,  49),
                         col.New3n ("U2",  241,  71,  28),
                         col.New3n ("U3",   54, 171, 148),
                         col.New3n ("U4",  253, 210,   5),
                         col.New3n ("U5",  120,  81,  60),
                         col.New3n ("U6",  119,  95, 176),
                         col.New3n ("U7",   57, 159, 223),
                         col.New3n ("U8",   26,  94, 164),
                         col.New3n ("U9",  242,  88,  48),
                         col.New3n ("S1",  119,  95, 176),
                         col.New3n ("S2",   19, 133,  75),
                         col.New3n ("S25", 160, 132,  73),
                         col.New3n ("S26", 160, 132,  73),
                         col.New3n ("S3",   21, 106, 184),
                         col.New3n ("S41", 166,  76,  53),
                         col.New3n ("S45", 196, 141,  66),
                         col.New3n ("S46", 196, 141,  66),
                         col.New3n ("S47", 196, 141,  66),
                         col.New3n ("S5",  244, 103,  23),
                         col.New3n ("S7",  119,  95, 176),
                         col.New3n ("S75", 119,  95, 176),
                         col.New3n ("S8",   85, 184,  49),
                         col.New3n ("S85",  85, 184,  49),
                         col.New3n ("S9",  148,  36,  77),
                         col.New3n ("Zoo",   0, 168,   0),
                         col.New3n ("BG",    0, 168,   0)}
)
