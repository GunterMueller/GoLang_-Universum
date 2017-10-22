package enum

// (c) Christian Maurer   v. 170419 - license see µU.go

var (
  lRecordLabel, sRecordLabel []string =
  []string { "", "2001", "Angel", "BMG", "CBS", "Decca", "Denon", "Deutsche Grammophon",
             "EMI", "Erato", "Harmonia mundi", "Melodia", "Philips", "Polygram", "Sony",
             "Supraphon", "Teldec", "UMG", "Warner", "Zyx" },
  lRecordLabel
  NRecordLabels = uint8(len(lRecordLabel))
)
func init() {
  l[RecordLabel], s[RecordLabel] = lRecordLabel, sRecordLabel
  if NRecordLabels != uint8(len(lRecordLabel)) { panic ("enum.NRecordLabels wrong") }
}
