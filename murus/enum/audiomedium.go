package enum

// (c) Christian Maurer   v. 170419 - license see murus.go

var (
  lAudioMedium, sAudioMedium []string =
  []string { "", "Single Play record", "Long Play record", "Composeract disk",
             "Digital versatile disc", "Super Audio CD", "Blu-ray disc" },
  []string { "", "SP", "LP", "CD", "DVD", "SACD", "BD" }
  NAudioMedia = uint8(len(lAudioMedium))
)

func init() {
  l[AudioMedium], s[AudioMedium] = lAudioMedium, sAudioMedium
}
