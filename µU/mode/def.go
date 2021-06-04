package mode

// (c) Christian Maurer   v. 210526 - license see µU.go

type
  Mode byte; const (
  None = Mode(iota) // lines x colums for 8x16-font
  Mini   //  192 x  160   10 x  24
  HQVGA  //  240 x  160   10 x  30
  QVGA   //  320 x  240   15 x  40
  HVGA   //  480 x  320   20 x  60
  TXT    //  640 x  400   25 x  80
  VGA    //  640 x  480   30 x  80
  PAL    //  768 x  576   36 x  96
  WVGA   //  800 x  480   30 x 100
  SVGA   //  800 x  600   37 x 100
  XGA    // 1024 x  768   48 x 128
  HD     // 1280 x  720   45 x 160
  WXGA   // 1280 x  800   50 x 160
  SXVGA  // 1280 x  960   60 x 160
  SXGA   // 1280 x 1024   64 x 160
  WXGA1  // 1366 x  768   48 x 171
  SXGAp  // 1400 x 1050   65 x 175
  WXGAp  // 1440 x  900   56 x 180
  WXGApp // 1600 x  900   56 x 200
  WSXGA  // 1600 x 1024   64 x 200
  UXGA   // 1600 x 1200   75 x 200
  WSXGAp // 1680 x 1050   65 x 210
  FHD    // 1920 x 1080   67 x 240
  WUXGA  // 1920 x 1200   75 x 240
  SUXGA  // 1920 x 1440   90 x 240
  QWXGA  // 2048 x 1152   72 x 256
  QXGA   // 2048 x 1536   96 x 256
  WSUXGA // 2560 x 1440   90 x 320
  WQXGA  // 2560 x 1600  100 x 320
  QSXGAp // 2800 x 2100  131 x 350
  QUXGA  // 3200 x 2400  150 x 400
  UHD    // 3840 x 2160  135 x 440
  HXGA   // 4096 x 3072  192 x 512
  WHXGA  // 5120 x 3200  200 x 640
  HSXGA  // 5120 x 4096  256 x 640
  HUXGA  // 6400 x 4800  300 x 800
  FUHD   // 7680 x 4320  270 x 960
  NEW
  NModes )

// Returns the pixelwidth of m.
func Wd (m Mode) uint { return x[m] }

// Returns the pixelheight of m.
func Ht (m Mode) uint { return y[m] }

// Returns the pixelwidth and -height of m.
func Res (m Mode) (uint, uint) { return x[m], y[m] }

// Returns the mode with (w, h) pixels for (width, height),
// if such exists; panics otherwise.
func ModeOf (w, h uint) Mode { return modeOf(w,h) }
