package enum

// (c) Christian Maurer   v. 170419 - license see murus.go

var (
  lTitle, sTitle []string =
  []string {
    "",
    "Dr.",
    "Dr.med.",
    "Dr.med.dent.",
    "Dr.rer.nat.",
    "Dr.phil.",
    "Dr.iur.",
    "Dr.med.vet.",
    "Dr.rer.pol.",
    "Dr.-Ing.",
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
    "M.Sc.",
  },
  lTitle
  NTitles = uint8(len(lTitle))
)


func init() {
  l[Title], s[Title] = lTitle, sTitle
}
