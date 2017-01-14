package enum

// (c) murus.org  v. 140522 - license see murus.go

var
  lRecordLabel, sRecordLabel []string =
  []string { "", "2001", "Angel", "BMG", "CBS", "Decca", "Denon", "Deutsche Grammophon",
             "EMI", "Erato", "Harmonia mundi", "Melodia", "Philips", "Polygram", "Sony",
             "Supraphon", "Teldec", "UMG", "Warner", "Zyx" },
  lRecordLabel
const
  NRecordLabels = 20
func init() {
  l[RecordLabel], s[RecordLabel] = lRecordLabel, sRecordLabel
  if NRecordLabels != uint8(len(lRecordLabel)) { panic ("enum.NRecordLabels wrong") }
}
