package enum

// (c) Christian Maurer   v. 201024 - license see µU.go

var
  title = []string {"",
                    "Dr.",
                    "Dr.med.",
                    "Dr.med.dent.",
                    "Dr.rer.nat.",
                    "Dr.phil.",
                    "Dr.iur.",
                    "Dr.med.vet.",
                    "Dr.rer.pol.",
                    "Dr.-Ing.",
                    "Prof.",
                    "Prof.Dr.",
                    "Prof.Dr.med.",
                    "Prof.Dr.-Ing.",
                    "Dipl.-Math.",
                    "Dipl.-Phys.",
                    "Dipl.-Chem.",
                    "Dipl.-Biol.",
                    "Dipl.-Geol.",
                    "Dipl.-Ing.",
                    "Dipl.-Jur.",
                    "Dipl.-Kfm.",
                    "B.Sc.",
                    "M.Sc.",
                    }

func init() {
  l[Title] = title
  s[Title] = title
}
