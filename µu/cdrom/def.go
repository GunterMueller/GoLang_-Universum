package cdrom

// (c) Christian Maurer   v. 150622 - license see µu.go

import (
  "os"
  "µu/clk"
)
const
  MaxVol = 1<<8 - 1
type
  Controller uint8; const (Left = iota; Right; Balance; All; NCtrl)
var (
  StartTime, Length []clk.Clocktime
  TotalTime, Time, TrackTime clk.Clocktime = clk.New(), clk.New(), clk.New()
  Ctrltext [NCtrl]string
)

// Returns the open file /dev/cdrom, iff there is a readable
// audio CD or DVD in that device ("cdrom" can be replaced by the
// 1st parameter in the call of the program, that uses this package).
// In this case TotalTime and the StartTime and Length of all
// entries of the CD/DVD are defined. Returns otherwise nil.
// The execution of this function with result != nil
// is the precondition for all further functions.
func Soundfile () *os.File { return soundfile() }

// Returns the number of tracks of the CD.
func NTracks () uint8 { return nTracks }

// Returns the actual track of the CD.
// Time and TrackTime are actualized.
func ActTrack() uint8 { return actTrack() }

// Returns the actual status of the CD.
func String() string { return string_() }

// Returns TODO
func CtrlString (c Controller) string { return ctrltext[c] }

// Returns the actual sound volume of controller c.
func Volume (c Controller) uint8 { return volume(c) }

// Pre: n >= 1.
// The CD plays for 1 <= n <= NTracks from the beginning of the n-th track,
// otherwise from beginning of the last track.
func PlayTrack (n uint8) { playTrack(n) }

// The actual track is for f == true/false one ahead/back.
func PlayTrack1 (f bool) { playTrack1(f) }

// The CD plays from start of the actual track.
func PlayTrack0 () { playTrack0() }

// The CD is playing for f == true/false for s = 0 from start of the next/previous track,
// otherwise from s seconds behind/before.
func PosTime1 (f bool, s uint) { posTime1(f, s) }

// The CD is playing for f == true/false for s <= total time/time of the actual track (in seconds)
// relatively to the beginning of the CD/of the actual track from second s,
// otherwise from two seconds before its end.
func PosTime (f bool, s uint) { posTime(f, s) }

// If the CD was playing before, it has now stopped, otherwise it is now playing.
func Switch() { switch_() }

// For l == true the sound volume of c is one unit higher than before,
// otherwise one unit lower, if that is possible.
func Ctrl1 (c Controller, l bool) { ctrl1(c,l) }

// The controller c has the sound volume v.
func Ctrl (c Controller, v uint8) { ctrl(c,v) }

// CD does not play any more.
func Fin() { fin() }

// CD does not play any more; the CD-ROM-tray is open.
func Fin1() { fin1() }
