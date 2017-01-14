package enum

// (c) murus.org  v. 140522 - license see murus.go

var
  lTitle, sTitle []string =
  []string {
    "",
    "Dr.", "Dr.med.", "Dr.med.dent.", "Dr.rer.nat.", "Dr.phil.",
    "Dr.iur.", "Dr.med.vet.", "Dr.rer.pol.", "Dr.-Ing.",
    "Prof.Dr.", "Prof.Dr.med.", "Prof.Dr.-Ing.",
  },
  lTitle
const
  NTitles = uint8(13)


func init() {
  l[Title], s[Title] = lTitle, sTitle
  if NTitles != uint8(len(lTitle)) { panic ("enum.Titles wrong") }
}
