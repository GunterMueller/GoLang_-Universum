package enum

// (c) murus.org  v. 140522 - license see murus.go

var
  lAudioMedium, sAudioMedium []string =
  []string { "", "Single Play record", "Long Play record", "Composeract disk",
             "Digital versatile disc", "Super Audio CD", "Blu-ray disc" },
  []string { "", "SP", "LP", "CD", "DVD", "SACD", "BD" }

const
  NAudioMedia = uint8(7)
func init() {
  l[AudioMedium], s[AudioMedium] = lAudioMedium, sAudioMedium
  if NAudioMedia != uint8(len(lAudioMedium)) { panic ("enum.NAudioMedia wrong") }
}
